package client

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	discovery "github.com/Clever/discovery-go"
	"github.com/Clever/wag/samples/gen-go/models"
	"github.com/afex/hystrix-go/hystrix"
	logger "gopkg.in/Clever/kayvee-go.v5/logger"
)

var _ = json.Marshal
var _ = strings.Replace
var _ = strconv.FormatInt
var _ = bytes.Compare

// WagClient is used to make requests to the swagger-test service.
type WagClient struct {
	basePath    string
	requestDoer doer
	transport   *http.Transport
	timeout     time.Duration
	// Keep the retry doer around so that we can set the number of retries
	retryDoer *retryDoer
	// Keep the circuit doer around so that we can turn it on / off
	circuitDoer    *circuitBreakerDoer
	defaultTimeout time.Duration
	logger         *logger.Logger
}

var _ Client = (*WagClient)(nil)

// New creates a new client. The base path and http transport are configurable.
func New(basePath string) *WagClient {
	base := baseDoer{}
	tracing := tracingDoer{d: base}
	// For the short-term don't use the default retry policy since its 5 retries can 5X
	// the traffic. Once we've enabled circuit breakers by default we can turn it on.
	retry := retryDoer{d: tracing, retryPolicy: SingleRetryPolicy{}}
	logger := logger.New("swagger-test-wagclient")
	circuit := &circuitBreakerDoer{
		d:     &retry,
		debug: true,
		// one circuit for each service + url pair
		circuitName: fmt.Sprintf("swagger-test-%s", shortHash(basePath)),
		logger:      logger,
	}
	circuit.init()
	client := &WagClient{requestDoer: circuit, retryDoer: &retry, circuitDoer: circuit, defaultTimeout: 10 * time.Second,
		transport: &http.Transport{}, basePath: basePath}
	client.SetCircuitBreakerSettings(DefaultCircuitBreakerSettings)
	return client
}

// NewFromDiscovery creates a client from the discovery environment variables. This method requires
// the three env vars: SERVICE_SWAGGER_TEST_HTTP_(HOST/PORT/PROTO) to be set. Otherwise it returns an error.
func NewFromDiscovery() (*WagClient, error) {
	url, err := discovery.URL("swagger-test", "default")
	if err != nil {
		url, err = discovery.URL("swagger-test", "http") // Added fallback to maintain reverse compatibility
		if err != nil {
			return nil, err
		}
	}
	return New(url), nil
}

// SetRetryPolicy sets a the given retry policy for all requests.
func (c *WagClient) SetRetryPolicy(retryPolicy RetryPolicy) {
	c.retryDoer.retryPolicy = retryPolicy
}

// SetCircuitBreakerDebug puts the circuit
func (c *WagClient) SetCircuitBreakerDebug(b bool) {
	c.circuitDoer.debug = b
}

// SetLogger allows for setting a custom logger
func (c *WagClient) SetLogger(logger *logger.Logger) {
	c.logger = logger
	c.circuitDoer.logger = logger
}

// CircuitBreakerSettings are the parameters that govern the client's circuit breaker.
type CircuitBreakerSettings struct {
	// MaxConcurrentRequests is the maximum number of concurrent requests
	// the client can make at the same time. Default: 100.
	MaxConcurrentRequests int
	// RequestVolumeThreshold is the minimum number of requests needed
	// before a circuit can be tripped due to health. Default: 20.
	RequestVolumeThreshold int
	// SleepWindow how long, in milliseconds, to wait after a circuit opens
	// before testing for recovery. Default: 5000.
	SleepWindow int
	// ErrorPercentThreshold is the threshold to place on the rolling error
	// rate. Once the error rate exceeds this percentage, the circuit opens.
	// Default: 90.
	ErrorPercentThreshold int
}

// DefaultCircuitBreakerSettings describes the default circuit parameters.
var DefaultCircuitBreakerSettings = CircuitBreakerSettings{
	MaxConcurrentRequests:  100,
	RequestVolumeThreshold: 20,
	SleepWindow:            5000,
	ErrorPercentThreshold:  90,
}

// SetCircuitBreakerSettings sets parameters on the circuit breaker. It must be
// called on application startup.
func (c *WagClient) SetCircuitBreakerSettings(settings CircuitBreakerSettings) {
	hystrix.ConfigureCommand(c.circuitDoer.circuitName, hystrix.CommandConfig{
		// redundant, with the timeout we set on the context, so set
		// this to something high and irrelevant
		Timeout:                100 * 1000,
		MaxConcurrentRequests:  settings.MaxConcurrentRequests,
		RequestVolumeThreshold: settings.RequestVolumeThreshold,
		SleepWindow:            settings.SleepWindow,
		ErrorPercentThreshold:  settings.ErrorPercentThreshold,
	})
}

// SetTimeout sets a timeout on all operations for the client. To make a single request
// with a timeout use context.WithTimeout as described here: https://godoc.org/golang.org/x/net/context#WithTimeout.
func (c *WagClient) SetTimeout(timeout time.Duration) {
	c.defaultTimeout = timeout
}

// GetBooks makes a GET request to /books
// Returns a list of books
// 200: []models.Book
// 400: *models.BadRequest
// 500: *models.InternalError
// default: client side HTTP errors, for example: context.DeadlineExceeded.
func (c *WagClient) GetBooks(ctx context.Context, i *models.GetBooksInput) ([]models.Book, error) {
	var body []byte

	path, err := i.Path()

	if err != nil {
		return nil, err
	}

	path = c.basePath + path

	client := &http.Client{Transport: c.transport}
	req, err := http.NewRequest("GET", path, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	// Add the opname for doers like tracing
	ctx = context.WithValue(ctx, opNameCtx{}, "getBooks")
	req = req.WithContext(ctx)
	// Don't add the timeout in a "doer" because we don't want to call "defer.cancel()"
	// until we've finished all the processing of the request object. Otherwise we'll cancel
	// our own request before we've finished it.
	if c.defaultTimeout != 0 {
		ctx, cancel := context.WithTimeout(req.Context(), c.defaultTimeout)
		defer cancel()
		req = req.WithContext(ctx)
	}
	resp, err := c.requestDoer.Do(client, req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	switch resp.StatusCode {

	case 200:

		var output []models.Book
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return output, nil

	case 400:

		var output models.BadRequest
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	case 500:

		var output models.InternalError
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	default:
		return nil, &models.InternalError{Message: "Unknown response"}
	}
}

// CreateBook makes a POST request to /books
// Creates a book
// 200: *models.Book
// 400: *models.BadRequest
// 500: *models.InternalError
// default: client side HTTP errors, for example: context.DeadlineExceeded.
func (c *WagClient) CreateBook(ctx context.Context, i *models.Book) (*models.Book, error) {
	var body []byte

	path := c.basePath + "/v1/books"

	if i != nil {

		var err error
		body, err = json.Marshal(i)

		if err != nil {
			return nil, err
		}

	}

	client := &http.Client{Transport: c.transport}
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	// Add the opname for doers like tracing
	ctx = context.WithValue(ctx, opNameCtx{}, "createBook")
	req = req.WithContext(ctx)
	// Don't add the timeout in a "doer" because we don't want to call "defer.cancel()"
	// until we've finished all the processing of the request object. Otherwise we'll cancel
	// our own request before we've finished it.
	if c.defaultTimeout != 0 {
		ctx, cancel := context.WithTimeout(req.Context(), c.defaultTimeout)
		defer cancel()
		req = req.WithContext(ctx)
	}
	resp, err := c.requestDoer.Do(client, req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	switch resp.StatusCode {

	case 200:

		var output models.Book
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return &output, nil

	case 400:

		var output models.BadRequest
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	case 500:

		var output models.InternalError
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	default:
		return nil, &models.InternalError{Message: "Unknown response"}
	}
}

// GetBookByID makes a GET request to /books/{book_id}
// Returns a book
// 200: *models.Book
// 400: *models.BadRequest
// 401: *models.Unathorized
// 404: *models.Error
// 500: *models.InternalError
// default: client side HTTP errors, for example: context.DeadlineExceeded.
func (c *WagClient) GetBookByID(ctx context.Context, i *models.GetBookByIDInput) (*models.Book, error) {
	var body []byte

	path, err := i.Path()

	if err != nil {
		return nil, err
	}

	path = c.basePath + path

	client := &http.Client{Transport: c.transport}
	req, err := http.NewRequest("GET", path, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", i.Authorization)

	// Add the opname for doers like tracing
	ctx = context.WithValue(ctx, opNameCtx{}, "getBookByID")
	req = req.WithContext(ctx)
	// Don't add the timeout in a "doer" because we don't want to call "defer.cancel()"
	// until we've finished all the processing of the request object. Otherwise we'll cancel
	// our own request before we've finished it.
	if c.defaultTimeout != 0 {
		ctx, cancel := context.WithTimeout(req.Context(), c.defaultTimeout)
		defer cancel()
		req = req.WithContext(ctx)
	}
	resp, err := c.requestDoer.Do(client, req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	switch resp.StatusCode {

	case 200:

		var output models.Book
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return &output, nil

	case 400:

		var output models.BadRequest
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	case 401:

		var output models.Unathorized
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	case 404:

		var output models.Error
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	case 500:

		var output models.InternalError
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	default:
		return nil, &models.InternalError{Message: "Unknown response"}
	}
}

// GetBookByID2 makes a GET request to /books2/{id}
// Retrieve a book
// 200: *models.Book
// 400: *models.BadRequest
// 404: *models.Error
// 500: *models.InternalError
// default: client side HTTP errors, for example: context.DeadlineExceeded.
func (c *WagClient) GetBookByID2(ctx context.Context, id string) (*models.Book, error) {
	var body []byte

	path, err := models.GetBookByID2InputPath(id)

	if err != nil {
		return nil, err
	}

	path = c.basePath + path

	client := &http.Client{Transport: c.transport}
	req, err := http.NewRequest("GET", path, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	// Add the opname for doers like tracing
	ctx = context.WithValue(ctx, opNameCtx{}, "getBookByID2")
	req = req.WithContext(ctx)
	// Don't add the timeout in a "doer" because we don't want to call "defer.cancel()"
	// until we've finished all the processing of the request object. Otherwise we'll cancel
	// our own request before we've finished it.
	if c.defaultTimeout != 0 {
		ctx, cancel := context.WithTimeout(req.Context(), c.defaultTimeout)
		defer cancel()
		req = req.WithContext(ctx)
	}
	resp, err := c.requestDoer.Do(client, req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	switch resp.StatusCode {

	case 200:

		var output models.Book
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return &output, nil

	case 400:

		var output models.BadRequest
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	case 404:

		var output models.Error
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	case 500:

		var output models.InternalError
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return nil, &output

	default:
		return nil, &models.InternalError{Message: "Unknown response"}
	}
}

// HealthCheck makes a GET request to /health/check
//
// 200: nil
// 400: *models.BadRequest
// 500: *models.InternalError
// default: client side HTTP errors, for example: context.DeadlineExceeded.
func (c *WagClient) HealthCheck(ctx context.Context) error {
	var body []byte

	path := c.basePath + "/v1/health/check"

	client := &http.Client{Transport: c.transport}
	req, err := http.NewRequest("GET", path, bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	// Add the opname for doers like tracing
	ctx = context.WithValue(ctx, opNameCtx{}, "healthCheck")
	req = req.WithContext(ctx)
	// Don't add the timeout in a "doer" because we don't want to call "defer.cancel()"
	// until we've finished all the processing of the request object. Otherwise we'll cancel
	// our own request before we've finished it.
	if c.defaultTimeout != 0 {
		ctx, cancel := context.WithTimeout(req.Context(), c.defaultTimeout)
		defer cancel()
		req = req.WithContext(ctx)
	}
	resp, err := c.requestDoer.Do(client, req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	switch resp.StatusCode {

	case 200:

		return nil

	case 400:

		var output models.BadRequest
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return err
		}
		return &output

	case 500:

		var output models.InternalError
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return err
		}
		return &output

	default:
		return &models.InternalError{Message: "Unknown response"}
	}
}

func shortHash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))[0:6]
}

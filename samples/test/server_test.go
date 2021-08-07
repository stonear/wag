package test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Clever/wag/samples/v8/gen-go/client"
	"github.com/Clever/wag/samples/v8/gen-go/models"
	"github.com/Clever/wag/samples/v8/gen-go/server"
	"github.com/Clever/wag/v8/swagger"

	"net/http"
	"net/http/httptest"
	"net/url"
)

func init() {
	swagger.InitCustomFormats()
}

func TestBasicEndToEnd(t *testing.T) {
	s, _ := setupServer()
	defer s.Close()

	c := client.New(s.URL)

	bookID := int64(124)
	bookName := "Test"

	createdBook, err := c.CreateBook(
		context.Background(), &models.Book{ID: bookID, Name: bookName})
	assert.NoError(t, err)
	assert.Equal(t, bookID, createdBook.ID)
	assert.Equal(t, bookName, createdBook.Name)

	booksOutput, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	require.Equal(t, 1, len(booksOutput))
	assert.Equal(t, bookID, (booksOutput)[0].ID)
	assert.Equal(t, bookName, (booksOutput)[0].Name)

	singleBook, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: bookID})
	require.NoError(t, err)
	assert.Equal(t, bookID, singleBook.ID)
	assert.Equal(t, bookName, singleBook.Name)
}

func TestNextPageHeader(t *testing.T) {
	s, controller := setupServer()
	defer s.Close()

	controller.pageSize = 2

	c := client.New(s.URL)

	_, err := c.CreateBook(context.Background(), &models.Book{
		ID: int64(1), Name: "Test",
	})
	require.NoError(t, err)

	// Make a raw HTTP request (i.e. don't use the client) so we can check the headers
	resp, err := http.Get(fmt.Sprintf("%s/v1/books", s.URL))
	assert.NoError(t, err)
	resp.Body.Close()
	assert.Equal(t, "", resp.Header.Get("X-Next-Page-Path"))

	controller.pageSize = 1
	_, err = c.CreateBook(context.Background(), &models.Book{
		ID: int64(2), Name: "Second",
	})
	require.NoError(t, err)

	// Now that there is a second page, X-Next-Page-Path should be set
	resp, err = http.Get(fmt.Sprintf("%s/v1/books?available=false", s.URL))
	assert.NoError(t, err)
	resp.Body.Close()

	// can't just use equality because of default parameters in URL
	nextPath := resp.Header.Get("X-Next-Page-Path")
	assert.True(t, strings.HasPrefix(nextPath, "/v1/books?"))
	assert.True(t, strings.Contains(nextPath, "startingAfter=1"))
	assert.True(t, strings.Contains(nextPath, "available=false"))
}

func TestNilArray(t *testing.T) {
	s, _ := setupServer()
	defer s.Close()

	c := client.New(s.URL)
	books, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	require.NoError(t, err)
	assert.Equal(t, 0, len(books))
	assert.False(t, books == nil)
}

func TestNilBody(t *testing.T) {
	s, _ := setupServer()
	defer s.Close()

	c := client.New(s.URL)
	name := "foo"
	_, err := c.GetAuthorsWithPut(context.Background(), &models.GetAuthorsWithPutInput{
		Name:          &name,
		FavoriteBooks: nil, // nil body param
	})
	require.NoError(t, err)
}

func TestNilBodySingleParam(t *testing.T) {
	s, controller := setupServer()
	defer s.Close()

	c := client.New(s.URL)
	_, err := c.PutBook(context.Background(), nil /* nil body param */)
	require.NoError(t, err)
	assert.Equal(t, 1, controller.nilPutBookCount)
}

func TestUserDefinedErrorResponse(t *testing.T) {
	// The 404 generated by the code
	s, _ := setupServer()
	defer s.Close()

	c := client.New(s.URL)

	_, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: 124})
	assert.Error(t, err)
	assert.IsType(t, &models.Error{}, err)
}

func TestDefaultErrorResponse(t *testing.T) {
	s, _ := setupServer()
	defer s.Close()

	c := client.New(s.URL)

	_, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: 400})
	assert.Error(t, err)
	badReq, ok := err.(*models.BadRequest)
	assert.True(t, ok)
	assert.Equal(t, "My 400 failure", badReq.Message)
}

func TestValidationErrorResponse(t *testing.T) {
	s, _ := setupServer()
	defer s.Close()

	c := client.New(s.URL)

	// Book ID should be a multiple of two
	_, err := c.GetBookByID(context.Background(), &models.GetBookByIDInput{BookID: 123})
	assert.Error(t, err)
	assert.IsType(t, &models.BadRequest{}, err)
}

func TestClientSideError(t *testing.T) {
	c := client.New("badServer")

	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	assert.Error(t, err)
	assert.IsType(t, &url.Error{}, err)
}

func TestHeaders(t *testing.T) {
	s, _ := setupServer()
	defer s.Close()

	bookID := int64(124)
	c := client.New(s.URL)
	_, err := c.CreateBook(context.Background(),
		&models.Book{ID: bookID, Name: "test"})
	assert.NoError(t, err)

	// Make a raw HTTP request (i.e. don't use the client) so we can check the headers
	resp, err := http.Get(fmt.Sprintf("%s/v1/books/%d", s.URL, bookID))
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestCustomStringValidation(t *testing.T) {
	s, _ := setupServer()
	defer s.Close()

	c := client.New(s.URL)

	bookID := int64(124)
	_, err := c.CreateBook(context.Background(),
		&models.Book{ID: bookID, Name: "test"})
	assert.NoError(t, err)

	badFormat := "nonMongoFormat"
	_, err = c.GetBookByID(context.Background(),
		&models.GetBookByIDInput{BookID: bookID, AuthorID: &badFormat})
	require.Error(t, err)
	assert.IsType(t, &models.BadRequest{}, err)

	validFormat := "012345678901234567890123"
	_, err = c.GetBookByID(context.Background(),
		&models.GetBookByIDInput{BookID: bookID, AuthorID: &validFormat})
	assert.NoError(t, err)
}

type LastCallServer struct {
	lastState     string
	lastAvailable bool
	lastMaxPages  float64
	lastMinPages  int32
	lastAuthors   []string
}

func (d *LastCallServer) GetBooks(ctx context.Context, input *models.GetBooksInput) ([]models.Book, int64, error) {
	d.lastState = *input.State
	d.lastAvailable = *input.Available
	d.lastMaxPages = *input.MaxPages
	d.lastMinPages = *input.MinPages
	d.lastAuthors = input.Authors
	return []models.Book{}, int64(0), nil
}
func (d *LastCallServer) GetBookByID(ctx context.Context, input *models.GetBookByIDInput) (*models.Book, error) {
	return nil, nil
}
func (d *LastCallServer) GetBookByID2(ctx context.Context, id string) (*models.Book, error) {
	return nil, nil
}
func (d *LastCallServer) CreateBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (d *LastCallServer) PutBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (d *LastCallServer) GetAuthors(ctx context.Context, i *models.GetAuthorsInput) (*models.AuthorsResponse, string, error) {
	return nil, "", nil
}
func (d *LastCallServer) GetAuthorsWithPut(ctx context.Context, i *models.GetAuthorsWithPutInput) (*models.AuthorsResponse, string, error) {
	return nil, "", nil
}
func (d *LastCallServer) HealthCheck(ctx context.Context) error {
	return nil
}

func TestDefaultValue(t *testing.T) {
	d := LastCallServer{}
	s := server.New(&d, ":8080")
	testServer := httptest.NewServer(s.Handler)
	c := client.New(testServer.URL)

	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	require.NoError(t, err)
	assert.Equal(t, "finished", d.lastState)
	assert.True(t, d.lastAvailable)
	assert.Equal(t, 500.5, d.lastMaxPages)
	assert.Equal(t, int32(5), d.lastMinPages)
}

func TestPassInArray(t *testing.T) {
	d := LastCallServer{}
	s := server.New(&d, ":8080")
	testServer := httptest.NewServer(s.Handler)
	c := client.New(testServer.URL)

	_, err := c.GetBooks(context.Background(),
		&models.GetBooksInput{Authors: []string{"author1", "author2"}})
	require.NoError(t, err)
	require.Equal(t, 2, len(d.lastAuthors))
	assert.Equal(t, "author1", d.lastAuthors[0])
	assert.Equal(t, "author2", d.lastAuthors[1])
}

type MiddlewareTest struct {
	foundKey string
}

func (m *MiddlewareTest) GetBooks(ctx context.Context, input *models.GetBooksInput) ([]models.Book, int64, error) {
	m.foundKey = ctx.Value(testContextKey{}).(string)
	return []models.Book{}, int64(0), nil
}
func (m *MiddlewareTest) GetBookByID(ctx context.Context, input *models.GetBookByIDInput) (*models.Book, error) {
	return nil, nil
}
func (m *MiddlewareTest) GetBookByID2(ctx context.Context, id string) (*models.Book, error) {
	return nil, nil
}
func (m *MiddlewareTest) CreateBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (m *MiddlewareTest) PutBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (m *MiddlewareTest) GetAuthors(ctx context.Context, i *models.GetAuthorsInput) (*models.AuthorsResponse, string, error) {
	return nil, "", nil
}
func (m *MiddlewareTest) GetAuthorsWithPut(ctx context.Context, i *models.GetAuthorsWithPutInput) (*models.AuthorsResponse, string, error) {
	return nil, "", nil
}
func (m *MiddlewareTest) HealthCheck(ctx context.Context) error {
	return nil
}

type testContextKey struct{}

func testContextMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newCtx := context.WithValue(r.Context(), testContextKey{}, "contextValue")
		h.ServeHTTP(w, r.WithContext(newCtx))
	})
}

type orderingMiddleware struct {
	overallCount int
	first        int
	second       int
}

func (f *orderingMiddleware) First(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.first = f.overallCount
		f.overallCount++
		h.ServeHTTP(w, r)
	})
}

func (f *orderingMiddleware) Second(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.second = f.overallCount
		f.overallCount++
		h.ServeHTTP(w, r)
	})
}

func TestMiddleware(t *testing.T) {
	controller := MiddlewareTest{}
	ordering := orderingMiddleware{}
	s := server.NewWithMiddleware(&controller, "", []func(http.Handler) http.Handler{
		ordering.First, ordering.Second, testContextMiddleware})
	testServer := httptest.NewServer(s.Handler)
	c := client.New(testServer.URL)
	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})

	// Setting context
	assert.NoError(t, err)
	assert.Equal(t, "contextValue", controller.foundKey)

	// Running in the right order
	assert.Equal(t, 0, ordering.first)
	assert.Equal(t, 1, ordering.second)
}

type TimeoutController struct{}

func (m *TimeoutController) GetBooks(ctx context.Context, input *models.GetBooksInput) ([]models.Book, int64, error) {
	var books []models.Book
	for i := 0; i < 1000; i++ {
		books = append(books, models.Book{Name: "testing"})
	}
	time.Sleep(100 * time.Millisecond)
	return books, int64(0), nil
}
func (m *TimeoutController) GetBookByID(ctx context.Context, input *models.GetBookByIDInput) (*models.Book, error) {
	return nil, nil
}
func (m *TimeoutController) GetBookByID2(ctx context.Context, id string) (*models.Book, error) {
	return nil, nil
}
func (m *TimeoutController) CreateBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (m *TimeoutController) PutBook(ctx context.Context, input *models.Book) (*models.Book, error) {
	return nil, nil
}
func (m *TimeoutController) GetAuthors(ctx context.Context, i *models.GetAuthorsInput) (*models.AuthorsResponse, string, error) {
	return nil, "", nil
}
func (m *TimeoutController) GetAuthorsWithPut(ctx context.Context, i *models.GetAuthorsWithPutInput) (*models.AuthorsResponse, string, error) {
	return nil, "", nil
}
func (m *TimeoutController) HealthCheck(ctx context.Context) error {
	return nil
}

func TestTimeout(t *testing.T) {
	s := server.New(&TimeoutController{}, "")
	testServer := httptest.NewServer(testContextMiddleware(s.Handler))
	c := client.New(testServer.URL)

	// Without a timeout, no error
	_, err := c.GetBooks(context.Background(), &models.GetBooksInput{})
	assert.NoError(t, err)

	// Add a per request context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err = c.GetBooks(ctx, &models.GetBooksInput{})
	assert.Error(t, err)
	assert.IsType(t, context.DeadlineExceeded, err)
	end := time.Now()
	assert.True(t, end.Sub(start) < 20*time.Millisecond)

	// Try with a global client setting
	c.SetTimeout(10 * time.Millisecond)
	_, err = c.GetBooks(context.Background(), &models.GetBooksInput{})
	require.Error(t, err)
	assert.IsType(t, context.DeadlineExceeded, err)

	// TODO: Ideally this would actually take effect
	// Adding a higher per request context timeout has no effect
	ctx, cancel = context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	start = time.Now()
	_, err = c.GetBooks(ctx, &models.GetBooksInput{})
	assert.Error(t, err)
	assert.IsType(t, context.DeadlineExceeded, err)
	end = time.Now()
	assert.True(t, end.Sub(start) < 20*time.Millisecond)
}

func TestOmitEmpty(t *testing.T) {
	m, _ := json.Marshal(models.OmitEmpty{
		ArrayFieldNotOmitted: nil,
		ArrayFieldOmitted:    nil,
	})
	require.Equal(t, `{"arrayFieldNotOmitted":null}`, string(m))
}

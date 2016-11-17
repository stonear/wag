package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Clever/wag/samples/gen-go-nils/models"
	"github.com/go-errors/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"gopkg.in/Clever/kayvee-go.v5/logger"
)

var _ = strconv.ParseInt
var _ = strfmt.Default
var _ = swag.ConvertInt32
var _ = errors.New
var _ = mux.Vars
var _ = bytes.Compare
var _ = ioutil.ReadAll

var formats = strfmt.Default
var _ = formats

// convertBase64 takes in a string and returns a strfmt.Base64 if the input
// is valid base64 and an error otherwise.
func convertBase64(input string) (strfmt.Base64, error) {
	temp, err := formats.Parse("byte", input)
	if err != nil {
		return strfmt.Base64{}, err
	}
	return *temp.(*strfmt.Base64), nil
}

// convertDateTime takes in a string and returns a strfmt.DateTime if the input
// is a valid DateTime and an error otherwise.
func convertDateTime(input string) (strfmt.DateTime, error) {
	temp, err := formats.Parse("date-time", input)
	if err != nil {
		return strfmt.DateTime{}, err
	}
	return *temp.(*strfmt.DateTime), nil
}

// convertDate takes in a string and returns a strfmt.Date if the input
// is a valid Date and an error otherwise.
func convertDate(input string) (strfmt.Date, error) {
	temp, err := formats.Parse("date", input)
	if err != nil {
		return strfmt.Date{}, err
	}
	return *temp.(*strfmt.Date), nil
}

func jsonMarshalNoError(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		// This should never happen
		return ""
	}
	return string(bytes)
}

// statusCodeForNilCheck returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForNilCheck(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	default:
		return -1
	}
}

func (h handler) NilCheckHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newNilCheckInput(r)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = h.NilCheck(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForNilCheck(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(""))

}

// newNilCheckInput takes in an http.Request an returns the input struct.
func newNilCheckInput(r *http.Request) (*models.NilCheckInput, error) {
	var input models.NilCheckInput

	var err error
	_ = err

	pathParam := mux.Vars(r)["id"]
	var idStrs []string
	if len(pathParam) > 0 {
		idStrs = []string{pathParam}
	}

	if len(idStrs) == 0 {
		return nil, errors.New("parameter must be specified")
	}

	if len(idStrs) > 0 {
		iDStr := idStrs[0]
		var iDTmp string
		iDTmp, err = iDStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.ID = iDTmp
	}

	queryStrs := r.URL.Query()["query"]

	if len(queryStrs) > 0 {
		queryStr := queryStrs[0]
		var queryTmp string
		queryTmp, err = queryStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Query = &queryTmp
	}

	headerStrs := r.Header["header"]

	if len(headerStrs) > 0 {
		headerStr := headerStrs[0]
		var headerTmp string
		headerTmp, err = headerStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Header = &headerTmp
	}

	data, err := ioutil.ReadAll(r.Body)

	if len(data) > 0 {
		input.Body = &models.NilFields{}
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(input.Body); err != nil {
			return nil, err
		}
	}

	return &input, nil
}

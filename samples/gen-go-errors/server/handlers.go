package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Clever/wag/samples/gen-go-errors/models"
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

// statusCodeForGetBook returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetBook(obj interface{}) int {

	switch obj.(type) {

	case *models.ExtendedError:
		return 400

	case *models.InternalError:
		return 500

	case *models.NotFoundError:
		return 404

	case models.ExtendedError:
		return 400

	case models.InternalError:
		return 500

	case models.NotFoundError:
		return 404

	default:
		return -1
	}
}

func (h handler) GetBookHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newGetBookInput(r)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.ExtendedError{Msg: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.ExtendedError{Msg: err.Error()}), http.StatusBadRequest)
		return
	}

	err = h.GetBook(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForGetBook(err)
		if statusCode == -1 {
			err = models.InternalError{Msg: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(""))

}

// newGetBookInput takes in an http.Request an returns the input struct.
func newGetBookInput(r *http.Request) (*models.GetBookInput, error) {
	var input models.GetBookInput

	var err error
	_ = err

	iDStr := mux.Vars(r)["id"]
	if len(iDStr) == 0 {
		return nil, errors.New("Parameter must be specified")
	}
	if len(iDStr) != 0 {
		var iDTmp int64
		iDTmp, err = swag.ConvertInt64(iDStr)
		if err != nil {
			return nil, err
		}
		input.ID = iDTmp

	}

	return &input, nil
}
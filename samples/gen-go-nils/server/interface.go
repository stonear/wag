package server

import (
	"context"

	"github.com/Clever/wag/samples/gen-go-nils/models/v9"
)

// Controller defines the interface for the nil-test service.
type Controller interface {

	// NilCheck handles POST requests to /check/{id}
	// Nil check tests
	// 200: nil
	// 400: *models.BadRequest
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	NilCheck(ctx context.Context, i *models.NilCheckInput) error
}

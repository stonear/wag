package server

import (
	"context"

	"github.com/Clever/wag/v5/samples/gen-go-deprecated/models"
)

//go:generate mockgen -source=$GOFILE -destination=mock_controller.go -package=server

// Controller defines the interface for the swagger-test service.
type Controller interface {

	// Health handles GET requests to /health
	//
	// 200: nil
	// 400: *models.BadRequest
	// 404: *models.NotFound
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	Health(ctx context.Context, i *models.HealthInput) error
}

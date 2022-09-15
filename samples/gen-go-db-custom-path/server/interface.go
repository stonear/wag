package server

import (
	"context"
)

// Controller defines the interface for the swagger-test service.
type Controller interface {

	// HealthCheck handles GET requests to /health/check
	//
	// 200: nil
	// 400: *models.BadRequest
	// 500: *models.InternalError
	// default: client side HTTP errors, for example: context.DeadlineExceeded.
	HealthCheck(ctx context.Context) error
}

package bookingdetail

import (
	"nakarin-studio/app/modules/net/httpx"

	"go.opentelemetry.io/otel/trace"
)

type Controller struct {
	tracer trace.Tracer
	svc    *Service
	cli    *httpx.Client
}

func newController(trace trace.Tracer, svc *Service) *Controller {
	return &Controller{tracer: trace, svc: svc, cli: httpx.NewClient()}
}

type IDUri struct {
	ID string `uri:"id" binding:"required"`
}

type CreateRequest struct {
	BookingID string  `json:"booking_id" binding:"required"`
	FirstName string  `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name"`
	Phone     string  `json:"phone" binding:"required"`
}

type UpdateRequest struct {
	BookingID *string `json:"booking_id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
}

type ListQuery struct {
	BookingID *string `form:"booking_id"`
}

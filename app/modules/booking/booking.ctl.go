package booking

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
	BookingNo string  `json:"booking_no" binding:"required"`
	Status    *string `json:"status"`
	Payment   *string `json:"payment"`
}

type UpdateRequest struct {
	BookingNo *string `json:"booking_no"`
	Status    *string `json:"status"`
	Payment   *string `json:"payment"`
}

type ListQuery struct {
	Status  *string `form:"status"`
	Payment *string `form:"payment"`
}

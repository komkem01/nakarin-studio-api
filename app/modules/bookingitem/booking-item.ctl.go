package bookingitem

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
	BookingID          string  `json:"booking_id" binding:"required"`
	ProductID          string  `json:"product_id" binding:"required"`
	ProductName        string  `json:"product_name" binding:"required"`
	UnitPriceAtBooking float64 `json:"unit_price_at_booking" binding:"required"`
	Quantity           int     `json:"quantity" binding:"required"`
	LineTotal          float64 `json:"line_total" binding:"required"`
	Note               *string `json:"note"`
	SortOrder          *int    `json:"sort_order"`
}

type UpdateRequest struct {
	BookingID          *string  `json:"booking_id"`
	ProductID          *string  `json:"product_id"`
	ProductName        *string  `json:"product_name"`
	UnitPriceAtBooking *float64 `json:"unit_price_at_booking"`
	Quantity           *int     `json:"quantity"`
	LineTotal          *float64 `json:"line_total"`
	Note               *string  `json:"note"`
	SortOrder          *int     `json:"sort_order"`
}

type ListQuery struct {
	BookingID *string `form:"booking_id"`
	ProductID *string `form:"product_id"`
}

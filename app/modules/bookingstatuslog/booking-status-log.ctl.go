package bookingstatuslog

import (
	"time"

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
	BookingID     string     `json:"booking_id" binding:"required"`
	FromStatus    *string    `json:"from_status"`
	ToStatus      string     `json:"to_status" binding:"required"`
	ChangedBy     *string    `json:"changed_by"`
	ChangedByRole *string    `json:"changed_by_role"`
	Reason        *string    `json:"reason"`
	ChangedAt     *time.Time `json:"changed_at"`
}

type UpdateRequest struct {
	BookingID     *string    `json:"booking_id"`
	FromStatus    *string    `json:"from_status"`
	ToStatus      *string    `json:"to_status"`
	ChangedBy     *string    `json:"changed_by"`
	ChangedByRole *string    `json:"changed_by_role"`
	Reason        *string    `json:"reason"`
	ChangedAt     *time.Time `json:"changed_at"`
}

type ListQuery struct {
	BookingID *string `form:"booking_id"`
	ToStatus  *string `form:"to_status"`
}

package memberbooking

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
	MemberID  string `json:"member_id" binding:"required"`
	BookingID string `json:"booking_id" binding:"required"`
}

type UpdateRequest struct {
	MemberID  *string `json:"member_id"`
	BookingID *string `json:"booking_id"`
}

type ListQuery struct {
	MemberID  *string `form:"member_id"`
	BookingID *string `form:"booking_id"`
}

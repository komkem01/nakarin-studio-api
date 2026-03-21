package payment

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
	Channel       *string    `json:"channel"`
	Amount        float64    `json:"amount" binding:"required"`
	DepositAmount float64    `json:"deposit_amount" binding:"required"`
	Status        *string    `json:"status"`
	ProofURL      *string    `json:"proof_url"`
	Note          *string    `json:"note"`
	PaidAt        *time.Time `json:"paid_at"`
}

type UpdateRequest struct {
	BookingID     *string    `json:"booking_id"`
	Channel       *string    `json:"channel"`
	Amount        *float64   `json:"amount"`
	DepositAmount *float64   `json:"deposit_amount"`
	Status        *string    `json:"status"`
	ProofURL      *string    `json:"proof_url"`
	Note          *string    `json:"note"`
	PaidAt        *time.Time `json:"paid_at"`
}

type ListQuery struct {
	BookingID *string `form:"booking_id"`
	Channel   *string `form:"channel"`
	Status    *string `form:"status"`
}

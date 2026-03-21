package member

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
	GenderID  string  `json:"gender_id" binding:"required"`
	PrefixID  *string `json:"prefix_id"`
	Role      *string `json:"role"`
	FirstName string  `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name"`
	Phone     string  `json:"phone" binding:"required"`
}

type UpdateRequest struct {
	GenderID  *string `json:"gender_id"`
	PrefixID  *string `json:"prefix_id"`
	Role      *string `json:"role"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
}

type ListQuery struct {
	GenderID *string `form:"gender_id"`
	PrefixID *string `form:"prefix_id"`
	Role     *string `form:"role"`
	Phone    *string `form:"phone"`
}

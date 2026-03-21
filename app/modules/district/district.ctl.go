package district

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
	ProvinceID string `json:"province_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	IsActive   *bool  `json:"is_active"`
}

type UpdateRequest struct {
	ProvinceID *string `json:"province_id"`
	Name       *string `json:"name"`
	IsActive   *bool   `json:"is_active"`
}

type ListQuery struct {
	ProvinceID *string `form:"province_id"`
	IsActive   *bool   `form:"is_active"`
}

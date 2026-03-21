package zipcode

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
	SubDistrictID string `json:"sub_district_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	IsActive      *bool  `json:"is_active"`
}

type UpdateRequest struct {
	SubDistrictID *string `json:"sub_district_id"`
	Name          *string `json:"name"`
	IsActive      *bool   `json:"is_active"`
}

type ListQuery struct {
	SubDistrictID *string `form:"sub_district_id"`
	IsActive      *bool   `form:"is_active"`
}

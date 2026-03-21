package productimage

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
	ProductID string  `json:"product_id" binding:"required"`
	ImageURL  string  `json:"image_url" binding:"required"`
	AltText   *string `json:"alt_text"`
	SortOrder *int    `json:"sort_order"`
	IsActive  *bool   `json:"is_active"`
}

type UpdateRequest struct {
	ProductID *string `json:"product_id"`
	ImageURL  *string `json:"image_url"`
	AltText   *string `json:"alt_text"`
	SortOrder *int    `json:"sort_order"`
	IsActive  *bool   `json:"is_active"`
}

type ListQuery struct {
	ProductID *string `form:"product_id"`
	IsActive  *bool   `form:"is_active"`
}

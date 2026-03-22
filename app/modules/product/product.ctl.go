package product

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
	CategoryID    string  `json:"category_id" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Description   *string `json:"description"`
	SuitableFor   *string `json:"suitable_for"`
	OnSite        *string `json:"on_site"`
	ReceivedItems *string `json:"received_items"`
	Note          *string `json:"note"`
	Price         float64 `json:"price" binding:"required"`
	IsActive      *bool   `json:"is_active"`
	IsAvailable   *bool   `json:"is_available"`
	PrepTime      *int    `json:"prep_time"`
	SortOrder     *int    `json:"sort_order"`
}

type UpdateRequest struct {
	CategoryID    *string  `json:"category_id"`
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	SuitableFor   *string  `json:"suitable_for"`
	OnSite        *string  `json:"on_site"`
	ReceivedItems *string  `json:"received_items"`
	Note          *string  `json:"note"`
	Price         *float64 `json:"price"`
	IsActive      *bool    `json:"is_active"`
	IsAvailable   *bool    `json:"is_available"`
	PrepTime      *int     `json:"prep_time"`
	SortOrder     *int     `json:"sort_order"`
}

type ListQuery struct {
	CategoryID  *string `form:"category_id"`
	Name        *string `form:"name"`
	IsActive    *bool   `form:"is_active"`
	IsAvailable *bool   `form:"is_available"`
}

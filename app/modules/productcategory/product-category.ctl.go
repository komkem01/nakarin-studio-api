package productcategory

import (
	"nakarin-studio/app/modules/net/httpx"
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
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
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

type UpdateRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

type ListQuery struct {
	IsActive *bool `form:"is_active"`
}

func (c *Controller) Create(ctx *gin.Context) {
	var req CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	_, err := c.svc.Create(ctx.Request.Context(), req.Name, req.Description, isActive)
	if err != nil {
		base.InternalServerError(ctx, "product-category-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

func (c *Controller) List(ctx *gin.Context) {
	var req ListQuery
	if err := ctx.ShouldBindQuery(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	items, err := c.svc.List(ctx.Request.Context(), req.IsActive)
	if err != nil {
		base.InternalServerError(ctx, "product-category-list-failed", nil)
		return
	}

	base.Success(ctx, items, "success")
}

func (c *Controller) Info(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	item, err := c.svc.GetByID(ctx.Request.Context(), uri.ID)
	if err != nil {
		base.InternalServerError(ctx, "product-category-info-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

func (c *Controller) Update(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	var req UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	item, err := c.svc.UpdateByID(ctx.Request.Context(), uri.ID, req.Name, req.Description, req.IsActive)
	if err != nil {
		base.InternalServerError(ctx, "product-category-update-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

func (c *Controller) Delete(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	if err := c.svc.DeleteByID(ctx.Request.Context(), uri.ID); err != nil {
		base.InternalServerError(ctx, "product-category-delete-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

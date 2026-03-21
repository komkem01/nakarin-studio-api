package product

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

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
	isAvailable := true
	if req.IsAvailable != nil {
		isAvailable = *req.IsAvailable
	}
	sortOrder := 0
	if req.SortOrder != nil {
		sortOrder = *req.SortOrder
	}
	prepTime := 0
	if req.PrepTime != nil {
		prepTime = *req.PrepTime
	}

	_, err := c.svc.Create(ctx.Request.Context(), req.Name, req.Description, req.Price, isActive, isAvailable, prepTime, sortOrder)
	if err != nil {
		base.InternalServerError(ctx, "product-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

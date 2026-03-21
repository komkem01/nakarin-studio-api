package productimage

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

	sortOrder := 0
	if req.SortOrder != nil {
		sortOrder = *req.SortOrder
	}
	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	_, err := c.svc.Create(ctx.Request.Context(), req.ProductID, req.ImageURL, req.AltText, sortOrder, isActive)
	if err != nil {
		base.InternalServerError(ctx, "product-image-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

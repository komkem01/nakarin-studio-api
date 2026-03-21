package province

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	IsActive *bool  `json:"is_active"`
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
	_, err := c.svc.Create(ctx.Request.Context(), req.Name, isActive)
	if err != nil {
		base.InternalServerError(ctx, "province-create-failed", nil)
		return
	}
	base.Success(ctx, nil, "success")
}

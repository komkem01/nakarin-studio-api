package admin

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

	_, err := c.svc.Create(ctx.Request.Context(), req.MemberID, req.Username, req.PasswordHash, req.DisplayName, req.LastLoginAt, isActive)
	if err != nil {
		base.InternalServerError(ctx, "admin-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

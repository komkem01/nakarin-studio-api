package admin

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	admin, err := c.svc.Authenticate(ctx.Request.Context(), req.Username, req.Password)
	if err != nil {
		base.Unauthorized(ctx, "admin-auth-failed", nil)
		return
	}

	base.Success(ctx, gin.H{"admin_id": admin.ID.String(), "username": admin.Username, "role": "admin"}, "success")
}

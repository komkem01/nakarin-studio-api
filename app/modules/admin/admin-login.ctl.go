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

	auth, err := c.svc.Login(ctx.Request.Context(), req.Username, req.Password)
	if err != nil {
		base.Unauthorized(ctx, "admin-auth-failed", nil)
		return
	}

	base.Success(ctx, gin.H{"access_token": auth.AccessToken, "refresh_token": auth.RefreshToken, "admin_id": auth.Admin.ID.String(), "username": auth.Admin.Username, "role": "admin"}, "success")
}

func (c *Controller) RefreshToken(ctx *gin.Context) {
	var req RefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	auth, err := c.svc.Refresh(ctx.Request.Context(), req.RefreshToken)
	if err != nil {
		base.Unauthorized(ctx, "admin-refresh-failed", nil)
		return
	}

	base.Success(ctx, gin.H{"access_token": auth.AccessToken, "refresh_token": auth.RefreshToken, "admin_id": auth.Admin.ID.String(), "username": auth.Admin.Username, "role": "admin"}, "success")
}

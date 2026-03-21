package admin

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

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

	item, err := c.svc.UpdateByID(ctx.Request.Context(), uri.ID, req.MemberID, req.Username, req.PasswordHash, req.DisplayName, req.LastLoginAt, req.IsActive)
	if err != nil {
		base.InternalServerError(ctx, "admin-update-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

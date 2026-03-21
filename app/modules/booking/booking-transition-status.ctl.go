package booking

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) TransitionStatus(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	var req TransitionStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	adminID, _ := ctx.Get("admin_id")
	adminIDStr, _ := adminID.(string)

	if err := c.svc.TransitionStatus(ctx.Request.Context(), uri.ID, *req.Status, req.Reason, stringOrNil(adminIDStr)); err != nil {
		base.InternalServerError(ctx, "booking-transition-status-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

func stringOrNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

package bookingstatuslog

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

	_, err := c.svc.Create(ctx.Request.Context(), req.BookingID, req.FromStatus, req.ToStatus, req.ChangedBy, req.ChangedByRole, req.Reason, req.ChangedAt)
	if err != nil {
		base.InternalServerError(ctx, "booking-status-log-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

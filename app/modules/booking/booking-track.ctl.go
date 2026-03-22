package booking

import (
	"database/sql"
	"errors"

	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Track(ctx *gin.Context) {
	var req TrackingQuery
	if err := ctx.ShouldBindQuery(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	item, err := c.svc.TrackByBookingNoAndPhone(ctx.Request.Context(), req.BookingNo, req.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, "booking-not-found", nil)
			return
		}
		base.InternalServerError(ctx, "booking-track-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

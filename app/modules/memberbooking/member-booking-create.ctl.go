package memberbooking

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

	_, err := c.svc.Create(ctx.Request.Context(), req.MemberID, req.BookingID)
	if err != nil {
		base.InternalServerError(ctx, "member-booking-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

package bookingdetail

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

	_, err := c.svc.Create(ctx.Request.Context(), req.BookingID, req.FirstName, req.LastName, req.Phone)
	if err != nil {
		base.InternalServerError(ctx, "booking-detail-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

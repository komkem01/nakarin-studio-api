package booking

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

	_, err := c.svc.Create(
		ctx.Request.Context(),
		req.BookingNo,
		req.Status,
		req.Payment,
		req.CancelledReason,
		req.InternalNote,
		req.TrackingAttemptCount,
		req.LastTrackingAt,
		req.DeliveryMemberAddressID,
		req.DeliveryFirstName,
		req.DeliveryLastName,
		req.DeliveryPhone,
		req.DeliveryNo,
		req.DeliveryVillage,
		req.DeliveryStreet,
		req.DeliveryProvinceID,
		req.DeliveryDistrictID,
		req.DeliverySubDistrictID,
		req.DeliveryZipcodeID,
		req.DeliveryNote,
	)
	if err != nil {
		base.InternalServerError(ctx, "booking-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

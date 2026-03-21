package booking

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

	item, err := c.svc.UpdateByID(
		ctx.Request.Context(),
		uri.ID,
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
		base.InternalServerError(ctx, "booking-update-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

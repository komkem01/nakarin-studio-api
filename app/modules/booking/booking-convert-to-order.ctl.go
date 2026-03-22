package booking

import (
	"errors"

	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ConvertToOrder(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	var req ConvertToOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	adminID, _ := ctx.Get("admin_id")
	adminIDStr, _ := adminID.(string)

	order, err := c.svc.ConvertToOrder(ctx.Request.Context(), uri.ID, req.Reason, stringOrNil(adminIDStr))
	if err != nil {
		if errors.Is(err, ErrCannotConvertCanceledBooking) {
			base.BadRequest(ctx, "booking-convert-to-order-invalid", nil)
			return
		}
		base.InternalServerError(ctx, "booking-convert-to-order-failed", nil)
		return
	}

	base.Success(ctx, ConvertToOrderResponse{
		OrderID: order.ID.String(),
		OrderNo: order.OrderNo,
	}, "success")
}

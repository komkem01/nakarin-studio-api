package payment

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

	item, err := c.svc.UpdateByID(ctx.Request.Context(), uri.ID, req.BookingID, req.Channel, req.Amount, req.DepositAmount, req.Status, req.ProofURL, req.Note, req.PaidAt)
	if err != nil {
		base.InternalServerError(ctx, "payment-update-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

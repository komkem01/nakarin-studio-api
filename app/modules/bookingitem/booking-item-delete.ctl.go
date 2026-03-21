package bookingitem

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Delete(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	if err := c.svc.DeleteByID(ctx.Request.Context(), uri.ID); err != nil {
		base.InternalServerError(ctx, "booking-item-delete-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

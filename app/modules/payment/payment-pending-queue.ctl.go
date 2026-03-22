package payment

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) PendingQueue(ctx *gin.Context) {
	items, err := c.svc.PendingQueue(ctx.Request.Context())
	if err != nil {
		base.InternalServerError(ctx, "payment-pending-queue-failed", nil)
		return
	}

	base.Success(ctx, items, "success")
}

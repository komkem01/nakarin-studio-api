package booking

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) AggregateCreate(ctx *gin.Context) {
	var req AggregateCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	if err := c.svc.AggregateCreate(ctx.Request.Context(), &req); err != nil {
		base.InternalServerError(ctx, "booking-aggregate-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

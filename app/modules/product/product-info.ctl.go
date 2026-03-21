package product

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Info(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	item, err := c.svc.GetByID(ctx.Request.Context(), uri.ID)
	if err != nil {
		base.InternalServerError(ctx, "product-info-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

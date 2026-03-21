package product

import (
	"nakarin-studio/app/utils/base"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Delete(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	if err := c.svc.DeleteByID(ctx.Request.Context(), uri.ID); err != nil {
		base.InternalServerError(ctx, "product-delete-failed", nil)
		return
	}

	ctx.Status(http.StatusNoContent)
}

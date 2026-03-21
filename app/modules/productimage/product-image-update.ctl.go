package productimage

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

	item, err := c.svc.UpdateByID(ctx.Request.Context(), uri.ID, req.ProductID, req.ImageURL, req.AltText, req.SortOrder, req.IsActive)
	if err != nil {
		base.InternalServerError(ctx, "product-image-update-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

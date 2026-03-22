package payment

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) ProofViewURL(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	proofURL, err := c.svc.ProofViewURL(ctx.Request.Context(), uri.ID)
	if err != nil {
		base.InternalServerError(ctx, "payment-proof-url-failed", nil)
		return
	}

	base.Success(ctx, gin.H{"url": proofURL}, "success")
}

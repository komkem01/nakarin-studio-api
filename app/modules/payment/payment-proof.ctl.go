package payment

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

type ProofRequest struct {
	ProofURL string `json:"proof_url" binding:"required"`
}

type ApproveRejectRequest struct {
	Note *string `json:"note"`
}

func (c *Controller) UploadProof(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}
	var req ProofRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}
	if err := c.svc.UploadProof(ctx.Request.Context(), uri.ID, req.ProofURL); err != nil {
		base.InternalServerError(ctx, "payment-upload-proof-failed", nil)
		return
	}
	base.Success(ctx, nil, "success")
}

func (c *Controller) Approve(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}
	var req ApproveRejectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}
	if err := c.svc.Approve(ctx.Request.Context(), uri.ID, req.Note); err != nil {
		base.InternalServerError(ctx, "payment-approve-failed", nil)
		return
	}
	base.Success(ctx, nil, "success")
}

func (c *Controller) Reject(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}
	var req ApproveRejectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}
	if err := c.svc.Reject(ctx.Request.Context(), uri.ID, req.Note); err != nil {
		base.InternalServerError(ctx, "payment-reject-failed", nil)
		return
	}
	base.Success(ctx, nil, "success")
}

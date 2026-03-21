package payment

import (
	"errors"
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

type ApproveRejectRequest struct {
	Note *string `json:"note"`
}

func (c *Controller) UploadProof(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	fileHeader, err := ctx.FormFile("proof")
	if err != nil {
		base.BadRequest(ctx, "proof-file-required", nil)
		return
	}

	proofURL, err := c.svc.UploadProof(ctx.Request.Context(), uri.ID, fileHeader)
	if err != nil {
		if errors.Is(err, ErrProofFileInvalid) {
			base.BadRequest(ctx, "invalid-proof-file", nil)
			return
		}
		base.InternalServerError(ctx, "payment-upload-proof-failed", nil)
		return
	}
	base.Success(ctx, gin.H{"proof_url": proofURL}, "success")
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

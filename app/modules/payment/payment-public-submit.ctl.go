package payment

import (
	"errors"
	"strconv"
	"time"

	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) PublicSubmit(ctx *gin.Context) {
	bookingNo := ctx.PostForm("booking_no")
	phone := ctx.PostForm("phone")
	channelValue := ctx.PostForm("channel")
	amountRaw := ctx.PostForm("amount")
	depositRaw := ctx.PostForm("deposit_amount")
	noteValue := ctx.PostForm("note")
	paidAtRaw := ctx.PostForm("paid_at")
	proofURLValue := ctx.PostForm("proof_url")

	if bookingNo == "" || phone == "" || amountRaw == "" || depositRaw == "" {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	amount, err := strconv.ParseFloat(amountRaw, 64)
	if err != nil {
		base.BadRequest(ctx, "invalid-amount", nil)
		return
	}

	depositAmount, err := strconv.ParseFloat(depositRaw, 64)
	if err != nil {
		base.BadRequest(ctx, "invalid-deposit-amount", nil)
		return
	}

	var channel *string
	if channelValue != "" {
		channel = &channelValue
	}

	var note *string
	if noteValue != "" {
		note = &noteValue
	}

	var proofURL *string
	if proofURLValue != "" {
		proofURL = &proofURLValue
	}

	var paidAt *time.Time
	if paidAtRaw != "" {
		parsed, err := time.Parse(time.RFC3339, paidAtRaw)
		if err != nil {
			base.BadRequest(ctx, "invalid-paid-at", nil)
			return
		}
		paidAt = &parsed
	}

	proofFile, err := ctx.FormFile("proof")
	if err != nil {
		proofFile = nil
	}

	item, err := c.svc.SubmitPublicPayment(
		ctx.Request.Context(),
		bookingNo,
		phone,
		channel,
		amount,
		depositAmount,
		note,
		paidAt,
		proofURL,
		proofFile,
	)
	if err != nil {
		if errors.Is(err, ErrProofFileInvalid) {
			base.BadRequest(ctx, "invalid-proof-file", nil)
			return
		}
		base.InternalServerError(ctx, "payment-submit-failed", nil)
		return
	}

	base.Success(ctx, gin.H{"payment_id": item.ID.String()}, "success")
}

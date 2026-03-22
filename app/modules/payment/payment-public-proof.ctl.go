package payment

import (
	"errors"

	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) PublicUploadProof(ctx *gin.Context) {
	bookingNo := ctx.PostForm("booking_no")
	phone := ctx.PostForm("phone")
	paymentID := ctx.PostForm("payment_id")

	if bookingNo == "" || phone == "" || paymentID == "" {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	booking, err := c.svc.booking.GetBookingByBookingNoAndPhone(ctx.Request.Context(), bookingNo, phone)
	if err != nil || booking == nil {
		base.BadRequest(ctx, "booking-not-found", nil)
		return
	}

	payment, err := c.svc.db.GetPaymentByID(ctx.Request.Context(), paymentID)
	if err != nil || payment == nil {
		base.BadRequest(ctx, "payment-not-found", nil)
		return
	}

	if payment.BookingID.String() != booking.ID.String() {
		base.BadRequest(ctx, "payment-booking-mismatch", nil)
		return
	}

	fileHeader, err := ctx.FormFile("proof")
	if err != nil {
		base.BadRequest(ctx, "proof-file-required", nil)
		return
	}

	proofURL, err := c.svc.UploadProof(ctx.Request.Context(), paymentID, fileHeader)
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

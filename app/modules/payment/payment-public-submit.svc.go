package payment

import (
	"context"
	"mime/multipart"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) SubmitPublicPayment(
	ctx context.Context,
	bookingNo string,
	phone string,
	channel *string,
	amount float64,
	depositAmount float64,
	note *string,
	paidAt *time.Time,
	proofURL *string,
	proofFile *multipart.FileHeader,
) (*ent.PaymentEntity, error) {
	booking, err := s.booking.GetBookingByBookingNoAndPhone(ctx, bookingNo, phone)
	if err != nil {
		return nil, err
	}

	status := "pending"
	item, err := s.db.CreatePayment(
		ctx,
		booking.ID.String(),
		channel,
		amount,
		depositAmount,
		&status,
		proofURL,
		note,
		paidAt,
	)
	if err != nil {
		return nil, err
	}

	if proofFile != nil {
		proof, err := s.UploadProof(ctx, item.ID.String(), proofFile)
		if err != nil {
			return nil, err
		}
		item.ProofURL = &proof
	}

	return item, nil
}

package payment

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, bookingID string, channel *string, amount float64, depositAmount float64, status *string, proofURL *string, note *string, paidAt *time.Time) (*ent.PaymentEntity, error) {
	return s.db.CreatePayment(ctx, bookingID, channel, amount, depositAmount, status, proofURL, note, paidAt)
}

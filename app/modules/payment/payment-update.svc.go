package payment

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, bookingID *string, channel *string, amount *float64, depositAmount *float64, status *string, proofURL *string, note *string, paidAt *time.Time) (*ent.PaymentEntity, error) {
	return s.db.UpdatePaymentByID(ctx, id, bookingID, channel, amount, depositAmount, status, proofURL, note, paidAt)
}

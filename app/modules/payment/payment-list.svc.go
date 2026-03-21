package payment

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, bookingID *string, channel *string, status *string) ([]*ent.PaymentEntity, error) {
	return s.db.ListPayments(ctx, bookingID, channel, status)
}

package payment

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) InfoByID(ctx context.Context, id string) (*ent.PaymentEntity, error) {
	return s.db.GetPaymentByID(ctx, id)
}

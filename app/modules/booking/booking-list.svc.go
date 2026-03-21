package booking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, status *string, payment *string) ([]*ent.BookingEntity, error) {
	return s.db.ListBookings(ctx, status, payment)
}

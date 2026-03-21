package booking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.BookingEntity, error) {
	return s.db.GetBookingByID(ctx, id)
}

package bookingitem

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) InfoByID(ctx context.Context, id string) (*ent.BookingItemEntity, error) {
	return s.db.GetBookingItemByID(ctx, id)
}

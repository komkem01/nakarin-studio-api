package bookingitem

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, bookingID *string, productID *string) ([]*ent.BookingItemEntity, error) {
	return s.db.ListBookingItems(ctx, bookingID, productID)
}

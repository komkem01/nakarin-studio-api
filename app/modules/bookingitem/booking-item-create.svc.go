package bookingitem

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, bookingID string, productID string, productName string, unitPriceAtBooking float64, quantity int, lineTotal float64, note *string, sortOrder int) (*ent.BookingItemEntity, error) {
	return s.db.CreateBookingItem(ctx, bookingID, productID, productName, unitPriceAtBooking, quantity, lineTotal, note, sortOrder)
}

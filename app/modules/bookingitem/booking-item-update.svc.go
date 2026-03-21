package bookingitem

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, bookingID *string, productID *string, productName *string, unitPriceAtBooking *float64, quantity *int, lineTotal *float64, note *string, sortOrder *int) (*ent.BookingItemEntity, error) {
	return s.db.UpdateBookingItemByID(ctx, id, bookingID, productID, productName, unitPriceAtBooking, quantity, lineTotal, note, sortOrder)
}

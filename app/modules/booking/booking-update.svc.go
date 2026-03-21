package booking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, bookingNo *string, status *string, payment *string) (*ent.BookingEntity, error) {
	return s.db.UpdateBookingByID(ctx, id, bookingNo, status, payment)
}

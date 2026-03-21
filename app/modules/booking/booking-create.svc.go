package booking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, bookingNo string, status *string, payment *string) (*ent.BookingEntity, error) {
	return s.db.CreateBooking(ctx, bookingNo, status, payment)
}

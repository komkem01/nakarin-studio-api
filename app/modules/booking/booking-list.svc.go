package booking

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, status *string, payment *string, bookingNo *string, phone *string, createdAtFrom *time.Time, createdAtTo *time.Time) ([]*ent.BookingEntity, error) {
	return s.db.ListBookings(ctx, status, payment, bookingNo, phone, createdAtFrom, createdAtTo)
}

package booking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) TrackByBookingNoAndPhone(ctx context.Context, bookingNo string, phone string) (*ent.BookingEntity, error) {
	return s.db.GetBookingByBookingNoAndPhone(ctx, bookingNo, phone)
}

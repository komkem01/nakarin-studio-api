package bookingdetail

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, bookingID *string) ([]*ent.BookingDetailEntity, error) {
	return s.db.ListBookingDetails(ctx, bookingID)
}

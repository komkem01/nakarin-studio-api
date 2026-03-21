package bookingdetail

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.BookingDetailEntity, error) {
	return s.db.GetBookingDetailByID(ctx, id)
}

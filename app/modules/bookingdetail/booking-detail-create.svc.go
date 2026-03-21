package bookingdetail

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, bookingID string, firstName string, lastName *string, phone string) (*ent.BookingDetailEntity, error) {
	return s.db.CreateBookingDetail(ctx, bookingID, firstName, lastName, phone)
}

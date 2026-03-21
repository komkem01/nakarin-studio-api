package bookingdetail

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, bookingID *string, firstName *string, lastName *string, phone *string) (*ent.BookingDetailEntity, error) {
	return s.db.UpdateBookingDetailByID(ctx, id, bookingID, firstName, lastName, phone)
}

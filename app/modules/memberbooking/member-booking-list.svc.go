package memberbooking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, memberID *string, bookingID *string) ([]*ent.MemberBookingEntity, error) {
	return s.db.ListMemberBookings(ctx, memberID, bookingID)
}

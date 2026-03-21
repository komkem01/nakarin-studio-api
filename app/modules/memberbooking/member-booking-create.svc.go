package memberbooking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, memberID string, bookingID string) (*ent.MemberBookingEntity, error) {
	return s.db.CreateMemberBooking(ctx, memberID, bookingID)
}

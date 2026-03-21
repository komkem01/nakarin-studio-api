package memberbooking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, memberID *string, bookingID *string) (*ent.MemberBookingEntity, error) {
	return s.db.UpdateMemberBookingByID(ctx, id, memberID, bookingID)
}

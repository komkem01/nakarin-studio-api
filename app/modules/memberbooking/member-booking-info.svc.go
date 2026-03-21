package memberbooking

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.MemberBookingEntity, error) {
	return s.db.GetMemberBookingByID(ctx, id)
}

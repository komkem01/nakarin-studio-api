package member

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, genderID *string, prefixID *string, role *string, phone *string) ([]*ent.MemberEntity, error) {
	return s.db.ListMembers(ctx, genderID, prefixID, role, phone)
}

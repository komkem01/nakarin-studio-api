package member

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, genderID *string, prefixID *string, role *string, firstName *string, lastName *string, phone *string) (*ent.MemberEntity, error) {
	return s.db.UpdateMemberByID(ctx, id, genderID, prefixID, role, firstName, lastName, phone)
}

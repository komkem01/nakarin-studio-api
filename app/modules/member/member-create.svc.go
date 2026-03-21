package member

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, genderID string, prefixID *string, role *string, firstName string, lastName *string, phone string) (*ent.MemberEntity, error) {
	return s.db.CreateMember(ctx, genderID, prefixID, role, firstName, lastName, phone)
}

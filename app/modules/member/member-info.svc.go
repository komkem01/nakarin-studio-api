package member

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.MemberEntity, error) {
	return s.db.GetMemberByID(ctx, id)
}

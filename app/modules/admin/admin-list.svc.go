package admin

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, memberID *string, username *string, isActive *bool) ([]*ent.AdminEntity, error) {
	return s.db.ListAdmins(ctx, memberID, username, isActive)
}

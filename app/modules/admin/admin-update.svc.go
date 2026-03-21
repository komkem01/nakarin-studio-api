package admin

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, memberID *string, username *string, passwordHash *string, displayName *string, lastLoginAt *time.Time, isActive *bool) (*ent.AdminEntity, error) {
	return s.db.UpdateAdminByID(ctx, id, memberID, username, passwordHash, displayName, lastLoginAt, isActive)
}

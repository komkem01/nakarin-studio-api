package admin

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, memberID *string, username string, passwordHash string, displayName *string, lastLoginAt *time.Time, isActive bool) (*ent.AdminEntity, error) {
	return s.db.CreateAdmin(ctx, memberID, username, passwordHash, displayName, lastLoginAt, isActive)
}

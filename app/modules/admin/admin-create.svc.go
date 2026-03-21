package admin

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, memberID *string, username string, passwordHash string, displayName *string, lastLoginAt *time.Time, isActive bool) (*ent.AdminEntity, error) {
	hashed, err := hashPassword(passwordHash)
	if err != nil {
		return nil, err
	}

	return s.db.CreateAdmin(ctx, memberID, username, hashed, displayName, lastLoginAt, isActive)
}

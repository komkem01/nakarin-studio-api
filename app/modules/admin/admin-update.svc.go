package admin

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, memberID *string, username *string, passwordHash *string, displayName *string, lastLoginAt *time.Time, isActive *bool) (*ent.AdminEntity, error) {
	if passwordHash != nil {
		hashed, err := hashPassword(*passwordHash)
		if err != nil {
			return nil, err
		}
		passwordHash = &hashed
	}

	return s.db.UpdateAdminByID(ctx, id, memberID, username, passwordHash, displayName, lastLoginAt, isActive)
}

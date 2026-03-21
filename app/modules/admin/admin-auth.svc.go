package admin

import (
	"context"
	"fmt"
	"strings"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Authenticate(ctx context.Context, username string, password string) (*ent.AdminEntity, error) {
	usernameValue := strings.TrimSpace(username)
	passwordValue := strings.TrimSpace(password)
	if usernameValue == "" || passwordValue == "" {
		return nil, fmt.Errorf("username and password are required")
	}

	admin, err := s.db.GetAdminByUsername(ctx, usernameValue)
	if err != nil {
		return nil, err
	}
	if !admin.IsActive {
		return nil, fmt.Errorf("admin is inactive")
	}
	if admin.PasswordHash != passwordValue {
		return nil, fmt.Errorf("invalid credentials")
	}

	now := time.Now()
	_, _ = s.db.UpdateAdminByID(ctx, admin.ID.String(), nil, nil, nil, nil, &now, nil)
	return admin, nil
}

package admin

import (
	"context"
	"fmt"
	"strings"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

type AuthResult struct {
	Admin        *ent.AdminEntity
	AccessToken  string
	RefreshToken string
}

func (s *Service) Login(ctx context.Context, username string, password string) (*AuthResult, error) {
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
	if !verifyPassword(admin.PasswordHash, passwordValue) {
		return nil, fmt.Errorf("invalid credentials")
	}

	accessToken, err := s.generateAccessToken(admin)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(admin)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	_, _ = s.db.UpdateAdminByID(ctx, admin.ID.String(), nil, nil, nil, nil, &now, nil)

	return &AuthResult{Admin: admin, AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) (*AuthResult, error) {
	claims, err := s.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	admin, err := s.db.GetAdminByID(ctx, claims.AdminID)
	if err != nil {
		return nil, err
	}
	if admin == nil || !admin.IsActive {
		return nil, fmt.Errorf("admin is inactive")
	}

	accessToken, err := s.generateAccessToken(admin)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.generateRefreshToken(admin)
	if err != nil {
		return nil, err
	}

	return &AuthResult{Admin: admin, AccessToken: accessToken, RefreshToken: newRefreshToken}, nil
}

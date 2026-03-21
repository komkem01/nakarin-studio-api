package entities

import (
	"context"
	"fmt"
	"strings"
	"time"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateAdmin(ctx context.Context, memberID *string, username string, passwordHash string, displayName *string, lastLoginAt *time.Time, isActive bool) (*ent.AdminEntity, error) {
	usernameValue := strings.TrimSpace(username)
	if usernameValue == "" {
		return nil, fmt.Errorf("username is required")
	}

	passwordHashValue := strings.TrimSpace(passwordHash)
	if passwordHashValue == "" {
		return nil, fmt.Errorf("password_hash is required")
	}

	memberUUID, err := parseOptionalUUID(memberID)
	if err != nil {
		return nil, err
	}

	model := &ent.AdminEntity{
		ID:           uuid.New(),
		MemberID:     memberUUID,
		Username:     usernameValue,
		PasswordHash: passwordHashValue,
		DisplayName:  normalizeOptionalString(displayName),
		LastLoginAt:  lastLoginAt,
		IsActive:     isActive,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetAdminByID(ctx context.Context, id string) (*ent.AdminEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.AdminEntity{}
	if err := s.db.NewSelect().Model(model).Where("admin.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetAdminByUsername(ctx context.Context, username string) (*ent.AdminEntity, error) {
	usernameValue := strings.TrimSpace(username)
	if usernameValue == "" {
		return nil, fmt.Errorf("username is required")
	}

	model := &ent.AdminEntity{}
	if err := s.db.NewSelect().Model(model).Where("admin.username = ?", usernameValue).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListAdmins(ctx context.Context, memberID *string, username *string, isActive *bool) ([]*ent.AdminEntity, error) {
	items := make([]*ent.AdminEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("admin.created_at DESC")

	if memberID != nil && strings.TrimSpace(*memberID) != "" {
		memberUUID, err := uuid.Parse(strings.TrimSpace(*memberID))
		if err != nil {
			return nil, err
		}
		q = q.Where("admin.member_id = ?", memberUUID)
	}

	if username != nil && strings.TrimSpace(*username) != "" {
		q = q.Where("admin.username ILIKE ?", "%"+strings.TrimSpace(*username)+"%")
	}

	if isActive != nil {
		q = q.Where("admin.is_active = ?", *isActive)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateAdminByID(ctx context.Context, id string, memberID *string, username *string, passwordHash *string, displayName *string, lastLoginAt *time.Time, isActive *bool) (*ent.AdminEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.AdminEntity{}
	if err := s.db.NewSelect().Model(model).Where("admin.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if memberID != nil {
		memberUUID, err := parseOptionalUUID(memberID)
		if err != nil {
			return nil, err
		}
		model.MemberID = memberUUID
	}
	if username != nil {
		usernameValue := strings.TrimSpace(*username)
		if usernameValue == "" {
			return nil, fmt.Errorf("username is required")
		}
		model.Username = usernameValue
	}
	if passwordHash != nil {
		passwordHashValue := strings.TrimSpace(*passwordHash)
		if passwordHashValue == "" {
			return nil, fmt.Errorf("password_hash is required")
		}
		model.PasswordHash = passwordHashValue
	}
	if displayName != nil {
		model.DisplayName = normalizeOptionalString(displayName)
	}
	if lastLoginAt != nil {
		model.LastLoginAt = lastLoginAt
	}
	if isActive != nil {
		model.IsActive = *isActive
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("member_id", "username", "password_hash", "display_name", "last_login_at", "is_active", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteAdminByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.AdminEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

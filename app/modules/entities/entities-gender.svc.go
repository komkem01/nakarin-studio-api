package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"
	entitiesinf "nakarin-studio/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.GenderEntity = (*Service)(nil)

func (s *Service) CreateGender(ctx context.Context, name string, isActive bool) (*ent.GenderEntity, error) {
	value := strings.TrimSpace(name)
	if value == "" {
		return nil, fmt.Errorf("name is required")
	}

	model := &ent.GenderEntity{
		ID:       uuid.New(),
		Name:     value,
		IsActive: isActive,
	}

	_, err := s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetGenderByID(ctx context.Context, id string) (*ent.GenderEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.GenderEntity{}
	if err := s.db.NewSelect().
		Model(model).
		Where("gender.id = ?", uid).
		Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListGenders(ctx context.Context, isActive *bool) ([]*ent.GenderEntity, error) {
	items := make([]*ent.GenderEntity, 0)

	q := s.db.NewSelect().
		Model(&items).
		Order("gender.created_at DESC")

	if isActive != nil {
		q = q.Where("gender.is_active = ?", *isActive)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateGenderByID(ctx context.Context, id string, name *string, isActive *bool) (*ent.GenderEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.GenderEntity{}
	if err := s.db.NewSelect().
		Model(model).
		Where("gender.id = ?", uid).
		Scan(ctx); err != nil {
		return nil, err
	}

	if name != nil {
		value := strings.TrimSpace(*name)
		if value == "" {
			return nil, fmt.Errorf("name is required")
		}
		model.Name = value
	}

	if isActive != nil {
		model.IsActive = *isActive
	}

	_, err = s.db.NewUpdate().
		Model(model).
		WherePK().
		Column("name", "is_active", "updated_at").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteGenderByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().
		Model(&ent.GenderEntity{}).
		Where("id = ?", uid).
		Exec(ctx)
	return err
}

package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreatePrefix(ctx context.Context, genderID string, name string, isActive bool) (*ent.PrefixEntity, error) {
	gid, err := uuid.Parse(genderID)
	if err != nil {
		return nil, err
	}

	value := strings.TrimSpace(name)
	if value == "" {
		return nil, fmt.Errorf("name is required")
	}

	model := &ent.PrefixEntity{
		ID:       uuid.New(),
		GenderID: gid,
		Name:     value,
		IsActive: isActive,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetPrefixByID(ctx context.Context, id string) (*ent.PrefixEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.PrefixEntity{}
	if err := s.db.NewSelect().
		Model(model).
		Where("prefix.id = ?", uid).
		Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListPrefixes(ctx context.Context, genderID *string, isActive *bool) ([]*ent.PrefixEntity, error) {
	items := make([]*ent.PrefixEntity, 0)

	q := s.db.NewSelect().
		Model(&items).
		Order("prefix.created_at DESC")

	if genderID != nil && strings.TrimSpace(*genderID) != "" {
		gid, err := uuid.Parse(strings.TrimSpace(*genderID))
		if err != nil {
			return nil, err
		}
		q = q.Where("prefix.gender_id = ?", gid)
	}

	if isActive != nil {
		q = q.Where("prefix.is_active = ?", *isActive)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdatePrefixByID(ctx context.Context, id string, genderID *string, name *string, isActive *bool) (*ent.PrefixEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.PrefixEntity{}
	if err := s.db.NewSelect().
		Model(model).
		Where("prefix.id = ?", uid).
		Scan(ctx); err != nil {
		return nil, err
	}

	if genderID != nil {
		gid, err := uuid.Parse(strings.TrimSpace(*genderID))
		if err != nil {
			return nil, err
		}
		model.GenderID = gid
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
		Column("gender_id", "name", "is_active", "updated_at").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeletePrefixByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().
		Model(&ent.PrefixEntity{}).
		Where("id = ?", uid).
		Exec(ctx)
	return err
}

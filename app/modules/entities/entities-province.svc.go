package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateProvince(ctx context.Context, name string, isActive bool) (*ent.ProvinceEntity, error) {
	value := strings.TrimSpace(name)
	if value == "" {
		return nil, fmt.Errorf("name is required")
	}

	model := &ent.ProvinceEntity{
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

func (s *Service) GetProvinceByID(ctx context.Context, id string) (*ent.ProvinceEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.ProvinceEntity{}
	if err := s.db.NewSelect().Model(model).Where("province.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) ListProvinces(ctx context.Context, isActive *bool) ([]*ent.ProvinceEntity, error) {
	items := make([]*ent.ProvinceEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("province.created_at DESC")
	if isActive != nil {
		q = q.Where("province.is_active = ?", *isActive)
	}
	if err := q.Scan(ctx); err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpdateProvinceByID(ctx context.Context, id string, name *string, isActive *bool) (*ent.ProvinceEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := &ent.ProvinceEntity{}
	if err := s.db.NewSelect().Model(model).Where("province.id = ?", uid).Scan(ctx); err != nil {
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

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("name", "is_active", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) DeleteProvinceByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = s.db.NewDelete().Model(&ent.ProvinceEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

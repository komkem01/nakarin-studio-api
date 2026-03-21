package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateDistrict(ctx context.Context, provinceID string, name string, isActive bool) (*ent.DistrictEntity, error) {
	pid, err := uuid.Parse(provinceID)
	if err != nil {
		return nil, err
	}
	value := strings.TrimSpace(name)
	if value == "" {
		return nil, fmt.Errorf("name is required")
	}

	model := &ent.DistrictEntity{ID: uuid.New(), ProvinceID: pid, Name: value, IsActive: isActive}
	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) GetDistrictByID(ctx context.Context, id string) (*ent.DistrictEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := &ent.DistrictEntity{}
	if err := s.db.NewSelect().Model(model).Where("district.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) ListDistricts(ctx context.Context, provinceID *string, isActive *bool) ([]*ent.DistrictEntity, error) {
	items := make([]*ent.DistrictEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("district.created_at DESC")
	if provinceID != nil && strings.TrimSpace(*provinceID) != "" {
		pid, err := uuid.Parse(strings.TrimSpace(*provinceID))
		if err != nil {
			return nil, err
		}
		q = q.Where("district.province_id = ?", pid)
	}
	if isActive != nil {
		q = q.Where("district.is_active = ?", *isActive)
	}
	if err := q.Scan(ctx); err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpdateDistrictByID(ctx context.Context, id string, provinceID *string, name *string, isActive *bool) (*ent.DistrictEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := &ent.DistrictEntity{}
	if err := s.db.NewSelect().Model(model).Where("district.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}
	if provinceID != nil {
		pid, err := uuid.Parse(strings.TrimSpace(*provinceID))
		if err != nil {
			return nil, err
		}
		model.ProvinceID = pid
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
	_, err = s.db.NewUpdate().Model(model).WherePK().Column("province_id", "name", "is_active", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) DeleteDistrictByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = s.db.NewDelete().Model(&ent.DistrictEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

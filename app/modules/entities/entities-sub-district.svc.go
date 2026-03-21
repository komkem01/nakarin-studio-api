package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateSubDistrict(ctx context.Context, districtID string, name string, isActive bool) (*ent.SubDistrictEntity, error) {
	did, err := uuid.Parse(districtID)
	if err != nil {
		return nil, err
	}
	value := strings.TrimSpace(name)
	if value == "" {
		return nil, fmt.Errorf("name is required")
	}
	model := &ent.SubDistrictEntity{ID: uuid.New(), DistrictID: did, Name: value, IsActive: isActive}
	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) GetSubDistrictByID(ctx context.Context, id string) (*ent.SubDistrictEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := &ent.SubDistrictEntity{}
	if err := s.db.NewSelect().Model(model).Where("sub_district.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) ListSubDistricts(ctx context.Context, districtID *string, isActive *bool) ([]*ent.SubDistrictEntity, error) {
	items := make([]*ent.SubDistrictEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("sub_district.created_at DESC")
	if districtID != nil && strings.TrimSpace(*districtID) != "" {
		did, err := uuid.Parse(strings.TrimSpace(*districtID))
		if err != nil {
			return nil, err
		}
		q = q.Where("sub_district.district_id = ?", did)
	}
	if isActive != nil {
		q = q.Where("sub_district.is_active = ?", *isActive)
	}
	if err := q.Scan(ctx); err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpdateSubDistrictByID(ctx context.Context, id string, districtID *string, name *string, isActive *bool) (*ent.SubDistrictEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := &ent.SubDistrictEntity{}
	if err := s.db.NewSelect().Model(model).Where("sub_district.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}
	if districtID != nil {
		did, err := uuid.Parse(strings.TrimSpace(*districtID))
		if err != nil {
			return nil, err
		}
		model.DistrictID = did
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
	_, err = s.db.NewUpdate().Model(model).WherePK().Column("district_id", "name", "is_active", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) DeleteSubDistrictByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = s.db.NewDelete().Model(&ent.SubDistrictEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

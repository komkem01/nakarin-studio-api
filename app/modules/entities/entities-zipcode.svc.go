package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateZipcode(ctx context.Context, subDistrictID string, name string, isActive bool) (*ent.ZipcodeEntity, error) {
	sid, err := uuid.Parse(subDistrictID)
	if err != nil {
		return nil, err
	}
	value := strings.TrimSpace(name)
	if value == "" {
		return nil, fmt.Errorf("name is required")
	}
	model := &ent.ZipcodeEntity{ID: uuid.New(), SubDistrictID: sid, Name: value, IsActive: isActive}
	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) GetZipcodeByID(ctx context.Context, id string) (*ent.ZipcodeEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := &ent.ZipcodeEntity{}
	if err := s.db.NewSelect().Model(model).Where("zipcode.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) ListZipcodes(ctx context.Context, subDistrictID *string, isActive *bool) ([]*ent.ZipcodeEntity, error) {
	items := make([]*ent.ZipcodeEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("zipcode.created_at DESC")
	if subDistrictID != nil && strings.TrimSpace(*subDistrictID) != "" {
		sid, err := uuid.Parse(strings.TrimSpace(*subDistrictID))
		if err != nil {
			return nil, err
		}
		q = q.Where("zipcode.sub_district_id = ?", sid)
	}
	if isActive != nil {
		q = q.Where("zipcode.is_active = ?", *isActive)
	}
	if err := q.Scan(ctx); err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpdateZipcodeByID(ctx context.Context, id string, subDistrictID *string, name *string, isActive *bool) (*ent.ZipcodeEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := &ent.ZipcodeEntity{}
	if err := s.db.NewSelect().Model(model).Where("zipcode.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}
	if subDistrictID != nil {
		sid, err := uuid.Parse(strings.TrimSpace(*subDistrictID))
		if err != nil {
			return nil, err
		}
		model.SubDistrictID = sid
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
	_, err = s.db.NewUpdate().Model(model).WherePK().Column("sub_district_id", "name", "is_active", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s *Service) DeleteZipcodeByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = s.db.NewDelete().Model(&ent.ZipcodeEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

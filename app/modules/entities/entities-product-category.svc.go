package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"
	entitiesinf "nakarin-studio/app/modules/entities/inf"

	"github.com/google/uuid"
)

var _ entitiesinf.ProductCategoryEntity = (*Service)(nil)

func (s *Service) CreateProductCategory(ctx context.Context, name string, description *string, isActive bool) (*ent.ProductCategoryEntity, error) {
	nameValue := strings.TrimSpace(name)
	if nameValue == "" {
		return nil, fmt.Errorf("name is required")
	}

	model := &ent.ProductCategoryEntity{
		ID:          uuid.New(),
		Name:        nameValue,
		Description: normalizeOptionalString(description),
		IsActive:    isActive,
	}

	_, err := s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetProductCategoryByID(ctx context.Context, id string) (*ent.ProductCategoryEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.ProductCategoryEntity{}
	if err := s.db.NewSelect().Model(model).Where("product_category.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListProductCategories(ctx context.Context, isActive *bool) ([]*ent.ProductCategoryEntity, error) {
	items := make([]*ent.ProductCategoryEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("product_category.created_at DESC")

	if isActive != nil {
		q = q.Where("product_category.is_active = ?", *isActive)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateProductCategoryByID(ctx context.Context, id string, name *string, description *string, isActive *bool) (*ent.ProductCategoryEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.ProductCategoryEntity{}
	if err := s.db.NewSelect().Model(model).Where("product_category.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if name != nil {
		nameValue := strings.TrimSpace(*name)
		if nameValue == "" {
			return nil, fmt.Errorf("name is required")
		}
		model.Name = nameValue
	}

	if description != nil {
		model.Description = normalizeOptionalString(description)
	}

	if isActive != nil {
		model.IsActive = *isActive
	}

	_, err = s.db.NewUpdate().
		Model(model).
		WherePK().
		Column("name", "description", "is_active", "updated_at").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteProductCategoryByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	model := &ent.ProductCategoryEntity{}
	if err := s.db.NewSelect().Model(model).Where("product_category.id = ?", uid).Scan(ctx); err != nil {
		return err
	}

	model.IsActive = false
	_, err = s.db.NewUpdate().
		Model(model).
		WherePK().
		Column("is_active", "updated_at").
		Exec(ctx)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(model).WherePK().Exec(ctx)
	return err
}

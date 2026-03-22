package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateProduct(ctx context.Context, name string, categoryID string, description *string, suitableFor *string, onSite *string, receivedItems *string, note *string, price float64, isActive bool, isAvailable bool, prepTime int, sortOrder int) (*ent.ProductEntity, error) {
	nameValue := strings.TrimSpace(name)
	if nameValue == "" {
		return nil, fmt.Errorf("name is required")
	}

	categoryIDValue := strings.TrimSpace(categoryID)
	if categoryIDValue == "" {
		return nil, fmt.Errorf("category_id is required")
	}

	categoryUUID, err := uuid.Parse(categoryIDValue)
	if err != nil {
		return nil, fmt.Errorf("invalid category_id")
	}

	if price < 0 {
		return nil, fmt.Errorf("price must be greater than or equal to 0")
	}

	if prepTime < 0 {
		return nil, fmt.Errorf("prep_time must be greater than or equal to 0")
	}

	model := &ent.ProductEntity{
		ID:            uuid.New(),
		CategoryID:    &categoryUUID,
		Name:          nameValue,
		Description:   normalizeOptionalString(description),
		SuitableFor:   normalizeOptionalString(suitableFor),
		OnSite:        normalizeOptionalString(onSite),
		ReceivedItems: normalizeOptionalString(receivedItems),
		Note:          normalizeOptionalString(note),
		Price:         price,
		IsActive:      isActive,
		IsAvailable:   isAvailable,
		PrepTime:      prepTime,
		SortOrder:     sortOrder,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetProductByID(ctx context.Context, id string) (*ent.ProductEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.ProductEntity{}
	if err := s.db.NewSelect().Model(model).Where("product.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListProducts(ctx context.Context, name *string, categoryID *string, isActive *bool, isAvailable *bool) ([]*ent.ProductEntity, error) {
	items := make([]*ent.ProductEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("product.sort_order ASC").Order("product.created_at DESC")

	if name != nil && strings.TrimSpace(*name) != "" {
		q = q.Where("product.name ILIKE ?", "%"+strings.TrimSpace(*name)+"%")
	}
	if categoryID != nil && strings.TrimSpace(*categoryID) != "" {
		categoryUUID, err := uuid.Parse(strings.TrimSpace(*categoryID))
		if err != nil {
			return nil, fmt.Errorf("invalid category_id")
		}
		q = q.Where("product.category_id = ?", categoryUUID)
	}
	if isActive != nil {
		q = q.Where("product.is_active = ?", *isActive)
	}
	if isAvailable != nil {
		q = q.Where("product.is_available = ?", *isAvailable)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateProductByID(ctx context.Context, id string, categoryID *string, name *string, description *string, suitableFor *string, onSite *string, receivedItems *string, note *string, price *float64, isActive *bool, isAvailable *bool, prepTime *int, sortOrder *int) (*ent.ProductEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.ProductEntity{}
	if err := s.db.NewSelect().Model(model).Where("product.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if name != nil {
		nameValue := strings.TrimSpace(*name)
		if nameValue == "" {
			return nil, fmt.Errorf("name is required")
		}
		model.Name = nameValue
	}
	if categoryID != nil {
		categoryIDValue := strings.TrimSpace(*categoryID)
		if categoryIDValue == "" {
			model.CategoryID = nil
		} else {
			categoryUUID, err := uuid.Parse(categoryIDValue)
			if err != nil {
				return nil, fmt.Errorf("invalid category_id")
			}
			model.CategoryID = &categoryUUID
		}
	}
	if description != nil {
		model.Description = normalizeOptionalString(description)
	}
	if suitableFor != nil {
		model.SuitableFor = normalizeOptionalString(suitableFor)
	}
	if onSite != nil {
		model.OnSite = normalizeOptionalString(onSite)
	}
	if receivedItems != nil {
		model.ReceivedItems = normalizeOptionalString(receivedItems)
	}
	if note != nil {
		model.Note = normalizeOptionalString(note)
	}
	if price != nil {
		if *price < 0 {
			return nil, fmt.Errorf("price must be greater than or equal to 0")
		}
		model.Price = *price
	}
	if isActive != nil {
		model.IsActive = *isActive
	}
	if isAvailable != nil {
		model.IsAvailable = *isAvailable
	}
	if prepTime != nil {
		if *prepTime < 0 {
			return nil, fmt.Errorf("prep_time must be greater than or equal to 0")
		}
		model.PrepTime = *prepTime
	}
	if sortOrder != nil {
		model.SortOrder = *sortOrder
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("category_id", "name", "description", "suitable_for", "on_site", "received_items", "note", "price", "is_active", "is_available", "prep_time", "sort_order", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteProductByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.ProductEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

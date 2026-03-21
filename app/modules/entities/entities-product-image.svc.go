package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateProductImage(ctx context.Context, productID string, imageURL string, altText *string, sortOrder int, isActive bool) (*ent.ProductImageEntity, error) {
	productUUID, err := uuid.Parse(strings.TrimSpace(productID))
	if err != nil {
		return nil, err
	}

	url := strings.TrimSpace(imageURL)
	if url == "" {
		return nil, fmt.Errorf("image_url is required")
	}

	model := &ent.ProductImageEntity{
		ID:        uuid.New(),
		ProductID: productUUID,
		ImageURL:  url,
		AltText:   normalizeOptionalString(altText),
		SortOrder: sortOrder,
		IsActive:  isActive,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetProductImageByID(ctx context.Context, id string) (*ent.ProductImageEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.ProductImageEntity{}
	if err := s.db.NewSelect().Model(model).Where("product_image.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListProductImages(ctx context.Context, productID *string, isActive *bool) ([]*ent.ProductImageEntity, error) {
	items := make([]*ent.ProductImageEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("product_image.sort_order ASC").Order("product_image.created_at DESC")

	if productID != nil && strings.TrimSpace(*productID) != "" {
		productUUID, err := uuid.Parse(strings.TrimSpace(*productID))
		if err != nil {
			return nil, err
		}
		q = q.Where("product_image.product_id = ?", productUUID)
	}
	if isActive != nil {
		q = q.Where("product_image.is_active = ?", *isActive)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateProductImageByID(ctx context.Context, id string, productID *string, imageURL *string, altText *string, sortOrder *int, isActive *bool) (*ent.ProductImageEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.ProductImageEntity{}
	if err := s.db.NewSelect().Model(model).Where("product_image.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if productID != nil {
		productUUID, err := uuid.Parse(strings.TrimSpace(*productID))
		if err != nil {
			return nil, err
		}
		model.ProductID = productUUID
	}
	if imageURL != nil {
		url := strings.TrimSpace(*imageURL)
		if url == "" {
			return nil, fmt.Errorf("image_url is required")
		}
		model.ImageURL = url
	}
	if altText != nil {
		model.AltText = normalizeOptionalString(altText)
	}
	if sortOrder != nil {
		model.SortOrder = *sortOrder
	}
	if isActive != nil {
		model.IsActive = *isActive
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("product_id", "image_url", "alt_text", "sort_order", "is_active", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteProductImageByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.ProductImageEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

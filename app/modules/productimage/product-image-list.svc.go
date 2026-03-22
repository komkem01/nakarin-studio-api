package productimage

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, productID *string, isActive *bool) ([]*ent.ProductImageEntity, error) {
	items, err := s.db.ListProductImages(ctx, productID, isActive)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item == nil {
			continue
		}
		item.ImageURL = s.storage.displayImageURL(ctx, item.ImageURL)
	}

	return items, nil
}

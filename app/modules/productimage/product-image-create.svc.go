package productimage

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, productID string, imageURL string, altText *string, sortOrder int, isActive bool) (*ent.ProductImageEntity, error) {
	return s.db.CreateProductImage(ctx, productID, imageURL, altText, sortOrder, isActive)
}

package productimage

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, productID *string, imageURL *string, altText *string, sortOrder *int, isActive *bool) (*ent.ProductImageEntity, error) {
	return s.db.UpdateProductImageByID(ctx, id, productID, imageURL, altText, sortOrder, isActive)
}

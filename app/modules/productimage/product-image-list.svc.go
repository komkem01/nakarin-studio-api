package productimage

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, productID *string, isActive *bool) ([]*ent.ProductImageEntity, error) {
	return s.db.ListProductImages(ctx, productID, isActive)
}

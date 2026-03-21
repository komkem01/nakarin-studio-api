package productimage

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.ProductImageEntity, error) {
	return s.db.GetProductImageByID(ctx, id)
}

package product

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.ProductEntity, error) {
	return s.db.GetProductByID(ctx, id)
}

package product

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, name *string, categoryID *string, isActive *bool, isAvailable *bool) ([]*ent.ProductEntity, error) {
	return s.db.ListProducts(ctx, name, categoryID, isActive, isAvailable)
}

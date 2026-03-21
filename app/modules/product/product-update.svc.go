package product

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, name *string, description *string, price *float64, isActive *bool, isAvailable *bool, sortOrder *int) (*ent.ProductEntity, error) {
	return s.db.UpdateProductByID(ctx, id, name, description, price, isActive, isAvailable, sortOrder)
}

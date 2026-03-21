package product

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, name string, description *string, price float64, isActive bool, isAvailable bool, prepTime int, sortOrder int) (*ent.ProductEntity, error) {
	return s.db.CreateProduct(ctx, name, description, price, isActive, isAvailable, prepTime, sortOrder)
}

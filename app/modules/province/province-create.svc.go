package province

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, name string, isActive bool) (*ent.ProvinceEntity, error) {
	return s.db.CreateProvince(ctx, name, isActive)
}

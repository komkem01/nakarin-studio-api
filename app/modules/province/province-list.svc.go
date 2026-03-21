package province

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, isActive *bool) ([]*ent.ProvinceEntity, error) {
	return s.db.ListProvinces(ctx, isActive)
}

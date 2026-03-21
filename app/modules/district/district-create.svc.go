package district

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, provinceID string, name string, isActive bool) (*ent.DistrictEntity, error) {
	return s.db.CreateDistrict(ctx, provinceID, name, isActive)
}

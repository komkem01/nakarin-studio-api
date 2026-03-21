package district

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, provinceID *string, isActive *bool) ([]*ent.DistrictEntity, error) {
	return s.db.ListDistricts(ctx, provinceID, isActive)
}

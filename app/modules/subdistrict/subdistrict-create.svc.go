package subdistrict

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, districtID string, name string, isActive bool) (*ent.SubDistrictEntity, error) {
	return s.db.CreateSubDistrict(ctx, districtID, name, isActive)
}

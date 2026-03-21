package subdistrict

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, districtID *string, isActive *bool) ([]*ent.SubDistrictEntity, error) {
	return s.db.ListSubDistricts(ctx, districtID, isActive)
}

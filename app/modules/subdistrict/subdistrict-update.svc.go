package subdistrict

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, districtID *string, name *string, isActive *bool) (*ent.SubDistrictEntity, error) {
	return s.db.UpdateSubDistrictByID(ctx, id, districtID, name, isActive)
}

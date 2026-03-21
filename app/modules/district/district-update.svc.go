package district

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, provinceID *string, name *string, isActive *bool) (*ent.DistrictEntity, error) {
	return s.db.UpdateDistrictByID(ctx, id, provinceID, name, isActive)
}

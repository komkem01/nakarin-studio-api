package province

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, name *string, isActive *bool) (*ent.ProvinceEntity, error) {
	return s.db.UpdateProvinceByID(ctx, id, name, isActive)
}

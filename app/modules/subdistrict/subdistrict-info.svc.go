package subdistrict

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.SubDistrictEntity, error) {
	return s.db.GetSubDistrictByID(ctx, id)
}

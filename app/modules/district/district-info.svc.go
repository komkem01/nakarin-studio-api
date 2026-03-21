package district

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.DistrictEntity, error) {
	return s.db.GetDistrictByID(ctx, id)
}

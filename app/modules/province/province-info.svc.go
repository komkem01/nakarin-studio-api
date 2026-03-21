package province

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.ProvinceEntity, error) {
	return s.db.GetProvinceByID(ctx, id)
}

package zipcode

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, subDistrictID string, name string, isActive bool) (*ent.ZipcodeEntity, error) {
	return s.db.CreateZipcode(ctx, subDistrictID, name, isActive)
}

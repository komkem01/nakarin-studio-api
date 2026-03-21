package zipcode

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, subDistrictID *string, name *string, isActive *bool) (*ent.ZipcodeEntity, error) {
	return s.db.UpdateZipcodeByID(ctx, id, subDistrictID, name, isActive)
}

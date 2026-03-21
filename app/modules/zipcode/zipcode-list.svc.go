package zipcode

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, subDistrictID *string, isActive *bool) ([]*ent.ZipcodeEntity, error) {
	return s.db.ListZipcodes(ctx, subDistrictID, isActive)
}

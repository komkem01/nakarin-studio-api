package gender

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, name *string, isActive *bool) (*ent.GenderEntity, error) {
	return s.db.UpdateGenderByID(ctx, id, name, isActive)
}

package gender

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, isActive *bool) ([]*ent.GenderEntity, error) {
	return s.db.ListGenders(ctx, isActive)
}

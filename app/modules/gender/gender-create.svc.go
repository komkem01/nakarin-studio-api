package gender

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, name string, isActive bool) (*ent.GenderEntity, error) {
	return s.db.CreateGender(ctx, name, isActive)
}

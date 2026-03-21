package gender

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.GenderEntity, error) {
	return s.db.GetGenderByID(ctx, id)
}

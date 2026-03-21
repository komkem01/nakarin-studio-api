package zipcode

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.ZipcodeEntity, error) {
	return s.db.GetZipcodeByID(ctx, id)
}

package prefix

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.PrefixEntity, error) {
	return s.db.GetPrefixByID(ctx, id)
}

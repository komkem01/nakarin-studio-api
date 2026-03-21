package prefix

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, genderID string, name string, isActive bool) (*ent.PrefixEntity, error) {
	return s.db.CreatePrefix(ctx, genderID, name, isActive)
}

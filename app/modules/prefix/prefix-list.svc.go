package prefix

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, genderID *string, isActive *bool) ([]*ent.PrefixEntity, error) {
	return s.db.ListPrefixes(ctx, genderID, isActive)
}

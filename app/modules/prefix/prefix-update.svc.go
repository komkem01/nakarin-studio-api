package prefix

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, genderID *string, name *string, isActive *bool) (*ent.PrefixEntity, error) {
	return s.db.UpdatePrefixByID(ctx, id, genderID, name, isActive)
}

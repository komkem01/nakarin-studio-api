package admin

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) InfoByID(ctx context.Context, id string) (*ent.AdminEntity, error) {
	return s.db.GetAdminByID(ctx, id)
}

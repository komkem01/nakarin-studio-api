package memberaddress

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) GetByID(ctx context.Context, id string) (*ent.MemberAddressEntity, error) {
	return s.db.GetMemberAddressByID(ctx, id)
}

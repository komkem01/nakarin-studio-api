package memberaddress

import "context"

func (s *Service) DeleteByID(ctx context.Context, id string) error {
	return s.db.DeleteMemberAddressByID(ctx, id)
}

package gender

import "context"

func (s *Service) DeleteByID(ctx context.Context, id string) error {
	return s.db.DeleteGenderByID(ctx, id)
}

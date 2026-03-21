package province

import "context"

func (s *Service) DeleteByID(ctx context.Context, id string) error {
	return s.db.DeleteProvinceByID(ctx, id)
}

package bookingdetail

import "context"

func (s *Service) DeleteByID(ctx context.Context, id string) error {
	return s.db.DeleteBookingDetailByID(ctx, id)
}

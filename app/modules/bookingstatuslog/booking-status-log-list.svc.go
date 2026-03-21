package bookingstatuslog

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) List(ctx context.Context, bookingID *string, toStatus *string) ([]*ent.BookingStatusLogEntity, error) {
	return s.db.ListBookingStatusLogs(ctx, bookingID, toStatus)
}

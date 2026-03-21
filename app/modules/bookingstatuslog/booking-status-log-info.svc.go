package bookingstatuslog

import (
	"context"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) InfoByID(ctx context.Context, id string) (*ent.BookingStatusLogEntity, error) {
	return s.db.GetBookingStatusLogByID(ctx, id)
}

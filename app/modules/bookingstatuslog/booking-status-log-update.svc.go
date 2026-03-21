package bookingstatuslog

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) UpdateByID(ctx context.Context, id string, bookingID *string, fromStatus *string, toStatus *string, changedBy *string, changedByRole *string, reason *string, changedAt *time.Time) (*ent.BookingStatusLogEntity, error) {
	return s.db.UpdateBookingStatusLogByID(ctx, id, bookingID, fromStatus, toStatus, changedBy, changedByRole, reason, changedAt)
}

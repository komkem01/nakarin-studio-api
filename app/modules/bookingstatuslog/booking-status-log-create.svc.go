package bookingstatuslog

import (
	"context"
	"time"

	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) Create(ctx context.Context, bookingID string, fromStatus *string, toStatus string, changedBy *string, changedByRole *string, reason *string, changedAt *time.Time) (*ent.BookingStatusLogEntity, error) {
	return s.db.CreateBookingStatusLog(ctx, bookingID, fromStatus, toStatus, changedBy, changedByRole, reason, changedAt)
}

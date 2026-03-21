package booking

import (
	"context"
	"fmt"
)

func (s *Service) TransitionStatus(ctx context.Context, bookingID string, toStatus string, reason *string, changedBy *string) error {
	current, err := s.db.GetBookingByID(ctx, bookingID)
	if err != nil {
		return err
	}

	if !canTransitionStatus(string(current.Status), toStatus) {
		return fmt.Errorf("invalid booking status transition")
	}

	_, err = s.db.UpdateBookingByID(
		ctx,
		bookingID,
		nil,
		&toStatus,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	_, err = s.statusLogDB.CreateBookingStatusLog(ctx, bookingID, stringPtr(string(current.Status)), toStatus, changedBy, stringPtr("admin"), reason, nil)
	return err
}

func canTransitionStatus(from string, to string) bool {
	if from == to {
		return true
	}

	switch from {
	case "pending":
		return to == "processing" || to == "canceled"
	case "processing":
		return to == "completed" || to == "canceled"
	case "completed", "canceled":
		return false
	default:
		return false
	}
}

func stringPtr(v string) *string { return &v }

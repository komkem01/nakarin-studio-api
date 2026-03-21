package entities

import (
	"context"
	"fmt"
	"strings"
	"time"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateBookingStatusLog(ctx context.Context, bookingID string, fromStatus *string, toStatus string, changedBy *string, changedByRole *string, reason *string, changedAt *time.Time) (*ent.BookingStatusLogEntity, error) {
	bookingUUID, err := uuid.Parse(strings.TrimSpace(bookingID))
	if err != nil {
		return nil, err
	}

	var fromStatusValue *ent.BookingStatus
	if fromStatus != nil && strings.TrimSpace(*fromStatus) != "" {
		value, err := normalizeBookingStatus(fromStatus)
		if err != nil {
			return nil, err
		}
		fromStatusValue = &value
	}

	toStatusValue, err := normalizeBookingStatus(&toStatus)
	if err != nil {
		return nil, err
	}

	changedByUUID, err := parseOptionalUUID(changedBy)
	if err != nil {
		return nil, err
	}

	changedByRoleValue, err := normalizeChangedByRole(changedByRole)
	if err != nil {
		return nil, err
	}

	changedAtValue := time.Now()
	if changedAt != nil {
		changedAtValue = *changedAt
	}

	model := &ent.BookingStatusLogEntity{
		ID:            uuid.New(),
		BookingID:     bookingUUID,
		FromStatus:    fromStatusValue,
		ToStatus:      toStatusValue,
		ChangedBy:     changedByUUID,
		ChangedByRole: changedByRoleValue,
		Reason:        normalizeOptionalString(reason),
		ChangedAt:     changedAtValue,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetBookingStatusLogByID(ctx context.Context, id string) (*ent.BookingStatusLogEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingStatusLogEntity{}
	if err := s.db.NewSelect().Model(model).Where("booking_status_log.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListBookingStatusLogs(ctx context.Context, bookingID *string, toStatus *string) ([]*ent.BookingStatusLogEntity, error) {
	items := make([]*ent.BookingStatusLogEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("booking_status_log.changed_at DESC")

	if bookingID != nil && strings.TrimSpace(*bookingID) != "" {
		bookingUUID, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		q = q.Where("booking_status_log.booking_id = ?", bookingUUID)
	}

	if toStatus != nil && strings.TrimSpace(*toStatus) != "" {
		statusValue, err := normalizeBookingStatus(toStatus)
		if err != nil {
			return nil, err
		}
		q = q.Where("booking_status_log.to_status = ?", statusValue)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateBookingStatusLogByID(ctx context.Context, id string, bookingID *string, fromStatus *string, toStatus *string, changedBy *string, changedByRole *string, reason *string, changedAt *time.Time) (*ent.BookingStatusLogEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingStatusLogEntity{}
	if err := s.db.NewSelect().Model(model).Where("booking_status_log.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if bookingID != nil {
		bookingUUID, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		model.BookingID = bookingUUID
	}
	if fromStatus != nil {
		if strings.TrimSpace(*fromStatus) == "" {
			model.FromStatus = nil
		} else {
			statusValue, err := normalizeBookingStatus(fromStatus)
			if err != nil {
				return nil, err
			}
			model.FromStatus = &statusValue
		}
	}
	if toStatus != nil {
		statusValue, err := normalizeBookingStatus(toStatus)
		if err != nil {
			return nil, err
		}
		model.ToStatus = statusValue
	}
	if changedBy != nil {
		changedByUUID, err := parseOptionalUUID(changedBy)
		if err != nil {
			return nil, err
		}
		model.ChangedBy = changedByUUID
	}
	if changedByRole != nil {
		changedByRoleValue, err := normalizeChangedByRole(changedByRole)
		if err != nil {
			return nil, err
		}
		model.ChangedByRole = changedByRoleValue
	}
	if reason != nil {
		model.Reason = normalizeOptionalString(reason)
	}
	if changedAt != nil {
		model.ChangedAt = *changedAt
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("booking_id", "from_status", "to_status", "changed_by", "changed_by_role", "reason", "changed_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteBookingStatusLogByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.BookingStatusLogEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

func normalizeChangedByRole(value *string) (*ent.MemberRole, error) {
	if value == nil || strings.TrimSpace(*value) == "" {
		return nil, nil
	}

	normalized := ent.MemberRole(strings.TrimSpace(*value))
	switch normalized {
	case ent.MemberRoleCustomer, ent.MemberRoleAdmin:
		return &normalized, nil
	default:
		return nil, fmt.Errorf("changed_by_role is invalid")
	}
}

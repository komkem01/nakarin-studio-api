package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateBooking(ctx context.Context, bookingNo string, status *string, payment *string) (*ent.BookingEntity, error) {
	value := strings.TrimSpace(bookingNo)
	if value == "" {
		return nil, fmt.Errorf("booking_no is required")
	}

	statusValue, err := normalizeBookingStatus(status)
	if err != nil {
		return nil, err
	}

	paymentValue, err := normalizePaymentType(payment)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingEntity{
		ID:        uuid.New(),
		BookingNo: value,
		Status:    statusValue,
		Payment:   paymentValue,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetBookingByID(ctx context.Context, id string) (*ent.BookingEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingEntity{}
	if err := s.db.NewSelect().Model(model).Where("booking.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListBookings(ctx context.Context, status *string, payment *string) ([]*ent.BookingEntity, error) {
	items := make([]*ent.BookingEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("booking.created_at DESC")

	if status != nil {
		statusValue, err := normalizeBookingStatus(status)
		if err != nil {
			return nil, err
		}
		q = q.Where("booking.status = ?", statusValue)
	}

	if payment != nil {
		paymentValue, err := normalizePaymentType(payment)
		if err != nil {
			return nil, err
		}
		q = q.Where("booking.payment = ?", paymentValue)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateBookingByID(ctx context.Context, id string, bookingNo *string, status *string, payment *string) (*ent.BookingEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingEntity{}
	if err := s.db.NewSelect().Model(model).Where("booking.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if bookingNo != nil {
		value := strings.TrimSpace(*bookingNo)
		if value == "" {
			return nil, fmt.Errorf("booking_no is required")
		}
		model.BookingNo = value
	}

	if status != nil {
		statusValue, err := normalizeBookingStatus(status)
		if err != nil {
			return nil, err
		}
		model.Status = statusValue
	}

	if payment != nil {
		paymentValue, err := normalizePaymentType(payment)
		if err != nil {
			return nil, err
		}
		model.Payment = paymentValue
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("booking_no", "status", "payment", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteBookingByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.BookingEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

func normalizeBookingStatus(value *string) (ent.BookingStatus, error) {
	if value == nil {
		return ent.BookingStatusPending, nil
	}

	raw := strings.TrimSpace(*value)
	if raw == "" {
		return "", fmt.Errorf("status is invalid")
	}

	normalized := ent.BookingStatus(raw)
	switch normalized {
	case ent.BookingStatusPending, ent.BookingStatusProcessing, ent.BookingStatusCompleted, ent.BookingStatusCanceled:
		return normalized, nil
	default:
		return "", fmt.Errorf("status is invalid")
	}
}

func normalizePaymentType(value *string) (ent.PaymentType, error) {
	if value == nil {
		return ent.PaymentTypeDeposit, nil
	}

	raw := strings.TrimSpace(*value)
	if raw == "" {
		return "", fmt.Errorf("payment is invalid")
	}

	normalized := ent.PaymentType(raw)
	switch normalized {
	case ent.PaymentTypeDeposit, ent.PaymentTypePaid:
		return normalized, nil
	default:
		return "", fmt.Errorf("payment is invalid")
	}
}

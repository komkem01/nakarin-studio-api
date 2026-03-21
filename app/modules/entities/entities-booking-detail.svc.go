package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateBookingDetail(ctx context.Context, bookingID string, firstName string, lastName *string, phone string) (*ent.BookingDetailEntity, error) {
	bid, err := uuid.Parse(strings.TrimSpace(bookingID))
	if err != nil {
		return nil, err
	}

	firstNameValue := strings.TrimSpace(firstName)
	if firstNameValue == "" {
		return nil, fmt.Errorf("first_name is required")
	}

	phoneValue := strings.TrimSpace(phone)
	if phoneValue == "" {
		return nil, fmt.Errorf("phone is required")
	}

	model := &ent.BookingDetailEntity{
		ID:        uuid.New(),
		BookingID: bid,
		FirstName: firstNameValue,
		LastName:  normalizeOptionalString(lastName),
		Phone:     phoneValue,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetBookingDetailByID(ctx context.Context, id string) (*ent.BookingDetailEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingDetailEntity{}
	if err := s.db.NewSelect().Model(model).Where("booking_detail.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListBookingDetails(ctx context.Context, bookingID *string) ([]*ent.BookingDetailEntity, error) {
	items := make([]*ent.BookingDetailEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("booking_detail.created_at DESC")

	if bookingID != nil && strings.TrimSpace(*bookingID) != "" {
		bid, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		q = q.Where("booking_detail.booking_id = ?", bid)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateBookingDetailByID(ctx context.Context, id string, bookingID *string, firstName *string, lastName *string, phone *string) (*ent.BookingDetailEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingDetailEntity{}
	if err := s.db.NewSelect().Model(model).Where("booking_detail.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if bookingID != nil {
		bid, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		model.BookingID = bid
	}

	if firstName != nil {
		firstNameValue := strings.TrimSpace(*firstName)
		if firstNameValue == "" {
			return nil, fmt.Errorf("first_name is required")
		}
		model.FirstName = firstNameValue
	}

	if lastName != nil {
		model.LastName = normalizeOptionalString(lastName)
	}

	if phone != nil {
		phoneValue := strings.TrimSpace(*phone)
		if phoneValue == "" {
			return nil, fmt.Errorf("phone is required")
		}
		model.Phone = phoneValue
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("booking_id", "first_name", "last_name", "phone", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteBookingDetailByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.BookingDetailEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

func normalizeOptionalString(value *string) *string {
	if value == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}

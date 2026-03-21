package entities

import (
	"context"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateMemberBooking(ctx context.Context, memberID string, bookingID string) (*ent.MemberBookingEntity, error) {
	memberUUID, err := uuid.Parse(strings.TrimSpace(memberID))
	if err != nil {
		return nil, err
	}

	bookingUUID, err := uuid.Parse(strings.TrimSpace(bookingID))
	if err != nil {
		return nil, err
	}

	model := &ent.MemberBookingEntity{
		ID:        uuid.New(),
		MemberID:  memberUUID,
		BookingID: bookingUUID,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetMemberBookingByID(ctx context.Context, id string) (*ent.MemberBookingEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.MemberBookingEntity{}
	if err := s.db.NewSelect().Model(model).Where("member_booking.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListMemberBookings(ctx context.Context, memberID *string, bookingID *string) ([]*ent.MemberBookingEntity, error) {
	items := make([]*ent.MemberBookingEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("member_booking.created_at DESC")

	if memberID != nil && strings.TrimSpace(*memberID) != "" {
		memberUUID, err := uuid.Parse(strings.TrimSpace(*memberID))
		if err != nil {
			return nil, err
		}
		q = q.Where("member_booking.member_id = ?", memberUUID)
	}

	if bookingID != nil && strings.TrimSpace(*bookingID) != "" {
		bookingUUID, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		q = q.Where("member_booking.booking_id = ?", bookingUUID)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateMemberBookingByID(ctx context.Context, id string, memberID *string, bookingID *string) (*ent.MemberBookingEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.MemberBookingEntity{}
	if err := s.db.NewSelect().Model(model).Where("member_booking.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if memberID != nil {
		memberUUID, err := uuid.Parse(strings.TrimSpace(*memberID))
		if err != nil {
			return nil, err
		}
		model.MemberID = memberUUID
	}

	if bookingID != nil {
		bookingUUID, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		model.BookingID = bookingUUID
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("member_id", "booking_id").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteMemberBookingByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.MemberBookingEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

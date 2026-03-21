package entities

import (
	"context"
	"fmt"
	"strings"
	"time"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateBooking(ctx context.Context, bookingNo string, status *string, payment *string, cancelledReason *string, internalNote *string, trackingAttemptCount *int, lastTrackingAt *time.Time, deliveryMemberAddressID *string, deliveryFirstName *string, deliveryLastName *string, deliveryPhone *string, deliveryNo *string, deliveryVillage *string, deliveryStreet *string, deliveryProvinceID *string, deliveryDistrictID *string, deliverySubDistrictID *string, deliveryZipcodeID *string, deliveryNote *string) (*ent.BookingEntity, error) {
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

	trackingValue := 0
	if trackingAttemptCount != nil {
		if *trackingAttemptCount < 0 {
			return nil, fmt.Errorf("tracking_attempt_count must be greater than or equal to 0")
		}
		trackingValue = *trackingAttemptCount
	}

	parsedDeliveryMemberAddressID, parsedDeliveryProvinceID, parsedDeliveryDistrictID, parsedDeliverySubDistrictID, parsedDeliveryZipcodeID, err := normalizeBookingDelivery(ctx, s, deliveryMemberAddressID, deliveryFirstName, deliveryLastName, deliveryPhone, deliveryNo, deliveryVillage, deliveryStreet, deliveryProvinceID, deliveryDistrictID, deliverySubDistrictID, deliveryZipcodeID)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingEntity{
		ID:                      uuid.New(),
		BookingNo:               value,
		Status:                  statusValue,
		Payment:                 paymentValue,
		CancelledReason:         normalizeOptionalString(cancelledReason),
		InternalNote:            normalizeOptionalString(internalNote),
		TrackingAttemptCount:    trackingValue,
		LastTrackingAt:          lastTrackingAt,
		DeliveryMemberAddressID: parsedDeliveryMemberAddressID,
		DeliveryFirstName:       normalizeOptionalString(deliveryFirstName),
		DeliveryLastName:        normalizeOptionalString(deliveryLastName),
		DeliveryPhone:           normalizeOptionalString(deliveryPhone),
		DeliveryNo:              normalizeOptionalString(deliveryNo),
		DeliveryVillage:         normalizeOptionalString(deliveryVillage),
		DeliveryStreet:          normalizeOptionalString(deliveryStreet),
		DeliveryProvinceID:      parsedDeliveryProvinceID,
		DeliveryDistrictID:      parsedDeliveryDistrictID,
		DeliverySubDistrictID:   parsedDeliverySubDistrictID,
		DeliveryZipcodeID:       parsedDeliveryZipcodeID,
		DeliveryNote:            normalizeOptionalString(deliveryNote),
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

func (s *Service) GetBookingByBookingNoAndPhone(ctx context.Context, bookingNo string, phone string) (*ent.BookingEntity, error) {
	bookingNoValue := strings.TrimSpace(bookingNo)
	if bookingNoValue == "" {
		return nil, fmt.Errorf("booking_no is required")
	}

	phoneValue := strings.TrimSpace(phone)
	if phoneValue == "" {
		return nil, fmt.Errorf("phone is required")
	}

	model := &ent.BookingEntity{}
	q := s.db.NewSelect().Model(model).
		Where("booking.booking_no = ?", bookingNoValue).
		Where("exists (select 1 from booking_details bd where bd.booking_id = booking.id and bd.phone = ? and bd.deleted_at is null)", phoneValue)

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListBookings(ctx context.Context, status *string, payment *string, bookingNo *string, phone *string, createdAtFrom *time.Time, createdAtTo *time.Time) ([]*ent.BookingEntity, error) {
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

	if bookingNo != nil && strings.TrimSpace(*bookingNo) != "" {
		q = q.Where("booking.booking_no ILIKE ?", "%"+strings.TrimSpace(*bookingNo)+"%")
	}

	if phone != nil && strings.TrimSpace(*phone) != "" {
		q = q.Where("exists (select 1 from booking_details bd where bd.booking_id = booking.id and bd.phone ILIKE ? and bd.deleted_at is null)", "%"+strings.TrimSpace(*phone)+"%")
	}

	if createdAtFrom != nil {
		q = q.Where("booking.created_at >= ?", *createdAtFrom)
	}

	if createdAtTo != nil {
		q = q.Where("booking.created_at <= ?", *createdAtTo)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateBookingByID(ctx context.Context, id string, bookingNo *string, status *string, payment *string, cancelledReason *string, internalNote *string, trackingAttemptCount *int, lastTrackingAt *time.Time, deliveryMemberAddressID *string, deliveryFirstName *string, deliveryLastName *string, deliveryPhone *string, deliveryNo *string, deliveryVillage *string, deliveryStreet *string, deliveryProvinceID *string, deliveryDistrictID *string, deliverySubDistrictID *string, deliveryZipcodeID *string, deliveryNote *string) (*ent.BookingEntity, error) {
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

	if cancelledReason != nil {
		model.CancelledReason = normalizeOptionalString(cancelledReason)
	}
	if internalNote != nil {
		model.InternalNote = normalizeOptionalString(internalNote)
	}
	if trackingAttemptCount != nil {
		if *trackingAttemptCount < 0 {
			return nil, fmt.Errorf("tracking_attempt_count must be greater than or equal to 0")
		}
		model.TrackingAttemptCount = *trackingAttemptCount
	}
	if lastTrackingAt != nil {
		model.LastTrackingAt = lastTrackingAt
	}

	if deliveryMemberAddressID != nil || deliveryFirstName != nil || deliveryLastName != nil || deliveryPhone != nil || deliveryNo != nil || deliveryVillage != nil || deliveryStreet != nil || deliveryProvinceID != nil || deliveryDistrictID != nil || deliverySubDistrictID != nil || deliveryZipcodeID != nil {
		currentDeliveryMemberAddressID := uuidToStringPtr(model.DeliveryMemberAddressID)
		currentDeliveryFirstName := model.DeliveryFirstName
		currentDeliveryLastName := model.DeliveryLastName
		currentDeliveryPhone := model.DeliveryPhone
		currentDeliveryNo := model.DeliveryNo
		currentDeliveryVillage := model.DeliveryVillage
		currentDeliveryStreet := model.DeliveryStreet
		currentDeliveryProvinceID := uuidToStringPtr(model.DeliveryProvinceID)
		currentDeliveryDistrictID := uuidToStringPtr(model.DeliveryDistrictID)
		currentDeliverySubDistrictID := uuidToStringPtr(model.DeliverySubDistrictID)
		currentDeliveryZipcodeID := uuidToStringPtr(model.DeliveryZipcodeID)

		if deliveryMemberAddressID != nil {
			currentDeliveryMemberAddressID = deliveryMemberAddressID
		}
		if deliveryFirstName != nil {
			currentDeliveryFirstName = deliveryFirstName
		}
		if deliveryLastName != nil {
			currentDeliveryLastName = deliveryLastName
		}
		if deliveryPhone != nil {
			currentDeliveryPhone = deliveryPhone
		}
		if deliveryNo != nil {
			currentDeliveryNo = deliveryNo
		}
		if deliveryVillage != nil {
			currentDeliveryVillage = deliveryVillage
		}
		if deliveryStreet != nil {
			currentDeliveryStreet = deliveryStreet
		}
		if deliveryProvinceID != nil {
			currentDeliveryProvinceID = deliveryProvinceID
		}
		if deliveryDistrictID != nil {
			currentDeliveryDistrictID = deliveryDistrictID
		}
		if deliverySubDistrictID != nil {
			currentDeliverySubDistrictID = deliverySubDistrictID
		}
		if deliveryZipcodeID != nil {
			currentDeliveryZipcodeID = deliveryZipcodeID
		}

		parsedDeliveryMemberAddressID, parsedDeliveryProvinceID, parsedDeliveryDistrictID, parsedDeliverySubDistrictID, parsedDeliveryZipcodeID, err := normalizeBookingDelivery(ctx, s, currentDeliveryMemberAddressID, currentDeliveryFirstName, currentDeliveryLastName, currentDeliveryPhone, currentDeliveryNo, currentDeliveryVillage, currentDeliveryStreet, currentDeliveryProvinceID, currentDeliveryDistrictID, currentDeliverySubDistrictID, currentDeliveryZipcodeID)
		if err != nil {
			return nil, err
		}

		model.DeliveryMemberAddressID = parsedDeliveryMemberAddressID
		model.DeliveryFirstName = normalizeOptionalString(currentDeliveryFirstName)
		model.DeliveryLastName = normalizeOptionalString(currentDeliveryLastName)
		model.DeliveryPhone = normalizeOptionalString(currentDeliveryPhone)
		model.DeliveryNo = normalizeOptionalString(currentDeliveryNo)
		model.DeliveryVillage = normalizeOptionalString(currentDeliveryVillage)
		model.DeliveryStreet = normalizeOptionalString(currentDeliveryStreet)
		model.DeliveryProvinceID = parsedDeliveryProvinceID
		model.DeliveryDistrictID = parsedDeliveryDistrictID
		model.DeliverySubDistrictID = parsedDeliverySubDistrictID
		model.DeliveryZipcodeID = parsedDeliveryZipcodeID
	}

	if deliveryNote != nil {
		model.DeliveryNote = normalizeOptionalString(deliveryNote)
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("booking_no", "status", "payment", "cancelled_reason", "internal_note", "tracking_attempt_count", "last_tracking_at", "delivery_member_address_id", "delivery_first_name", "delivery_last_name", "delivery_phone", "delivery_no", "delivery_village", "delivery_street", "delivery_province_id", "delivery_district_id", "delivery_sub_district_id", "delivery_zipcode_id", "delivery_note", "updated_at").Exec(ctx)
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

func normalizeBookingDelivery(ctx context.Context, s *Service, deliveryMemberAddressID *string, deliveryFirstName *string, deliveryLastName *string, deliveryPhone *string, deliveryNo *string, deliveryVillage *string, deliveryStreet *string, deliveryProvinceID *string, deliveryDistrictID *string, deliverySubDistrictID *string, deliveryZipcodeID *string) (*uuid.UUID, *uuid.UUID, *uuid.UUID, *uuid.UUID, *uuid.UUID, error) {
	firstNameValue := strings.TrimSpace(getPointerValue(deliveryFirstName))
	phoneValue := strings.TrimSpace(getPointerValue(deliveryPhone))
	noValue := strings.TrimSpace(getPointerValue(deliveryNo))

	if firstNameValue == "" {
		return nil, nil, nil, nil, nil, fmt.Errorf("delivery_first_name is required")
	}
	if phoneValue == "" {
		return nil, nil, nil, nil, nil, fmt.Errorf("delivery_phone is required")
	}
	if noValue == "" {
		return nil, nil, nil, nil, nil, fmt.Errorf("delivery_no is required")
	}

	provinceUUID, err := parseOptionalUUIDRequired(deliveryProvinceID, "delivery_province_id")
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	districtUUID, err := parseOptionalUUIDRequired(deliveryDistrictID, "delivery_district_id")
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	subDistrictUUID, err := parseOptionalUUIDRequired(deliverySubDistrictID, "delivery_sub_district_id")
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	zipcodeUUID, err := parseOptionalUUIDRequired(deliveryZipcodeID, "delivery_zipcode_id")
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	memberAddressUUID, err := parseOptionalUUID(deliveryMemberAddressID)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	if memberAddressUUID != nil {
		address := &ent.MemberAddressEntity{}
		if err := s.db.NewSelect().Model(address).Where("member_address.id = ?", *memberAddressUUID).Scan(ctx); err != nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("delivery_member_address_id not found")
		}
	}

	return memberAddressUUID, provinceUUID, districtUUID, subDistrictUUID, zipcodeUUID, nil
}

func parseOptionalUUIDRequired(value *string, field string) (*uuid.UUID, error) {
	if value == nil || strings.TrimSpace(*value) == "" {
		return nil, fmt.Errorf("%s is required", field)
	}

	uid, err := uuid.Parse(strings.TrimSpace(*value))
	if err != nil {
		return nil, err
	}

	return &uid, nil
}

func getPointerValue(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}

func uuidToStringPtr(value *uuid.UUID) *string {
	if value == nil {
		return nil
	}

	s := value.String()
	return &s
}

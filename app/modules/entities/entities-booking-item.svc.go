package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateBookingItem(ctx context.Context, bookingID string, productID string, productName string, unitPriceAtBooking float64, quantity int, lineTotal float64, note *string, sortOrder int) (*ent.BookingItemEntity, error) {
	bookingUUID, productUUID, err := parseBookingItemRequiredUUIDs(bookingID, productID)
	if err != nil {
		return nil, err
	}

	nameValue := strings.TrimSpace(productName)
	if nameValue == "" {
		return nil, fmt.Errorf("product_name is required")
	}
	if unitPriceAtBooking < 0 {
		return nil, fmt.Errorf("unit_price_at_booking must be greater than or equal to 0")
	}
	if quantity <= 0 {
		return nil, fmt.Errorf("quantity must be greater than 0")
	}
	if lineTotal < 0 {
		return nil, fmt.Errorf("line_total must be greater than or equal to 0")
	}

	model := &ent.BookingItemEntity{
		ID:                 uuid.New(),
		BookingID:          bookingUUID,
		ProductID:          productUUID,
		ProductName:        nameValue,
		UnitPriceAtBooking: unitPriceAtBooking,
		Quantity:           quantity,
		LineTotal:          lineTotal,
		Note:               normalizeOptionalString(note),
		SortOrder:          sortOrder,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetBookingItemByID(ctx context.Context, id string) (*ent.BookingItemEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingItemEntity{}
	if err := s.db.NewSelect().Model(model).Where("booking_item.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListBookingItems(ctx context.Context, bookingID *string, productID *string) ([]*ent.BookingItemEntity, error) {
	items := make([]*ent.BookingItemEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("booking_item.sort_order ASC").Order("booking_item.created_at DESC")

	if bookingID != nil && strings.TrimSpace(*bookingID) != "" {
		bookingUUID, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		q = q.Where("booking_item.booking_id = ?", bookingUUID)
	}

	if productID != nil && strings.TrimSpace(*productID) != "" {
		productUUID, err := uuid.Parse(strings.TrimSpace(*productID))
		if err != nil {
			return nil, err
		}
		q = q.Where("booking_item.product_id = ?", productUUID)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdateBookingItemByID(ctx context.Context, id string, bookingID *string, productID *string, productName *string, unitPriceAtBooking *float64, quantity *int, lineTotal *float64, note *string, sortOrder *int) (*ent.BookingItemEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.BookingItemEntity{}
	if err := s.db.NewSelect().Model(model).Where("booking_item.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if bookingID != nil {
		bookingUUID, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		model.BookingID = bookingUUID
	}
	if productID != nil {
		productUUID, err := uuid.Parse(strings.TrimSpace(*productID))
		if err != nil {
			return nil, err
		}
		model.ProductID = productUUID
	}
	if productName != nil {
		nameValue := strings.TrimSpace(*productName)
		if nameValue == "" {
			return nil, fmt.Errorf("product_name is required")
		}
		model.ProductName = nameValue
	}
	if unitPriceAtBooking != nil {
		if *unitPriceAtBooking < 0 {
			return nil, fmt.Errorf("unit_price_at_booking must be greater than or equal to 0")
		}
		model.UnitPriceAtBooking = *unitPriceAtBooking
	}
	if quantity != nil {
		if *quantity <= 0 {
			return nil, fmt.Errorf("quantity must be greater than 0")
		}
		model.Quantity = *quantity
	}
	if lineTotal != nil {
		if *lineTotal < 0 {
			return nil, fmt.Errorf("line_total must be greater than or equal to 0")
		}
		model.LineTotal = *lineTotal
	}
	if note != nil {
		model.Note = normalizeOptionalString(note)
	}
	if sortOrder != nil {
		model.SortOrder = *sortOrder
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("booking_id", "product_id", "product_name", "unit_price_at_booking", "quantity", "line_total", "note", "sort_order", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeleteBookingItemByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.BookingItemEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

func parseBookingItemRequiredUUIDs(bookingID string, productID string) (uuid.UUID, uuid.UUID, error) {
	bookingUUID, err := uuid.Parse(strings.TrimSpace(bookingID))
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}
	productUUID, err := uuid.Parse(strings.TrimSpace(productID))
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}
	return bookingUUID, productUUID, nil
}

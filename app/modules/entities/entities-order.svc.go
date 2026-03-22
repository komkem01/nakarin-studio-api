package entities

import (
	"context"
	"fmt"
	"strings"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreateOrder(ctx context.Context, bookingID string, orderNo string, status *string, totalAmount *float64) (*ent.OrderEntity, error) {
	bookingUUID, err := uuid.Parse(strings.TrimSpace(bookingID))
	if err != nil {
		return nil, err
	}

	orderNoValue := strings.TrimSpace(orderNo)
	if orderNoValue == "" {
		return nil, fmt.Errorf("order_no is required")
	}

	statusValue, err := normalizeOrderStatus(status)
	if err != nil {
		return nil, err
	}

	totalAmountValue := 0.0
	if totalAmount != nil {
		if *totalAmount < 0 {
			return nil, fmt.Errorf("total_amount must be greater than or equal to 0")
		}
		totalAmountValue = *totalAmount
	}

	model := &ent.OrderEntity{
		ID:          uuid.New(),
		BookingID:   bookingUUID,
		OrderNo:     orderNoValue,
		Status:      statusValue,
		TotalAmount: totalAmountValue,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetOrderByID(ctx context.Context, id string) (*ent.OrderEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.OrderEntity{}
	if err := s.db.NewSelect().Model(model).Where("order.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetOrderByBookingID(ctx context.Context, bookingID string) (*ent.OrderEntity, error) {
	uid, err := uuid.Parse(strings.TrimSpace(bookingID))
	if err != nil {
		return nil, err
	}

	model := &ent.OrderEntity{}
	if err := s.db.NewSelect().Model(model).Where("order.booking_id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func normalizeOrderStatus(value *string) (ent.OrderStatus, error) {
	if value == nil || strings.TrimSpace(*value) == "" {
		return ent.OrderStatusNew, nil
	}

	normalized := ent.OrderStatus(strings.TrimSpace(*value))
	switch normalized {
	case ent.OrderStatusNew, ent.OrderStatusPreparing, ent.OrderStatusReady, ent.OrderStatusCompleted:
		return normalized, nil
	default:
		return "", fmt.Errorf("order status is invalid")
	}
}

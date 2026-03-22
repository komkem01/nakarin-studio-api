package booking

import (
	"context"
	"database/sql"
	"fmt"

	"nakarin-studio/app/modules/entities"
	"nakarin-studio/app/modules/entities/ent"
)

func (s *Service) ConvertToOrder(ctx context.Context, bookingID string, reason *string, changedBy *string) (*ent.OrderEntity, error) {
	if s.txDB == nil {
		return nil, fmt.Errorf("transaction service is not configured")
	}

	var createdOrder *ent.OrderEntity
	err := s.txDB.RunInTx(ctx, func(ctx context.Context, txSvc *entities.Service) error {
		booking, err := txSvc.GetBookingByID(ctx, bookingID)
		if err != nil {
			return err
		}

		if booking.Status == ent.BookingStatusCanceled {
			return fmt.Errorf("cannot convert canceled booking")
		}

		existingOrder, err := txSvc.GetOrderByBookingID(ctx, bookingID)
		if err == nil {
			createdOrder = existingOrder
			return nil
		}
		if err != sql.ErrNoRows {
			return err
		}

		items, err := txSvc.ListBookingItems(ctx, &bookingID, nil)
		if err != nil {
			return err
		}

		totalAmount := 0.0
		for _, item := range items {
			totalAmount += item.LineTotal
		}

		orderNo := fmt.Sprintf("ORD-%s", booking.BookingNo)
		createdOrder, err = txSvc.CreateOrder(ctx, bookingID, orderNo, stringPtr(string(ent.OrderStatusNew)), &totalAmount)
		if err != nil {
			return err
		}

		if booking.Status != ent.BookingStatusCompleted {
			targetStatus := string(ent.BookingStatusCompleted)
			if _, err := txSvc.UpdateBookingByID(
				ctx,
				bookingID,
				nil,
				&targetStatus,
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
			); err != nil {
				return err
			}

			logReason := reason
			if logReason == nil {
				defaultReason := fmt.Sprintf("convert to order %s", orderNo)
				logReason = &defaultReason
			}
			if _, err := txSvc.CreateBookingStatusLog(ctx, bookingID, stringPtr(string(booking.Status)), targetStatus, changedBy, stringPtr("admin"), logReason, nil); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}

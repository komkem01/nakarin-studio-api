package booking

import (
	"context"
	"fmt"

	"nakarin-studio/app/modules/entities"
)

func (s *Service) AggregateCreate(ctx context.Context, req *AggregateCreateRequest) error {
	if req == nil {
		return fmt.Errorf("request is required")
	}
	if len(req.Items) == 0 {
		return fmt.Errorf("items is required")
	}
	if s.txDB == nil {
		return fmt.Errorf("transaction database is not initialized")
	}

	return s.txDB.RunInTx(ctx, func(ctx context.Context, txSvc *entities.Service) error {
		createdBooking, err := txSvc.CreateBooking(
			ctx,
			req.Booking.BookingNo,
			req.Booking.Status,
			req.Booking.Payment,
			req.Booking.CancelledReason,
			req.Booking.InternalNote,
			req.Booking.TrackingAttemptCount,
			req.Booking.LastTrackingAt,
			req.Booking.DeliveryMemberAddressID,
			req.Booking.DeliveryFirstName,
			req.Booking.DeliveryLastName,
			req.Booking.DeliveryPhone,
			req.Booking.DeliveryNo,
			req.Booking.DeliveryVillage,
			req.Booking.DeliveryStreet,
			req.Booking.DeliveryProvinceID,
			req.Booking.DeliveryDistrictID,
			req.Booking.DeliverySubDistrictID,
			req.Booking.DeliveryZipcodeID,
			req.Booking.DeliveryNote,
		)
		if err != nil {
			return err
		}

		_, err = txSvc.CreateBookingDetail(ctx, createdBooking.ID.String(), req.Detail.FirstName, req.Detail.LastName, req.Detail.Phone)
		if err != nil {
			return err
		}

		for i, item := range req.Items {
			sortOrder := i
			if item.SortOrder != nil {
				sortOrder = *item.SortOrder
			}
			_, err := txSvc.CreateBookingItem(ctx, createdBooking.ID.String(), item.ProductID, item.ProductName, item.UnitPriceAtBooking, item.Quantity, item.LineTotal, item.Note, sortOrder)
			if err != nil {
				return err
			}
		}

		if req.Payment != nil {
			_, err := txSvc.CreatePayment(ctx, createdBooking.ID.String(), req.Payment.Channel, req.Payment.Amount, req.Payment.DepositAmount, req.Payment.Status, req.Payment.ProofURL, req.Payment.Note, req.Payment.PaidAt)
			if err != nil {
				return err
			}
		}

		_, err = txSvc.CreateBookingStatusLog(ctx, createdBooking.ID.String(), nil, string(createdBooking.Status), nil, stringPtr("customer"), nil, nil)
		return err
	})
}

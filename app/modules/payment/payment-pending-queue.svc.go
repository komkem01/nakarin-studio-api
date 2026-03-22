package payment

import (
	"context"
	"time"
)

type PendingQueueItem struct {
	PaymentID      string     `json:"payment_id"`
	BookingID      string     `json:"booking_id"`
	BookingNo      string     `json:"booking_no"`
	CustomerName   string     `json:"customer_name"`
	Phone          string     `json:"phone"`
	Channel        string     `json:"channel"`
	Amount         float64    `json:"amount"`
	DepositAmount  float64    `json:"deposit_amount"`
	Status         string     `json:"status"`
	ProofURL       *string    `json:"proof_url"`
	Note           *string    `json:"note"`
	PaidAt         *time.Time `json:"paid_at"`
	PaymentCreated time.Time  `json:"payment_created_at"`
}

func (s *Service) PendingQueue(ctx context.Context) ([]*PendingQueueItem, error) {
	pending := "pending"
	items, err := s.db.ListPayments(ctx, nil, nil, &pending)
	if err != nil {
		return nil, err
	}

	out := make([]*PendingQueueItem, 0, len(items))
	for _, item := range items {
		booking, err := s.booking.GetBookingByID(ctx, item.BookingID.String())
		if err != nil {
			continue
		}

		bookingID := booking.ID.String()
		details, err := s.detail.ListBookingDetails(ctx, &bookingID)
		if err != nil {
			details = nil
		}

		customerName := "ไม่ระบุชื่อ"
		phone := "-"
		if len(details) > 0 && details[0] != nil {
			firstName := details[0].FirstName
			lastName := ""
			if details[0].LastName != nil {
				lastName = *details[0].LastName
			}
			if firstName != "" || lastName != "" {
				customerName = firstName
				if lastName != "" {
					if customerName != "" {
						customerName += " "
					}
					customerName += lastName
				}
			}
			if details[0].Phone != "" {
				phone = details[0].Phone
			}
		}

		out = append(out, &PendingQueueItem{
			PaymentID:      item.ID.String(),
			BookingID:      booking.ID.String(),
			BookingNo:      booking.BookingNo,
			CustomerName:   customerName,
			Phone:          phone,
			Channel:        string(item.Channel),
			Amount:         item.Amount,
			DepositAmount:  item.DepositAmount,
			Status:         string(item.Status),
			ProofURL:       item.ProofURL,
			Note:           item.Note,
			PaidAt:         item.PaidAt,
			PaymentCreated: item.CreatedAt,
		})
	}

	return out, nil
}

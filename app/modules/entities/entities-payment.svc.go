package entities

import (
	"context"
	"fmt"
	"strings"
	"time"

	"nakarin-studio/app/modules/entities/ent"

	"github.com/google/uuid"
)

func (s *Service) CreatePayment(ctx context.Context, bookingID string, channel *string, amount float64, depositAmount float64, status *string, proofURL *string, note *string, paidAt *time.Time) (*ent.PaymentEntity, error) {
	bookingUUID, err := uuid.Parse(strings.TrimSpace(bookingID))
	if err != nil {
		return nil, err
	}
	if amount < 0 {
		return nil, fmt.Errorf("amount must be greater than or equal to 0")
	}
	if depositAmount < 0 {
		return nil, fmt.Errorf("deposit_amount must be greater than or equal to 0")
	}
	if err := validatePaymentAmounts(amount, depositAmount); err != nil {
		return nil, err
	}

	channelValue, err := normalizePaymentChannel(channel)
	if err != nil {
		return nil, err
	}
	statusValue, err := normalizePaymentStatus(status)
	if err != nil {
		return nil, err
	}

	model := &ent.PaymentEntity{
		ID:            uuid.New(),
		BookingID:     bookingUUID,
		Channel:       channelValue,
		Amount:        amount,
		DepositAmount: depositAmount,
		Status:        statusValue,
		ProofURL:      normalizeOptionalString(proofURL),
		Note:          normalizeOptionalString(note),
		PaidAt:        paidAt,
	}

	_, err = s.db.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) GetPaymentByID(ctx context.Context, id string) (*ent.PaymentEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.PaymentEntity{}
	if err := s.db.NewSelect().Model(model).Where("payment.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) ListPayments(ctx context.Context, bookingID *string, channel *string, status *string) ([]*ent.PaymentEntity, error) {
	items := make([]*ent.PaymentEntity, 0)
	q := s.db.NewSelect().Model(&items).Order("payment.created_at DESC")

	if bookingID != nil && strings.TrimSpace(*bookingID) != "" {
		bookingUUID, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		q = q.Where("payment.booking_id = ?", bookingUUID)
	}
	if channel != nil && strings.TrimSpace(*channel) != "" {
		channelValue, err := normalizePaymentChannel(channel)
		if err != nil {
			return nil, err
		}
		q = q.Where("payment.channel = ?", channelValue)
	}
	if status != nil && strings.TrimSpace(*status) != "" {
		statusValue, err := normalizePaymentStatus(status)
		if err != nil {
			return nil, err
		}
		q = q.Where("payment.status = ?", statusValue)
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Service) UpdatePaymentByID(ctx context.Context, id string, bookingID *string, channel *string, amount *float64, depositAmount *float64, status *string, proofURL *string, note *string, paidAt *time.Time) (*ent.PaymentEntity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	model := &ent.PaymentEntity{}
	if err := s.db.NewSelect().Model(model).Where("payment.id = ?", uid).Scan(ctx); err != nil {
		return nil, err
	}

	if bookingID != nil {
		bookingUUID, err := uuid.Parse(strings.TrimSpace(*bookingID))
		if err != nil {
			return nil, err
		}
		model.BookingID = bookingUUID
	}
	if channel != nil {
		channelValue, err := normalizePaymentChannel(channel)
		if err != nil {
			return nil, err
		}
		model.Channel = channelValue
	}
	if amount != nil {
		if *amount < 0 {
			return nil, fmt.Errorf("amount must be greater than or equal to 0")
		}
		model.Amount = *amount
	}
	if depositAmount != nil {
		if *depositAmount < 0 {
			return nil, fmt.Errorf("deposit_amount must be greater than or equal to 0")
		}
		model.DepositAmount = *depositAmount
	}
	if err := validatePaymentAmounts(model.Amount, model.DepositAmount); err != nil {
		return nil, err
	}
	if status != nil {
		statusValue, err := normalizePaymentStatus(status)
		if err != nil {
			return nil, err
		}
		model.Status = statusValue
	}
	if proofURL != nil {
		model.ProofURL = normalizeOptionalString(proofURL)
	}
	if note != nil {
		model.Note = normalizeOptionalString(note)
	}
	if paidAt != nil {
		model.PaidAt = paidAt
	}

	_, err = s.db.NewUpdate().Model(model).WherePK().Column("booking_id", "channel", "amount", "deposit_amount", "status", "proof_url", "note", "paid_at", "updated_at").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Service) DeletePaymentByID(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.NewDelete().Model(&ent.PaymentEntity{}).Where("id = ?", uid).Exec(ctx)
	return err
}

func normalizePaymentChannel(value *string) (ent.PaymentChannel, error) {
	if value == nil {
		return ent.PaymentChannelBankTransfer, nil
	}

	raw := strings.TrimSpace(*value)
	if raw == "" {
		return "", fmt.Errorf("channel is invalid")
	}

	normalized := ent.PaymentChannel(raw)
	switch normalized {
	case ent.PaymentChannelBankTransfer, ent.PaymentChannelPromptPay, ent.PaymentChannelCash, ent.PaymentChannelCreditCard, ent.PaymentChannelOther:
		return normalized, nil
	default:
		return "", fmt.Errorf("channel is invalid")
	}
}

func normalizePaymentStatus(value *string) (ent.PaymentStatus, error) {
	if value == nil {
		return ent.PaymentStatusPending, nil
	}

	raw := strings.TrimSpace(*value)
	if raw == "" {
		return "", fmt.Errorf("status is invalid")
	}

	normalized := ent.PaymentStatus(raw)
	switch normalized {
	case ent.PaymentStatusPending, ent.PaymentStatusPaid, ent.PaymentStatusFailed, ent.PaymentStatusRefunded:
		return normalized, nil
	default:
		return "", fmt.Errorf("status is invalid")
	}
}

func validatePaymentAmounts(amount float64, depositAmount float64) error {
	if depositAmount > amount {
		return fmt.Errorf("deposit_amount must be less than or equal to amount")
	}

	return nil
}

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaymentChannel string

const (
	PaymentChannelBankTransfer PaymentChannel = "bank_transfer"
	PaymentChannelPromptPay    PaymentChannel = "promptpay"
	PaymentChannelCash         PaymentChannel = "cash"
	PaymentChannelCreditCard   PaymentChannel = "credit_card"
	PaymentChannelOther        PaymentChannel = "other"
)

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusPaid     PaymentStatus = "paid"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

type PaymentEntity struct {
	bun.BaseModel `bun:"table:payments,alias:payment"`

	ID            uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	BookingID     uuid.UUID      `bun:"booking_id,type:uuid,notnull" json:"booking_id"`
	Channel       PaymentChannel `bun:"channel,type:payment_channel,notnull,default:'bank_transfer'" json:"channel"`
	Amount        float64        `bun:"amount,type:numeric(12,2),notnull,default:0" json:"amount"`
	DepositAmount float64        `bun:"deposit_amount,type:numeric(12,2),notnull,default:0" json:"deposit_amount"`
	Status        PaymentStatus  `bun:"status,type:payment_status,notnull,default:'pending'" json:"status"`
	ProofURL      *string        `bun:"proof_url" json:"proof_url"`
	Note          *string        `bun:"note" json:"note"`
	PaidAt        *time.Time     `bun:"paid_at" json:"paid_at"`
	CreatedAt     time.Time      `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time      `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt     *time.Time     `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

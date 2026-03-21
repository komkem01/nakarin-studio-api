package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BookingStatus string

const (
	BookingStatusPending    BookingStatus = "pending"
	BookingStatusProcessing BookingStatus = "processing"
	BookingStatusCompleted  BookingStatus = "completed"
	BookingStatusCanceled   BookingStatus = "canceled"
)

type PaymentType string

const (
	PaymentTypeDeposit PaymentType = "deposit"
	PaymentTypePaid    PaymentType = "paid"
)

type BookingEntity struct {
	bun.BaseModel `bun:"table:bookings,alias:booking"`

	ID        uuid.UUID     `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	BookingNo string        `bun:"booking_no,notnull,unique" json:"booking_no"`
	Status    BookingStatus `bun:"status,type:booking_status,notnull,default:'pending'" json:"status"`
	Payment   PaymentType   `bun:"payment,type:payment_type,notnull,default:'deposit'" json:"payment"`
	CreatedAt time.Time     `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time     `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time    `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

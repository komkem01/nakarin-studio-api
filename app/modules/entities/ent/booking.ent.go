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

	ID                      uuid.UUID     `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	BookingNo               string        `bun:"booking_no,notnull,unique" json:"booking_no"`
	Status                  BookingStatus `bun:"status,type:booking_status,notnull,default:'pending'" json:"status"`
	Payment                 PaymentType   `bun:"payment,type:payment_type,notnull,default:'deposit'" json:"payment"`
	CancelledReason         *string       `bun:"cancelled_reason" json:"cancelled_reason"`
	InternalNote            *string       `bun:"internal_note" json:"internal_note"`
	TrackingAttemptCount    int           `bun:"tracking_attempt_count,notnull,default:0" json:"tracking_attempt_count"`
	LastTrackingAt          *time.Time    `bun:"last_tracking_at" json:"last_tracking_at"`
	DeliveryMemberAddressID *uuid.UUID    `bun:"delivery_member_address_id,type:uuid" json:"delivery_member_address_id"`
	DeliveryFirstName       *string       `bun:"delivery_first_name" json:"delivery_first_name"`
	DeliveryLastName        *string       `bun:"delivery_last_name" json:"delivery_last_name"`
	DeliveryPhone           *string       `bun:"delivery_phone" json:"delivery_phone"`
	DeliveryNo              *string       `bun:"delivery_no" json:"delivery_no"`
	DeliveryVillage         *string       `bun:"delivery_village" json:"delivery_village"`
	DeliveryStreet          *string       `bun:"delivery_street" json:"delivery_street"`
	DeliveryProvinceID      *uuid.UUID    `bun:"delivery_province_id,type:uuid" json:"delivery_province_id"`
	DeliveryDistrictID      *uuid.UUID    `bun:"delivery_district_id,type:uuid" json:"delivery_district_id"`
	DeliverySubDistrictID   *uuid.UUID    `bun:"delivery_sub_district_id,type:uuid" json:"delivery_sub_district_id"`
	DeliveryZipcodeID       *uuid.UUID    `bun:"delivery_zipcode_id,type:uuid" json:"delivery_zipcode_id"`
	DeliveryNote            *string       `bun:"delivery_note" json:"delivery_note"`
	CreatedAt               time.Time     `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt               time.Time     `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt               *time.Time    `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

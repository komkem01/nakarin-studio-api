package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MemberBookingEntity struct {
	bun.BaseModel `bun:"table:member_booking,alias:member_booking"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	MemberID  uuid.UUID `bun:"member_id,type:uuid,notnull" json:"member_id"`
	BookingID uuid.UUID `bun:"booking_id,type:uuid,notnull" json:"booking_id"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
}

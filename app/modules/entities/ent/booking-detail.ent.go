package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BookingDetailEntity struct {
	bun.BaseModel `bun:"table:booking_details,alias:booking_detail"`

	ID        uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	BookingID uuid.UUID  `bun:"booking_id,type:uuid,notnull" json:"booking_id"`
	FirstName string     `bun:"first_name,notnull" json:"first_name"`
	LastName  *string    `bun:"last_name" json:"last_name"`
	Phone     string     `bun:"phone,notnull" json:"phone"`
	CreatedAt time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

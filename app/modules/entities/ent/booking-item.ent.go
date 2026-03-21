package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BookingItemEntity struct {
	bun.BaseModel `bun:"table:booking_items,alias:booking_item"`

	ID                 uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	BookingID          uuid.UUID  `bun:"booking_id,type:uuid,notnull" json:"booking_id"`
	ProductID          uuid.UUID  `bun:"product_id,type:uuid,notnull" json:"product_id"`
	ProductName        string     `bun:"product_name,notnull" json:"product_name"`
	UnitPriceAtBooking float64    `bun:"unit_price_at_booking,type:numeric(12,2),notnull,default:0" json:"unit_price_at_booking"`
	Quantity           int        `bun:"quantity,notnull,default:1" json:"quantity"`
	LineTotal          float64    `bun:"line_total,type:numeric(12,2),notnull,default:0" json:"line_total"`
	Note               *string    `bun:"note" json:"note"`
	SortOrder          int        `bun:"sort_order,notnull,default:0" json:"sort_order"`
	CreatedAt          time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt          time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt          *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

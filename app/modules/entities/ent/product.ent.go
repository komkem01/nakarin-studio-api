package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductEntity struct {
	bun.BaseModel `bun:"table:products,alias:product"`

	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CategoryID    *uuid.UUID `bun:"category_id,type:uuid" json:"category_id"`
	Name          string     `bun:"name,notnull" json:"name"`
	Description   *string    `bun:"description" json:"description"`
	SuitableFor   *string    `bun:"suitable_for" json:"suitable_for"`
	OnSite        *string    `bun:"on_site" json:"on_site"`
	ReceivedItems *string    `bun:"received_items" json:"received_items"`
	Note          *string    `bun:"note" json:"note"`
	Price         float64    `bun:"price,type:numeric(12,2),notnull,default:0" json:"price"`
	IsActive      bool       `bun:"is_active,notnull,default:true" json:"is_active"`
	IsAvailable   bool       `bun:"is_available,notnull,default:true" json:"is_available"`
	PrepTime      int        `bun:"prep_time,notnull,default:0" json:"prep_time"`
	SortOrder     int        `bun:"sort_order,notnull,default:0" json:"sort_order"`
	CreatedAt     time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductImageEntity struct {
	bun.BaseModel `bun:"table:product_images,alias:product_image"`

	ID        uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	ProductID uuid.UUID  `bun:"product_id,type:uuid,notnull" json:"product_id"`
	ImageURL  string     `bun:"image_url,notnull" json:"image_url"`
	AltText   *string    `bun:"alt_text" json:"alt_text"`
	SortOrder int        `bun:"sort_order,notnull,default:0" json:"sort_order"`
	IsActive  bool       `bun:"is_active,notnull,default:true" json:"is_active"`
	CreatedAt time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

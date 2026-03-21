package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DistrictEntity struct {
	bun.BaseModel `bun:"table:districts,alias:district"`

	ID         uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	ProvinceID uuid.UUID  `bun:"province_id,type:uuid,notnull" json:"province_id"`
	Name       string     `bun:"name,notnull" json:"name"`
	IsActive   bool       `bun:"is_active,notnull,default:true" json:"is_active"`
	CreatedAt  time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt  *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

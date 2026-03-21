package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ZipcodeEntity struct {
	bun.BaseModel `bun:"table:zipcodes,alias:zipcode"`

	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	SubDistrictID uuid.UUID  `bun:"sub_district_id,type:uuid,notnull" json:"sub_district_id"`
	Name          string     `bun:"name,notnull" json:"name"`
	IsActive      bool       `bun:"is_active,notnull,default:true" json:"is_active"`
	CreatedAt     time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

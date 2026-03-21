package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MemberAddressEntity struct {
	bun.BaseModel `bun:"table:member_addresses,alias:member_address"`

	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	MemberID      uuid.UUID  `bun:"member_id,type:uuid,notnull" json:"member_id"`
	FirstName     string     `bun:"first_name,notnull" json:"first_name"`
	LastName      *string    `bun:"last_name" json:"last_name"`
	Phone         string     `bun:"phone,notnull" json:"phone"`
	No            *string    `bun:"no" json:"no"`
	Village       *string    `bun:"village" json:"village"`
	Street        *string    `bun:"street" json:"street"`
	ProvinceID    uuid.UUID  `bun:"province_id,type:uuid,notnull" json:"province_id"`
	DistrictID    uuid.UUID  `bun:"district_id,type:uuid,notnull" json:"district_id"`
	SubDistrictID uuid.UUID  `bun:"sub_district_id,type:uuid,notnull" json:"sub_district_id"`
	ZipcodeID     uuid.UUID  `bun:"zipcode_id,type:uuid,notnull" json:"zipcode_id"`
	CreatedAt     time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

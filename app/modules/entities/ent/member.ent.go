package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MemberRole string

const (
	MemberRoleCustomer MemberRole = "customer"
	MemberRoleAdmin    MemberRole = "admin"
)

type MemberEntity struct {
	bun.BaseModel `bun:"table:members,alias:member"`

	ID        uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	GenderID  uuid.UUID  `bun:"gender_id,type:uuid,notnull" json:"gender_id"`
	PrefixID  *uuid.UUID `bun:"prefix_id,type:uuid" json:"prefix_id"`
	Role      MemberRole `bun:"role,type:member_role,notnull,default:'customer'" json:"role"`
	FirstName string     `bun:"first_name,notnull" json:"first_name"`
	LastName  *string    `bun:"last_name" json:"last_name"`
	Phone     string     `bun:"phone,notnull" json:"phone"`
	CreatedAt time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

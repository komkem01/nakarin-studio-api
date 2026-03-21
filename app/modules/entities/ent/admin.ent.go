package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AdminEntity struct {
	bun.BaseModel `bun:"table:admins,alias:admin"`

	ID           uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	MemberID     *uuid.UUID `bun:"member_id,type:uuid" json:"member_id"`
	Username     string     `bun:"username,notnull,unique" json:"username"`
	PasswordHash string     `bun:"password_hash,notnull" json:"password_hash"`
	DisplayName  *string    `bun:"display_name" json:"display_name"`
	LastLoginAt  *time.Time `bun:"last_login_at" json:"last_login_at"`
	IsActive     bool       `bun:"is_active,notnull,default:true" json:"is_active"`
	CreatedAt    time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt    *time.Time `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

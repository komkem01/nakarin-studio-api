package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BookingStatusLogEntity struct {
	bun.BaseModel `bun:"table:booking_status_logs,alias:booking_status_log"`

	ID            uuid.UUID      `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	BookingID     uuid.UUID      `bun:"booking_id,type:uuid,notnull" json:"booking_id"`
	FromStatus    *BookingStatus `bun:"from_status,type:booking_status" json:"from_status"`
	ToStatus      BookingStatus  `bun:"to_status,type:booking_status,notnull" json:"to_status"`
	ChangedBy     *uuid.UUID     `bun:"changed_by,type:uuid" json:"changed_by"`
	ChangedByRole *MemberRole    `bun:"changed_by_role,type:member_role" json:"changed_by_role"`
	Reason        *string        `bun:"reason" json:"reason"`
	ChangedAt     time.Time      `bun:"changed_at,nullzero,notnull,default:current_timestamp" json:"changed_at"`
}

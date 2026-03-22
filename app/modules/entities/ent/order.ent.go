package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type OrderStatus string

const (
	OrderStatusNew       OrderStatus = "new"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusCompleted OrderStatus = "completed"
)

type OrderEntity struct {
	bun.BaseModel `bun:"table:orders,alias:order"`

	ID          uuid.UUID   `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	BookingID   uuid.UUID   `bun:"booking_id,type:uuid,notnull" json:"booking_id"`
	OrderNo     string      `bun:"order_no,notnull,unique" json:"order_no"`
	Status      OrderStatus `bun:"status,notnull,default:'new'" json:"status"`
	TotalAmount float64     `bun:"total_amount,type:numeric(12,2),notnull,default:0" json:"total_amount"`
	CreatedAt   time.Time   `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time   `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	DeletedAt   *time.Time  `bun:"deleted_at,soft_delete" json:"deleted_at"`
}

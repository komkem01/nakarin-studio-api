package admin

import (
	"time"

	"nakarin-studio/app/modules/net/httpx"

	"go.opentelemetry.io/otel/trace"
)

type Controller struct {
	tracer trace.Tracer
	svc    *Service
	cli    *httpx.Client
}

func newController(trace trace.Tracer, svc *Service) *Controller {
	return &Controller{tracer: trace, svc: svc, cli: httpx.NewClient()}
}

type IDUri struct {
	ID string `uri:"id" binding:"required"`
}

type CreateRequest struct {
	MemberID     *string    `json:"member_id"`
	Username     string     `json:"username" binding:"required"`
	PasswordHash string     `json:"password_hash" binding:"required"`
	DisplayName  *string    `json:"display_name"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	IsActive     *bool      `json:"is_active"`
}

type UpdateRequest struct {
	MemberID     *string    `json:"member_id"`
	Username     *string    `json:"username"`
	PasswordHash *string    `json:"password_hash"`
	DisplayName  *string    `json:"display_name"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	IsActive     *bool      `json:"is_active"`
}

type ListQuery struct {
	MemberID *string `form:"member_id"`
	Username *string `form:"username"`
	IsActive *bool   `form:"is_active"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

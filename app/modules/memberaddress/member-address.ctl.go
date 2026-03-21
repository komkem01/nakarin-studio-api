package memberaddress

import (
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
	MemberID      string  `json:"member_id" binding:"required"`
	FirstName     string  `json:"first_name" binding:"required"`
	LastName      *string `json:"last_name"`
	Phone         string  `json:"phone" binding:"required"`
	No            *string `json:"no"`
	Village       *string `json:"village"`
	Street        *string `json:"street"`
	ProvinceID    string  `json:"province_id" binding:"required"`
	DistrictID    string  `json:"district_id" binding:"required"`
	SubDistrictID string  `json:"sub_district_id" binding:"required"`
	ZipcodeID     string  `json:"zipcode_id" binding:"required"`
}

type UpdateRequest struct {
	MemberID      *string `json:"member_id"`
	FirstName     *string `json:"first_name"`
	LastName      *string `json:"last_name"`
	Phone         *string `json:"phone"`
	No            *string `json:"no"`
	Village       *string `json:"village"`
	Street        *string `json:"street"`
	ProvinceID    *string `json:"province_id"`
	DistrictID    *string `json:"district_id"`
	SubDistrictID *string `json:"sub_district_id"`
	ZipcodeID     *string `json:"zipcode_id"`
}

type ListQuery struct {
	MemberID      *string `form:"member_id"`
	ProvinceID    *string `form:"province_id"`
	DistrictID    *string `form:"district_id"`
	SubDistrictID *string `form:"sub_district_id"`
	ZipcodeID     *string `form:"zipcode_id"`
	Phone         *string `form:"phone"`
}

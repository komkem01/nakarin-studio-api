package booking

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
	BookingNo               string     `json:"booking_no" binding:"required"`
	Status                  *string    `json:"status"`
	Payment                 *string    `json:"payment"`
	CancelledReason         *string    `json:"cancelled_reason"`
	InternalNote            *string    `json:"internal_note"`
	TrackingAttemptCount    *int       `json:"tracking_attempt_count"`
	LastTrackingAt          *time.Time `json:"last_tracking_at"`
	DeliveryMemberAddressID *string    `json:"delivery_member_address_id"`
	DeliveryFirstName       *string    `json:"delivery_first_name"`
	DeliveryLastName        *string    `json:"delivery_last_name"`
	DeliveryPhone           *string    `json:"delivery_phone"`
	DeliveryNo              *string    `json:"delivery_no"`
	DeliveryVillage         *string    `json:"delivery_village"`
	DeliveryStreet          *string    `json:"delivery_street"`
	DeliveryProvinceID      *string    `json:"delivery_province_id"`
	DeliveryDistrictID      *string    `json:"delivery_district_id"`
	DeliverySubDistrictID   *string    `json:"delivery_sub_district_id"`
	DeliveryZipcodeID       *string    `json:"delivery_zipcode_id"`
	DeliveryNote            *string    `json:"delivery_note"`
}

type UpdateRequest struct {
	BookingNo               *string    `json:"booking_no"`
	Status                  *string    `json:"status"`
	Payment                 *string    `json:"payment"`
	CancelledReason         *string    `json:"cancelled_reason"`
	InternalNote            *string    `json:"internal_note"`
	TrackingAttemptCount    *int       `json:"tracking_attempt_count"`
	LastTrackingAt          *time.Time `json:"last_tracking_at"`
	DeliveryMemberAddressID *string    `json:"delivery_member_address_id"`
	DeliveryFirstName       *string    `json:"delivery_first_name"`
	DeliveryLastName        *string    `json:"delivery_last_name"`
	DeliveryPhone           *string    `json:"delivery_phone"`
	DeliveryNo              *string    `json:"delivery_no"`
	DeliveryVillage         *string    `json:"delivery_village"`
	DeliveryStreet          *string    `json:"delivery_street"`
	DeliveryProvinceID      *string    `json:"delivery_province_id"`
	DeliveryDistrictID      *string    `json:"delivery_district_id"`
	DeliverySubDistrictID   *string    `json:"delivery_sub_district_id"`
	DeliveryZipcodeID       *string    `json:"delivery_zipcode_id"`
	DeliveryNote            *string    `json:"delivery_note"`
}

type ListQuery struct {
	Status        *string    `form:"status"`
	Payment       *string    `form:"payment"`
	BookingNo     *string    `form:"booking_no"`
	Phone         *string    `form:"phone"`
	CreatedAtFrom *time.Time `form:"created_at_from"`
	CreatedAtTo   *time.Time `form:"created_at_to"`
}

type TrackingQuery struct {
	BookingNo string `form:"booking_no" binding:"required"`
	Phone     string `form:"phone" binding:"required"`
}

type TransitionStatusRequest struct {
	Status *string `json:"status" binding:"required"`
	Reason *string `json:"reason"`
}

type ConvertToOrderRequest struct {
	Reason *string `json:"reason"`
}

type ConvertToOrderResponse struct {
	OrderID string `json:"order_id"`
	OrderNo string `json:"order_no"`
}

type AggregateBookingDetailRequest struct {
	FirstName string  `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name"`
	Phone     string  `json:"phone" binding:"required"`
}

type AggregateBookingItemRequest struct {
	ProductID          string  `json:"product_id" binding:"required"`
	ProductName        string  `json:"product_name" binding:"required"`
	UnitPriceAtBooking float64 `json:"unit_price_at_booking" binding:"required"`
	Quantity           int     `json:"quantity" binding:"required"`
	LineTotal          float64 `json:"line_total" binding:"required"`
	Note               *string `json:"note"`
	SortOrder          *int    `json:"sort_order"`
}

type AggregatePaymentRequest struct {
	Channel       *string    `json:"channel"`
	Amount        float64    `json:"amount" binding:"required"`
	DepositAmount float64    `json:"deposit_amount" binding:"required"`
	Status        *string    `json:"status"`
	ProofURL      *string    `json:"proof_url"`
	Note          *string    `json:"note"`
	PaidAt        *time.Time `json:"paid_at"`
}

type AggregateCreateRequest struct {
	Booking CreateRequest                 `json:"booking" binding:"required"`
	Detail  AggregateBookingDetailRequest `json:"detail" binding:"required"`
	Items   []AggregateBookingItemRequest `json:"items" binding:"required"`
	Payment *AggregatePaymentRequest      `json:"payment"`
}

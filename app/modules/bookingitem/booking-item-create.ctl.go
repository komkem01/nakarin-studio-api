package bookingitem

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Create(ctx *gin.Context) {
	var req CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	sortOrder := 0
	if req.SortOrder != nil {
		sortOrder = *req.SortOrder
	}

	_, err := c.svc.Create(ctx.Request.Context(), req.BookingID, req.ProductID, req.ProductName, req.UnitPriceAtBooking, req.Quantity, req.LineTotal, req.Note, sortOrder)
	if err != nil {
		base.InternalServerError(ctx, "booking-item-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

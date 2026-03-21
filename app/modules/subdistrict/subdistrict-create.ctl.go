package subdistrict

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

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	_, err := c.svc.Create(ctx.Request.Context(), req.DistrictID, req.Name, isActive)
	if err != nil {
		base.InternalServerError(ctx, "sub-district-create-failed", nil)
		return
	}

	base.Success(ctx, nil, "success")
}

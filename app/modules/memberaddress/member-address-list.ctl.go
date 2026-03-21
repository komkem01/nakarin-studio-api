package memberaddress

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) List(ctx *gin.Context) {
	var req ListQuery
	if err := ctx.ShouldBindQuery(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	items, err := c.svc.List(ctx.Request.Context(), req.MemberID, req.ProvinceID, req.DistrictID, req.SubDistrictID, req.ZipcodeID, req.Phone)
	if err != nil {
		base.InternalServerError(ctx, "member-address-list-failed", nil)
		return
	}

	base.Success(ctx, items, "success")
}

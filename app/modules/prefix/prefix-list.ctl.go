package prefix

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

type ListQuery struct {
	GenderID *string `form:"gender_id"`
	IsActive *bool   `form:"is_active"`
}

func (c *Controller) List(ctx *gin.Context) {
	var req ListQuery
	if err := ctx.ShouldBindQuery(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	items, err := c.svc.List(ctx.Request.Context(), req.GenderID, req.IsActive)
	if err != nil {
		base.InternalServerError(ctx, "prefix-list-failed", nil)
		return
	}

	base.Success(ctx, items, "success")
}

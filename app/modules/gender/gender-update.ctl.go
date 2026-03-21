package gender

import (
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Name     *string `json:"name"`
	IsActive *bool   `json:"is_active"`
}

type IDUri struct {
	ID string `uri:"id" binding:"required"`
}

func (c *Controller) Update(ctx *gin.Context) {
	var uri IDUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, "invalid-id", nil)
		return
	}

	var req UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid-request", nil)
		return
	}

	item, err := c.svc.UpdateByID(ctx.Request.Context(), uri.ID, req.Name, req.IsActive)
	if err != nil {
		base.InternalServerError(ctx, "gender-update-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

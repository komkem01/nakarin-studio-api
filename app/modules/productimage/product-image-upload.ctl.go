package productimage

import (
	"errors"
	"strconv"

	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Upload(ctx *gin.Context) {
	productID := ctx.PostForm("product_id")
	if productID == "" {
		base.BadRequest(ctx, "product-id-required", nil)
		return
	}

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		base.BadRequest(ctx, "image-file-required", nil)
		return
	}

	var altText *string
	if v := ctx.PostForm("alt_text"); v != "" {
		altText = &v
	}

	sortOrder := 0
	if v := ctx.PostForm("sort_order"); v != "" {
		parsed, parseErr := strconv.Atoi(v)
		if parseErr != nil {
			base.BadRequest(ctx, "invalid-sort-order", nil)
			return
		}
		sortOrder = parsed
	}

	isActive := true
	if v := ctx.PostForm("is_active"); v != "" {
		parsed, parseErr := strconv.ParseBool(v)
		if parseErr != nil {
			base.BadRequest(ctx, "invalid-is-active", nil)
			return
		}
		isActive = parsed
	}

	item, err := c.svc.UploadAndCreate(ctx.Request.Context(), productID, fileHeader, altText, sortOrder, isActive)
	if err != nil {
		if errors.Is(err, ErrProductImageFileInvalid) {
			base.BadRequest(ctx, "invalid-image-file", nil)
			return
		}
		base.InternalServerError(ctx, "product-image-upload-failed", nil)
		return
	}

	base.Success(ctx, item, "success")
}

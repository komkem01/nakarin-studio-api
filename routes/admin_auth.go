package routes

import (
	"strings"

	"nakarin-studio/app/modules"
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func adminAuth(mod *modules.Modules) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		adminID := strings.TrimSpace(ctx.GetHeader("X-Admin-ID"))
		if adminID == "" {
			_ = base.Unauthorized(ctx, "admin-id-required", nil)
			ctx.Abort()
			return
		}

		admin, err := mod.Admin.Svc.InfoByID(ctx.Request.Context(), adminID)
		if err != nil || admin == nil || !admin.IsActive {
			_ = base.Unauthorized(ctx, "admin-auth-failed", nil)
			ctx.Abort()
			return
		}

		ctx.Set("admin_id", admin.ID.String())
		ctx.Next()
	}
}

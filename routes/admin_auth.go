package routes

import (
	"strings"

	"nakarin-studio/app/modules"
	"nakarin-studio/app/utils/base"

	"github.com/gin-gonic/gin"
)

func adminAuth(mod *modules.Modules) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := strings.TrimSpace(ctx.GetHeader("Authorization"))
		if authHeader == "" {
			_ = base.Unauthorized(ctx, "admin-bearer-required", nil)
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
			_ = base.Unauthorized(ctx, "admin-bearer-invalid", nil)
			ctx.Abort()
			return
		}

		claims, err := mod.Admin.Svc.ValidateAccessToken(strings.TrimSpace(parts[1]))
		if err != nil {
			_ = base.Unauthorized(ctx, "admin-auth-failed", nil)
			ctx.Abort()
			return
		}

		adminID := claims.AdminID
		if strings.TrimSpace(adminID) == "" {
			_ = base.Unauthorized(ctx, "admin-auth-failed", nil)
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

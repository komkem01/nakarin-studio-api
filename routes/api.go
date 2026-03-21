package routes

import (
	"fmt"
	"net/http"

	"nakarin-studio/app/modules"

	"github.com/gin-gonic/gin"
)

func WarpH(router *gin.RouterGroup, prefix string, handler http.Handler) {
	router.Any(fmt.Sprintf("%s/*w", prefix), gin.WrapH(http.StripPrefix(fmt.Sprintf("%s%s", router.BasePath(), prefix), handler)))
}

func api(r *gin.RouterGroup, mod *modules.Modules) {
	r.GET("/example/:id", mod.Example.Ctl.Get)
	r.GET("/example-http", mod.Example.Ctl.GetHttpReq)
	r.POST("/example", mod.Example.Ctl.Create)
}

func apiSystem(r *gin.RouterGroup, mod *modules.Modules) {
	system := r.Group("/system")
	{
		gender := system.Group("/genders")
		{
			gender.POST("", mod.Gender.Ctl.Create)
			gender.GET("", mod.Gender.Ctl.List)
			gender.GET("/:id", mod.Gender.Ctl.Info)
			gender.PATCH("/:id", mod.Gender.Ctl.Update)
			gender.DELETE("/:id", mod.Gender.Ctl.Delete)
		}
	}
}

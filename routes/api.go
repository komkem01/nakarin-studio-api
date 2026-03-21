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

		prefix := system.Group("/prefixes")
		{
			prefix.POST("", mod.Prefix.Ctl.Create)
			prefix.GET("", mod.Prefix.Ctl.List)
			prefix.GET("/:id", mod.Prefix.Ctl.Info)
			prefix.PATCH("/:id", mod.Prefix.Ctl.Update)
			prefix.DELETE("/:id", mod.Prefix.Ctl.Delete)
		}

		province := system.Group("/provinces")
		{
			province.POST("", mod.Province.Ctl.Create)
			province.GET("", mod.Province.Ctl.List)
			province.GET("/:id", mod.Province.Ctl.Info)
			province.PATCH("/:id", mod.Province.Ctl.Update)
			province.DELETE("/:id", mod.Province.Ctl.Delete)
		}

		district := system.Group("/districts")
		{
			district.POST("", mod.District.Ctl.Create)
			district.GET("", mod.District.Ctl.List)
			district.GET("/:id", mod.District.Ctl.Info)
			district.PATCH("/:id", mod.District.Ctl.Update)
			district.DELETE("/:id", mod.District.Ctl.Delete)
		}

		subDistrict := system.Group("/sub-districts")
		{
			subDistrict.POST("", mod.SubDistrict.Ctl.Create)
			subDistrict.GET("", mod.SubDistrict.Ctl.List)
			subDistrict.GET("/:id", mod.SubDistrict.Ctl.Info)
			subDistrict.PATCH("/:id", mod.SubDistrict.Ctl.Update)
			subDistrict.DELETE("/:id", mod.SubDistrict.Ctl.Delete)
		}

		zipcode := system.Group("/zipcodes")
		{
			zipcode.POST("", mod.Zipcode.Ctl.Create)
			zipcode.GET("", mod.Zipcode.Ctl.List)
			zipcode.GET("/:id", mod.Zipcode.Ctl.Info)
			zipcode.PATCH("/:id", mod.Zipcode.Ctl.Update)
			zipcode.DELETE("/:id", mod.Zipcode.Ctl.Delete)
		}

		booking := system.Group("/bookings")
		{
			booking.POST("", mod.Booking.Ctl.Create)
			booking.GET("", mod.Booking.Ctl.List)
			booking.GET("/:id", mod.Booking.Ctl.Info)
			booking.PATCH("/:id", mod.Booking.Ctl.Update)
			booking.DELETE("/:id", mod.Booking.Ctl.Delete)
		}

		bookingDetail := system.Group("/booking-details")
		{
			bookingDetail.POST("", mod.BookingDetail.Ctl.Create)
			bookingDetail.GET("", mod.BookingDetail.Ctl.List)
			bookingDetail.GET("/:id", mod.BookingDetail.Ctl.Info)
			bookingDetail.PATCH("/:id", mod.BookingDetail.Ctl.Update)
			bookingDetail.DELETE("/:id", mod.BookingDetail.Ctl.Delete)
		}

		member := system.Group("/members")
		{
			member.POST("", mod.Member.Ctl.Create)
			member.GET("", mod.Member.Ctl.List)
			member.GET("/:id", mod.Member.Ctl.Info)
			member.PATCH("/:id", mod.Member.Ctl.Update)
			member.DELETE("/:id", mod.Member.Ctl.Delete)
		}

		memberBooking := system.Group("/member-bookings")
		{
			memberBooking.POST("", mod.MemberBooking.Ctl.Create)
			memberBooking.GET("", mod.MemberBooking.Ctl.List)
			memberBooking.GET("/:id", mod.MemberBooking.Ctl.Info)
			memberBooking.PATCH("/:id", mod.MemberBooking.Ctl.Update)
			memberBooking.DELETE("/:id", mod.MemberBooking.Ctl.Delete)
		}

		memberAddress := system.Group("/member-addresses")
		{
			memberAddress.POST("", mod.MemberAddress.Ctl.Create)
			memberAddress.GET("", mod.MemberAddress.Ctl.List)
			memberAddress.GET("/:id", mod.MemberAddress.Ctl.Info)
			memberAddress.PATCH("/:id", mod.MemberAddress.Ctl.Update)
			memberAddress.DELETE("/:id", mod.MemberAddress.Ctl.Delete)
		}

		product := system.Group("/products")
		{
			product.POST("", mod.Product.Ctl.Create)
			product.GET("", mod.Product.Ctl.List)
			product.GET("/:id", mod.Product.Ctl.Info)
			product.PATCH("/:id", mod.Product.Ctl.Update)
			product.DELETE("/:id", mod.Product.Ctl.Delete)
		}

		productImage := system.Group("/product-images")
		{
			productImage.POST("", mod.ProductImage.Ctl.Create)
			productImage.GET("", mod.ProductImage.Ctl.List)
			productImage.GET("/:id", mod.ProductImage.Ctl.Info)
			productImage.PATCH("/:id", mod.ProductImage.Ctl.Update)
			productImage.DELETE("/:id", mod.ProductImage.Ctl.Delete)
		}
	}
}

package routes

import (
	"registrasi-tamu/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/guests", controllers.CreateGuest)
	r.GET("/guests", controllers.GetGuests)
	r.GET("/guests/today", controllers.GetGuestsToday)
	r.GET("/guests/csv", controllers.ExportCSV)
	r.GET("/guests/pdf", controllers.ExportPDF)

	r.POST("/login", controllers.Login)
}

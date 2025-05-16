package routes

import (
	"registrasi-tamu/controllers"
	"registrasi-tamu/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.POST("/guests", controllers.CreateGuest)
        auth.GET("/guests", controllers.GetGuests)
        auth.GET("/guests/today", controllers.GetGuestsToday)
        auth.GET("/guests/csv", controllers.ExportCSV)
        auth.GET("/guests/pdf", controllers.ExportPDF)
    }
}

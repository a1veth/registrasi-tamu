package main

import (
	
	"registrasi-tamu/routes"
	"registrasi-tamu/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(":8080")
}

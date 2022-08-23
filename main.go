package main

import (
	"api/router"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(cors.Default())

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "I am healthy! 🚀",
		})
	})

	router.SetupRoutes(app)

	app.Run(":3050")
}

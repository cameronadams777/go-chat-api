package router

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	// Group all endpoints under the /api path
	api := app.Group("/api")

	api.GET("/ws", controllers.HandleWS)
}

package routes

import (
	"boolean-as-a-service/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter Configure routes, returns the router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/", controllers.CreateBoolean)
	r.GET("/:id", controllers.GetBoolean)
	r.PATCH("/:id", controllers.UpdateBoolean)
	r.DELETE("/:id", controllers.DeleteBoolean)

	return r
}

package routes

import(
	controller "01.JWT/controllers"
	"01.JWT/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incoming Routes *gin.Engine){
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("users", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}
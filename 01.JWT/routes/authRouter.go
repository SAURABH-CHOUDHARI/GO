package routes

import(
	controller "01.JWT/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())

}
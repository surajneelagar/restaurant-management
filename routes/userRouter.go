package routes

import (
	"restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/user", controllers.GetUsers())
	incomingRoutes.GET("/user/:user_id", controllers.GetUser())
	incomingRoutes.POST("/signup", controllers.SignUp())
	incomingRoutes.PATCH("/login", controllers.Login())
}

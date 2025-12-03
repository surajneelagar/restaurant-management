package routes

import (
	"restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("menu", controllers.GetMenu())
	incomingRoutes.GET("menu/:menu_id", controllers.GetMenu())
	incomingRoutes.POST("menu", controllers.CreateMenu())
	incomingRoutes.PATCH("menu/:menu_id", controllers.UpdateMenu())
	incomingRoutes.DELETE("menu/:menu_id", controllers.DeleteMenu())
}

package routes

import (
	"restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("invoices", controllers.GetFoods())
	incomingRoutes.GET("invoices/:invoice_id", controllers.GetFood())
	incomingRoutes.POST("invoices", controllers.CreateFood())
	incomingRoutes.PATCH("invoices/:invoice_id", controllers.UpdateFood())
}

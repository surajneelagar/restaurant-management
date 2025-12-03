package main

import (
	"os"
	"restaurent-management/middleware"
	"restaurent-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	post := os.Getenv("PORT")

	if post == "" {
		post = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TableRoutes(router)
	routes.NoteRoutes(router)

}

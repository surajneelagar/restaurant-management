package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return fmt.Println("Authentication")
}

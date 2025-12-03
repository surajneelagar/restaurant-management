package controllers

import "github.com/gin-gonic/gin"

func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func CreateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func DeleteFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func round(num float64) int {
	return int(num + 0.5)
}

func toFixed(num float64, precision int) float64 {
	return float64(round(num*float64(precision))) / float64(precision)
}

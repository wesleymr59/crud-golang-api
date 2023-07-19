package routers

import (
	"crud-golang-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/healthy_check", controllers.HealthyCheck)
	r.GET("/orders", controllers.GetAllOrders)
	r.POST("/orders", controllers.CreateNewOrder)
	r.PUT("/orders/:id", controllers.UpdateNewOrder)
	r.Run(":5000")
}

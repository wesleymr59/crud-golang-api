package routers

import (
	"crud-golang-api/controllers"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
)

func HandleRequests() {
	// Crie um novo Tracer APM
	tracer, err := apm.NewTracer("crud-golang-api", "")
	if err != nil {
		panic(err)
	}
	defer tracer.Close()

	// Inicie o aplicativo Gin
	r := gin.Default()

	// Middleware APM Gin para rastreamento de solicitações HTTP
	r.Use(apmgin.Middleware(r, apmgin.WithTracer(tracer)))

	r.GET("/healthy_check", controllers.HealthyCheck)
	r.GET("/orders", controllers.GetAllOrders)
	r.POST("/orders", controllers.CreateNewOrder)
	r.PUT("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)

	r.Run(":8080")
}

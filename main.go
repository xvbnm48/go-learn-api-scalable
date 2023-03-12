package main

import (
	"github.com/xvbnm48/go-learn-api-scalable/app"
	"github.com/xvbnm48/go-learn-api-scalable/controller"
	"github.com/xvbnm48/go-learn-api-scalable/repository"
	"github.com/xvbnm48/go-learn-api-scalable/service"
)

func main() {
	db := app.NewDB()
	orderRepsitory := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepsitory)
	orderController := controller.NewOrderController(orderService)

	router := app.OrderRoute(orderController)

	router.Run()
}

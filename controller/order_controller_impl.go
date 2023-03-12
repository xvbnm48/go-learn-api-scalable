package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-learn-api-scalable/helper"
	"github.com/xvbnm48/go-learn-api-scalable/model/web"
	"github.com/xvbnm48/go-learn-api-scalable/model/web/request"
	"github.com/xvbnm48/go-learn-api-scalable/service"
	"strconv"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func (o OrderControllerImpl) Create(ctx *gin.Context) {
	//TODO implement me
	var request request.OrderCreate

	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	response := o.OrderService.Create(request)
	ctx.JSON(200, web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (o OrderControllerImpl) Update(ctx *gin.Context) {
	var request request.OrderUpdate
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&request)
	helper.PanicIfError(err)

	response := o.OrderService.Update(request, id)
	ctx.JSON(200, web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (o OrderControllerImpl) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := o.OrderService.Delete(id)
	helper.PanicIfError(err)

	ctx.JSON(200, web.Response{
		Code:   200,
		Status: "OK",
	})
}

func (o OrderControllerImpl) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	response, err := o.OrderService.FindById(id)
	helper.PanicIfError(err)
	ctx.JSON(200, web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (o OrderControllerImpl) FindAll(ctx *gin.Context) {
	response, err := o.OrderService.FindAll()
	helper.PanicIfError(err)
	ctx.JSON(200, web.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{OrderService: orderService}
}

package service

import (
	"github.com/xvbnm48/go-learn-api-scalable/model/domain"
	"github.com/xvbnm48/go-learn-api-scalable/model/web/request"
)

type OrderService interface {
	Create(request request.OrderCreate) error
	Update(request request.OrderUpdate, id int) error
	Delete(id int) error
	FindById(id int) (domain.Order, error)
	FindAll() ([]domain.Order, error)
}

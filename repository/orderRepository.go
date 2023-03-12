package repository

import "github.com/xvbnm48/go-learn-api-scalable/model/domain"

type OrderRepository interface {
	Create(order domain.Order) error
	Update(order domain.Order, id int) error
	Delete(id int) error
	FindById(id int) (domain.Order, error)
	FindAll() ([]domain.Order, error)
}

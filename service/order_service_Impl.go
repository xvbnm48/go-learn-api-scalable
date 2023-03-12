package service

import (
	"github.com/xvbnm48/go-learn-api-scalable/model/domain"
	"github.com/xvbnm48/go-learn-api-scalable/model/web/request"
	"github.com/xvbnm48/go-learn-api-scalable/repository"
	"time"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &OrderServiceImpl{OrderRepository: orderRepository}
}

func (service *OrderServiceImpl) Create(request request.OrderCreate) error {
	var items []domain.Item
	for _, item := range request.Items {
		items = append(items, domain.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	order := domain.Order{
		CustomerName: request.CustomerName,
		OrderAt:      time.Now(),
		Items:        items,
	}

	err := service.OrderRepository.Create(order)
	return err
}

func (service *OrderServiceImpl) Update(request request.OrderUpdate, id int) error {
	var items []domain.Item
	for _, item := range items {
		items = append(items, domain.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     uint(id),
		})
	}

	order := domain.Order{
		CustomerName: request.CustomerName,
		OrderAt:      time.Now(),
		Items:        items,
	}

	err := service.OrderRepository.Update(order, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *OrderServiceImpl) FindById(id int) (domain.Order, error) {
	order, err := service.OrderRepository.FindById(id)

	return order, err
}

func (service *OrderServiceImpl) FindAll() ([]domain.Order, error) {
	orders, err := service.OrderRepository.FindAll()

	return orders, err
}

func (service *OrderServiceImpl) Delete(id int) error {
	err := service.OrderRepository.Delete(id)
	return err
}

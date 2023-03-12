package repository

import (
	"github.com/xvbnm48/go-learn-api-scalable/model/domain"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{DB: db}
}

func (repository *OrderRepositoryImpl) Create(order domain.Order) error {
	err := repository.DB.Create(&order).Error
	return err
}

func (repository *OrderRepositoryImpl) Update(order domain.Order, id int) error {
	err := repository.DB.Model(&domain.Order{}).Where("id = ?", order.ID).Updates(order).Error
	if err != nil {
		return err
	}

	err = repository.DB.Where("order_id = ?", id).Delete(&domain.Item{}).Error
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		repository.DB.Create(&item)
	}

	return nil
}

func (repository *OrderRepositoryImpl) FindById(id int) (domain.Order, error) {
	var order domain.Order
	err := repository.DB.Model(&domain.Order{}).Preload("Items").Where("id = ?", id).First(&order).Error
	return order, err
}

func (repository *OrderRepositoryImpl) FindAll() ([]domain.Order, error) {
	var order []domain.Order

	err := repository.DB.Model(&domain.Order{}).Preload("Items").Find(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (repository *OrderRepositoryImpl) Delete(id int) error {
	err := repository.DB.Where("order_id = ?", id).Delete(&domain.Item{}).Error
	if err != nil {
		return err
	}

	err = repository.DB.Delete(&domain.Order{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

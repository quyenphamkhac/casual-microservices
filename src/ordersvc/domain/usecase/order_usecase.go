package usecase

import (
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/dto"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/model"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/repository"
)

type OrderUsecase interface {
	PlaceOrder(data *dto.PlaceOrderDto) (*model.Order, error)
	CancelOrder(data *dto.CancelOrderDto) (*model.Order, error)
}

type orderUsecaseImpl struct {
	repo repository.OrderRepository
}

func NewOrderUsecaseImpl(repo repository.OrderRepository) *orderUsecaseImpl {
	return &orderUsecaseImpl{
		repo: repo,
	}
}

func (u *orderUsecaseImpl) PlaceOrder(data *dto.PlaceOrderDto) (*model.Order, error) {
	insertData := &dto.InsertOrderDto{
		CustomerId: data.CustomerId,
		Total:      data.Total,
		Status:     "ORDER_PROCESSING_COMPLETED",
		OrderItems: data.OrderItems,
	}
	return u.repo.Insert(insertData)
}

func (u *orderUsecaseImpl) CancelOrder(data *dto.CancelOrderDto) (*model.Order, error) {
	return nil, nil
}

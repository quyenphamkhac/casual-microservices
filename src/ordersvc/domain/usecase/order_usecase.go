package usecase

import (
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/dto"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/event"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/model"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/repository"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/internal/mq"
)

type OrderUsecase interface {
	PlaceOrder(data *dto.PlaceOrderDto) (*model.Order, error)
	CancelOrder(data *dto.CancelOrderDto) (*model.Order, error)
	EmitProductsValidationEvent(data *event.ProductsValidationEvent) error
}

type orderUsecaseImpl struct {
	repo      repository.OrderRepository
	publisher mq.Publisher
}

func NewOrderUsecaseImpl(repo repository.OrderRepository, publisher mq.Publisher) *orderUsecaseImpl {
	return &orderUsecaseImpl{
		repo:      repo,
		publisher: publisher,
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

func (u *orderUsecaseImpl) EmitProductsValidationEvent(data *event.ProductsValidationEvent) error {
	_, err := u.publisher.Publish(data)
	return err
}

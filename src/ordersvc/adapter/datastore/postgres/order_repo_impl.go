package postgres

import (
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/dto"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/model"
)

type orderRepoImpl struct {
	pg *pg.DB
}

func NewOrderRepoImpl(pgDB *pg.DB) *orderRepoImpl {
	return &orderRepoImpl{
		pg: pgDB,
	}
}

func (r *orderRepoImpl) Find(search string) ([]model.Order, error) {
	var orders []model.Order
	err := r.pg.Model(&orders).Select()
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (r *orderRepoImpl) FindById(id interface{}) (*model.Order, error) {
	order := &model.Order{}
	r.pg.Model(order).WherePK()
	return nil, nil
}

func (r *orderRepoImpl) FindByIds(ids []interface{}) ([]model.Order, error) {
	return nil, nil
}

func (r *orderRepoImpl) Insert(data interface{}) (*model.Order, error) {
	orderData, ok := data.(*dto.InsertOrderDto)
	if !ok {
		return nil, errors.New("can't parse 'order payload' to value")
	}
	order := &model.Order{
		CustomerId: orderData.CustomerId,
		Total:      orderData.Total,
		Status:     orderData.Status,
		OrderItems: []*model.OrderItem{},
	}
	tx, err := r.pg.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Close()

	if _, err := r.pg.Model(order).Insert(); err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	for _, v := range orderData.OrderItems {
		orderItem := &model.OrderItem{
			OrderId:     order.Id,
			ProductId:   v.ProductId,
			Quantity:    v.Quantity,
			ProductName: v.ProductName,
			Price:       v.Price,
		}
		_, err := r.pg.Model(orderItem).Insert()
		order.OrderItems = append(order.OrderItems, orderItem)
		if err != nil {
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepoImpl) Update(id interface{}, data interface{}) (*model.Order, error) {
	return nil, nil
}

func (r *orderRepoImpl) Delete(id interface{}) (*model.Order, error) {
	return nil, nil
}

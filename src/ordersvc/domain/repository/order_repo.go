package repository

import (
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/model"
)

type OrderRepository interface {
	Find(search string) ([]model.Order, error)
	FindById(id interface{}) (*model.Order, error)
	FindByIds(ids []interface{}) ([]model.Order, error)
	Insert(data interface{}) (*model.Order, error)
	Update(id interface{}, data interface{}) (*model.Order, error)
	Delete(id interface{}) (*model.Order, error)
}

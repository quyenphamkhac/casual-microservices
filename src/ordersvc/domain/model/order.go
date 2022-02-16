package model

import "time"

type Order struct {
	Id         int          `json:"id"`
	CustomerId int          `json:"customer_id"`
	OrderItems []*OrderItem `json:"order_items" pg:"rel:has-many"`
	Total      int          `json:"total"`
	Status     string       `json:"status"`
	CreatedAt  time.Time    `json:"created_at" pg:"default:now()"`
	UpdatedAt  time.Time    `json:"updated_at" pg:"default:now()"`
}

type OrderItem struct {
	ProductId   int    `json:"product_id" pg:",pk"`
	OrderId     int    `json:"order_id,omitempty" pg:",pk"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

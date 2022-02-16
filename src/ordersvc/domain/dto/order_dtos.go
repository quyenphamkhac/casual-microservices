package dto

type InsertOrderDto struct {
	CustomerId int `json:"customer_id" binding:"required"`
	Total      int `json:"total" binding:"required"`
	Status     string
	OrderItems []*OrderItemDto
}

type InsertOrderItemDto struct {
	ProductId   int
	OrderId     int
	ProductName string
	Quantity    int
	Price       int
}

type UpdateOrderDto struct{}

type PlaceOrderDto struct {
	CustomerId int             `json:"customer_id" binding:"required"`
	Total      int             `json:"total" binding:"required"`
	OrderItems []*OrderItemDto `json:"order_items" binding:"min=1,max=20,required,dive"`
}

type OrderItemDto struct {
	ProductId   int    `json:"product_id" binding:"required"`
	ProductName string `json:"product_name" binding:"max=20,required"`
	Quantity    int    `json:"quantity" binding:"max=50,required"`
	Price       int    `json:"price" binding:"required"`
}

type CancelOrderDto struct{}

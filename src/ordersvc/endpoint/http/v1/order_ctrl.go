package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/dto"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/event"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/usecase"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/pkg/httperrors"
)

type orderCtrl struct {
	usecase usecase.OrderUsecase
}

func NewOrderCtrl(usecase usecase.OrderUsecase) *orderCtrl {
	return &orderCtrl{
		usecase: usecase,
	}
}

func (ctrl *orderCtrl) PlaceOrderEndpoint(c *gin.Context) {
	var data dto.PlaceOrderDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(httperrors.New(http.StatusBadRequest, err.Error()))
		return
	}
	order, err := ctrl.usecase.PlaceOrder(&data)
	if err != nil {
		c.Error(httperrors.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func (ctrl *orderCtrl) CancelOrderEndpoint(c *gin.Context) {
}

func (ctrl *orderCtrl) GetOrders(c *gin.Context) {
	err := ctrl.usecase.EmitProductsValidationEvent(&event.ProductsValidationEvent{
		Message: "Em la bup mang non",
	})
	if err != nil {
		c.Error(httperrors.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "OK"})
}

package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCtrl struct{}

func NewHealthCtrl() *healthCtrl {
	return &healthCtrl{}
}

func (h *healthCtrl) HealthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "OK"})
}

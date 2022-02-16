package transport

import "github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"

type mqTransport struct {
	cfg *config.Config
}

func NewMQTransport(cfg *config.Config) *mqTransport {
	return &mqTransport{
		cfg: cfg,
	}
}

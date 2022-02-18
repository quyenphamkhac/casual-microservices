package transport

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/adapter/datastore/postgres"
	rmqimpl "github.com/quyenphamkhac/casual-microservices/src/ordersvc/adapter/mq/rabbitmq"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/pkg/postgresql"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/pkg/rabbitmq"

	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/usecase"
	httpv1 "github.com/quyenphamkhac/casual-microservices/src/ordersvc/endpoint/http/v1"
	httpmdw "github.com/quyenphamkhac/casual-microservices/src/ordersvc/middleware/http"
)

type httpServer struct {
	cfg *config.Config
}

func NewHttpServer(cfg *config.Config) *httpServer {
	return &httpServer{
		cfg: cfg,
	}
}

func (s *httpServer) Run(addr string) {
	db := postgresql.NewPostgresqlDBConn(s.cfg)
	defer db.Close()
	err := postgresql.CreatePostgresqlDBSchema(db)
	if err != nil {
		panic(err)
	}

	rmq := rabbitmq.NewRabbitMQConn(&s.cfg.RabbitMQ)
	err = rmq.Dial()
	if err != nil {
		panic(err)
	}

	rmqConn := rmq.Conn()
	defer rmqConn.Close()

	ch, err := rmqConn.Channel()
	if err != nil {
		panic(err)
	}

	publisher := rmqimpl.NewPublisher(s.cfg, ch)

	r := gin.Default()
	r.Use(httpmdw.ErrorsMiddleware(gin.ErrorTypeAny))
	orderRepository := postgres.NewOrderRepoImpl(db)
	orderUsecase := usecase.NewOrderUsecaseImpl(orderRepository, publisher)

	healthCtrl := httpv1.NewHealthCtrl()
	orderCtrl := httpv1.NewOrderCtrl(orderUsecase)

	v1 := r.Group("/v1")
	{
		v1.GET("/health", healthCtrl.HealthEndpoint)
		v1.POST("/orders", orderCtrl.PlaceOrderEndpoint)
	}
	go func() {
		if err := r.Run(addr); err != nil {
			log.Println("run http server failed")
		}
	}()
}

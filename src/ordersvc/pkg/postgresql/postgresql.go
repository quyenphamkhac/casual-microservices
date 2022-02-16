package postgresql

import (
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/domain/model"
)

var initDBConnOnce sync.Once

func NewPostgresqlDBConn(cfg *config.Config) *pg.DB {
	var conn *pg.DB
	initDBConnOnce.Do(func() {
		conn = pg.Connect(&pg.Options{
			User:     cfg.PostgresDB.User,
			Password: cfg.PostgresDB.Password,
			Addr:     cfg.PostgresDB.Addr,
			Database: cfg.PostgresDB.Database,
		})
	})
	return conn
}

func CreatePostgresqlDBSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.Order)(nil),
		(*model.OrderItem)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

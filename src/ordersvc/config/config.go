package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgresDB
	RabbitMQ
}

type PostgresDB struct {
	Addr     string `mapstructure:"ORDERSVC_POSTGRESDB_ADDR"`
	Password string `mapstructure:"ORDERSVC_POSTGRESDB_PASSWORD"`
	User     string `mapstructure:"ORDERSVC_POSTGRESDB_USER"`
	Database string `mapstructure:"ORDERSVC_POSTGRESDB_DATABASE"`
}

type RabbitMQ struct {
	User     string `mapstructure:"ORDERSVC_RABBITMQ_USER"`
	Password string `mapstructure:"ORDERSVC_RABBITMQ_PASSWORD"`
	Host     string `mapstructure:"ORDERSVC_RABBITMQ_HOST"`
	Port     string `mapstructure:"ORDERSVC_RABBITMQ_PORT"`
}

func NewServiceConfig() (*Config, error) {
	v, err := loadConfig()
	v.SetDefault("rabbitmq.user", "")
	if err != nil {
		return nil, err
	}
	c, err := parseConfig(v)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func loadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetEnvPrefix("ordersvc")
	v.AutomaticEnv()
	return v, nil
}
func parseConfig(v *viper.Viper) (*Config, error) {
	c := &Config{
		PostgresDB: PostgresDB{
			Addr:     v.GetString("postgresdb_addr"),
			User:     v.GetString("postgresdb_user"),
			Database: v.GetString("postgresdb_database"),
			Password: v.GetString("postgresdb_password"),
		},
	}
	return c, nil
}

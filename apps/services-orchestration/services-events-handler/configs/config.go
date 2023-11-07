package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type conf struct {
	ServiceName               string `mapstructure:"SERVICE_NAME"`
	ContextEnv                string `mapstructure:"CONTEXT_ENV"`
	RabbitMQProtocol          string `mapstructure:"RABBITMQ_DEFAULT_PROTOCOL"`
	RabbitMQHost              string `mapstructure:"RABBITMQ_DEFAULT_HOST"`
	RabbitMQPort              string `mapstructure:"RABBITMQ_DEFAULT_PORT"`
	RabbitMQUser              string `mapstructure:"RABBITMQ_DEFAULT_USER"`
	RabbitMQPassword          string `mapstructure:"RABBITMQ_DEFAULT_PASS"`
	RabbitMQVhost             string `mapstructure:"RABBITMQ_DEFAULT_VHOST"`
	RabbitMQConsumerQueueName string `mapstructure:"RABBITMQ_CONSUMER_QUEUE_NAME"`
	RabbitMQConsumerName      string `mapstructure:"RABBITMQ_CONSUMER_NAME"`
	RabbitMQExchange          string `mapstructure:"RABBITMQ_DEFAULT_EXCHANGE"`
	RabbitMQExchangeType      string `mapstructure:"RABBITMQ_DEFAULT_EXCHANGE_TYPE"`
	RabbitMQDlxName           string `mapstructure:"RABBITMQ_DEFAULT_DLX_NAME"`
}

type Config *conf

func LoadConfig(path string, env string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("services-events-handler")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(fmt.Sprintf(".env.%s", env))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}

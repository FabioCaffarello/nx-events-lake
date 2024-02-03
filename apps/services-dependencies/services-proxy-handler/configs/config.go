package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type conf struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

type Config *conf

func LoadConfig(path string, env string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("services-proxy-handler")
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

package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment    string `mapstructure:"ENVIRONMENT"`
	HTTPServerURL  string `mapstructure:"HTTP_SERVER_URL"`
	HTTPBFFAddress string `mapstructure:"HTTP_BFF_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("bff")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

package viper

import (
	"testing"

	"github.com/spf13/viper"
)

type Config struct {
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
}

func TestViper(t *testing.T) {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()

	c := Config{}
	viper.Unmarshal(&c)
	t.Log(c)
}
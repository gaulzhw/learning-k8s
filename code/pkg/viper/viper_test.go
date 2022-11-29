package viper

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type Config struct {
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
}

var (
	yaml = []byte(`
host: 127.0.0.1
name: test
`)
)

func TestViper(t *testing.T) {
	// viper.SetConfigFile("config.yaml")
	// viper.ReadInConfig()

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	assert.NoError(t, err)

	c := &Config{}
	err = viper.Unmarshal(c)
	assert.NoError(t, err)

	t.Log(c)
}

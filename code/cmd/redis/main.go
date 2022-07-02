package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gaulzhw/learning-k8s/internal/config"
	"github.com/gaulzhw/learning-k8s/internal/tcp"
)

const (
	configFile = "redis.conf"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func main() {
	if fileExists(configFile) {
		config.SetupConfig(configFile)
	}

	err := tcp.ListenAndServeWithSignal(
		&tcp.Config{
			Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
		},
		tcp.NewEchoHandler(),
	)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/gaulzhw/learning-k8s/pkg/restful"
)

func main() {
	restful.StartContainer()
	restful.StartServer()
}

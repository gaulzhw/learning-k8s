package main

import (
	"github.com/gaulzhw/learning-k8s/sample/pkg/restful"
)

func main() {
	restful.StartContainer()
	restful.StartServer()
}

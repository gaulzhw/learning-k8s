package main

import (
	"os"

	"k8s.io/component-base/cli"
)

func main() {
	command := NewRootCommand()
	code := cli.Run(command)
	os.Exit(code)
}

package main

import (
	"os"

	"k8s.io/component-base/cli"

	"github.com/gaulzhw/learning-k8s/pkg/cobra"
)

func main() {
	command := cobra.NewRootCommand()
	code := cli.Run(command)
	os.Exit(code)
}

package main

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
	"k8s.io/component-base/cli"
)

func TestCobraCommand(t *testing.T) {
	cmd := &cobra.Command{
		Use:          "cobra",
		Long:         "Test kubernetes-base command.",
		SilenceUsage: true,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			// silence client-go warnings.
			rest.SetDefaultWarningHandler(rest.NoWarnings{})
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			t.Log("run...")
			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				t.Log(arg)
			}
			return nil
		},
	}

	code := cli.Run(cmd)
	assert.Equal(t, code, 0)
}

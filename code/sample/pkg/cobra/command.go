package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
)

func NewRootCommand() *cobra.Command {
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
			fmt.Println("run...")
			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				fmt.Println(arg)
			}
			return nil
		},
	}

	return cmd
}

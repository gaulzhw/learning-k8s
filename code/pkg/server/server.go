package server

import (
	"io"

	"github.com/spf13/cobra"
	"k8s.io/apiextensions-apiserver/pkg/apiserver"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	genericapiserver "k8s.io/apiserver/pkg/server"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
)

// APIServerOptions contains state for master/api server
type APIServerOptions struct {
	SharedInformerFactory informers.SharedInformerFactory
	StdOut                io.Writer
	StdErr                io.Writer
}

// NewAPIServerOptions returns a new APIServerOptions
func NewAPIServerOptions(out, errOut io.Writer) *APIServerOptions {
	o := &APIServerOptions{
		StdOut: out,
		StdErr: errOut,
	}
	return o
}

// NewCommandStartAPIServer provides a CLI handler for 'start master' command
// with a default APIServerOptions.
func NewCommandStartAPIServer(defaults *APIServerOptions, stopCh <-chan struct{}) *cobra.Command {
	o := *defaults
	cmd := &cobra.Command{
		Short: "Launch API server",
		Long:  "Launch API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	utilfeature.DefaultMutableFeatureGate.AddFlag(flags)

	return cmd
}

// Validate validates APIServerOptions
func (o APIServerOptions) Validate(args []string) error {
	errors := []error{}
	return utilerrors.NewAggregate(errors)
}

// Complete fills in fields required to have valid data
func (o *APIServerOptions) Complete() error {
	return nil
}

// Config returns config for the api server given APIServerOptions
func (o *APIServerOptions) Config() (*apiserver.Config, error) {
	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)

	config := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig:   apiserver.ExtraConfig{},
	}
	return config, nil
}

// RunServer starts a new APIServer given APIServerOptions
func (o APIServerOptions) RunServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New(genericapiserver.NewEmptyDelegate())
	if err != nil {
		return err
	}

	server.GenericAPIServer.AddPostStartHookOrDie("start-server-informers", func(context genericapiserver.PostStartHookContext) error {
		config.GenericConfig.SharedInformerFactory.Start(context.StopCh)
		o.SharedInformerFactory.Start(context.StopCh)
		return nil
	})

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}

package app

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/jwcesign/federation-metrics-adapter/cmd/app/options"
)

func NewFederationMetricsAdaptorCommand(ctx context.Context) *cobra.Command {
	opts := options.NewOptions()

	cmd := &cobra.Command{
		Short: "Launch federation-metrics-adaptor",
		Long:  "Launch federation-metrics-adaptor",
		RunE: func(c *cobra.Command, args []string) error {
			if err := runCommand(opts, ctx); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}

func runCommand(o *options.Options, ctx context.Context) error {

	serverConfig, err := o.ServerConfig()
	if err != nil {
		return err
	}

	server, err := serverConfig.Complete()
	if err != nil {
		return err
	}

	return server.RunUntil(ctx.Done())
}

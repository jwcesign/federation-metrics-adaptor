package main

import (
	pkgserver "k8s.io/apiserver/pkg/server"

	"github.com/jwcesign/federation-metrics-adapter/cmd/app"
)

func main() {
	ctx := pkgserver.SetupSignalContext()
	cmd := app.NewFederationMetricsAdaptorCommand(ctx)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

package server

import (
	"fmt"
	"time"

	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/jwcesign/federation-metrics-adapter/pkg/provider"
	"github.com/jwcesign/federation-metrics-adapter/pkg/provider/karmada"
)

var MetricsProvider = map[string]provider.StartMetricsProviderFunc{
	"karmada": karmada.NewKarmadaMetricsProvider,
}

type Config struct {
	Apiserver     *genericapiserver.Config
	ScrapeTimeout time.Duration

	Provider string

	MetricsProvider provider.FederationMetricsProvider
}

func (c *Config) Complete() (*server, error) {
	genericServer, err := c.Apiserver.Complete(nil).New("metrics-server", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}
	genericServer.Handler.NonGoRestfulMux.HandleFunc("/metrics", http_handler)

	if c.Provider == "karmada" {
		c.MetricsProvider = karmada.NewKarmadaMetricsProvider()
	}
	if c.MetricsProvider == nil {
		return nil, fmt.Errorf("metrics provider not supported")
	}

	return &server{
		GenericAPIServer: genericServer,
	}, nil
}

package options

import (
	"fmt"
	"net"
	"time"

	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"sigs.k8s.io/metrics-server/pkg/api"

	"github.com/jwcesign/federation-metrics-adapter/pkg/server"
	"k8s.io/component-base/version"
)

type Options struct {
	SecureServing  *genericoptions.SecureServingOptionsWithLoopback
	Authentication *genericoptions.DelegatingAuthenticationOptions
	Authorization  *genericoptions.DelegatingAuthorizationOptions
	Audit          *genericoptions.AuditOptions
	Features       *genericoptions.FeatureOptions

	Provider string
}

func (o *Options) Validate() error {
	return nil
}

func NewOptions() *Options {
	return &Options{
		SecureServing:  genericoptions.NewSecureServingOptions().WithLoopback(),
		Authentication: genericoptions.NewDelegatingAuthenticationOptions(),
		Authorization:  genericoptions.NewDelegatingAuthorizationOptions(),
		Features:       genericoptions.NewFeatureOptions(),
		Audit:          genericoptions.NewAuditOptions(),
		Provider:       "karmada",
	}
}

func (o Options) ServerConfig() (*server.Config, error) {
	apiserver, err := o.ApiserverConfig()
	if err != nil {
		return nil, err
	}
	return &server.Config{
		Apiserver:     apiserver,
		ScrapeTimeout: 30 * time.Second,
		Provider:      o.Provider,
	}, nil
}

func (o Options) ApiserverConfig() (*genericapiserver.Config, error) {
	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewConfig(api.Codecs)
	if err := o.SecureServing.ApplyTo(&serverConfig.SecureServing, &serverConfig.LoopbackClientConfig); err != nil {
		return nil, err
	}

	if err := o.Authentication.ApplyTo(&serverConfig.Authentication, serverConfig.SecureServing, nil); err != nil {
		return nil, err
	}
	if err := o.Authorization.ApplyTo(&serverConfig.Authorization); err != nil {
		return nil, err
	}

	if err := o.Audit.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	versionGet := version.Get()
	serverConfig.Version = &versionGet

	return serverConfig, nil
}

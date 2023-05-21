package provider

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

type StartMetricsProviderFunc func() FederationMetricsProvider

type FederationMetricsProvider interface {
	Init(ctx context.Context, client kubernetes.Interface)

	Run()

	GetMetricsForDeployment(deployment string)

	GeCustomMetricsForDeployment(deployment string)

	GetExternalMetricsForDeployment(deployment string)
}

package karmada

import (
	"context"

	"k8s.io/client-go/kubernetes"

	"github.com/jwcesign/federation-metrics-adapter/pkg/provider"
)

type KarmadaMetricsProvider struct {
}

func NewKarmadaMetricsProvider() provider.FederationMetricsProvider {
	return &KarmadaMetricsProvider{}
}

func (c *KarmadaMetricsProvider) Init(ctx context.Context, client kubernetes.Interface) {

}

func (c *KarmadaMetricsProvider) Run() {

}

func (c *KarmadaMetricsProvider) GetMetricsForDeployment(deployment string) {

}

func (c *KarmadaMetricsProvider) GeCustomMetricsForDeployment(deployment string) {

}

func (c *KarmadaMetricsProvider) GetExternalMetricsForDeployment(deployment string) {

}

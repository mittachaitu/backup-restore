package client

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Config returns a *rest.Config, using either the kubeconfig (if specified) or an in-cluster
// configuration.
func Config(kubeconfig, kubecontext string, qps float32, burst int) (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = kubeconfig
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: kubecontext}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	clientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error finding Kubernetes API server config in --kubeconfig, $KUBECONFIG, or in-cluster configuration")
	}
	if qps > 0.0 {
		clientConfig.QPS = qps
	}
	if burst > 0 {
		clientConfig.Burst = burst
	}
	return clientConfig, nil
}

package utils

import (
	"context"

	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/commands"
	"k8s.io/kops/pkg/kubeconfig"
)

func GetKubeConfigBuilder(name string, clientset simple.Clientset) (*kubeconfig.KubeconfigBuilder, error) {
	cluster, err := clientset.GetCluster(context.Background(), name)
	if err != nil {
		return nil, err
	}
	keyStore, err := clientset.KeyStore(cluster)
	if err != nil {
		return nil, err
	}
	secretStore, err := clientset.SecretStore(cluster)
	if err != nil {
		return nil, err
	}
	conf, err := kubeconfig.BuildKubecfg(cluster, keyStore, secretStore, &commands.CloudDiscoveryStatusStore{}, 0, "", false, "", false)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

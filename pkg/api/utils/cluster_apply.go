package utils

import (
	"context"

	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

func ClusterApply(clientset simple.Clientset, clusterName string) error {
	kc, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return err
	}
	apply := &cloudup.ApplyClusterCmd{
		Cluster:    kc,
		Clientset:  clientset,
		TargetName: cloudup.TargetDirect,
	}
	return apply.Run(context.Background())
}

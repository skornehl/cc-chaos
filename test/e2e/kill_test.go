package e2e

import (
	"testing"

	framework "github.com/operator-framework/operator-sdk/pkg/test"
	apis "github.com/skornehl/cc-chaos/pkg/apis"
	operator "github.com/skornehl/cc-chaos/pkg/apis/cache/v1alpha1"
)

func TestMain(m *testing.M) {
	framework.MainEntry(m)
}

func TestCCChaos(t *testing.T) {
	memcachedList := &operator.Chaosservice
	err := framework.AddToFrameworkScheme(apis.AddToScheme, memcachedList)
	if err != nil {
		t.Fatalf("failed to add custom resource scheme to framework: %v", err)
	}
	// run subtests
	// t.Run("memcached-group", func(t *testing.T) {
	// 	t.Run("Cluster", MemcachedCluster)
	// 	t.Run("Cluster2", MemcachedCluster)
	// })
}

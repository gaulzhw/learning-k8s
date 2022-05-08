package client

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func TestDiscoveryClient(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	assert.NoError(t, err)

	_, apiResourceLists, err := discoveryClient.ServerGroupsAndResources()
	assert.NoError(t, err)
	for _, apiResourceList := range apiResourceLists {
		groupVersionStr := apiResourceList.GroupVersion
		_, err = schema.ParseGroupVersion(groupVersionStr)
		assert.NoError(t, err)
		for _, singleResource := range apiResourceList.APIResources {
			t.Logf("%v\n", singleResource.Name)
		}
	}
}

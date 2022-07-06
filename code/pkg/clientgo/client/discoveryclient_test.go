package client

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

func TestRestMapping(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	restMapper, err := apiutil.NewDynamicRESTMapper(config)
	assert.NoError(t, err)

	mapping, err := restMapper.RESTMapping(schema.GroupKind{
		Group: "",
		Kind:  "Node",
	})
	assert.NoError(t, err)
	t.Log(mapping.GroupVersionKind)

	mapping, err = restMapper.RESTMapping(schema.GroupKind{
		Group: "",
		Kind:  "Node",
	})
	assert.NoError(t, err)
	t.Log(mapping.GroupVersionKind)
}

func TestMemCachedDiscoveryClient(t *testing.T) {
	client, err := newDiscoveryClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)

	cachedClient := memory.NewMemCacheClient(client)
	groups, err := cachedClient.ServerGroups()
	assert.NoError(t, err)
	for _, group := range groups.Groups {
		t.Logf("%+v", group)
	}
}

func TestCachedDiscoveryClient(t *testing.T) {
	client, err := newDiscoveryClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// discovery cache client
	cachedDiscoClient := NewMemCacheClient(client)
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedDiscoClient)

	mapping, err := restMapper.RESTMapping(corev1.SchemeGroupVersion.WithKind("Pod").GroupKind(), corev1.SchemeGroupVersion.Version)
	assert.NoError(t, err)
	t.Logf("%+v", mapping)
}

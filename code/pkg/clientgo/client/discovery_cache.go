package client

import (
	openapi_v2 "github.com/googleapis/gnostic/openapiv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	restclient "k8s.io/client-go/rest"
)

type MemCacheClient struct {
	client discovery.CachedDiscoveryInterface
}

var _ discovery.CachedDiscoveryInterface = &MemCacheClient{}

func NewMemCacheClient(delegate discovery.DiscoveryInterface) *MemCacheClient {
	return &MemCacheClient{
		client: memory.NewMemCacheClient(delegate),
	}
}

func (c *MemCacheClient) RESTClient() restclient.Interface {
	return c.client.RESTClient()
}

func (c *MemCacheClient) ServerGroups() (*metav1.APIGroupList, error) {
	return c.client.ServerGroups()
}

func (c *MemCacheClient) ServerResourcesForGroupVersion(groupVersion string) (*metav1.APIResourceList, error) {
	return c.client.ServerResourcesForGroupVersion(groupVersion)
}

func (c *MemCacheClient) ServerResources() ([]*metav1.APIResourceList, error) {
	return c.client.ServerResources()
}

func (c *MemCacheClient) ServerGroupsAndResources() ([]*metav1.APIGroup, []*metav1.APIResourceList, error) {
	return c.client.ServerGroupsAndResources()
}

func (c *MemCacheClient) ServerPreferredResources() ([]*metav1.APIResourceList, error) {
	return c.client.ServerPreferredResources()
}

func (c *MemCacheClient) ServerPreferredNamespacedResources() ([]*metav1.APIResourceList, error) {
	return c.client.ServerPreferredNamespacedResources()
}

func (c *MemCacheClient) ServerVersion() (*version.Info, error) {
	return c.client.ServerVersion()
}

func (c *MemCacheClient) OpenAPISchema() (*openapi_v2.Document, error) {
	return c.client.OpenAPISchema()
}

// Fresh is supposed to tell the caller whether or not to retry if the cache
// fails to find something (false = retry, true = no need to retry).
func (c *MemCacheClient) Fresh() bool {
	return true
}

// Invalidate enforces that no cached data that is older than the current time
// is used.
func (c *MemCacheClient) Invalidate() {}

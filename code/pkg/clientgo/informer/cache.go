package informer

import (
	"context"

	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	indexName = "labelIndex"
)

type labelIndexCache struct {
	cache.Cache
}

func newCache(config *rest.Config, opts cache.Options) (cache.Cache, error) {
	cache, err := cache.New(config, opts)
	if err != nil {
		return nil, err
	}

	indexCache := &labelIndexCache{
		Cache: cache,
	}

	return indexCache, err
}

func (c *labelIndexCache) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	// convert labelSelector to fieldSelector
	for _, opt := range opts {
		switch instance := opt.(type) {
		case client.MatchingLabels:
			for k, v := range instance {
				opts = append(opts, client.MatchingFields{indexName: indexKeyFun(k, v)})
			}
		case *client.ListOptions:
			labelSelector := instance.LabelSelector
			if labelSelector == nil || labelSelector.Empty() {
				break
			}
			requirements, _ := labelSelector.Requirements()
			for _, requirement := range requirements {
				if requirement.Operator() != selection.Equals {
					continue
				}
				for v := range requirement.Values() {
					opts = append(opts, client.MatchingFields{indexName: indexKeyFun(requirement.Key(), v)})
				}
			}
		}
	}

	return c.Cache.List(ctx, list, opts...)
}

func indexKeyFun(k, v string) string {
	return k + "/" + v
}

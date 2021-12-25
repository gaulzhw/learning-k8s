package clientgo

import (
	"fmt"
	"path/filepath"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// https://xinchen.blog.csdn.net/article/details/113800054
func InitDiscoveryClient() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return
	}

	apiGroup, apiResourceLists, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		return
	}

	fmt.Printf("APIGroup: \n\n %v\n\n\n\n", apiGroup)

	for _, apiResourceList := range apiResourceLists {
		groupVersionStr := apiResourceList.GroupVersion
		gv, err := schema.ParseGroupVersion(groupVersionStr)
		if err != nil {
			return
		}

		fmt.Println("*******************************************************")
		fmt.Printf("GV string [%v]\nGV struct [%#v]\nresources: \n\n", groupVersionStr, gv)

		for _, singleResource := range apiResourceList.APIResources {
			fmt.Printf("%v\n", singleResource.Name)
		}
	}
}

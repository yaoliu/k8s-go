//__author__ = "YaoYao"
//Date: 2020/8/13
package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("https://kubernetes.docker.internal:6443", *kubeconfig)
	if err != nil {
		panic(err)
	}
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	//_, APIResourceList, err := discoveryClient.ServerGroupsAndResources()
	//if err != nil {
	//	panic(err)
	//}
	//for _, list := range APIResourceList {
	//	gv, err := schema.ParseGroupVersion(list.GroupVersion)
	//	if err != nil {
	//		panic(err)
	//	}
	//	for _, resource := range list.APIResources {
	//		fmt.Printf("name: %v,group:%v, version:%v\n", resource.Name, gv.Group, gv.Version)
	//	}
	//}
	preferredResources, err := discoveryClient.ServerPreferredResources()
	deletableResources := discovery.FilteredBy(discovery.SupportsAllVerbs{Verbs: []string{"delete", "list", "watch"}}, preferredResources)
	for _, rl := range deletableResources {
		gv, err := schema.ParseGroupVersion(rl.GroupVersion)
		if err == nil {
			for i := range rl.APIResources{
				fmt.Println(gv.Group, gv.Version,rl.APIResources[i].Name)
			}
		}
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return ""
}

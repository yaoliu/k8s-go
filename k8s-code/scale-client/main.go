//__author__ = "YaoYao"
//Date: 2020/8/14
package main

import (
	"flag"
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

	_, err := clientcmd.BuildConfigFromFlags("https://kubernetes.docker.internal:6443", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}


}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return ""
}

//__author__ = "YaoYao"
//Date: 2020/7/4
package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("https://apiserver.demo:6443", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		for _, pod := range pods.Items {
			fmt.Printf("Pod:%s Namespace:%s\n", pod.Name, pod.Namespace)
		}
		time.Sleep(10 * time.Second)
		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		//_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		//if errors.IsNotFound(err) {
		//	fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		//} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		//	fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		//} else if err != nil {
		//	panic(err.Error())
		//} else {
		//	fmt.Printf("Found example-xxxxx pod in default namespace\n")
		//}
		//
	}
	//deployment, err := clientset.AppsV1beta1().Deployments("default").Get(context.Background(), *deploymentName, metav1.GetOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//if errors.IsNotFound(err) {
	//	fmt.Printf("Deployment not found\n")
	//} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	//	fmt.Printf("Error getting deployment%v\n", statusError.ErrStatus.Message)
	//} else if err != nil {
	//	panic(err.Error())
	//} else {
	//	fmt.Printf("Found deployment\n")
	//	name := deployment.GetName()
	//	fmt.Println("name ->", name)
	//	containers := &deployment.Spec.Template.Spec.Containers
	//	found := false
	//	for i := range *containers {
	//		c := *containers
	//		if c[i].Name == *appName {
	//			found = true
	//			fmt.Println("Old image ->", c[i].Image)
	//			fmt.Println("New image ->", *imageName)
	//			c[i].Image = *imageName
	//		}
	//	}
	//	if found == false {
	//		fmt.Println("The application container not exist in the deployment pods.")
	//		os.Exit(0)
	//	}
	//_, err := clientset.AppsV1beta1().Deployments("default").Update(context.Background(), deployment, metav1.UpdateOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return ""
}

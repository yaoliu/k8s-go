//__author__ = "YaoYao"
//Date: 2020/11/22
package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	// kubeconfig文件默认存在用户家目录的.kube/config中
	if home := os.Getenv("HOME"); home != "" {
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
		panic(err)
	}
	hpas, err := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, hpa := range hpas.Items {
		for _, metric := range hpa.Spec.Metrics {
			fmt.Println(hpa.Name, hpa.Spec.ScaleTargetRef.Name, metric.Type)
			fmt.Println(metric.Resource.Name, metric.Resource.Target.Type, metric.Resource.Target.AverageUtilization)
			fmt.Printf("%d\n", metric.Resource.Target.AverageUtilization)
		}
	}
	podList, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	requests, err := calculatePodRequests(podList.Items, v1.ResourceCPU)
	for k, v := range requests {
		fmt.Println(k, v)
	}
}

func calculatePodRequests(pods []v1.Pod, resource v1.ResourceName) (map[string]int64, error) {
	requests := make(map[string]int64, len(pods))
	// 遍历所有pod及pod.spec.Containers 将匹配搭配的resource对应的值进行累加
	for _, pod := range pods {
		podSum := int64(0)
		for _, container := range pod.Spec.Containers {
			if containerRequest, ok := container.Resources.Requests[resource]; ok {
				podSum += containerRequest.MilliValue()
			} else {
				return nil, fmt.Errorf("missing request for %s", resource)
			}
		}
		requests[pod.Name] = podSum
	}
	return requests, nil
}

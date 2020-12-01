//__author__ = "YaoYao"
//Date: 2020/11/19
package main

import (
	"context"
	"flag"
	"fmt"

	//"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	resourceclient "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
	"os"
	"path/filepath"
	"time"
)

type PodMetric struct {
	Timestamp time.Time
	Window    time.Duration
	Value     int64
}

type PodMetricsInfo map[string]PodMetric

func main() {
	var kubeconfig *string
	if home := os.Getenv("HOME"); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	clientConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientConfig = restclient.AddUserAgent(clientConfig, "Mertic-Client")
	rsclint := resourceclient.NewForConfigOrDie(clientConfig)
	metrics, err := rsclint.PodMetricses("default").List(context.TODO(), metav1.ListOptions{LabelSelector: labels.Everything().String()})
	if err != nil {
		panic(err.Error())
	}
	res := make(PodMetricsInfo, len(metrics.Items))

	for _, m := range metrics.Items {
		podSum := int64(0)
		missing := len(m.Containers) == 0
		for _, c := range m.Containers {
			resValue, found := c.Usage[v1.ResourceName("cpu")]
			if !found {
				missing = true
				break // containers loop
			}
			fmt.Println(m.Name,resValue.MilliValue())
			podSum += resValue.MilliValue()
		}

		if !missing {
			res[m.Name] = PodMetric{
				Timestamp: m.Timestamp.Time,
				Window:    m.Window.Duration,
				Value:     int64(podSum),
			}
		}
	}
	for name, pm := range res {
		fmt.Println(name, pm.Value,pm.Window,pm.Timestamp)
	}
}

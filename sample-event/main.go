//__author__ = "YaoYao"
//Date: 2020/8/30
package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var kubeconfig *string
	if home := os.Getenv("HOME"); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//创建事件消费者/事件广播器
	eventBroadcaster := record.NewBroadcaster()
	fmt.Println(eventBroadcaster)
	//设置事件写入日志
	eventBroadcaster.StartLogging(klog.Infof)
	//设置事件上传到Kubernetes API Server
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: clientset.CoreV1().Events("")})
	//创建事件生产者
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: "sample-event"})
	newPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx-events",
			Namespace: "default",
		},
		TypeMeta: metav1.TypeMeta{
			Kind: "Pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx-01",
					Image: "nginx",
				},
			},
		},
	}
	pod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), newPod, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	//添加一个事件信息
	recorder.Event(pod, corev1.EventTypeNormal, "Start", "create pod event")
	time.Sleep(10 * time.Second)
}
//__author__ = "YaoYao"
//Date: 2020/9/6
package main

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"time"
)

var (
	maxIntervalInSeconds = 600
)

type Event struct {
	Name    string
	Type    string
	Message string
}

func main() {
	maxInterval := time.Duration(maxIntervalInSeconds) * time.Second
	fmt.Println(maxInterval)
	newEvent := &Event{
		Name:    "create",
		Type:    "error",
		Message: "create error event",
	}
	oldEvent := &Event{
		Name:    "create",
		Type:    "error1",
		Message: "create error event",
	}
	newData, _ := json.Marshal(newEvent)
	oldData, _ := json.Marshal(oldEvent)
	patch, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, newEvent)
	fmt.Println(string(patch), err)
}

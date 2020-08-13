//__author__ = "YaoYao"
//Date: 2020/8/12
package workqueue

import (
	"container/heap"
	"fmt"
	"testing"
	"time"
)

func TestIntHeap(t *testing.T) {
	p := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(p)
	for len(*p) > 0 {
		fmt.Printf("%d ", heap.Pop(p))
	}
}

func TestPriorityTimeQueue(t *testing.T) {
	p := &PriorityTimeQueue{}
	*p = append(*p, &PriorityTimeType{index: 0, value: 1, readyAt: time.Now()})
	*p = append(*p, &PriorityTimeType{index: 0, value: 2, readyAt: time.Now().Add(1000)})
	*p = append(*p, &PriorityTimeType{index: 0, value: 2, readyAt: time.Now().Add(2000)})
	*p = append(*p, &PriorityTimeType{index: 0, value: 2, readyAt: time.Now().Add(3000)})
	heap.Init(p)
	for len(*p) > 0 {
		x := heap.Pop(p).(*PriorityTimeType)
		fmt.Println(x.readyAt.Second())
	}
}

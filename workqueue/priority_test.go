//__author__ = "YaoYao"
//Date: 2020/8/12
package workqueue

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriority(t *testing.T) {
	p := &PrioritySimple{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(p)
	for len(*p) > 0 {
		fmt.Printf("%d ", heap.Pop(p))
	}
}

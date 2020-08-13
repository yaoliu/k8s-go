//__author__ = "YaoYao"
//Date: 2020/8/12
package workqueue

import (
	"container/heap"
	"time"
)

type PriorityType struct {
	index    int
	priority int
	data     string
}

type PriorityTimeType struct {
	readyAt time.Time
	value   interface{}
	index   int
}

type PrioritySimple []int                  //堆

func (p PrioritySimple) Len() int {
	panic("implement me")
}

func (p PrioritySimple) Less(i, j int) bool {
	panic("implement me")
}

func (p PrioritySimple) Swap(i, j int) {
	panic("implement me")
}

func (p PrioritySimple) Push(x interface{}) {
	panic("implement me")
}

func (p PrioritySimple) Pop() interface{} {
	panic("implement me")
}

type PriorityQueue []*PriorityType         //优先级队列
type PriorityTimeQueue []*PriorityTimeType //根据时间的优先级队列

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].priority < p[j].priority
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index, p[j].index = i, j
}

func (p *PriorityQueue) Push(x interface{}) {
	n := len(*p)
	item := x.(*PriorityType)
	item.index = n
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	item.index = -1
	return item
}

func (p *PriorityQueue) update(item *PriorityType, data string, priority int) {
	item.data = data
	item.priority = priority
	heap.Fix(p, item.index)
}

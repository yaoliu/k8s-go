//__author__ = "YaoYao"
//Date: 2020/8/12
package workqueue

import (
	"container/heap"
	"time"
)

type PriorityType struct { //优先级队列 实现container/heap的Interface
	index    int    //在堆中堆位置
	priority int    //优先级
	value    string //数据
}

type PriorityTimeType struct {
	readyAt time.Time
	value   interface{}
	index   int
}

type IntHeap []int //堆

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	item := x.(int)
	*h = append(*h, item)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[1:]
	return x
}

type PriorityQueue []*PriorityType         //优先级队列
type PriorityTimeQueue []*PriorityTimeType //根据时间的优先级队列

func (p PriorityTimeQueue) Len() int {
	return len(p)
}

func (p PriorityTimeQueue) Less(i, j int) bool {
	return p[i].readyAt.Before(p[j].readyAt)
}

func (p PriorityTimeQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j

}

func (p *PriorityTimeQueue) Push(x interface{}) {
	n := len(*p)
	item := x.(PriorityTimeType)
	item.index = n
	*p = append(*p, &item)
}

func (p *PriorityTimeQueue) Pop() interface{} {
	old := *p
	n := len(*p)
	item := old[n-1]
	item.index = -1
	*p = old[0:(n - 1)]
	return item
}

//func (p *PriorityTimeQueue) Peek() *PriorityTimeType {
//	return p[0]
//}

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool { //小根堆
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

func (p *PriorityQueue) update(item *PriorityType, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(p, item.index)
}

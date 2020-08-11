//__author__ = "YaoYao"
//Date: 2020/8/11
package workqueue

import "sync"

//需要掌握cond heap set 等知识
type Interface interface {
	Add(item interface{})                   //给队列添加元素 item是接口类型 所以可以是任何类型的元素
	Len() int                               //返回当前队列的长度
	Get() (item interface{}, shutdown bool) //获取队列元素及当前队列状态
	Done(item interface{})                  //处理完的元素需要告知队列 该队列需要标记元素已经被处理
	ShutDown()                              //关闭队列 并且通知所有wait状态的G
	ShuttingDown() bool                     //查询当前队列关闭状态
}

type t interface{}
type empty struct{}
type set map[t]struct{}

func (s set) has(item t) bool {
	_, exists := s[item]
	return exists
}

func (s set) insert(item t) {
	if !s.has(item) {
		s[item] = empty{}
	}
}

func (s set) delete(item t) {
	delete(s, item)
}

type Type struct {
	queue        []t
	dirty        set        //集合 用于去重 底层使用map实现
	processing   set        //处理中的
	cond         *sync.Cond //用于加锁 唤醒G 等待
	shuttingDown bool
}

func NewQueue() *Type {
	return &Type{
		dirty:        set{},
		processing:   set{},
		cond:         sync.NewCond(&sync.Mutex{}),
		shuttingDown: false,
	}
}

func (q *Type) Add(item t) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if q.shuttingDown { //检查是否关闭
		return
	}
	if q.dirty.has(item) { //检查dirtry是否已经存在 如果已经存在 不能添加
		return
	}
	q.dirty.insert(item) //插入到dirty

	if q.processing.has(item) { //检查是否正常处理
		return
	}

	q.queue = append(q.queue, item)
	q.cond.Signal()
}

func (q *Type) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.queue)
}

func (q *Type) Get() (item interface{}, shutdown bool) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.queue) == 0 && !q.shuttingDown {
		q.cond.Wait()
	}

	if len(q.queue) == 0 {
		return nil, true
	}

	item, q.queue = q.queue[0], q.queue[1:]

	q.processing.insert(item)
	q.dirty.delete(item)
	return item, false
}

func (q *Type) ShutDown() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.shuttingDown = true
	q.cond.Broadcast() //告诉所有等待的消费者
}

func (q *Type) ShuttingDown() bool {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return q.shuttingDown
}

func (q *Type) Done(item t) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.processing.delete(item)
	if q.dirty.has(item) {
		q.queue = append(q.queue, item)
		q.cond.Signal() //添加完 唤醒消费者
	}
}

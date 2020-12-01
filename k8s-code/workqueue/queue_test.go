//__author__ = "YaoYao"
//Date: 2020/8/11
package workqueue

import (
	"sync"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	q := NewQueue()
	const producers = 50
	producerWG := &sync.WaitGroup{}
	producerWG.Add(producers)
	for i := 0; i < producers; i++ {
		go func(i int) {
			defer producerWG.Done()
			for j := 0; j < 50; j++ {
				q.Add(j)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}

	const consumers = 10
	consumerWG := &sync.WaitGroup{}
	consumerWG.Add(consumers)
	for i := 0; i < consumers; i++ {
		go func(i int) {
			defer consumerWG.Done()
			for {
				item, quit := q.Get()
				if item == "added after shutdown!" {
					t.Errorf("Got an item added after shutdown.")
				}
				if quit {
					return
				}
				t.Logf("Worker %v: begin processing %v", i, item)
				time.Sleep(3 * time.Millisecond)
				t.Logf("Worker %v: done processing %v", i, item)
				q.Done(item)
			}
		}(i)
	}
	producerWG.Wait()
	q.ShutDown()
	q.Add("added after shutdown!")
	consumerWG.Wait()
}

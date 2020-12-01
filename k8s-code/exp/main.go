//__author__ = "YaoYao"
//Date: 2020/10/7
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type ControlleeExpectations struct {
	add       int64
	del       int64
	key       string
	timestamp time.Time
}

func (exp *ControlleeExpectations) Add(add, del int64) {
	atomic.AddInt64(&exp.add, add)
	atomic.AddInt64(&exp.del, del)
}

func main() {
	exp := &ControlleeExpectations{}
	add := 1
	del := 0
	exp.Add(int64(-add), int64(-del))
	fmt.Println(exp.add)
}

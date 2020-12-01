//__author__ = "YaoYao"
//Date: 2020/8/11
package workqueue

import (
	"sync"
	"time"
)

type DelayingInterface interface {
	Interface
	AddAfter(item interface{},duration time.Duration)
}

type delayingType struct {
	Interface
	stopCh chan struct{}
	stopOnce sync.Once
}


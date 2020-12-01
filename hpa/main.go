//__author__ = "YaoYao"
//Date: 2020/12/1
package main

import "fmt"

type Notice interface {
	Notice(message string) bool
}

type workChat struct{}

func (notice *workChat) Notice(message string) bool {
	fmt.Println("message", message)
	return true
}

type dingTalk struct{}

func (notice *dingTalk) Notice(message string) bool {
	fmt.Println("message", message)
	return true
}

func NewNotice() Notice {
	return &dingTalk{}
}

type Monitor struct {
	notice Notice
}

func (m *Monitor) Alert() {
	m.notice.Notice("hello world")
}

func NewMonitor(notice Notice) *Monitor {
	return &Monitor{notice: notice}
}

func main() {
	notice := NewNotice()
	monitor := NewMonitor(notice)
	monitor.Alert()
}

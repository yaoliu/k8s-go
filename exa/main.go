//__author__ = "YaoYao"
//Date: 2020/8/16
package main

import (
	"fmt"
	"strings"
)

func main() {
	//now := time.Now()
	//fmt.Println(now.Format("2006-01-02 15:04:05"))
	////mm, _ := time.ParseDuration("10s")
	//nowafter := now.Add(time.Second * 10)
	//fmt.Println(nowafter.Format("2006-01-02 15:04:05"))
	//fmt.Println(now.Before(nowafter))
	//fmt.Println(nowafter.After(now))
	s1 := "18000000041"
	fmt.Println(s1)
	fmt.Println(s1[3:7])
	s2 := strings.Replace(s1, s1[3:7], "****", 0)
	fmt.Println(s2)
}

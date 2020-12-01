//__author__ = "YaoYao"
//Date: 2020/10/4
package main

import (
	"fmt"
	"github.com/robfig/cron"
	"strconv"
	"time"
)

func main() {
	ss := "*/1 * * * *"
	t := time.Now()
	sc, err := cron.ParseStandard(ss)
	if err != nil {
		fmt.Println(err)
	}
	next := sc.Next(t)
	fmt.Println(next)
	annotation := "1"
	generation, err := strconv.ParseInt(annotation, 10, 64)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(generation)
}

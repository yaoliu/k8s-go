//__author__ = "YaoYao"
//Date: 2020/9/1
package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
)

func main() {
	service := new(restful.WebService)
	service.Route(service.GET("/hello").To(hello))
	restful.Add(service)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func hello(request *restful.Request, response *restful.Response) {
	_, _ = io.WriteString(response, "hello world")
}

package main

import (
	"fmt"
	"net/http"
)

type HelloHandlerStruct struct {
	content string
}

// 实现了handler 接口
func (handler *HelloHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.content)
}

func main() {
	//handler 直接实现的serveHTTP
	http.Handle("/", &HelloHandlerStruct{
		content: "",
	})

	http.ListenAndServe(":8081", nil)

}

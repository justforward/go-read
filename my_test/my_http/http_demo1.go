package main

import (
	"fmt"
	"net/http"
)


// handler func(ResponseWriter, *Request)
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

//实现最简单的调用
func main() {
	//这里使用的是 DefaultServeMux
	// 在根路由上 注册了一个helloHandler
	http.HandleFunc("/",HelloHandler)
	//然后启动服务监听本机的8080 端口
	http.ListenAndServe(":8080",nil)
}

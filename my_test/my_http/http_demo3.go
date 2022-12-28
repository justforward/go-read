package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, request *http.Request) {
	// 设置一个 301 重定向
	//实现的跳转的www.2345.com 无效 必须是https://www.2345.com
	//w.Header().Set("Location", "https://www.2345.com/")
	//w.WriteHeader(301)
	fmt.Println("----")
}

type DemoHandler struct {
}

func (demo DemoHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Println("//////")
	str := "////"
	w.Write([]byte(str))
}

type Greeting struct {
	val string
}
type JSONHandler struct {
}

func (jsonHandler JSONHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	greeting := Greeting{
		"欢迎访问学院君个人网站?",
	}
	message, _ := json.Marshal(greeting)
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func main() {
	// 自定义的ServeMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorld)
	mux.Handle("/hello", &DemoHandler{})
	// 传入一个handler 里面创建 server
	http.ListenAndServe(":8080", mux)
}

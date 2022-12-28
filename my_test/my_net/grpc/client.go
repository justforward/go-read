package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {

	listen, err := net.Listen("tpc", ":1234")
	if err != nil {

	}
	// 构建 链接的
	client := make(chan *rpc.Client)
	for {
		// 从服务端得到
		conn, err := listen.Accept()
		if err != nil {

		}
		// 一直等待着接受传入的数据
		// 当我们想要向 Channel 发送数据时，就需要使用 ch <- i 语句
		client <- rpc.NewClient(conn)
	}
	// 在管道得到数据之后，调用对应的函数
	doClientWork(client)
}

func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()

	// 使用上诉构建的链接进行调用对应的函数
	var reply string
	err := client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

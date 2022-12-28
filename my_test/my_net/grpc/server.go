package main

import (
	"net"
	"net/rpc"
	"time"
)

func main() {

	rpc.Register(new(HelloService))

	for {
		// 主动链接 客户端
		// 反向 RPC 的内网服务将不再主动提供 TCP 监听服务，而是首先主动连接到对方的 TCP 服务器。然后基于每个建立的 TCP 连接向对方提供 RPC 服务。
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
}


type HelloService struct{}

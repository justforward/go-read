package main

func main() {
	// 其中 rpc.Register 函数调用会将对象类型中所有满足 RPC 规则的对象方法注册为 RPC 函数，
	//// 所有注册的方法会放在 “HelloService” 服务空间之下
	//rpc.RegisterName("helloService", HelloService{})
	//
	//// 构建Tcp的唯一链接
	//listen, err := net.Listen("tcp", ":1234")
	//if err != nil {
	//
	//}
	//
	//// conn io.ReadWriteCloser
	//// 返回的conn是io.ReadWithCloser 类型
	//conn, err := listen.Accept()
	//if err != nil {
	//
	//}
	//
	////  rpc.ServeConn 函数在该 TCP 连接上为对方提供 RPC 服务。
	//rpc.ServeConn(conn)

}

type HelloService struct{}

// Hello 其中 Hello 方法必须满足
// Go 语言的 RPC 规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个 error 类型，同时必须是公开的方法
func (p *HelloService) Hello(request string, replay *string) error {
	*replay = ".." + request
	return nil
}

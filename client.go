package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"os"
	"thrift_example/gen-go/timerpc"
)

func main() {
	// 先建立和服务器的连接的socket，再通过socket建立Transport
	socket, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		fmt.Println("Error opening socket:", err)
		os.Exit(1)
	}
	transport := thrift.NewTFramedTransport(socket)

	// 创建二进制协议
	protocol := thrift.NewTBinaryProtocolTransport(transport)

	// 打开Transport，与服务器进行连接
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+"localhost"+":"+"9090", err)
		os.Exit(1)
	}
	defer transport.Close()

	// 接口需要context，以便在长操作时用户可以取消RPC调用
	ctx := context.Background()

	// 使用Mytime服务
	MyTimeProtocol := thrift.NewTMultiplexedProtocol(protocol, "mytime")

	// 创建代理客户端，使用TMultiplexedProtocol访问对应的服务
	c := thrift.NewTStandardClient(MyTimeProtocol, MyTimeProtocol)

	client := timerpc.NewTimeServeClient(c)
	res, err := client.GetCurrtentTime(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)

	// 使用其他服务
	// 步骤与上面的相同
	// 使用Mytime服务
	MyTime2Protocol := thrift.NewTMultiplexedProtocol(protocol, "mytime2")
	c2 := thrift.NewTStandardClient(MyTime2Protocol, MyTime2Protocol)
	client2 := timerpc.NewTimeServeClient(c2)
	res, err = client2.GetCurrtentTime(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}

package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"os"
	"thrift_example/gen-go/timerpc"
	"time"
)

type MyTime struct{}

func (s *MyTime) GetCurrtentTime(_ context.Context) (r int32, err error) {
	t := int32(time.Now().Unix())
	fmt.Printf("come on:%d\n", t)
	return t, nil
}

type MyTime2 struct{}

func (s *MyTime2) GetCurrtentTime(_ context.Context) (r int32, err error) {
	t := int32(time.Now().Unix())
	fmt.Printf("come on2:%d\n", t)
	return t, nil
}

func main() {
	// 创建服务器
	serverTransport, err := thrift.NewTServerSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	// 创建二进制协议
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	// 创建Processor，用一个端口处理多个服务
	multiProcessor := thrift.NewTMultiplexedProcessor()
	MyTimeProcessor := timerpc.NewTimeServeProcessor(new(MyTime))
	MyTime2Processor := timerpc.NewTimeServeProcessor(new(MyTime2))

	// 给每个service起一个名字
	multiProcessor.RegisterProcessor("mytime", MyTimeProcessor)
	multiProcessor.RegisterProcessor("mytime2", MyTime2Processor)

	server := thrift.NewTSimpleServer4(multiProcessor, serverTransport, transportFactory, protocolFactory)

	fmt.Println("start")
	if err := server.Serve(); err != nil {
		panic(err)
	}

	// 退出时停止服务器
	defer server.Stop()
}

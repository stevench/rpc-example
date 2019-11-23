package main

import (
    "net"
    "fmt"
    "google.golang.org/grpc"
    pt "grpc/myproto"
    "context"
)

const (
    post = "127.0.0.1:18881"
)

// 对象要和proto内定义的服务一样
type server struct{}
// 实现RPC SayHello 接口
func (this *server) SayHello(ctx context.Context, in *pt.HelloRequest)(*pt.HelloReplay, error) {
    return &pt.HelloReplay{Message: "hello" + in.Name}, nil
}

// 实现 RPC GetHelloMsg 接口
func (this *server) GetHelloMsg(ctx context.Context, in *pt.HelloRequest)(*pt.HelloMessage, error) {
    return &pt.HelloMessage{Msg: "this is from server HaHa!"}, nil
}

func main() {
    // 监听网路
    ln, err := net.Listen("tcp", post)
    if err != nil {
        fmt.Println("网络异常", err)
    }

    // 创建一个grpc的句柄
    srv := grpc.NewServer()
    pt.RegisterHelloServerServer(srv, &server{})

    // 监听grpc服务
    err = srv.Serve(ln)
    if err != nil {
        fmt.Println("网络启动异常", err)
    }
}



package main

import (
    "google.golang.org/grpc"
    pt "grpc/myproto"
    "fmt"
    "context"
)

const (
    post = "127.0.0.1:18881"
)

func main() {
    conn, err := grpc.Dial(post, grpc.WithInsecure())
    if err != nil {
        fmt.Println("连接服务器失败", err)
    }

    defer conn.Close()

    // 获得grpc句柄
    c := pt.NewHelloServerClient(conn)

    // 远程调用 SayHello 接口
    r1, err := c.SayHello(context.Background(), &pt.HelloRequest{Name: "panda"})
    if err != nil {
        fmt.Println("cloud not get Hello server ..", err)
	return
    }

    fmt.Println("HelloServer resp: ", r1.Message)

    // 远程调用 GetHelloMsg 接口
    r2, err := c.GetHelloMsg(context.Background(), &pt.HelloRequest{Name: "panda"})
    if err != nil {
        fmt.Println("cloud not get hello msg ..", err)
	return
    }

    fmt.Println("HelloServer resp: ", r2.Msg)
}

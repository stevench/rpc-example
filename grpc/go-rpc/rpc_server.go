package main

import (
    "net/http"
    "net/rpc"
    "net"
    "github.com/astaxie/beego"
    "io"
)

type Panda int

func (this *Panda) Getinfo(argType int, replyType *int) error {
    beego.Info("RPC Request!")
    beego.Info(argType)
    *replyType = 1 + argType
    return nil
}

func main() {
    http.HandleFunc("/panda", pandatext)
    pd := new(Panda)
    // 注册服务
    // Register在默认服务中注册并公布 接收服务 pd对象 的方法
    rpc.Register(pd)
    rpc.HandleHTTP()
    // 建立网络监听
    ln, err := net.Listen("tcp", "127.0.0.1:10086")
    if err != nil {
        beego.Info("网络连接失败")
    }
    beego.Info("正在监听10086")
    http.Serve(ln, nil)
}

// 用来实现网页的web函数
func pandatext(w http.ResponseWriter, r *http.Request) {
   beego.Info("HTTP Request!")
    io.WriteString(w, "panda")
}

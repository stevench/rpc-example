package main

import (
    "net/rpc"
    "github.com/astaxie/beego"
)

func main() {
    cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
    if err != nil {
        beego.Info("网络连接失败")
    }

    var val int
    err = cli.Call("Panda.Getinfo", 123, &val)
    if err != nil {
        beego.Info("打Call失败")
    }

    beego.Info("返回结果：", val)
}

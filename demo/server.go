package main

import (
    "fmt"
    "net"

    "github.com/hollson/keep"
    "github.com/sirupsen/logrus"
)

var server *keep.Server

func main() {
    // 创建默认服务或带有初始项的服务
    server = keep.Default()
    // server = keep.NewServer(keep.WithKeepAlive(10, 3, 5))

    // 添加中间件
    server.Use(func(ctx *keep.Context) {
        fmt.Printf("Request ACT =\033[0;34m %d \033[0m\n", ctx.Query.Action)
    })

    // 注册路由
    server.AddRoute(101, LoginHandler)

    // 添加事件 - 客户端连接时触发
    server.Accepted = func(conn net.Conn) {
        logrus.Infof("用户「%s」已上线️.\n", conn.RemoteAddr().String())
    }

    // 添加事件 - 客户端中断时触发
    server.Aborted = func(c net.Conn, m string) {
        logrus.Infof("用户「%s」已掉线️,%s .\n", c.RemoteAddr().String(), m)
    }

    // 启动服务
    server.Run(":4321")
}

// 处理程序 - 用户登录
func LoginHandler(c *keep.Context) {
    // 请求参数
    fmt.Printf("「%v」\n", c.Query)

    // 响应消息
    var msg = &keep.Message{
        Action: 333,
        Code:   keep.Status(33),
        Data:   c.Query.Data,
    }

    // 给当前用户发送
    c.Self(msg)

    // 给某个用户发送响应(包括广播)
    // c.Send(msg,c.SelfConn(),c.SelfConn())

    // 仅响应动作/指令(无需数据)
    // c.Action(1002, c.SelfConn())

    // 仅响应状态码(Action为0)
    // c.Code(3)
}

// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package keep

import (
    "fmt"
    "io"
    "net"
    "time"

    "github.com/hollson/keep/lib/keepalive"
    "github.com/sirupsen/logrus"
)

var (
    defaultPort uint = 9980 // 默认端口
)

// 服务引擎
type Server struct {
    router     Router                       // 路由表
    middleware []HandlerFunc                // 中间件
    keepAlive  keepAlive                    // Alive探测
    Accepted   func(c net.Conn)             // 客户端连接事件
    Aborted    func(c net.Conn, msg string) // 客户端异常中断
    Level      int                          // 日志级别
}

type keepAlive struct {
    idle     time.Duration // 「TCP探测」空闲等待时间
    count    int           // 「TCP探测」探测次数
    interval time.Duration // 「TCP探测」探测间隔
}

// 创建Server服务
func NewServer(opts ...Option) *Server {
    var server = &Server{
        router:     make(map[uint32][]HandlerFunc),
        middleware: make([]HandlerFunc, 0),
        keepAlive: keepAlive{
            idle:     10,
            count:    3,
            interval: 3,
        },
    }
    for _, f := range opts {
        f(server)
    }
    return server
}

// 添加路由
//  act：动作编号(0为系统保留)
//  fs：处理查询
func (s *Server) AddRoute(act uint32, fs ...HandlerFunc) {
    // if act < 101 {
    //     logrus.Errorf("action[%s] can not be used, 0~100 is system retention action", act)
    //     return
    // }
    if act > 0 {
        s.router[act] = append(s.router[act], fs...)
        return
    }
    logrus.Panic("action must be large then 0")
}

// 加载中间件
func (s *Server) Use(f HandlerFunc) {
    s.middleware = append(s.middleware, f)
}

// // 开启Zip压缩(默认zip使用zip压缩)
// func (s *server) SetCompress(f HandlerFunc) {
//     compress = true
// }

// 启动服务，如: s.Run(":9980"), 默认端口:9980
func (s *Server) Run(addr ...string) error {
    var _addr string
    switch len(addr) {
    case 0:
        _addr = fmt.Sprintf(":%d", defaultPort)
    case 1:
        _addr = addr[0]
    default:
        return fmt.Errorf("too many parameters")
    }

    tcpAddr, err := net.ResolveTCPAddr("tcp4", _addr)
    if err != nil {
        return err
    }

    // 监听器
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        return err
    }
    defer listener.Close()

    fmt.Printf(" ⚽️ Keep Server 已启动: %s\n", listener.Addr().String())

    for {
        // 创建客户端连接
        conn, err := listener.Accept()
        if err != nil {
            logrus.Errorf("accept err, %v", err)
            time.Sleep(time.Second)
            continue
        }

        // 触发连接事件
        if s.Accepted != nil {
            s.Accepted(conn)
        }

        go s.exchange(conn)
    }
}

// 处理会话
func (s *Server) exchange(conn net.Conn) {
    defer conn.Close()
    var ctx = newContext(conn)

    // Tcp客户端探测
    if err := keepalive.SetKeepAlive(conn, s.keepAlive.idle, s.keepAlive.count, s.keepAlive.interval); err != nil {
        logrus.Errorf("keepalive err, %v", err)
        return
    }

    go s.response(ctx)
    if err := s.request(ctx); err != nil {
        s.Aborted(conn, err.Error())
        return
    }
}

// 处理响应
func (s *Server) response(ctx *Context) {
    // 接收消息
    for req := range ctx.chanel {
        // 执行中间件
        for _, f := range s.middleware {
            f(ctx)
        }

        // 处理请求
        if fs, ok := s.router[req.Action]; ok {
            for _, f := range fs {
                f(ctx)
            }
            continue
        }
        ctx.Code(MSG_NOT_FOUND)
    }
}

// 处理请求
func (s *Server) request(ctx *Context) error {
    for {
        msg, err := UnPack(ctx.conn)
        if err != nil {
            if err == io.EOF {
                // 中断连接
                return err
            }
            // 解包错误，继续接收
            ctx.Code(MSG_ERR_PACKET)
            // 防止客户端自旋
            time.Sleep(time.Second)
            continue
        }

        ctx.Query = msg
        ctx.chanel <- msg
    }
}

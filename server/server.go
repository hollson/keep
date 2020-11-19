// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
    "fmt"
    "net"
    "time"

    "github.com/hollson/team/lib/tcpkeepalive"
    "github.com/sirupsen/logrus"
)

var _running = false

// 监听器
type TeamListener interface {
    // Route()
    Run()
    Close() error
}

type Group interface {
    AllotGroupHash() int // 分配心跳监听组的端口号
}

// TeamServer
type server struct {
    Listener          *net.TCPListener   // 监听地址
    Clients           map[string]*client // 客户端
    KeepAliveIdle     int                // 空闲多长时间开始探测(以秒为单位),默认为30s。
    KeepAliveCount    int                // 探测总次数,默认为4次。
    KeepAliveInterval int                // 探测时间间隔(以秒为单位),默认为5秒。
}

// 连接客户端
type client struct {
    uuid string
    conn *tcpkeepalive.Conn
}

//  创建TCP监听器
//  默认TcpAddr为： 0.0.0.0:54321
func NewServer(option ...Option) (*server, error) {
    tcpAddr, err := net.ResolveTCPAddr("tcp4", ":4321")
    if err != nil {
        return nil, fmt.Errorf("resolve tcp addr error,%v\n", err)
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        return nil, fmt.Errorf("create tcp listener error,%v\n", err)
    }
    // defer listener.Close()

    _server := &server{
        Listener:          listener,
        Clients:           make(map[string]*client),
        KeepAliveCount:    4,
        KeepAliveIdle:     30,
        KeepAliveInterval: 5,
    }

    for _, f := range option {
        if err := f(_server); err != nil {
            return nil, err
        }
    }
    return _server, nil
}

// 启动监听服务
func (s *server) Run() {
    logrus.Infoln("「Team Editor」TCP监听服务已启动：", s.Listener.Addr().String())
    for {
        // 创建客户端连接
        conn, err := s.Listener.AcceptTCP()
        if err != nil {
            logrus.Errorf("accept client err,%v", err)
            continue
        }

        kaConn, err := tcpkeepalive.EnableKeepAlive(conn)
        if err != nil {
            s.Offline(kaConn, 2)
            break
        }

        // todo
        kaConn.SetKeepAliveIdle(time.Duration(s.KeepAliveIdle) * time.Second)
        kaConn.SetKeepAliveCount(s.KeepAliveCount)
        kaConn.SetKeepAliveInterval(time.Duration(s.KeepAliveInterval) * time.Second)
        //
        // s.Clients[conn.RemoteAddr().String()] = &client{conn: kaConn}
        // logrus.Infof("%s 已上线,当前在线人数：%d", conn.RemoteAddr().String(), len(s.Clients))
        go s.messageHandler(kaConn)
    }
}

// 路由
func Route() {

}

// 关闭监听器(一般情况下只需中断应用)
func (s *server) Close() error {
    return s.Listener.Close()
}

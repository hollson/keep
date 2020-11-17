// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package team

import (
    "fmt"
    "io"
    "net"

    "github.com/sirupsen/logrus"
)

var _running = false

// 心跳
type server struct {
    Listener *net.TCPListener   // 监听地址
    Client   map[string]*client // 客户端
    // conn.SetKeepAlive(true)                   // 启用 TCP keepalive
    // conn.SetKeepAlivePeriod(time.Second * 30) //
    // conn.SetDeadline(time.Now().Add(time.Second * 10))
}

// 连接客户端
type client struct {
    conn *net.TCPConn
}

// 创建监听器
func NewServer(option ...Option) (*server, error) {
    tcpAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:54321")
    if err != nil {
        return nil, fmt.Errorf("resolve tcp addr error,%v\n", err)
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        return nil, fmt.Errorf("create tcp listener error,%v\n", err)
    }
    // defer listener.Close()

    _server := &server{
        Listener: listener,
        Client:   make(map[string]*client),
    }

    for _, f := range option {
        if err := f(_server); err != nil {
            return nil, err
        }
    }
    return _server, nil
}

func (s *server) Run() {
    logrus.Infoln("「Team Editor」TCP服务已启动：", s.Listener.Addr().String())
    for {
        // 接受TCP客户端
        conn, err := s.Listener.AcceptTCP()
        if err != nil {
            logrus.Errorf("accept client err,%v", err)
            continue
        }
        //
        // conn.SetKeepAlive(true)                   // 启用 TCP keepalive
        // conn.SetKeepAlivePeriod(time.Second * 30) //
        // conn.SetDeadline(time.Now().Add(time.Second * 10))
        s.Client[conn.RemoteAddr().String()] = &client{conn: conn}
        logrus.Infof("%s 已上线,当前在线人数：%d", conn.RemoteAddr().String(), len(s.Client))
        go s.messageHandler(conn)
    }
}

func (s *server) Close() {
    s.Listener.Close()
}

// 处理消息
func (s *server) messageHandler(conn *net.TCPConn) {
    // defer conn.Close() // 记得关闭连接
    // 循环接收用户数据
    for {
        data := make([]byte, 256)
        total, err := conn.Read(data)
        if err == io.EOF {
            logrus.Errorf("%s主动退出", conn.RemoteAddr().String())
            delete(s.Client, conn.RemoteAddr().String())
            break
        }

        s.broadCast(data[:total], conn)
    }
}

// 广播
func (s *server) broadCast(data []byte, conn *net.TCPConn) {
    for _, c := range s.Client {
        if c.conn.RemoteAddr().String() == conn.RemoteAddr().String() {
            continue
        }
        fmt.Printf("%s => %s\n", conn.RemoteAddr().String(), string(data))
        c.conn.Write(data) // todo 客户端下线
    }
}

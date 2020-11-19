// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package teamar

import (
    "context"
    "fmt"
    "net"

    "github.com/hollson/team/team"
    "github.com/sirupsen/logrus"
)

// 监听器
type TeamServer interface {
    Run()
    Close() error
}

// Teamar处理程序
type TeamarHandlerFunc func(ctx context.Context)

// team框架
type teamar struct {
    Listener          *net.TCPListener // 监听地址
    Team              team.Team        // 团队字典
    RemoteAddrsCache  team.RemoteAddrs // 缓存客户远程地址
    KeepAliveIdle     int              // 空闲多长时间开始探测(以秒为单位),默认为30s。
    KeepAliveCount    int              // 探测总次数,默认为4次。
    KeepAliveInterval int              // 探测时间间隔(以秒为单位),默认为5秒。
}

//  创建Team监听服务器
//  默认TcpAddr为： 0.0.0.0:54321
func NewTeamar(option ...Option) (*teamar, error) {
    tcpAddr, err := net.ResolveTCPAddr("tcp4", ":4321")
    if err != nil {
        return nil, fmt.Errorf("resolve tcp addr error,%v\n", err)
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        return nil, fmt.Errorf("create tcp listener error,%v\n", err)
    }
    // defer listener.Close()

    _server := &teamar{
        Listener:          listener,
        Team:              make(team.Team),
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
func (s *teamar) Run() {
    logrus.Infoln("「Team Editor」TCP监听服务已启动：", s.Listener.Addr().String())
    s.TcpAllot() // 客户长连接分配器
}

// 关闭监听器(一般情况下只需中断应用)
func (s *teamar) Close() error {
    return s.Listener.Close()
}

// 路由
func Route() {

}

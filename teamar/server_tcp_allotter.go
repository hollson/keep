// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package teamar

import (
    "time"

    "github.com/hollson/team/lib/tcpkeepalive"
    "github.com/sirupsen/logrus"
)

// 客户长连接分配器
func (s *teamar) TcpAllot() {
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
        // s.Team[conn.RemoteAddr().String()] = &client{conn: kaConn}
        // logrus.Infof("%s 已上线,当前在线人数：%d", conn.RemoteAddr().String(), len(s.Team))
        go s.TcpRoute(kaConn)
    }
}

// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package team

import (
    "fmt"
    "math/rand"
    "net"
    "testing"
    "time"

    "github.com/sirupsen/logrus"
)

// 服务端：监听心跳
func TestServerHeartMonitor(t *testing.T) {
    ht, err := NewServer(WithTcpAddr("", 9999))
    if err != nil {
        fmt.Println(err)
        return
    }
    if err := ht.MonitorHeartBeat(); err != nil {
        fmt.Println(err)
        return
    }
}

// 客户端：发送心跳
func TestClientSendHeartBeat(t *testing.T) {
    go sendHeartBeat("aaa")
    go sendHeartBeat("bbb")
    sendHeartBeat("ccc")
}

// 发送心跳
func sendHeartBeat(uuid string) {
    rand.Seed(time.Now().UnixNano())
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))

    // 解析TCP地址
    tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")
    if err != nil {
        fmt.Printf("resolve tCP addr error,%v\n", err)
        return
    }

    for {
        // 拨号(三次失败取消)s
        conn, err := net.DialTCP("tcp", nil, tcpAddr)
        if err != nil {
            fmt.Printf("dial tcp server server err,%v\n", err)
        }

        u := fmt.Sprintf("%s", uuid)
        // 发送心跳
        conn.Write([]byte(u))
        logrus.Printf("「%s」发送了心跳", u)
        <-time.After(time.Second * 5)
        conn.Close()
    }
}

// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
    "fmt"
    "net"
)

type Option func(h *server) error

// host为空时,监听地址为0.0.0.0
func WithTcpAddr(ip string, port int) Option {
    var addr string
    if ip == "" {
        addr = fmt.Sprintf("0.0.0.0:%d", port)
    } else {
        addr = fmt.Sprintf("%s:%d", ip, port)
    }

    return func(h *server) error {
        tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
        if err != nil {
            return fmt.Errorf("resolve tCP addr error,%v\n", err)
        }

        listener, err := net.ListenTCP("tcp", tcpAddr)
        if err != nil {
            return fmt.Errorf("create tcp listener error,%v\n", err)
        }
        h.Listener = listener
        return nil
    }
}

// 空闲多长时间开始探测(以秒为单位),默认为30s。
func WithKeepAliveIdle(aliveIdle int) Option {
    return func(h *server) error {
        if aliveIdle < 1 || aliveIdle > 3600 {
            return fmt.Errorf("keepaliveidle shuold be in 1 and 3600")
        }
        h.KeepAliveIdle = aliveIdle
        return nil
    }
}

// 探测总次数,默认为4次。
func WithKeepAliveCount(aliveCount int) Option {
    return func(h *server) error {
        if aliveCount < 1 || aliveCount > 10 {
            return fmt.Errorf("keepalivecount shuold be in 1 and 10")
        }
        return nil
    }
}

// 探测时间间隔(以秒为单位),默认为5秒。
func WithKeepAliveInterval(interval int) Option {
    return func(h *server) error {
        if interval < 3 || interval > 300 {
            return fmt.Errorf("keepaliveinterval shuold be in 3 and 300")
        }
        return nil
    }
}

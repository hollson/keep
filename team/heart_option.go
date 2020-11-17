// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package team

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

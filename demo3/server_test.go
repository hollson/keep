// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package demo3

import (
"fmt"
"net"
"strings"
    "testing"
)

var ConnMap map[string]net.Conn = make(map[string]net.Conn)

//ConnMap := make(map[string]net.Conn)

func TestServer(t *testing.T) {
    listen_socket, err := net.Listen("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println("server start error")
    }

    defer listen_socket.Close()
    fmt.Println("server is wating ....")

    for {
        conn, err := listen_socket.Accept()
        if err != nil {
            fmt.Println("conn fail ...")
        }
        fmt.Println(conn.RemoteAddr(), "connect successed")

        go handle(conn)
    }
}

func handle(conn net.Conn) {
    for {
        data := make([]byte, 255)
        msg_read, err := conn.Read(data)
        if msg_read == 0 || err != nil {
            continue
        }

        //解析协议
        msg_str := strings.Split(string(data[0:msg_read]), "|")

        switch msg_str[0] {
        case "nick":
            fmt.Println(conn.RemoteAddr(), "-->", msg_str[1])
            for k, v := range ConnMap {
                if k != msg_str[1] {
                    v.Write([]byte("[" + msg_str[1] + "]: join..."))
                }
            }
            ConnMap[msg_str[1]] = conn
        case "say":
            for k, v := range ConnMap {
                if k != msg_str[1] {
                    fmt.Println("Send "+msg_str[2]+" to ", k)
                    v.Write([]byte("[" + msg_str[1] + "]: " + msg_str[2]))
                }
            }
        case "quit":
            for k, v := range ConnMap {
                if k != msg_str[1] {
                    v.Write([]byte("[" + msg_str[1] + "]: quit"))
                }
            }
            delete(ConnMap, msg_str[1])
        }

    }

}
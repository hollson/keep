// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
    "net"
    "os"
)

var ch = make(chan int)
var nickname string

//go:generate go run client.go
func main() {
    TcpAdd, _ := net.ResolveTCPAddr("tcp", ":54321")
    conn, err := net.DialTCP("tcp", nil, TcpAdd)
    if err != nil {
        fmt.Println("服务没打开")
        os.Exit(1)
    }
    defer conn.Close()

    go tcpReader(conn)
    tcpWriter(conn)
}

func tcpReader(conn *net.TCPConn) {
    buff := make([]byte, 256)
    for {
        j, err := conn.Read(buff)
        if err != nil {
            ch <- 1
            break
        }
        fmt.Printf("%s\n", buff[0:j])
    }
}

func tcpWriter(conn *net.TCPConn) {
    fmt.Println("请输入昵称")
    fmt.Scanln(&nickname)
    fmt.Println("你的昵称是：", nickname)
    for {
        var msg string
        fmt.Scan(&msg)
        fmt.Print("<" + nickname + ">" + "说:")
        // for i, _ := range msg {
        //	fmt.Printf("%c", msg[i])
        // }
        fmt.Println(msg)
        b := []byte("<" + nickname + ">" + "说：" + msg)
        conn.Write(b)
        select {
        case <-ch:
            fmt.Println("server发生错误，请重新连接")
            os.Exit(2)
        default:
        }
    }
}
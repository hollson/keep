// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
    "encoding/json"
    "fmt"
    "net"
    "os"
    "time"

    "github.com/hollson/team/pb"
    "github.com/sirupsen/logrus"
)

var ch = make(chan int)
var nickname string
var action uint8

//go:generate go run client.go
func main() {
    TcpAdd, _ := net.ResolveTCPAddr("tcp", ":4321")
    conn, err := net.DialTCP("tcp", nil, TcpAdd)
    if err != nil {
        fmt.Println("TCP拨号失败,", err)
        os.Exit(1)
    }
    defer conn.Close()

    go TcpReader(conn) // 读数据
    TcpWriter(conn)    // 写数据
}

// 监控服务端消息
func TcpReader(conn *net.TCPConn) {
    buff := make([]byte, 1024*32)
    for {
        j, err := conn.Read(buff)
        if err != nil {
            ch <- 1
            break
        }
        resp := &pb.Response{}
        err = json.Unmarshal(buff[0:j], resp)
        if err != nil {
            logrus.Errorln("json.Unmarshal error ", err)
            logrus.Println(string(buff[0:j]))
            continue
        }
        fmt.Printf("%+v\n", resp)
    }
}


// 数据、xx上线、xx下线、通知


func TcpWriter(conn *net.TCPConn) {
    fmt.Println("请输入Account")
    fmt.Scanln(&nickname)
    fmt.Println("你的UUID是 ：", nickname)

    fmt.Println("请输入ActionID")
    fmt.Scanln(&action)
    for {
        var msg string
        fmt.Scan(&msg)
        req := pb.Request{
            UUID:      nickname,
            Group:     111,
            Timestamp: time.Now().Unix(),
            Sign:      "xxxx",
            Action:    pb.ReqAction(action),
            Receiver:  "BBB",
            Size:      0,
            Data:      []byte(msg),
        }

        data, err := json.Marshal(req)
        if err != nil {
            fmt.Println("Marshal错误")
            continue
        }
        conn.Write(data)
        select {
        case <-ch:
            fmt.Println("server发生错误，请重新连接")
            os.Exit(2)
        default:
        }
    }
}

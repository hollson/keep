// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "os"

    "github.com/hollson/goox/memory"
    "github.com/hollson/team/lib/tcpkeepalive"
    "github.com/hollson/team/pb"
    "github.com/sirupsen/logrus"
)



// 广播
func (s *server) broadCast(req *pb.Request, conn *tcpkeepalive.Conn) {
    // logrus.Println("广播")
    for _, c := range s.Clients {
        if c.conn.RemoteAddr().String() == conn.RemoteAddr().String() {
            continue
        }
        fmt.Printf("%s => %s\n", req.UUID, string(req.Data))
        resp := pb.Response{
            Code:   1,
            Action: 1,
            Size:   1,
            Data:   req.Data,
        }

        data, err := json.Marshal(resp)
        if err != nil {
            logrus.Errorln("json.Marshal error")
            continue
        }
        fmt.Println("发送数据")
        c.conn.Write(data)
    }
}

// 私发
func (s *server) privateSend(req *pb.Request, conn *tcpkeepalive.Conn) {
    if other, ok := s.Clients[req.Receiver]; !ok {
        conn.Write([]byte("用户不存在或用户不在线"))
    } else if req.Receiver == req.UUID {
        conn.Write([]byte("不能自我发送"))
    } else {
        fmt.Printf("%s => %s\n", conn.RemoteAddr().String(), string(req.Data))
        other.conn.Write(req.Data)
    }
}

// 溯源
func (s *server) retrospect(req *pb.Request, conn *tcpkeepalive.Conn) {
    conn.Write([]byte("溯源"))
}

// 用户下线，1. 主动退出，2. 探测退出
func (s *server) Offline(conn *tcpkeepalive.Conn, _type int) {
    logrus.Errorf("%s已退出[%d]", conn.RemoteAddr().String(), _type)
    delete(s.Clients, conn.RemoteAddr().String())
}

func Read2() string {
    fi, err := os.Open("file/test")
    if err != nil {
        panic(err)
    }
    defer fi.Close()

    r := bufio.NewReader(fi)
    var chunks []byte

    buf := make([]byte, 1024)

    for {
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if 0 == n {
            break
        }
        // fmt.Println(string(buf))
        chunks = append(chunks, buf...)
    }
    return string(chunks)
    // fmt.Println(string(chunks))
}

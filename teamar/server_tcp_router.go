// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package teamar

import (
    "encoding/json"
    "fmt"
    "io"

    "github.com/hollson/goox/memory"
    "github.com/hollson/team/lib/tcpkeepalive"
    "github.com/hollson/team/pb"
    "github.com/sirupsen/logrus"
)

// 处理消息
func (s *teamar) TcpRoute(conn *tcpkeepalive.Conn) {
    // 客户路由
    for {
        var buf = make([]byte, memory.MB*4) // todo
        n, err := conn.Read(buf)
        if err == io.EOF {
            s.Offline(conn, 1)
            break
        }
        if err != nil {
            logrus.Errorln(err)
            break
        }

        var req *pb.Request = new(pb.Request)
        if err := json.Unmarshal(buf[:n], req); err == nil {
            if _, ok := s.Team[req.UUID]; !ok {
                s.Team[req.UUID] = &client{ // 验证未知用户
                    uuid: req.UUID,
                    conn: conn,
                }
                logrus.Infof("【%s】%s已上线,当前在线人数：%d", conn.RemoteAddr().String(), req.UUID, len(s.Team))
            }

            switch req.Uri {
            case "pb.BROADCAST":
                s.broadCast(req, conn)
            case "pb.PRIVATE":
                s.privateSend(req, conn)
            case "pb.RETROSPECT":
                s.retrospect(req, conn)
            default:
                fmt.Println("未知的操作：", req.Uri)
            }
        } else {
            logrus.Errorln("Unmarshal", err)
        }
    }
}



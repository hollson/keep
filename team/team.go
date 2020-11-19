// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package team

import (
    "github.com/hollson/team/lib/tcpkeepalive"
)

// 团队字典(map[team_id]members)
type Team map[int64]Members //

// 成员字典(map[uuid]：remote_addrs)
type Members map[string]string

// 远程地址(map[teamId_uuid]remote_addrs)
type RemoteAddrs map[string]string

// 连接客户端
type client struct {
    uuid string
    // tcpid  //TcpAddrKey TCP地址标识符
    conn *tcpkeepalive.Conn
}

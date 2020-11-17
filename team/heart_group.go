// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// 心跳监听负载器(负载均衡)
package team

type Group interface {
    AllotGroupHash() int // 分配心跳监听组的端口号
}

// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package keep

import (
    "time"
)

type Option func(s *Server)

// 设置KeepAlive客户端探测
//  idle：空闲多长时间开始探测,默认为10秒，单位:秒。
//  count：探测总次数,默认为3次。
//  interval：探测时间间隔,默认为3秒，单位:秒。
func WithKeepAlive(idle, count, interval int) Option {
    if idle < 1 || idle > 3600 {
        idle = 10
    }
    if count < 1 || count > 10 {
        count = 3
    }
    if interval < 1 || interval > 300 {
        interval = 3
    }
    return func(s *Server) {
        s.keepAlive.idle = time.Duration(idle) * time.Second
        s.keepAlive.count = count
        s.keepAlive.interval = time.Duration(interval) * time.Second
    }
}

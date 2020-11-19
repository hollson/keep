// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
    "github.com/hollson/team/pb"
)

// 消息/错误对照表
var CodeMsg = map[pb.CODE]string{
    1: "OK",
    2: "签名验证错误",
    3: "无权限",
}


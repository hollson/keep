// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pb

// 消息状态码
type CODE int8
type ReqAction int8

const (
    LOGIN      ReqAction = iota + 1 // 登录
    BROADCAST                       // 广播
    PRIVATE                         // 私发
    RETROSPECT                      // 溯源
    LOGOUT                          // 退出
)

type Request struct {
    Uri       string `json:"uri"`       // ReqAction `json:"action"`    // 操作类型（登录，群发，私发，溯源，op1，op2）
    UUID      string `json:"uuid"`      // 用户编号(OpenId)
    Group     int64  `json:"group"`     // 群组ID(多人协作项目ID)
    Timestamp int64  `json:"timestamp"` // 时间戳
    Serial    int64  `json:"serial"`    // 序列号(请求与响应配对)
    Sign      string `json:"sign"`      // 签名(校验身份)
    Receiver  string `json:"receiver"`  // 接收者(群组ID[群发]或个人UUID[私发])
    Size      int32  `json:"size"`      // 报文大小
    Data      []byte `json:"data"`      // 数据内容
}

// https://studygolang.com/articles/3981
// https://studygolang.com/articles/5062
// https://blog.csdn.net/fengfengdiandia/article/details/79986237
// https://github.com/vmihailenco/msgpack
type Response struct {
    Code   CODE   `json:"code"`   // 状态码
    Action int8   `json:"action"` // 数据类型，数据、xx上线、xx下线、通知
    Size   int32  `json:"size"`   // 数据大小
    Data   []byte `json:"data"`   // 数据内容
}

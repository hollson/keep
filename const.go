// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package keep

// 响应状态码
type Status uint32

// 通用消息状态
const (
    MSG_NULL        Status = iota // 无状态码
    MSG_OK                        // 响应成功
    MSG_NOT_FOUND                 // 路由错误
    MSG_ERR_PACKET                // 参数错误 - 解包失败
    MSG_ERR_CRC32                 // 校验失败 - (解包成功)crc失败
    MSG_ERR_LENGHT                // 长度错误 - (解包成功)消息长度错误
    MSG_ERR_GROUP                 // 群号错误 - (解包成功)
    MSG_ERR_UUID                  // 身份错误 - (解包成功)未登录或账号错误等
    MSG_ERR_LIMIT                 // 权限错误
    MSG_ERR_SERVER                // 服务错误
    MSG_ERR_TIMEOUT               // 服务超时
    MSG_CMD_ABORT                 // 强制中断/下线
)

func (a Format) Status() uint8 {
    return uint8(a)
}

// 对象序列化格式
type Format uint8

const (
    FMT_BIN  Format = iota // 二进制序列化(小端模式)
    FMT_PB                 // Proto序列化(大端模式)
    FMT_JSON               // Json序列化
)

func (a Format) Uint8() uint8 {
    return uint8(a)
}

// 动作类型
type Act uint32

const (
    // ACT_TIP Act = iota+1 // 提示
    // ACT_ERR            // 错误
    // ACT_CMD            // 命令
)

func (a Act) Uint32() uint32 {
    return uint32(a)
}

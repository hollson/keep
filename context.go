// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package keep

import (
    "context"
    "net"

    "github.com/sirupsen/logrus"
)

type H map[string]interface{}

type Context struct {
    Ctx    context.Context
    conn   net.Conn
    Query  *ReqMessage      // 请求参数
    chanel chan *ReqMessage // 参数通道
}

type HandlerFunc func(ctx *Context)

type Router map[uint32][]HandlerFunc

// 创建上下文
//  c：当前连接者连接对象
func newContext(c net.Conn) *Context {
    return &Context{
        Ctx:    context.Background(),
        conn:   c,
        Query:  new(ReqMessage),
        chanel: make(chan *ReqMessage),
    }
}

// 获取当前客户的连接对象
func (c *Context) SelfConn() net.Conn {
    return c.conn
}

// 动作必须大于100
// 仅响应动作
func (c *Context) Action(actorId uint32, receivers ...net.Conn) {
    msg := &Message{
        Action: actorId,
        Code:   MSG_OK,
        Length: 0,
        // crc:    0,
    }
    sendMsg(msg, receivers...)
}

// 给当前连接客户发送消息
func (c *Context) Self(msg *Message) {
    sendMsg(msg, c.conn)
}

// 发送消息(广播)
func (c *Context) Send(msg *Message, receivers ...net.Conn) {
    sendMsg(msg, receivers...)
}

// 发送中断/强制下线通知
func (c *Context) Abort(receivers ...net.Conn) {
    msg := &Message{
        Action: 0,
        Code:   MSG_CMD_ABORT,
        Length: 0,
    }
    if err := sendMsg(msg, receivers...); err != nil {
        logrus.Errorln("Abort Response: ", err)
    }
}

// 发送错误状态消息
func (c *Context) Code(status Status) {
    msg := &Message{
        Action: 0,
        Code:   status,
        Length: 0,
        // crc:    0,
    }
    if err := sendMsg(msg, c.conn); err != nil {
        logrus.Errorln("Error Response: ", err.Error())
    }
}

func sendMsg(msg *Message, receivers ...net.Conn) error {
    data, err := msg.Pack()
    if err != nil {
        return err
    }

    for _, c := range receivers {
        if _, err := c.Write(data); err != nil {
            return err
        }
    }
    return nil
}

// 客户端地址
func (c *Context) ClientAddr() net.Addr {
    return c.conn.RemoteAddr()
}

// 客户端地址
func (c *Context) Dispose() error {
    return nil
}

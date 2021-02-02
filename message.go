// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package keep

import (
    "bufio"
    "bytes"
    "encoding/binary"
    "fmt"
    "hash/crc32"
    "io"
)

// https://studygolang.com/articles/3903  序列化
// 压缩

// var (
//     compress = true
// )

// TCP报文 - 封包器
type Packer interface {
    Pack() ([]byte, error)
}

// TCP报文 - 解包器
type UnPacker interface {
    UnPack(r io.Reader) error
}

// 上行消息
type ReqMessage struct {
    Action uint32 // 消息类型(路由标识)
    Group  uint32 // 群组编号
    UUID   uint32 // 用户标识
    Length uint32 // 消息长度(只读，用于数据校验)
    crc32  uint32 // (私有)校验码
    Data   []byte // 消息内容
}

func (m *ReqMessage) String() string {
    type dump struct {
        Action uint32 // 消息类型(路由标识)
        Group  uint32 // 群组编号
        UUID   uint32 // 用户标识
        Length uint32 // 消息长度(只读，用于数据校验)
        crc32  uint32 // (私有)校验码
        Data   string // 消息内容
    }

    req := dump{
        Action: m.Action,
        Group:  m.Group,
        UUID:   m.UUID,
        Length: m.Length,
        crc32:  m.crc32,
        Data:   string(m.Data), // 忽略内容
    }
    if len(m.Data) > 1024 {
        req.Data = "Too long, omitted..."
    }
    return fmt.Sprintf("%+v", req)
}

// Socket报文封包
func (m *ReqMessage) Pack() ([]byte, error) {
    var buf bytes.Buffer

    m.Length = uint32(len(m.Data))
    m.crc32 = crc32.ChecksumIEEE(m.Data)

    // TCP/IP协议RFC1700规定使用“大端”字节序为网络字节序
    if err := binary.Write(&buf, binary.BigEndian, m.Action); err != nil {
        return nil, fmt.Errorf("pack of action error: %w", err)
    }

    if err := binary.Write(&buf, binary.BigEndian, m.Group); err != nil {
        return nil, fmt.Errorf("pack of group error: %w", err)
    }

    if err := binary.Write(&buf, binary.BigEndian, m.UUID); err != nil {
        return nil, fmt.Errorf("pack of uuid error: %w", err)
    }

    if err := binary.Write(&buf, binary.BigEndian, m.Length); err != nil {
        return nil, fmt.Errorf("pack of length error: %w", err)
    }

    if err := binary.Write(&buf, binary.BigEndian, m.crc32); err != nil {
        return nil, fmt.Errorf("pack of crc32 error: %w", err)
    }

    // 写入消息内容
    if err := binary.Write(&buf, binary.BigEndian, m.Data); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// scaner ????
// Socket报文解包
func UnPack(r io.Reader) (*ReqMessage, error) {
    var (
        req     = new(ReqMessage)
        _peek   = 5 * 4
        _reader = bufio.NewReader(r)
    )

    // 获取报文头
    head, err := _reader.Peek(_peek)
    if err != nil {
        // 包含EOF错误
        return nil, err
    }

    if err := binary.Read(bytes.NewBuffer(head[:4]), binary.BigEndian, &req.Action); err != nil {
        return nil, fmt.Errorf("unpack of actor error: %w", err)
    }

    if err := binary.Read(bytes.NewBuffer(head[4:8]), binary.BigEndian, &req.Group); err != nil {
        return nil, fmt.Errorf("unpack of actor group: %w", err)
    }

    if err := binary.Read(bytes.NewBuffer(head[8:12]), binary.BigEndian, &req.UUID); err != nil {
        return nil, fmt.Errorf("unpack of actor uuid: %w", err)
    }

    if err := binary.Read(bytes.NewBuffer(head[12:16]), binary.BigEndian, &req.Length); err != nil {
        return nil, fmt.Errorf("unpack of crc32 group: %w", err)
    }

    if err := binary.Read(bytes.NewBuffer(head[16:]), binary.BigEndian, &req.crc32); err != nil {
        return nil, fmt.Errorf("unpack of crc32 group: %w", err)
    }

    // 跳过前16个字节
    if n, err := _reader.Discard(_peek); err != nil {
        return nil, fmt.Errorf("unpack discard peek(%d) error:%w", n, err)
    }

    if int(req.Length) != _reader.Buffered() {
        return nil, fmt.Errorf("unpack of length[%d] match buffred[%d] error", req.Length, _reader.Buffered())
    }

    // 消息体为空，则短路返回
    if _reader.Buffered() == 0 {
        return req, nil
    }
    // fmt.Println("收到报文===》》》7", req)
    buffer := make([]byte, 1024)
    for _reader.Buffered() > 0 {
        n, err := _reader.Read(buffer)
        if err != nil {
            return nil, fmt.Errorf("unpack read error:%w", err)
        }
        req.Data = append(req.Data, buffer[:n]...)
    }
    return req, nil
}

// 下行消息
type Message struct {
    Action uint32 // 消息类型(路由标识)
    Code   Status // 状态码
    // format uint8  // 数据格式
    Length uint32 // (私有)消息长度
    // crc    uint32 // (私有)校验码
    Data []byte // 消息数据
}

// Socket报文封包
func (msg *Message) Pack() ([]byte, error) {
    var buf bytes.Buffer

    msg.Length = uint32(len(msg.Data))
    // msg.crc = 111

    // TCP/IP协议RFC1700规定使用“大端”字节序为网络字节序
    if err := binary.Write(&buf, binary.BigEndian, msg.Action); err != nil {
        return nil, fmt.Errorf("pack of actor error: %w", err)
    }

    if err := binary.Write(&buf, binary.BigEndian, msg.Code); err != nil {
        return nil, fmt.Errorf("pack of code error: %w", err)
    }

    if err := binary.Write(&buf, binary.BigEndian, msg.Length); err != nil {
        return nil, fmt.Errorf("pack of length error: %w", err)
    }

    // if err := binary.Write(&buf, binary.BigEndian, msg.crc); err != nil {
    //     return nil, fmt.Errorf("pack of crc32 error: %w", err)
    // }

    // 写入消息内容
    if err := binary.Write(&buf, binary.BigEndian, msg.Data); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// scaner ????
// Socket报文解包
func MessageUnPack(r io.Reader) (*Message, error) {
    var (
        req     = new(Message)
        _peek   = 3 * 4
        _reader = bufio.NewReader(r)
    )
    fmt.Println("bufed: ", _reader.Buffered())
    // 获取报文头
    head, err := _reader.Peek(_peek)
    if err != nil {
        return nil, fmt.Errorf("reply unpack error： %w", err)
    }

    if err := binary.Read(bytes.NewBuffer(head[:4]), binary.BigEndian, &req.Action); err != nil {
        return nil, fmt.Errorf("unpack of actor error: %w", err)
    }

    if err := binary.Read(bytes.NewBuffer(head[4:8]), binary.BigEndian, &req.Code); err != nil {
        return nil, fmt.Errorf("unpack of actor group: %w", err)
    }

    if err := binary.Read(bytes.NewBuffer(head[8:12]), binary.BigEndian, &req.Length); err != nil {
        return nil, fmt.Errorf("unpack of actor uuid: %w", err)
    }

    // if err := binary.Read(bytes.NewBuffer(head[12:16]), binary.BigEndian, &req.Data); err != nil {
    //     return nil, fmt.Errorf("unpack of crc32 group: %w", err)
    // }
    fmt.Println("bufed: ", _reader.Buffered())
    // 跳过前16个字节
    if n, err := _reader.Discard(_peek); err != nil {
        return nil, fmt.Errorf("unpack discard peek(%d) error:%w", n, err)
    }

    if int(req.Length) != _reader.Buffered() {
        fmt.Println("===>>>", req)
        return nil, fmt.Errorf("reply unpack of length[%d] match buffred[%d] error", req.Length, _reader.Buffered())
    }

    // 消息体为空，则短路返回
    if _reader.Buffered() == 0 {
        return req, nil
    }
    // fmt.Println("收到报文===》》》7", req)
    buffer := make([]byte, 1024)
    for _reader.Buffered() > 0 {
        n, err := _reader.Read(buffer)
        if err != nil {
            return nil, fmt.Errorf("unpack read error:%w", err)
        }
        req.Data = append(req.Data, buffer[:n]...)
    }
    return req, nil
}

func (m *Message) String() string {
    type dump struct {
        Action uint32 // 消息类型(路由标识)
        Code   Status // 状态码
        // format uint8  // 数据格式
        length uint32 // (私有)消息长度
        // crc    uint32 // (私有)校验码
        Data string // 消息数据
    }

    req := dump{
        Action: m.Action,
        Code:   m.Code,
        length: m.Length,
        Data:   string(m.Data), // 忽略内容
    }
    if len(m.Data) > 1024 {
        req.Data = "Too long, omitted..."
    }
    return fmt.Sprintf("%+v", req)
}

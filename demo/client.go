package main

import (
    "fmt"
    "net"
    "os"

    "github.com/hollson/keep"
    "github.com/sirupsen/logrus"
)

var (
    uuid    uint32
    teamdId uint32
    actId   uint32
    recev   string
    data    string
)

//go:generate go run client.go
func main() {
    TcpAdd, _ := net.ResolveTCPAddr("tcp", ":4321")
    conn, err := net.DialTCP("tcp", nil, TcpAdd)
    if err != nil {
        fmt.Println("TCP拨号失败,", err)
        os.Exit(1)
    }
    defer conn.Close()

    go TcpReader(conn) // 读数据
    TcpWriter(conn)    // 写数据
}

// 监控服务端消息
func TcpReader(conn *net.TCPConn) {
    for {
        m, err := keep.MessageUnPack(conn)
        if err != nil {
            logrus.Errorf("Server错误，%v，请重新连接.", err)
            os.Exit(1)
        }
        fmt.Println("<<= ", m)
    }
}

// 数据、xx上线、xx下线、通知

func TcpWriter(conn *net.TCPConn) {
    fmt.Print("用户账号 : ")
    fmt.Scanln(&uuid)
    fmt.Print("团队编号 : ")
    fmt.Scanln(&teamdId)
    fmt.Println()

    for {
        fmt.Print("操作类型(101~200): ")
        fmt.Scanln(&actId)

        req := &keep.ReqMessage{
            Action: actId,
            Group:  teamdId,
            UUID:   uuid,
            Data:   nil,
        }
        // fmt.Print("内容 :")

        // fmt.Scanln(&data)
        // req.Data = []byte(data)

        req.Action = 101
        req.Data = testText()

        // 封包
        b, err := req.Pack()
        if err != nil {
            logrus.Errorln(err)
        }

        // 发送
        _, err = conn.Write(b)
        if err != nil {
            logrus.Errorln(err)
            continue
        }
        fmt.Printf("=> 发送数据:  %s\n\n", req)
    }
}

// 测试内容
func testText() []byte {
    mps := map[int]string{
        1: "hello world",
        2: "logrus.Errorf",
        3: "发送数据 😊",
        4: "12345",
        5: "Unmarshal",
        6: "We walked on the main road",
        7: "fmt.Println(😂😂😂)",
        8: "return []byte(s)",
        9: "🚗 ✈✏️ 🌹 😊️",
    }
    for _, s := range mps {
        return []byte(s)
    }
    panic("done")
}

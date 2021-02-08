// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package bin

import (
    "fmt"
    "testing"
)

type Student struct {
    Id   int32
    Name string
}

func TestMashal(t *testing.T) {
    s := &Student{
        Id:   1001,
        Name: "Tom",
    }
    data, err := Encode(s)
    if err != nil {
        panic(fmt.Errorf("==>%v", err))
    }

    var stu *Student
    if err := Decode(data, stu); err != nil {
        panic(err)
    }

    fmt.Println(stu)
}

type Person struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

// func TestJsonMashal(t *testing.T) {
//     p := Person{
//         Id:   1001,
//         Name: "jack",
//     }
//
//     buf := new(bytes.Buffer)
//     // var pi float64 = math.Pi
//     err := binary.Write(buf, binary.LittleEndian, p)
//     if err != nil {
//         fmt.Println("binary.Write failed:", err)
//     }
//     fmt.Printf("% x", buf.Bytes())
//
//
//     // p := Person{
//     //     Id:   1001,
//     //     Name: "jack",
//     // }
//     // // 方式1：
//     // r, err := json.MarshalIndent(p, "", "\t")
//     // if err != nil {
//     //     log.Println(err)
//     // }
//     // fmt.Println(string(r))
//     //
//     // // 方式2：
//     // r, err = Encode(p)
//     // if err != nil {
//     //     log.Println(err)
//     // }
//     // fmt.Println(string(r))
//
//     // var dd="hello"
//     // _buf := new(bytes.Buffer)
//     // if err := binary.Write(_buf, binary.LittleEndian, p); err != nil {
//     //     log.Println(err)
//     // }
//     // fmt.Println(_buf.Bytes())
//
//     // _buf = new(bytes.Buffer)
//     // if err := binary.Write(_buf, binary.BigEndian, p); err != nil {
//     //     log.Println(err)
//     // }
//     // fmt.Println(_buf.Bytes())
//
// }

// func Encode(v interface{}) ([]byte, error) {
//     return json.Marshal(v)
// }

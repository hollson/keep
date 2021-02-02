// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lib

import (
    "encoding/json"
    "fmt"
    "testing"
)

type Student struct {
    Id   int32
    Name string
}


func TestMashal(t *testing.T) {
    s := Student{
        Id:   1001,
        Name: "Tom",
    }
    data, err := BinEncode(s)
    if err != nil {
        panic(fmt.Errorf("==>%v", err))
    }

    var stu *Student
    if err := BinDecode(data, stu); err != nil {
        panic(err)
    }

    fmt.Println(stu)
}

type Person struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func TestJsonMashal(t *testing.T) {
    p := Person{
        Id:   1001,
        Name: "jack",
    }
    // 方式1：
    r, _ := json.MarshalIndent(p, "", "\t")
    fmt.Println(string(r))

    // 方式2：
    r, _ = Encode(p)
    fmt.Println(string(r))
}

func Encode(v interface{}) ([]byte, error) {
    return json.Marshal(v)
}

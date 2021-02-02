// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lib

import (
    "bytes"
    "encoding/binary"
    "encoding/json"
)

// 二进制编码:
//  order: 默认为LittleEndian
func BinEncode(obj interface{}, order ...binary.ByteOrder) ([]byte, error) {
    var (
        _buf                    = new(bytes.Buffer)
        _order binary.ByteOrder = binary.LittleEndian
    )
    if len(order) > 0 {
        _order = order[0]
    }
    if err := binary.Write(_buf, _order, obj); err != nil {
        return nil, err
    }
    return _buf.Bytes(), nil
}

// 二进制解码:
//  order: 默认为LittleEndian
func BinDecode(b []byte, v interface{}, order ...binary.ByteOrder) error {
    var (
        _buf                    = bytes.NewBuffer(b)
        _order binary.ByteOrder = binary.LittleEndian
    )
    if len(order) > 0 {
        _order = order[0]
    }
    if err := binary.Read(_buf, _order, &v); err != nil {
        return err
    }
    return nil
}


func JsonEncode(obj interface{}, order ...binary.ByteOrder) ([]byte, error) {
    return json.Marshal(obj)
}

// func JsonDecode(b []byte, v interface{}, order ...binary.ByteOrder) error {
//      return json.Unmarshal(nil)
// }
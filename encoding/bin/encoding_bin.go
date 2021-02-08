// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package bin

import (
    "bytes"
    "encoding/binary"
)

// 二进制编码(ç: LittleEndian)
func Encode(obj interface{}) ([]byte, error) {
    var (
        _buf                    = new(bytes.Buffer)
        _order binary.ByteOrder = binary.LittleEndian
    )
    if err := binary.Write(_buf, _order, obj); err != nil {
        return nil, err
    }
    return _buf.Bytes(), nil
}

// 二进制解码(ByteOrder: LittleEndian)
func Decode(b []byte, v interface{}) error {
    var (
        _buf                    = bytes.NewBuffer(b)
        _order binary.ByteOrder = binary.LittleEndian
    )

    if err := binary.Read(_buf, _order, &v); err != nil {
        return err
    }
    return nil
}

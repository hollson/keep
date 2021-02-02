// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package keep

import (
    "bytes"
    "encoding/binary"
    "encoding/json"
)

type Encode func(m *ReqMessage) ([]byte, error)

func BinEncode(m *ReqMessage) ([]byte, error) {
    _buf := new(bytes.Buffer)
    if err := binary.Write(_buf, binary.LittleEndian, m); err != nil {
        return nil, err
    }
    return _buf.Bytes(), nil
}

func JsonEncode(m *ReqMessage) ([]byte, error) {
    return json.Marshal(m)
}

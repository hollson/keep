// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package encoding

type Encoding interface {
    Encode(obj interface{}) ([]byte, error)
    Decode(b []byte, v interface{}) error
}
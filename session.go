// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package keep

// 群组成员
type Member interface {
    Group() int64 // 群组编号
    Ident() int64 // 成员标识
}

type Session struct {
    Group map[int64]int32
}

// 在线群组
func (s *Session) Groups() {

}

// 在线成员
func (s *Session) Members(groupId int64) {

}

// 成员信息
func (s *Session) Member(groupId int64, uuid int64) {

}



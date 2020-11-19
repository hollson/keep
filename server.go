// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
    "os"

    "github.com/hollson/team/server"
    "github.com/sirupsen/logrus"
)

func main() {
    server, err := server.NewServer()
    if err != nil {
        logrus.Errorln(err)
        os.Exit(1)
    }
    server.Run()
    server.Close()
}

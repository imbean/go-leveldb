// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"code.google.com/p/go-leveldb"
)

func main() {
	fmt.Printf("leveldb-%d.%d\n", leveldb.MajorVersion, leveldb.MinorVersion)
}

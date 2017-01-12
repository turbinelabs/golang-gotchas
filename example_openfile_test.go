/*
Copyright 2017 Turbine Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package golanggotchas_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	filename = fmt.Sprintf("/tmp/testfile.%d", os.Getpid())
)

func testOpenFile(data string, flags int) {
	f, _ := os.OpenFile(filename, flags, 0660)
	defer f.Close()
	fmt.Fprintln(f, data)
}

func dumpFile() {
	bytes, _ := ioutil.ReadFile(filename)
	fmt.Println(string(bytes))
}

func Example_openfile() {
	testOpenFile(strings.Repeat("x", 10), os.O_CREATE|os.O_WRONLY)
	dumpFile()

	testOpenFile(strings.Repeat("y", 5), os.O_CREATE|os.O_WRONLY)
	dumpFile()

	// Surprise! The file now contains "yyyyy\nxxxx\n" -- need to
	// either set os.O_TRUNC to truncate the file on open or
	// os.O_APPEND to open with the file pointer at EOF.
	// See: https://golang.org/pkg/os/#pkg-constants
	// Related: man 2 open

	// Output:
	// xxxxxxxxxx
	//
	// yyyyy
	// xxxx
}

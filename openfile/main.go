package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	filename = fmt.Sprintf("testfile.%d", os.Getpid())
)

func test(data string, flags int) {
	f, _ := os.OpenFile(filename, flags, 0660)
	defer f.Close()
	fmt.Fprintln(f, data)
}

func main() {
	test(strings.Repeat("x", 10), os.O_CREATE|os.O_WRONLY)
	test(strings.Repeat("y", 5), os.O_CREATE|os.O_WRONLY)
	// Surprise! The file now contains "yyyyy\nxxxx\n" -- need to
	// either set os.O_TRUNC to truncate the file on open or
	// os.O_APPEND to open with the file pointer at EOF.
	// See: https://golang.org/pkg/os/#pkg-constants
	// Related: man 2 open
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("let's read this message\n")
	b := bytes.NewBuffer(nil)
	w := io.MultiWriter(b, os.Stdout)
	io.Copy(w, r)
	fmt.Println(b.String())
}

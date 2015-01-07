package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/kr/pretty"
)

func main() {
	a := bufio.NewReader(bytes.NewReader([]byte("var i")))
	r := NewReader(a)
	t := NewLexer(r)
	s, err := t.Scan()
	if err != io.EOF {
		fmt.Println(err)
	} else {
		fmt.Printf("%# v", pretty.Formatter(s))
	}
}

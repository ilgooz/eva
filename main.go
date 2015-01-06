package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kr/pretty"
	"io"
)

func main() {
	a := bufio.NewReader(bytes.NewReader([]byte("var i = 0; print(i)")))
	r := NewReader(a)
	t := NewLexer(r)
	s, err := t.Scan()
	if err != io.EOF {
		fmt.Println(err)
	} else {
		fmt.Printf("%# v", pretty.Formatter(s))
	}
}

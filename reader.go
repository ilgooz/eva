package main

import (
	"bufio"
)

type Reader struct {
	r    *bufio.Reader
	line int
	col  int
}

func NewReader(r *bufio.Reader) *Reader {
	return &Reader{
		r:    r,
		line: 1,
		col:  1,
	}
}

func (r *Reader) next() (string, error) {
	var l string
	b, err := r.r.ReadByte()
	if err != nil {
		return l, err
	}
	l = string(b)
	if l == "\\n" || l == "\\r\\n" || l == "\\r" {
		r.line++
		r.col = 1
	} else {
		r.col++
	}
	return l, nil
}

func (r *Reader) jump(w int) error {
	for ; w > 0; w-- {
		_, err := r.next()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reader) peek(w int) (string, error) {
	b, err := r.r.Peek(w)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

package main

import (
	"fmt"
	"io"
)

const (
	t_whitespace int = iota + 1
	t_var
	t_number
	t_equal
	t_ident
	t_lparan
	t_rparan
	t_semi
)

type Token struct {
	typ int
	lex string
}

type Lexer struct {
	r *Reader
}

func NewLexer(r *Reader) *Lexer {
	return &Lexer{
		r: r,
	}
}

func (l *Lexer) Scan() ([]*Token, error) {
	var tokens []*Token
	for {
		token, err := l.scan()
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (l *Lexer) scan() (*Token, error) {
	token := &Token{}

	c, err := l.r.peek(1)
	if err != nil {
		return token, err
	}

	switch c {
	case " ", "\\t":
		token.typ = t_whitespace
	case "=":
		token.typ = t_equal
	case ";":
		token.typ = t_semi
	case "(":
		token.typ = t_lparan
	case ")":
		token.typ = t_rparan
	default:
		if digit := isDigit(c); digit {
			return l.scanNumeric()
		}

		char, err := isChar(c)
		if err != nil {
			return token, err
		}
		if char {
			return l.scanKI()
		}

		return token, newError(fmt.Sprintf("eva error: Illegal char %s at %d:%d", c, l.r.line, l.r.col))
	}

	err = l.r.jump(1)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (l *Lexer) scanIdentifier() {

}

func (l *Lexer) scanNumeric() (*Token, error) {
	token := &Token{
		typ: t_number,
	}

	for {
		c, err := l.r.peek(1)
		if err != nil {
			return token, err
		}
		digit := isDigit(c)
		if digit {
			token.lex += c
			err = l.r.jump(1)
			if err != nil {
				return token, err
			}
		} else {
			break
		}
	}

	return token, nil
}

// Keyword or Identifier
func (l *Lexer) scanKI() (*Token, error) {
	token := &Token{}

	// keyword
	var matched bool
	var err error

	matched, err = l.match("var")
	if err != nil {
		return token, err
	}
	if matched {
		token.typ = t_var
		return token, nil
	}

	// identifier
	token.typ = t_ident
	for {
		c, err := l.r.peek(1)
		if err != nil && err != io.EOF {
			return token, err
		}
		char, err := isChar(c)
		if err != nil {
			return token, err
		}
		if char {
			token.lex += c
			err = l.r.jump(1)
			if err != nil {
				return token, err
			}
		} else {
			break
		}
	}

	return token, nil
}

func (l *Lexer) match(s string) (bool, error) {
	ln := len(s)
	cs, err := l.r.peek(ln)
	if err != nil && io.EOF != err {
		return false, err
	}
	if io.EOF == err || cs != s {
		return false, nil
	}
	err = l.r.jump(ln)
	if err != nil {
		return false, err
	}
	return true, nil
}

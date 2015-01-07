package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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
	case " ":
		token.typ = t_whitespace
	case "\\t":
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
		digit := isDigit(c)
		if digit {
			return l.numeric()
		}

		char, err := isChar(c)
		if err != nil {
			return token, err
		}
		if char {
			return l.ki()
		}
		return token, errors.New(fmt.Sprintf("eva error: Illegal char %s at %d:%d", c, l.r.line, l.r.col))
	}

	err = l.r.jump(1)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (l *Lexer) numeric() (*Token, error) {
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

// Dedect keyword/identifier
func (l *Lexer) ki() (*Token, error) {
	token := &Token{}

	// is keyword: var
	chars, err := l.r.peek(3)
	if err != nil {
		return token, err
	}

	if chars == "var" {
		token.typ = t_var
		err = l.r.jump(3)
		if err != nil {
			return token, err
		}
		return token, nil
	}

	// is identifier
	token.typ = t_ident
	for {
		c, err := l.r.peek(1)
		if err != nil {
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

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func isChar(s string) (bool, error) {
	return regexp.MatchString("[a-zA-Z]", s)
}

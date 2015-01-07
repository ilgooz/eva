package main

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/bmizerany/assert"
)

var programs = map[string][]*Token{
	"var": []*Token{
		&Token{t_var, ""},
	},

	"var ": []*Token{
		&Token{t_var, ""},
		&Token{t_whitespace, ""},
	},

	"var a": []*Token{
		&Token{t_var, ""},
		&Token{t_whitespace, ""},
		&Token{t_ident, "a"},
	},
}

func tokenize(s string) ([]*Token, error) {
	reader := NewReader(bufio.NewReader(bytes.NewReader([]byte(s))))
	lexer := NewLexer(reader)
	tokens, err := lexer.Scan()
	if err != io.EOF {
		return tokens, err
	}
	return tokens, nil
}

func TestPrograms(t *testing.T) {
	for program, expected := range programs {
		token, err := tokenize(program)
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, token)
	}
}

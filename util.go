package main

import (
	"regexp"
	"strconv"
)

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func isChar(s string) (bool, error) {
	return regexp.MatchString("[a-zA-Z]", s)
}

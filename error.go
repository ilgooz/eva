package main

import "errors"

func newError(s string) error {
	return errors.New("eva error: " + s)
}

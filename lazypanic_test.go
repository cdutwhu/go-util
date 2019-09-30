package util

import (
	"testing"
	// "github.com/go-errors/errors"
)

//
func makeErr(a int, b string) (int, string, error) {
	return a, b, nil // errors.New("here")
}

func TestPanicOnError(t *testing.T) {
	defer func() { ph(recover(), "./log.txt") }()
	pe(fEf("test#NOFATAL"))
}

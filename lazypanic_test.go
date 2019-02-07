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
	defer func() { PH(recover(), "./log.txt") }()
	PE(fEf("test#NOFATAL"))
}

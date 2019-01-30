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
	b, c := Must2(makeErr(100, "1we200"))
	fPln(b, c)
}

package util

import (
	"fmt"
	"testing"
)

var pln = fmt.Println

func TestRemovePrefix(t *testing.T) {
	pln(Str("sif.abc").RemovePrefix("sif."))
}

func TestRemoveSuffix(t *testing.T) {
	pln(Str("sif.abc").RemoveSuffix("abc"))
}

func TestRemoveTailFromLast(t *testing.T) {
	pln(Str("a.sif.abc").RemoveTailFromLast("."))
}

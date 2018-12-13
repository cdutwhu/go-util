package util

import (
	"fmt"
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	pln := fmt.Println
	pln(Str("sif.abc").RemovePrefix("sif."))
}

func TestRemoveSuffix(t *testing.T) {
	pln := fmt.Println
	pln(Str("sif.abc").RemoveSuffix("abc"))
}

package util

import (
	"fmt"
	"testing"
)

func TestNearest(t *testing.T) {
	pln := fmt.Println
	pln(I64(10).Nearest(100, 30, 12, 3, -4, 110))
}

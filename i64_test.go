package util

import (
	"testing"
)

func TestNearest(t *testing.T) {
	fPln(I64(10).Nearest(100, 30, 12, 3, -4, 110))
}

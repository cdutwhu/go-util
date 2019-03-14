package util

import (
	"testing"
)

func TestPrintMaxMin(t *testing.T) {
	fPln(MaxInt)
	fPln(MaxUint)
	fPln(MinInt)
	fPln(MinUint)

	fPln(MinPosOf(-1, 0, 1, 5, 3))
}

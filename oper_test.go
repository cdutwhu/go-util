package util

import "testing"

func TestTerOper(t *testing.T) {
	fPln(TerOper(1 == 2, "abc", "def").(string))
}

package util

import "testing"

func TestTerOp(t *testing.T) {
	fPln(IF(1 == 1, "abc", "def").(string))
	fPln(IF(1 == 2, "abc", "def").(string))
}

func TestMatchAssign(t *testing.T) {
	fPln(MatchAssign("22", "1", "22", 3, 4, 5).(int))
	fPln(MatchAssign("22", "1", "222", 3, 4, 5).(int))
}

func TestTrueAssign(t *testing.T) {
	a := 1
	fPln(TrueAssign(a+1 == 2, a > 1, a > 0, "a", "b", "c", "dft"))
	fPln(TrueAssign(a+1 == 20, a > 1, a > 10, "a", "b", "c", "dft"))
}

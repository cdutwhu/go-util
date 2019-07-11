package util

import "testing"

func TestTerOp(t *testing.T) {
	fPln(IF(1 == 2, "abc", "def").(string))
}

func TestMatchAssign(t *testing.T) {
	fPln(MatchAssign("22", "1", "22", 3, 4).(int))
}

func TestConditionAssign(t *testing.T) {
	a := 1
	fPln(ConditionAssign(a+1 == 2, a > 1, a > 0, "a", "b", "c", "dft"))
}

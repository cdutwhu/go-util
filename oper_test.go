package util

import "testing"

func TestTerOp(t *testing.T) {
	fPln(TerOp(1 == 2, "abc", "def").(string))
}

func TestCaseAssign(t *testing.T) {
	fPln(CaseAssign("22", "1", "22", 3, 4).(int))
}

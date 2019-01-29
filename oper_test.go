package util

import "testing"

func TestTerOper(t *testing.T) {
	fPln(TerOper(1 == 2, "abc", "def").(string))
}

func TestCaseAssign(t *testing.T) {
	fPln(CaseAssign("22", "1", "22", 3, 4).(int))
}

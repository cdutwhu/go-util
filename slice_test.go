package util

import (
	"testing"
)

func TestSearch(t *testing.T) {
	arr := ToGA(1, 2, 3, 4, 5)
	r, i, ok := arr.Search(func(i int, a interface{}) bool { return i == 0 })
	fPln(r, i, ok)
}

func TestInsertBefore(t *testing.T) {
	arr := ToGA(1, 2, 3, 4, 5)
	r, _, _ := arr.InsertBefore(300, func(i int, a interface{}) bool { return i == 3 })
	fPln(r)
}

func TestInsertAfter(t *testing.T) {
	arr := ToGA(1, 2, 3, 4, 5)
	r, _, _ := arr.InsertAfter(44, func(i int, a interface{}) bool { return i == 4 })
	fPln(r)
}

func TestAdd(t *testing.T) {
	arr := ToGA(1, 2, 3, 4, 5)
	arr, i := arr.Add(600)
	fPln(arr, i)
}

func TestRemove(t *testing.T) {
	arr := ToGA(1, 2, 3, 4, 5)
	arr, _ = arr.Add(600)
	r, _, _, _ := arr.Remove(func(i int, a interface{}) bool { return a == 600 || i == 2 })
	fPln(r)
}

func TestReplace(t *testing.T) {
	arr := ToGA(1, 2, 3, 4, 5)
	arr, _ = arr.Add(600)
	r, _, _, _ := arr.Replace(func(i int, a interface{}) (interface{}, bool) { return a.(int) + 1, a == 600 || i == 2 })
	fPln(r)
}

func TestMoveItemAfter(t *testing.T) {
	arr := ToGA(1, 2, 3, 4, 5, 100)
	r, _, _, _ := arr.MoveItemAfter(
		func(move interface{}) bool { return move == 1 },
		func(after interface{}) bool { return after == 3 })
	fPln(r)
}

func TestRemoveRep(t *testing.T) {
	arr := ToGA("abc", "::", "abc", "de", "de")
	fPln(arr.RemoveRep())
}

func TestContain(t *testing.T) {
	arr := ToGA("::", "abc", "de", "mn", "")
	arr1 := ToGA("abc", "mn", "de", "::")
	fPln(arr.Contain(arr1))
}

func TestSeqContain(t *testing.T) {
	arr := ToGA("::", "abc", "de", "mn", "")
	arr1 := ToGA("abc", "de", "::")
	fPln(arr.SeqContain(arr1))
}

func TestAllAreIdentical(t *testing.T) {
	arr := ToGA("abc", "abc", "abc")
	fPln(arr.AllAreIdentical())
}

func TestInterSecANDUnion(t *testing.T) {
	arr := ToGA("::", "abc", "def", "mn", "A")
	r1 := arr.InterSec("abcd", "def", "::", "A")
	fPln(r1)
	r2 := arr.Union("abcd", "def", "::", "B")
	fPln(r2)
}

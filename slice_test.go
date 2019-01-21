package util

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := AI2AG(1, 2, 3, 4, 5)
	r, i, ok := arr.Search(func(i int, a interface{}) bool { return i == 0 })
	fmt.Println(r, i, ok)
}

func TestInsertBefore(t *testing.T) {
	arr := AI2AG(1, 2, 3, 4, 5)
	r, _, _ := arr.InsertBefore(300, func(i int, a interface{}) bool { return i == 3 })
	fmt.Println(r)
}

func TestInsertAfter(t *testing.T) {
	arr := AI2AG(1, 2, 3, 4, 5)
	r, _, _ := arr.InsertAfter(44, func(i int, a interface{}) bool { return i == 4 })
	fmt.Println(r)
}

func TestRemove(t *testing.T) {
	arr := AI2AG(1, 2, 3, 4, 5)
	r, _, _, _ := arr.Remove(func(i int, a interface{}) bool { return a == 1 || i == 3 })
	fmt.Println(r)
}

func TestMoveItemAfter(t *testing.T) {
	arr := AI2AG(1, 2, 3, 4, 5, 100)
	r, _, _, _ := arr.MoveItemAfter(
		func(move interface{}) bool { return move == 1 },
		func(after interface{}) bool { return after == 3 })
	fmt.Println(r)
}

func TestRemoveRepByLoop(t *testing.T) {
	arr := AS2AG("abc", "::", "abc", "de", "de")
	fmt.Println(arr.RemoveRep())
}

func TestContain(t *testing.T) {
	arr := AS2AG("::", "abc", "de", "mn", "")
	arr1 := AS2AG("abc", "mn", "de", "::")
	fPln(arr.Contain(arr1))
}

func TestSeqContain(t *testing.T) {
	arr := AS2AG("::", "abc", "de", "mn", "")
	arr1 := AS2AG("abc", "de", "::")
	fPln(arr.SeqContain(arr1))
}

func TestAllAreIdentical(t *testing.T) {
	arr := AS2AG("abc", "abc", "ab")
	fPln(arr.AllAreIdentical())
}

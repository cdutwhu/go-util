package util

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	r, i, ok := AI2AG(arr).Search(func(a interface{}) bool { return a == 10 })
	fmt.Println(r, i, ok)
}

func TestInsertBefore(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	r, _, _ := AI2AG(arr).InsertBefore(300, func(a interface{}) bool { return a == 1 })
	fmt.Println(r)
}

func TestInsertAfter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	r, _, _ := AI2AG(arr).InsertAfter(44, func(a interface{}) bool { return a == 5 })
	fmt.Println(r)
}

func TestRemove(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	r, _, _, _ := AI2AG(arr).Remove(func(a interface{}) bool { return a == 55 })
	fmt.Println(r)
}

func TestMoveItemAfter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 100}
	r, _, _, _ := AI2AG(arr).MoveItemAfter(
		func(move interface{}) bool { return move == 1 },
		func(after interface{}) bool { return after == 0 })
	fmt.Println(r)
}

func TestRemoveRepByLoop(t *testing.T) {
	arr := []string{"abc", "::", "abc", "de", "de"}
	fmt.Println(AS2AG(arr).RemoveRep())
}

func TestContain(t *testing.T) {
	arr := AS2AG([]string{"::", "abc", "de", "mn", ""})
	arr1 := AS2AG([]string{"abc", "mn", "de", "::"})
	fPln(arr.Contain(arr1))
}

func TestSeqContain(t *testing.T) {
	arr := AS2AG([]string{"::", "abc", "de", "mn", ""})
	arr1 := AS2AG([]string{"abc", "de", "::"})
	fPln(arr.SeqContain(arr1))
}

func TestAllAreIdentical(t *testing.T) {
	arr := AS2AG([]string{"abc", "abc", "abc"})
	fPln(arr.AllAreIdentical())
}

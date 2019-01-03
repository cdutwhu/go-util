package util

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	r, i, ok := Search(arr, func(a interface{}) bool {
		return a == 1
	})
	fmt.Println(r, i, ok)
}

func TestInsertBefore(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	b := 5
	InsertBefore(&arr, b, func(a interface{}) bool { return a.(int) == b-2 })
	fmt.Println(arr)
}

func TestInsertAfter(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	b := 5
	InsertAfter(&arr, b, func(a interface{}) bool { return a.(int) == b-2 })
	fmt.Println(arr)
}

func TestRemove(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	Remove(&arr, func(a interface{}) bool {
		return a == 3
	})
	fmt.Println(arr)
}

func TestMoveItemAfter(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5, 100}
	MoveItemAfter(&arr, func(after, move interface{}) bool { return after == 3 && move == 100 })
	fmt.Println(arr)
}

func TestRemoveRepByLoop(t *testing.T) {
	arr := []interface{}{"abc", "::", "abc", "de", "de"}
	fmt.Println(RemoveRep(arr))
}

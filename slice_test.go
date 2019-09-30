package util

import (
	"fmt"
	"testing"
)

func TestSliceCover(t *testing.T) {

	fmt.Println(XIn("11", []string{"a", "b", "1"}))
	fmt.Println(XIn("11", []string{"a", "b", "11"}))

	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{14, 15}
	ab := SliceAttach(a, b, 7).([]int)
	fmt.Println(ab)

	a = []int{1, 2, 3, 4, 5, 6, 7}
	b = []int{14, 15}
	ab = SliceAttach(a, b, 6).([]int)
	fmt.Println(ab)

	a = []int{1, 2}
	b = []int{13, 14, 15, 16}
	ab = SliceAttach(a, b, 1).([]int)
	fmt.Println(ab)

	A := []string{"a", "b"}
	B := []string{"1", "2", "3"}
	AB := SliceAttach(A, B, 1)
	fmt.Println(AB)

	A = []string{"a", "b", "c", "d"}
	B = []string{"1", "2", "3"}
	C := []string{"11"}
	ABC := SliceCover(A, B, C)
	fmt.Println(ABC)

	D := []string{"111"}
	DD := SliceCover(D)
	fmt.Println(DD)

	EE := SliceCover()
	fmt.Println(EE)
}

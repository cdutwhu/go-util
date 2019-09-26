package util

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]string{"a": "b", "c": "d", "e": "f"}
	fPln(MapKeys(m).([]string))
	m1 := map[int]string{1: "B", 2: "D", 3: "F"}
	fPln(MapKeys(m1).([]int))

	fPln(" ---------------------------------------------- ")

	k, v := MapKVs(m)
	K, V := k.([]string), v.([]string)
	fPln(K)
	fPln(V)

	fPln(" ---------------------------------------------- ")

	k, v = MapKVs(m1)
	K1, V1 := k.([]int), v.([]string)
	fPln(K1)
	fPln(V1)

	fPln(" ---------------------------------------------- ")

	// m2 := map[string]string{"aa": "bb", "cc": "dd", "ee": "ff"}
	m3 := map[int]string{4: "BB", 5: "DD", 1: "FF"}
	m02 := MapsJoin(m1, m3).(map[int]string)
	fPln(m02)

	m02 = MapsJoin(m3, m1).(map[int]string)
	fPln(m02)

	m4 := map[int]string{7: "BBB", 8: "DDD", 1: "FFF"}
	mm := MapsMerge(m1, m3, m4)
	fPln(mm)
}

func TestSliceCover(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{14, 15}
	ab := SliceCover(a, b).([]int)
	fmt.Println(ab)

	a = []int{1, 2}
	b = []int{13, 14, 15, 16}
	ab = SliceCover(a, b).([]int)
	fmt.Println(ab)

	// A := []string{"a", "b"}
	// B := []int{1, 2, 3}
	// AB := SliceCover(A, B)
	// fmt.Println(AB)
}

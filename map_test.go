package util

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]string{"zabc": "b", "f": "d", "e": "f"}
	fPln(MapKeys(m).([]string))
	m1 := map[int]string{11: "B", 2: "D", 3: "F", -1: "-1"}
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
	m3 := map[int]string{2: "B B   B C", -5: "DD", -1: "FF"}
	m13 := MapsJoin(m1, m3).(map[int]string)
	fPln(m13)

	fPln(" ********************************************** ")
	MapPrint(m13)

	m02 := MapsJoin(m3, m1).(map[int]string)
	fPln(m02)

	m4 := map[int]string{7: "BBB", 8: "DDD", 1: "FFF"}
	mm := MapsMerge(m1, m3, m4)
	fPln(mm)

	m4 = map[int]string{7: "BBB"}
	mm = MapsMerge(m4)
	fPln(mm)

	EE := MapsMerge()
	fPln(EE)
}

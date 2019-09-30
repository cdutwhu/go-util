package util

import (
	"math"
	ref "reflect"
)

// XIn :
func XIn(e, s interface{}) bool {
	v := ref.ValueOf(s)
	pc(v.Kind() != ref.Slice, fEf("s is NOT A SLICE!"))
	l := v.Len()
	for i := 0; i < l; i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

// SliceAttach :
func SliceAttach(s1, s2 interface{}, pos int) interface{} {
	v1, v2 := ref.ValueOf(s1), ref.ValueOf(s2)
	pc(v1.Kind() != ref.Slice, fEf("s1 is NOT A SLICE!"))
	pc(v2.Kind() != ref.Slice, fEf("s2 is NOT A SLICE!"))
	l1, l2 := v1.Len(), v2.Len()
	if l1 > 0 && l2 > 0 {
		if pos > l1 {
			return s1
		}
		lm := int(math.Max(float64(l1), float64(l2+pos)))
		v := ref.AppendSlice(v1.Slice(0, pos), v2)
		return v.Slice(0, lm).Interface()
	}
	if l1 > 0 && l2 == 0 {
		return s1
	}
	if l1 == 0 && l2 > 0 {
		return s2
	}
	return s1
}

// SliceCover :
func SliceCover(ss ...interface{}) interface{} {
	if len(ss) == 0 {
		return nil
	}
	attached := ss[0]
	for i, s := range ss {
		if i >= 1 {
			attached = SliceAttach(attached, s, 0)
		}
	}
	return attached
}

package util

import (
	"math"
	ref "reflect"
)

// MapKeys :
func MapKeys(m interface{}) interface{} {
	v := ref.ValueOf(m)
	PC(v.Kind() != ref.Map, fEf("NOT A MAP!"))
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := ref.TypeOf(keys[0].Interface())
		rst := ref.MakeSlice(ref.SliceOf(kType), L, L)
		for i, k := range keys {
			rst.Index(i).Set(ref.ValueOf(k.Interface()))
		}
		return rst.Interface()
	}
	return nil
}

// MapKVs :
func MapKVs(m interface{}) (interface{}, interface{}) {
	v := ref.ValueOf(m)
	PC(v.Kind() != ref.Map, fEf("NOT A MAP!"))
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := ref.TypeOf(keys[0].Interface())
		kRst := ref.MakeSlice(ref.SliceOf(kType), L, L)
		vType := ref.TypeOf(v.MapIndex(keys[0]).Interface())
		vRst := ref.MakeSlice(ref.SliceOf(vType), L, L)
		for i, k := range keys {
			kRst.Index(i).Set(ref.ValueOf(k.Interface()))
			vRst.Index(i).Set(ref.ValueOf(v.MapIndex(k).Interface()))
		}
		return kRst.Interface(), vRst.Interface()
	}
	return nil, nil
}

// MapsJoin : overwrited by the 2nd params
func MapsJoin(m1, m2 interface{}) interface{} {
	v1, v2 := ref.ValueOf(m1), ref.ValueOf(m2)
	PC(v1.Kind() != ref.Map, fEf("m1 is NOT A MAP!"))
	PC(v2.Kind() != ref.Map, fEf("m2 is NOT A MAP!"))
	keys1, keys2 := v1.MapKeys(), v2.MapKeys()
	if len(keys1) > 0 && len(keys2) > 0 {
		k1, k2 := keys1[0], keys2[0]
		k1Type, k2Type := ref.TypeOf(k1.Interface()), ref.TypeOf(k2.Interface())
		v1Type, v2Type := ref.TypeOf(v1.MapIndex(k1).Interface()), ref.TypeOf(v2.MapIndex(k2).Interface())
		PC(k1Type != k2Type, fEf("different maps' key type!"))
		PC(v1Type != v2Type, fEf("different maps' value type!"))
		aMap := ref.MakeMap(ref.MapOf(k1Type, v1Type))
		for _, k := range keys1 {
			aMap.SetMapIndex(ref.ValueOf(k.Interface()), ref.ValueOf(v1.MapIndex(k).Interface()))
		}
		for _, k := range keys2 {
			aMap.SetMapIndex(ref.ValueOf(k.Interface()), ref.ValueOf(v2.MapIndex(k).Interface()))
		}
		return aMap.Interface()
	}
	if len(keys1) > 0 && len(keys2) == 0 {
		return m1
	}
	if len(keys1) == 0 && len(keys2) > 0 {
		return m2
	}
	return m1
}

// MapsMerge : overwrited by the later params
func MapsMerge(ms ...interface{}) interface{} {
	if len(ms) == 1 {
		return ms[0]
	}
	mm := ms[0]
	for i, m := range ms {
		if i >= 1 {
			mm = MapsJoin(mm, m)
		}
	}
	return mm
}

// SliceCover :
func SliceCover(s1, s2 interface{}) interface{} {
	v1, v2 := ref.ValueOf(s1), ref.ValueOf(s2)
	PC(v1.Kind() != ref.Slice, fEf("s1 is NOT A SLICE!"))
	PC(v2.Kind() != ref.Slice, fEf("s2 is NOT A SLICE!"))
	l1, l2 := v1.Len(), v2.Len()
	if l1 > 0 && l2 > 0 {
		lm := int(math.Max(float64(l1), float64(l2)))
		v := ref.AppendSlice(v1.Slice(0, 0), v2)
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

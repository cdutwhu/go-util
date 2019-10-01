package util

import (
	ref "reflect"
	"regexp"
	"sort"
)

// MapPrint : Key Sorted Print
func MapPrint(m interface{}) {
	re := regexp.MustCompile(`^[+-]?[0-9]*\.?[0-9]+:`)
	mstr := fSp(m)
	mstr = mstr[4 : len(mstr)-1]
	fPln(mstr)
	I := 0
	rmIlist := []int{}
	ss := sSpl(mstr, " ")
	for i, s := range ss {
		if re.MatchString(s) {
			I = i
		} else {
			ss[I] += " " + s
			rmIlist = append(rmIlist, i) // to be deleted (i)
		}
	}
	for i, s := range ss {
		if !XIn(i, rmIlist) {
			fPln(i, s)
		}
	}
}

// MapKeys :
func MapKeys(m interface{}) interface{} {
	v := ref.ValueOf(m)
	pc(v.Kind() != ref.Map, fEf("NOT A MAP!"))
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := ref.TypeOf(keys[0].Interface())
		rstValue := ref.MakeSlice(ref.SliceOf(kType), L, L)
		for i, k := range keys {
			rstValue.Index(i).Set(ref.ValueOf(k.Interface()))
		}
		// sort keys if keys are int or float64 or string
		rst := rstValue.Interface()
		switch keys[0].Interface().(type) {
		case int:
			sort.Ints(rst.([]int))
		case float64:
			sort.Float64s(rst.([]float64))
		case string:
			sort.Strings(rst.([]string))
		}
		return rst
	}
	return nil
}

// MapKVs :
func MapKVs(m interface{}) (interface{}, interface{}) {
	v := ref.ValueOf(m)
	pc(v.Kind() != ref.Map, fEf("NOT A MAP!"))
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
	pc(v1.Kind() != ref.Map, fEf("m1 is NOT A MAP!"))
	pc(v2.Kind() != ref.Map, fEf("m2 is NOT A MAP!"))
	keys1, keys2 := v1.MapKeys(), v2.MapKeys()
	if len(keys1) > 0 && len(keys2) > 0 {
		k1, k2 := keys1[0], keys2[0]
		k1Type, k2Type := ref.TypeOf(k1.Interface()), ref.TypeOf(k2.Interface())
		v1Type, v2Type := ref.TypeOf(v1.MapIndex(k1).Interface()), ref.TypeOf(v2.MapIndex(k2).Interface())
		pc(k1Type != k2Type, fEf("different maps' key type!"))
		pc(v1Type != v2Type, fEf("different maps' value type!"))
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
	if len(ms) == 0 {
		return nil
	}
	mm := ms[0]
	for i, m := range ms {
		if i >= 1 {
			mm = MapsJoin(mm, m)
		}
	}
	return mm
}

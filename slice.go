package util

// Search :
func Search(arr []interface{}, check func(each interface{}) bool) (interface{}, int, bool) {
	for i, a := range arr {
		if check(a) {
			return a, i, true
		}
	}
	return nil, -1, false
}

// InsertBefore :
func InsertBefore(arr *[]interface{}, item interface{}, check func(each interface{}) bool) (int, bool) {
	for i, a := range *arr {
		if check(a) {
			*arr = append(*arr, 0)
			copy((*arr)[i+1:], (*arr)[i:])
			(*arr)[i] = item
			return i, true
		}
	}
	return -1, false
}

// InsertAfter :
func InsertAfter(arr *[]interface{}, item interface{}, check func(each interface{}) bool) (int, bool) {
	for i, a := range *arr {
		if check(a) {
			*arr = append(*arr, 0)
			copy((*arr)[i+2:], (*arr)[i+1:])
			(*arr)[i+1] = item
			return i, true
		}
	}
	return -1, false
}

// Remove :
func Remove(arr *[]interface{}, check func(each interface{}) bool) (interface{}, int, bool) {
	for i, a := range *arr {
		if check(a) {
			r := (*arr)[i]
			copy((*arr)[i:], (*arr)[i+1:])
			(*arr)[len(*arr)-1] = nil
			(*arr) = (*arr)[:len(*arr)-1]
			return r, i, true
		}
	}
	return nil, -1, false
}

// MoveItemAfter :
func MoveItemAfter(arr *[]interface{}, check func(after, move interface{}) bool) (int, bool) {
	for i, a0 := range *arr {
		for _, a1 := range *arr {
			if check(a0, a1) {
				if deleted, _, ok := Remove(arr, func(each interface{}) bool { return each == a1 }); ok {
					if _, ok = InsertAfter(arr, deleted, func(each interface{}) bool { return each == a0 }); ok {
						return i, true
					}
				}
			}
		}
	}
	return -1, false
}

// RemoveRep :
func RemoveRep(arr []interface{}) (result []interface{}) {
OUTER:
	for i := range arr {
		for j := range result {
			if arr[i] == result[j] {
				continue OUTER
			}
		}
		result = append(result, arr[i])
	}
	return result
}

// SA2GA :
func SA2GA(strs []string) (result []interface{}) {
	for _, s := range strs {
		result = append(result, s)
	}
	return
}

// GA2SA :
func GA2SA(arr []interface{}) (result []string) {
	for _, a := range arr {
		result = append(result, a.(string))
	}
	return
}

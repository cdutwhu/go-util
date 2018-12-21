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

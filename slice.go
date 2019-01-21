package util

// GArr : interface{} array
type GArr []interface{}

// // Search :
// func Search(arr []interface{}, check func(each interface{}) bool) (interface{}, int, bool) {
// 	for i, a := range arr {
// 		if check(a) {
// 			return a, i, true
// 		}
// 	}
// 	return nil, -1, false
// }

// InsertBefore :
// func InsertBefore(arr *[]interface{}, item interface{}, check func(each interface{}) bool) (int, bool) {
// 	for i, a := range *arr {
// 		if check(a) {
// 			*arr = append(*arr, 0)
// 			copy((*arr)[i+1:], (*arr)[i:])
// 			(*arr)[i] = item
// 			return i, true
// 		}
// 	}
// 	return -1, false
// }

// // InsertAfter :
// func InsertAfter(arr *[]interface{}, item interface{}, check func(each interface{}) bool) (int, bool) {
// 	for i, a := range *arr {
// 		if check(a) {
// 			*arr = append(*arr, 0)
// 			copy((*arr)[i+2:], (*arr)[i+1:])
// 			(*arr)[i+1] = item
// 			return i, true
// 		}
// 	}
// 	return -1, false
// }

// // Remove :
// func Remove(arr *[]interface{}, check func(each interface{}) bool) (interface{}, int, bool) {
// 	for i, a := range *arr {
// 		if check(a) {
// 			r := (*arr)[i]
// 			copy((*arr)[i:], (*arr)[i+1:])
// 			(*arr)[len(*arr)-1] = nil
// 			(*arr) = (*arr)[:len(*arr)-1]
// 			return r, i, true
// 		}
// 	}
// 	return nil, -1, false
// }

// MoveItemAfter :
// func MoveItemAfter(arr *[]interface{}, check func(after, move interface{}) bool) (int, bool) {
// 	for i, a0 := range *arr {
// 		for _, a1 := range *arr {
// 			if check(a0, a1) {
// 				if deleted, _, ok := Remove(arr, func(each interface{}) bool { return each == a1 }); ok {
// 					if _, ok = InsertAfter(arr, deleted, func(each interface{}) bool { return each == a0 }); ok {
// 						return i, true
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return -1, false
// }

// RemoveRep :
// func RemoveRep(arr []interface{}) (result []interface{}) {
// OUTER:
// 	for i := range arr {
// 		for j := range result {
// 			if arr[i] == result[j] {
// 				continue OUTER
// 			}
// 		}
// 		result = append(result, arr[i])
// 	}
// 	return result
// }

// func sg2ga(strs []string) (result []interface{}) {
// 	for _, s := range strs {
// 		result = append(result, s)
// 	}
// 	return
// }

// GA2SA : General Array to String Array
// func ga2sa(arr []interface{}) (result []string) {
// 	for _, a := range arr {
// 		result = append(result, a.(string))
// 	}
// 	return
// }

// L : ga's length
func (ga GArr) L() int {
	return len(ga)
}

// Search :
func (ga GArr) Search(check func(each interface{}) bool) (interface{}, int, bool) {
	for i, a := range ga {
		if check(a) {
			return a, i, true
		}
	}
	return nil, -1, false
}

// InsertBefore :
func (ga GArr) InsertBefore(item interface{}, check func(each interface{}) bool) (r []interface{}, idx int, did bool) {
	idx = -1
	for i, a := range ga {
		if check(a) {
			r, idx, did = append(r, item), i, true
		}
		r = append(r, a)
	}
	return
}

// InsertAfter :
func (ga GArr) InsertAfter(item interface{}, check func(each interface{}) bool) (r []interface{}, idx int, did bool) {
	idx = -1
	for i, a := range ga {
		r = append(r, a)
		if check(a) {
			r, idx, did = append(r, item), i, true
		}
	}
	return
}

// Remove :
func (ga GArr) Remove(check func(each interface{}) bool) (r []interface{}, del interface{}, idx int, did bool) {
	idx = -1
	for i, a := range ga {
		r = append(r, a)
		if check(a) {
			del, r, idx, did = a, r[:len(r)-1], i, true
		}
	}
	return
}

// MoveItemAfter :
func (ga GArr) MoveItemAfter(checkMove func(move interface{}) bool, checkAfter func(after interface{}) bool) (r []interface{}, idxMove, idxAfter int, did bool) {
	idxAfter, idxMove = -1, -1
	for i, a0 := range ga {
		if checkMove(a0) {
			if rst0, del, _, ok := ga.Remove(func(each interface{}) bool { return each == a0 }); ok {
				idxMove = i
				ga1 := GArr(rst0)
				for j, a1 := range ga1 {
					if checkAfter(a1) {
						if rst1, _, ok := ga1.InsertAfter(del, func(each interface{}) bool { return each == a1 }); ok {
							return rst1, i, j, true
						}
					}
				}
			}
		}
	}
	return ga, -1, -1, false
}

// RemoveRep :
func (ga GArr) RemoveRep() (r []interface{}) {
OUTER:
	for i := range ga {
		for j := range r {
			if ga[i] == r[j] {
				continue OUTER
			}
		}
		r = append(r, ga[i])
	}
	return r
}

// Contain :
func (ga GArr) Contain(arr GArr) bool {
	for _, a := range arr {
		if _, _, ok := ga.Search(func(each interface{}) bool { return each == a }); !ok {
			return false
		}
	}
	return true
}

// SeqContain :
func (ga GArr) SeqContain(arr GArr) bool {
	idxPrev := -1
	for _, a := range arr {
		if _, idx, ok := ga.Search(func(each interface{}) bool { return each == a }); ok {
			if idx > idxPrev {
				idxPrev = idx
				continue
			}
			return false
		}
		return false
	}
	return true
}

// AllAreIdentical :
func (ga GArr) AllAreIdentical() bool {
	if ga.L() > 1 {
		for _, a := range ga {
			if ga[0] != a {
				return false
			}
		}
	}
	return true
}

/**********************************************************/

// ToStrs :
func (ga GArr) ToStrs() (r []string) {
	for _, a := range ga {
		r = append(r, a.(string))
	}
	return
}

// ToInts :
func (ga GArr) ToInts() (r []int) {
	for _, a := range ga {
		r = append(r, a.(int))
	}
	return
}

/**********************************************************/

// AS2AG : String Array to General Array Type
func AS2AG(strs []string) (r GArr) {
	for _, s := range strs {
		r = append(r, s)
	}
	return
}

// AI2AG : Int Array to General Array Type
func AI2AG(ints []int) (r GArr) {
	for _, a := range ints {
		r = append(r, a)
	}
	return
}

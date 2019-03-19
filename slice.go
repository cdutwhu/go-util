package util

type (
	GArr []interface{}
	Strs []string
	I32s []int
	I64s []int64
	F32s []float32
	F64s []float64
	C32s []rune
)

func (ss Strs) ToG() (r GArr) {
	for _, s := range ss {
		r = append(r, s)
	}
	return
}

func (is I32s) ToG() (r GArr) {
	for _, i := range is {
		r = append(r, i)
	}
	return
}

func (is I64s) ToG() (r GArr) {
	for _, i := range is {
		r = append(r, i)
	}
	return
}

func (fs F32s) ToG() (r GArr) {
	for _, f := range fs {
		r = append(r, f)
	}
	return
}

func (fs F64s) ToG() (r GArr) {
	for _, f := range fs {
		r = append(r, f)
	}
	return
}

func (cs C32s) ToG() (r GArr) {
	for _, c := range cs {
		r = append(r, c)
	}
	return
}

/**********************************************************/

func ToGA(arrs ...interface{}) (r GArr) {
	for _, a := range arrs {
		r = append(r, a)
	}
	return
}

/*******************************************************/

// L : ga's length
func (ga GArr) L() int {
	return len(ga)
}

// Search :
func (ga GArr) Search(check func(idx int, each interface{}) bool) (interface{}, int, bool) {
	for i, a := range ga {
		if check(i, a) {
			return a, i, true
		}
	}
	return nil, -1, false
}

// InsertBefore :
func (ga GArr) InsertBefore(item interface{}, check func(idx int, each interface{}) bool) (r []interface{}, idx int, did bool) {
	idx = -1
	for i, a := range ga {
		if check(i, a) {
			r, idx, did = append(r, item), i, true
		}
		r = append(r, a)
	}
	return
}

// InsertAfter :
func (ga GArr) InsertAfter(item interface{}, check func(idx int, each interface{}) bool) (r []interface{}, idx int, did bool) {
	idx = -1
	for i, a := range ga {
		r = append(r, a)
		if check(i, a) {
			r, idx, did = append(r, item), i, true
		}
	}
	return
}

// Add :
func (ga GArr) Add(add interface{}) (r []interface{}, idx int) {
	for _, a := range ga {
		r = append(r, a)
	}
	return append(r, add), ga.L()
}

// Remove :
func (ga GArr) Remove(check func(idx int, each interface{}) bool) (r []interface{}, del interface{}, idx int, did bool) {
	idx = -1
	for i, a := range ga {
		r = append(r, a)
		if check(i, a) {
			del, r, idx, did = a, r[:len(r)-1], i, true
		}
	}
	return
}

// Replace :
func (ga GArr) Replace(check func(idx int, each interface{}) (interface{}, bool)) (r []interface{}, del interface{}, idx int, did bool) {
	idx = -1
	for i, a := range ga {
		r = append(r, a)
		if newItem, ok := check(i, a); ok {
			del, r, idx, did = a, r[:len(r)-1], i, true
			r = append(r, newItem)
		}
	}
	return
}

// MoveItemAfter :
func (ga GArr) MoveItemAfter(checkMove func(move interface{}) bool, checkAfter func(after interface{}) bool) (r []interface{}, idxMove, idxAfter int, did bool) {
	idxAfter, idxMove = -1, -1
	for i, a0 := range ga {
		if checkMove(a0) {
			if rst0, del, _, ok := ga.Remove(func(idx int, each interface{}) bool { return each == a0 }); ok {
				idxMove = i
				ga1 := GArr(rst0)
				for j, a1 := range ga1 {
					if checkAfter(a1) {
						if rst1, _, ok := ga1.InsertAfter(del, func(idx int, each interface{}) bool { return each == a1 }); ok {
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
		if _, _, ok := ga.Search(func(idx int, each interface{}) bool { return each == a }); !ok {
			return false
		}
	}
	return true
}

// SeqContain :
func (ga GArr) SeqContain(arr GArr) bool {
	idxPrev := -1
	for _, a := range arr {
		if _, idx, ok := ga.Search(func(idx int, each interface{}) bool { return each == a }); ok {
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

// InterSec :
func (ga GArr) InterSec(items ...interface{}) (r []interface{}) {
NEXT:
	for _, g := range ga {
		for _, item := range items {
			if g == item {
				r = append(r, g)
				continue NEXT
			}
		}
	}
	return
}

// Union :
func (ga GArr) Union(items ...interface{}) (r []interface{}) {
	for _, g := range ga {
		r = append(r, g)
	}
	rOri := make([]interface{}, len(r))
	copy(rOri, r)
NEXT:
	for _, item := range items {
		for _, rItem := range rOri {
			if item == rItem {
				continue NEXT
			}
		}
		r = append(r, item)
	}
	return
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

// ToInt64s :
func (ga GArr) ToInt64s() (r []int64) {
	for _, a := range ga {
		r = append(r, a.(int64))
	}
	return
}

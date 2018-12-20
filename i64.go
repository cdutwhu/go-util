package util

import "math"

// I64 is int 'class'
type I64 int64

// V : get int value
func (i *I64) V() int64 {
	return int64(*i)
}

// InArr : check if at least a same value exists in int array
func (i I64) InArr(arr ...int64) bool {
	for _, a := range arr {
		if i.V() == a {
			return true
		}
	}
	return false
}

// Nearest :
func (i I64) Nearest(arr ...int64) (int64, int) {
	minDis, minIdx := int64(MaxInt), -1
	for idx, a := range arr {
		dis := int64(math.Abs(float64(a - i.V())))
		if dis < minDis {
			minDis = dis
			minIdx = idx
		}
	}
	return arr[minIdx], minIdx
}

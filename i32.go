package util

//
// const MaxUint, MinUint, MaxInt, MinInt = ^uint(0), 0, int(MaxUint >> 1), -MaxInt - 1

const (
	// MaxUint : max uint
	MaxUint = ^uint(0)

	// MinUint : min uint
	MinUint = 0

	// MaxInt : max int
	MaxInt = int(MaxUint >> 1)

	// MinInt : min int
	MinInt = -MaxInt - 1
)

// I32 is int 'class'
type I32 int

// V : get int value
func (i *I32) V() int {
	return int(*i)
}

// InArr : check if at least a same value exists in int array
func (i I32) InArr(arr ...int) bool {
	for _, a := range arr {
		if i.V() == a {
			return true
		}
	}
	return false
}

// MinOf :
func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if i < min {
			min = i
		}
	}
	return min
}

// MinPosOf :
func MinPosOf(vars ...int) int {
	min := MaxInt
	for _, i := range vars {
		if i < min && i > 0 {
			min = i
		}
	}
	return TerOp(min != MaxInt, min, -1).(int)
}

// MinNoNegOf :
func MinNoNegOf(vars ...int) int {
	min := MaxInt
	for _, i := range vars {
		if i < min && i >= 0 {
			min = i
		}
	}
	return TerOp(min != MaxInt, min, -1).(int)
}

// MaxOf :
func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if i > max {
			max = i
		}
	}
	return max
}

// ToStr :
func (i I32) ToStr() string {
	return fSf("%d", i.V())
}

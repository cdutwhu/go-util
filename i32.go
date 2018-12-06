package util

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

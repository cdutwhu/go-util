package util

// C32 is rune 'class'
type C32 rune

// V : get rune value
func (c *C32) V() rune {
	return rune(*c)
}

// InArr : check if at least a same value exists in rune array
func (c C32) InArr(arr ...rune) bool {
	for _, a := range arr {
		if c.V() == a {
			return true
		}
	}
	return false
}

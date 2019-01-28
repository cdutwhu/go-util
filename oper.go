package util

// TerOper : Ternary Operator < ? : >
func TerOper(condition bool, block1, block2 interface{}) interface{} {
	if condition {
		return block1
	}
	return block2
}

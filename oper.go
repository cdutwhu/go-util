package util

// TerOp : Ternary Operator LIKE < ? : >, BUT NO S/C, so block1 and block2 MUST all valid. e.g. type assert, nil pointer, out of index
func TerOp(condition bool, block1, block2 interface{}) interface{} {
	if condition {
		return block1
	}
	return block2
}

// CaseAssign : NO S/C, MUST all valid, e.g. type assert, nil pointer, out of index
func CaseAssign(checkCasesValues ...interface{}) interface{} {
	l := len(checkCasesValues)
	PC(l < 3 || l%2 == 0, fEf("Invalid parameters"))
	_, l1, l2 := 1, (l-1)/2, (l-1)/2
	check := checkCasesValues[0]
	cases := checkCasesValues[1 : 1+l1]
	values := checkCasesValues[1+l1 : 1+l1+l2]
	for i, c := range cases {
		if check == c {
			return values[i]
		}
	}
	PC(true, fEf("Invalid parameters"))
	return nil
}

package util

// IF : Ternary Operator LIKE < ? : >, BUT NO S/C, so block1 and block2 MUST all valid. e.g. type assert, nil pointer, out of index
func IF(condition bool, block1, block2 interface{}) interface{} {
	if condition {
		return block1
	}
	return block2
}

// MatchAssign : NO ShortCut, MUST all valid, e.g. type assert, nil pointer, out of index
func MatchAssign(chkCasesValues ...interface{}) interface{} {
	l := len(chkCasesValues)
	PC(l < 4 || l%2 == 1, fEf("Invalid parameters"))
	_, l1, l2 := 1, (l-1)/2, (l-1)/2
	check := chkCasesValues[0]
	cases := chkCasesValues[1 : 1+l1]
	values := chkCasesValues[1+l1 : 1+l1+l2]
	for i, c := range cases {
		if check == c {
			return values[i]
		}
	}
	return chkCasesValues[l-1]
}

// TrueAssign : NO ShortCut, MUST all valid, e.g. type assert, nil pointer, out of index
func TrueAssign(condsValuesDft ...interface{}) interface{} {
	l := len(condsValuesDft)
	PC(l < 3 || l%2 == 0, fEf("Invalid parameters"))
	l1, l2 := l/2, l/2
	conds := condsValuesDft[0:l1]
	values := condsValuesDft[l1 : l1+l2]
	dft := condsValuesDft[l-1]
	for i, c := range conds {
		switch c.(type) {
		case bool:
			if c.(bool) == true {
				return values[i]
			}
		default:
			panic(fSf("The <%d> condition param : '%s' MUST be <bool> expression", i+1, c))
		}
	}
	return dft
}

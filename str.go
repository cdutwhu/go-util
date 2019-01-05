package util

import (
	"strings"
)

// QFlag : Flag for Quotes, single or double
type QFlag int

// BFlag : Flag for Brackets
type BFlag int

const (
	// QSingle : single quotes
	QSingle QFlag = 1
	// QDouble : double quotes
	QDouble QFlag = 2

	// BRound : round brackets
	BRound BFlag = 1
	// BBox : box brackets
	BBox BFlag = 2
	// BSquare : square brackets
	BSquare BFlag = 2
	// BCurly : curly brackets
	BCurly BFlag = 3
	// BAngle : angle brackets
	BAngle BFlag = 4
)

// Str is string 'class'
type Str string

// V : get string value
func (s *Str) V() string {
	return string(*s)
}

// L : get string length
func (s Str) L() int {
	return len(s.V())
}

// DefValue : if s is blank, assign it with input string value, otherwise keep its current value
func (s Str) DefValue(def string) string {
	if len(s) == 0 {
		return def
	}
	return s.V()
}

// HasAny :
func (s Str) HasAny(cks ...rune) bool {
	for _, c := range s.V() {
		for _, ck := range cks {
			if c == ck {
				return true
			}
		}
	}
	return false
}

// InArr : check if at least one same value exists in string array
func (s Str) InArr(arr ...string) (bool, int) {
	for i, a := range arr {
		if s.V() == a {
			return true, i
		}
	}
	return false, -1
}

// InMapSIKeys : check if at least one same value exists in string-key map
func (s Str) InMapSIKeys(m map[string]int) (bool, int) {
	for k, v := range m {
		if s.V() == k {
			return true, v
		}
	}
	return false, -1
}

// InMapSSValues : check if at least a same value exists in string-value map
func (s Str) InMapSSValues(m map[string]string) (bool, string) {
	for k, v := range m {
		if s.V() == v {
			return true, k
		}
	}
	return false, ""
}

// BeCoveredInMapSIKeys : check if at least one map(string)key value can cover the calling string
func (s Str) BeCoveredInMapSIKeys(m map[string]int) (bool, int) {
	for k, v := range m {
		if sC(k, s.V()) {
			return true, v
		}
	}
	return false, -1
}

// CoverAnyKeyInMapSI :
func (s Str) CoverAnyKeyInMapSI(m map[string]int) (bool, int) {
	for k, v := range m {
		if sC(s.V(), k) {
			return true, v
		}
	}
	return false, -1
}

// MakeBrackets :
func (s Str) MakeBrackets(f BFlag) string {
	bracketL, bracketR := ' ', ' '
	switch f {
	case BRound:
		bracketL, bracketR = '(', ')'
	case BBox:
		bracketL, bracketR = '[', ']'
	// case BSquare:
	// 	bracketL, bracketR = '[', ']'
	case BCurly:
		bracketL, bracketR = '{', '}'
	case BAngle:
		bracketL, bracketR = '<', '>'
	default:
		panic("error brackets flag")
	}
	if sHP(s.V(), string(bracketL)) && sHS(s.V(), string(bracketR)) {
		return s.V()
	}
	return string(bracketL) + s.V() + string(bracketR)
}

// RemoveBrackets :
func (s Str) RemoveBrackets() string {
	if (sHP(s.V(), "(") && sHS(s.V(), ")")) ||
		(sHP(s.V(), "[") && sHS(s.V(), "]")) ||
		(sHP(s.V(), "{") && sHS(s.V(), "}")) ||
		(sHP(s.V(), "<") && sHS(s.V(), ">")) {
		return s.V()[1 : len(s.V())-1]
	}
	return s.V()
}

// BracketsPos :
func (s Str) BracketsPos(f BFlag, level, index int) (left, right int) {
	bracketL, bracketR := ' ', ' '
	switch f {
	case BRound:
		bracketL, bracketR = '(', ')'
	case BBox:
		bracketL, bracketR = '[', ']'
	// case BSquare:
	// 	bracketL, bracketR = '[', ']'
	case BCurly:
		bracketL, bracketR = '{', '}'
	case BAngle:
		bracketL, bracketR = '<', '>'
	default:
		panic("error brackets flag")
	}

	curLevel, curIndex := 0, 0
	for i, c := range s.V() {
		if c == bracketL {
			curLevel++
		}
		if c == bracketR {
			curLevel--
		}
		if curLevel == level && c == bracketL {
			left = i
		}
		if curLevel == level-1 && c == bracketR {
			right = i
			curIndex++
			if curIndex == index {
				break
			}
		}
	}
	return
}

// BracketPairCount :
func (s Str) BracketPairCount(f BFlag) (count int) {
	bracketL, bracketR := ' ', ' '
	switch f {
	case BRound:
		bracketL, bracketR = '(', ')'
	case BBox:
		bracketL, bracketR = '[', ']'
	// case BSquare:
	// 	bracketL, bracketR = '[', ']'
	case BCurly:
		bracketL, bracketR = '{', '}'
	case BAngle:
		bracketL, bracketR = '<', '>'
	default:
		panic("error brackets flag")
	}

	level, inflag := 0, false
	for _, c := range s.V() {
		if c == bracketL {
			level++
		}
		if c == bracketR {
			level--
			if level == 0 {
				inflag = false
			}
		}
		if level == 1 {
			if !inflag {
				count++
				inflag = true
			}
		}
	}

	return count
}

// MakeQuotes :
func (s Str) MakeQuotes(f QFlag) string {
	quote := ""
	switch f {
	case QSingle:
		quote = "'"
	case QDouble:
		quote = "\""
	default:
		panic("error quotes flag")
	}
	if sHP(s.V(), quote) && sHS(s.V(), quote) {
		return s.V()
	}
	return quote + s.V() + quote
}

// RemoveQuotes : Remove single or double Quotes from a string. If no quotations, do nothing
func (s Str) RemoveQuotes() string {
	if (sHP(s.V(), "\"") && sHS(s.V(), "\"")) ||
		(sHP(s.V(), "'") && sHS(s.V(), "'")) {
		return s.V()[1 : len(s.V())-1]
	}
	return s.V()
}

// MakePrefix :
func (s Str) MakePrefix(prefix string) string {
	if !sHP(s.V(), prefix) {
		return prefix + s.V()
	}
	return s.V()
}

// RemovePrefix :
func (s Str) RemovePrefix(prefix string) string {
	if sHP(s.V(), prefix) {
		return s.V()[len(prefix):len(s.V())]
	}
	return s.V()
}

// MakeSuffix :
func (s Str) MakeSuffix(suffix string) string {
	if !sHS(s.V(), suffix) {
		return s.V() + suffix
	}
	return s.V()
}

// RemoveSuffix :
func (s Str) RemoveSuffix(suffix string) string {
	if sHS(s.V(), suffix) {
		return s.V()[:len(s.V())-len(suffix)]
	}
	return s.V()
}

// RemoveTailFromLast :
func (s Str) RemoveTailFromLast(tail string) string {
	if i := sLI(s.V(), tail); i >= 0 {
		return s.V()[:i]
	}
	return s.V()
}

// RemoveBlankBefore :
func (s Str) RemoveBlankBefore(strs ...string) string {
	whole := s.V()
	for _, str := range strs {
		str0, str1 := " "+str, "\t"+str
	NEXT:
		if p := sI(whole, str0); p >= 0 {
			whole = whole[:p] + whole[p+1:]
			goto NEXT
		}
		if p := sI(whole, str1); p >= 0 {
			whole = whole[:p] + whole[p+1:]
			goto NEXT
		}
	}
	return whole
}

// RemoveBlankAfter :
func (s Str) RemoveBlankAfter(strs ...string) string {
	whole := s.V()
	for _, str := range strs {
		str0, str1 := str+" ", str+"\t"
	NEXT:
		if p := sI(whole, str0); p >= 0 {
			whole = whole[:p+len(str0)-1] + whole[p+len(str0):]
			goto NEXT
		}
		if p := sI(whole, str1); p >= 0 {
			whole = whole[:p+len(str0)-1] + whole[p+len(str0):]
			goto NEXT
		}
	}
	return whole
}

// RemoveBlankNear :
func (s Str) RemoveBlankNear(strs ...string) string {
	s0 := s.RemoveBlankBefore(strs...)
	return Str(s0).RemoveBlankAfter(strs...)
}

// RemoveBlankNBefore :
func (s Str) RemoveBlankNBefore(n int, str string) string {
	// whole, left, right, strs := s.V(), "", "", []string{}
	// for i := 0; i < n; i++ {
	// 	if p := sI(whole, str); p >= 0 {
	// 		left, right = whole[:p+1], whole[p+1:]
	// 		left, whole = Str(left).RemoveBlankBefore(str), right
	// 		strs = append(strs, left)
	// 		if i == n-1 {
	// 			strs = append(strs, right)
	// 		}
	// 	} else {
	// 		if right != "" {
	// 			strs = append(strs, right)
	// 		}
	// 		break
	// 	}
	// }
	// return sJ(strs, "")

	segs, strs := sS(s.V(), str), []string{}
	for i, seg := range segs {
		if i == len(segs)-1 {
			strs = append(strs, seg)
			break
		}
		if i >= 0 && i < n {
			seg = sTR(seg, " \t")
		}
		strs = append(strs, seg)
	}
	return sJ(strs, str)
}

// RemoveBlankNAfter :
func (s Str) RemoveBlankNAfter(n int, str string) string {
	strs := []string{}
	for i, seg := range sS(s.V(), str) {
		if i == 0 {
			strs = append(strs, seg)
			continue
		}
		if i >= 1 && i <= n {
			seg = sTL(seg, " \t")
		}
		strs = append(strs, seg)
	}
	return sJ(strs, str)
}

// RemoveBlankNNear :
func (s Str) RemoveBlankNNear(n int, str string) string {
	s0 := s.RemoveBlankNBefore(n, str)
	return Str(s0).RemoveBlankNAfter(n, str)
}

// KeyValueMap :
func (s Str) KeyValueMap(delimiter, assign, terminator rune) (r map[string]string) {
	r = make(map[string]string)
	str := s.RemoveBlankNear(string(assign))
	if pt := sI(str, string(terminator)); pt > 0 {
		str = str[:pt]
	}
	for _, kv := range strings.FieldsFunc(str, func(c rune) bool { return c == delimiter }) {
		if strings.Contains(kv, string(assign)) {
			kvpair := sS(kv, string(assign))
			r[kvpair[0]] = Str(kvpair[1]).RemoveQuotes()
		}
	}
	return
}

// KeyValuePair :
func (s Str) KeyValuePair(assign, terminatorK, terminatorV rune, rmQuotes, trimBlank bool) (k, v string) {
	str := s.RemoveBlankNNear(1, string(assign))
	if p := sI(str, string(assign)); p >= 0 {
		k, v = str[:p], str[p+1:]
		if pk := sLI(k, string(terminatorK)); pk >= 0 {
			k = str[pk+1 : p]
		}
		if pv := sI(v, string(terminatorV)); pv >= 0 {
			v = str[p+1 : p+1+pv]
		}
	}
	if trimBlank {
		k, v = sT(k, " \t"), sT(v, " \t")
	}
	if rmQuotes {
		k, v = Str(k).RemoveQuotes(), Str(v).RemoveQuotes()
	}
	return
}

// LooseSearch :
// func (s Str) LooseSearch(aim string, ignore ...rune) (bool, int) {
// 	if len(aim) != 2 {
// 		pln("<aim> should only have 2 chars")
// 		return false, -1
// 	}
// }

// AllAreIdentical : check all the input strings are identical
func AllAreIdentical(arr ...string) bool {
	if len(arr) > 1 {
		for _, a := range arr {
			if arr[0] != a {
				return false
			}
		}
	}
	return true
}

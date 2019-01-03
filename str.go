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

// BeCoveredInMapSIKeys : check if at least one map(string)key value can cover the calling string
func (s Str) BeCoveredInMapSIKeys(m map[string]int) (bool, int) {
	for k, v := range m {
		if sI(k, s.V()) >= 0 {
			return true, v
		}
	}
	return false, -1
}

// CoverAnyKeyInMapSI :
func (s Str) CoverAnyKeyInMapSI(m map[string]int) (bool, int) {
	for k, v := range m {
		if sI(s.V(), k) >= 0 {
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

// KeyValueMap :
func (s Str) KeyValueMap(delimiter, assign, terminator rune) (r map[string]string) {
	r = make(map[string]string)
	str := s.RemoveBlankNear(string(assign))
	if pt := sI(str, string(terminator)); pt > 0 {
		str = str[:pt]
	}
	for _, kv := range strings.FieldsFunc(str, func(c rune) bool { return c == delimiter }) {
		if strings.Contains(kv, string(assign)) {
			kvpair := strings.Split(kv, string(assign))
			r[kvpair[0]] = Str(kvpair[1]).RemoveQuotes()
		}
	}
	return
}

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

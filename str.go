package util

import (
	"encoding/json"

	"github.com/google/uuid"
)

type (
	// QFlag : Flag for Quotes, single or double
	QFlag int
	// BFlag : Flag for Brackets
	BFlag int
)

const (

	// QSingle : single quotes ''
	QSingle QFlag = 1
	// QDouble : double quotes ""
	QDouble QFlag = 2

	// BRound : round brackets ()
	BRound BFlag = 1
	// BBox : box brackets []
	BBox BFlag = 2
	// BSquare : square brackets []
	BSquare BFlag = 2
	// BCurly : curly brackets {}
	BCurly BFlag = 3
	// BAngle : angle brackets <>
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

// ToInt64 :
func (s Str) ToInt64() int64 {
	return Must(sc2Int(s.V(), 10, 64)).(int64)
}

// ToInt :
func (s Str) ToInt() int {
	return int(Must(sc2Int(s.V(), 10, 64)).(int64))
}

// DefValue : if s is blank, assign it with input string value, otherwise keep its current value
func (s Str) DefValue(def string) string {
	if len(s) == 0 {
		return def
	}
	return s.V()
}

// Repeat : e.g. "ABC"(2) => "ABCABC"
func (s Str) Repeat(n int) (r string) {
	for i := 0; i < n; i++ {
		r += s.V()
	}
	return r
}

// HasAny : e.g. "ABC"('A', 'M') => true
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

// IsMadeFrom : e.g. "ABC"('C','B','A','D') => true
func (s Str) IsMadeFrom(chars ...rune) bool {
NEXT:
	for _, c := range s.V() {
		for _, ck := range chars {
			if c == ck {
				continue NEXT
			}
		}
		return false
	}
	return true
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
		if sCtn(k, s.V()) {
			return true, v
		}
	}
	return false, -1
}

// CoverAnyKeyInMapSI :
func (s Str) CoverAnyKeyInMapSI(m map[string]int) (bool, int) {
	for k, v := range m {
		if sCtn(s.V(), k) {
			return true, v
		}
	}
	return false, -1
}

// MkBrackets : e.g. "ABC"(BRound) => "(ABC)"
func (s Str) MkBrackets(f BFlag) string {
	bracketL := CaseAssign(f, BRound, BBox, BSquare, BCurly, BAngle, "(", "[", "[", "{", "<").(string)
	bracketR := CaseAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ")", "]", "]", "}", ">").(string)
	if sHP(s.V(), bracketL) && sHS(s.V(), bracketR) {
		return s.V()
	}
	return bracketL + s.V() + bracketR
}

// RmBrackets : e.g. "(ABC)" => "ABC"
func (s Str) RmBrackets() string {
	if (sHP(s.V(), "(") && sHS(s.V(), ")")) ||
		(sHP(s.V(), "[") && sHS(s.V(), "]")) ||
		(sHP(s.V(), "{") && sHS(s.V(), "}")) ||
		(sHP(s.V(), "<") && sHS(s.V(), ">")) {
		return s.V()[1 : len(s.V())-1]
	}
	return s.V()
}

// QuotesPos : index from 1
func (s Str) QuotesPos(f QFlag, index int) (str string, left, right int) {
	quote := CaseAssign(f, QSingle, QDouble, '\'', '"').(rune)
	cnt, left, right := 0, -1, -1
	for i, c := range s.V() {
		if left != -1 && right != -1 {
			break
		}
		if c == quote {
			cnt++
		}
		if (cnt+1)/2 == index {
			left = TerOp(cnt%2 == 1 && left == -1, i, left).(int)
			right = TerOp(cnt%2 == 0 && right == -1, i, right).(int)
		}
	}
	return s.V()[left : right+1], left, right
}

// BracketsPos : level from 1, index from 1, if index > count, get the last one
func (s Str) BracketsPos(f BFlag, level, index int) (str string, left, right int) {
	bracketL := CaseAssign(f, BRound, BBox, BSquare, BCurly, BAngle, '(', '[', '[', '{', '<').(rune)
	bracketR := CaseAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ')', ']', ']', '}', '>').(rune)

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
	return s.V()[left : right+1], left, right
}

// BracketPairCount :
func (s Str) BracketPairCount(f BFlag) (count int) {
	bracketL := CaseAssign(f, BRound, BBox, BSquare, BCurly, BAngle, '(', '[', '[', '{', '<').(rune)
	bracketR := CaseAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ')', ']', ']', '}', '>').(rune)
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

// BracketDepth :
func (s Str) BracketDepth(f BFlag, pos int) int {
	if pos >= s.L() {
		return -1
	}
	bracketL := CaseAssign(f, BRound, BBox, BSquare, BCurly, BAngle, '(', '[', '[', '{', '<').(rune)
	bracketR := CaseAssign(f, BRound, BBox, BSquare, BCurly, BAngle, ')', ']', ']', '}', '>').(rune)
	level := 0
	for i, c := range s.V() {
		if c == bracketL {
			level++
		}
		if i == pos {
			break
		}
		if c == bracketR {
			level--
		}
	}
	return level
}

// MkQuotes : e.g. "ABC"(QSingle) => "'ABC'"
func (s Str) MkQuotes(f QFlag) string {
	quote := CaseAssign(f, QSingle, QDouble, "'", "\"").(string)
	if sHP(s.V(), quote) && sHS(s.V(), quote) {
		return s.V()
	}
	return quote + s.V() + quote
}

// RmQuotes : Remove single or double Quotes from a string. If no quotations, do nothing
func (s Str) RmQuotes() string {
	if (sHP(s.V(), "\"") && sHS(s.V(), "\"")) ||
		(sHP(s.V(), "'") && sHS(s.V(), "'")) {
		return s.V()[1 : len(s.V())-1]
	}
	return s.V()
}

// MkPrefix : e.g. "ABC"("abc") => "abcABC"
func (s Str) MkPrefix(prefix string) string {
	if !sHP(s.V(), prefix) {
		return prefix + s.V()
	}
	return s.V()
}

// RmPrefix : e.g. "abcABC"("abc") => "ABC"
func (s Str) RmPrefix(prefix string) string {
	if sHP(s.V(), prefix) {
		return s.V()[len(prefix):len(s.V())]
	}
	return s.V()
}

// MkSuffix : e.g. "ABC"("abc") => "ABCabc"
func (s Str) MkSuffix(suffix string) string {
	if !sHS(s.V(), suffix) {
		return s.V() + suffix
	}
	return s.V()
}

// RmSuffix : e.g. "ABCabc"("abc") => "ABC"
func (s Str) RmSuffix(suffix string) string {
	if sHS(s.V(), suffix) {
		return s.V()[:len(s.V())-len(suffix)]
	}
	return s.V()
}

// RmTailFromLast : e.g. "AB.CD.EF"(".") => "AB.CD"
func (s Str) RmTailFromLast(tail string) string {
	if i := sLI(s.V(), tail); i >= 0 {
		return s.V()[:i]
	}
	return s.V()
}

// RmBlankBefore :
func (s Str) RmBlankBefore(strs ...string) string {
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

// RmBlankAfter :
func (s Str) RmBlankAfter(strs ...string) string {
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

// RmBlankNear :
func (s Str) RmBlankNear(strs ...string) string {
	s0 := s.RmBlankBefore(strs...)
	return Str(s0).RmBlankAfter(strs...)
}

// RmBlankNBefore :
func (s Str) RmBlankNBefore(n int, str string) string {
	// whole, left, right, strs := s.V(), "", "", []string{}
	// for i := 0; i < n; i++ {
	// 	if p := sI(whole, str); p >= 0 {
	// 		left, right = whole[:p+1], whole[p+1:]
	// 		left, whole = Str(left).RmBlankBefore(str), right
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

	segs, strs := sSpl(s.V(), str), []string{}
	for i, seg := range segs {
		if i < n {
			seg = sTR(seg, " \t")
		}
		strs = append(strs, seg)
	}
	return sJ(strs, str)
}

// RmBlankNAfter :
func (s Str) RmBlankNAfter(n int, str string) string {
	strs := []string{}
	for i, seg := range sSpl(s.V(), str) {
		if i >= 1 && i <= n {
			seg = sTL(seg, " \t")
		}
		strs = append(strs, seg)
	}
	return sJ(strs, str)
}

// RmBlankNNear :
func (s Str) RmBlankNNear(n int, str string) string {
	// s0 := s.RmBlankNBefore(n, str)
	// return Str(s0).RmBlankNAfter(n, str)

	segs, strs := sSpl(s.V(), str), []string{}
	for i, seg := range segs {
		if i == 0 && i != n {
			seg = sTR(seg, " \t")
		} else if i == n {
			seg = sTL(seg, " \t")
		} else if i >= 1 && i < n {
			seg = sT(seg, " \t")
		}
		strs = append(strs, seg)
	}
	return sJ(strs, str)
}

// TrimInternal :
func (s Str) TrimInternal(cutset rune, nkeep int) (r string) {
	pos, lens, strs := []int{}, []int{}, []string{}
	for p, c := range s.V() {
		if p < s.L()-1 {
			cNext := rune(s.V()[p+1])
			if c != cutset && cNext == cutset {
				pos = append(pos, p+1)
			}
		}
	}
NEXT:
	for _, p := range pos {
		l := s.V()[p:]
		for i, c := range l {
			if c != cutset {
				lens = append(lens, MinOf(i, nkeep))
				continue NEXT
			}
		}
	}
	for _, str := range sFF(s.V(), func(c rune) bool { return c == cutset }) {
		strs = append(strs, str)
	}
	cntL, cntR := 0, 0
	for p, c := range s.V() {
		if c != cutset {
			cntL = p
			break
		}
	}
	for p := s.L() - 1; p >= 0; p-- {
		if rune(s.V()[p]) != cutset {
			cntR = s.L() - p - 1
			break
		}
	}

	r += Str(cutset).Repeat(cntL)
	for i, str := range strs {
		r += str
		if i < len(strs)-1 {
			r += Str(cutset).Repeat(lens[i])
		}
	}
	r += Str(cutset).Repeat(cntR)
	return r
}

// TrimInternalEachLine :
func (s Str) TrimInternalEachLine(cutset rune, nkeep int) (r string) {
	strs := []string{}
	lns := sFF(s.V(), func(c rune) bool { return c == '\n' })
	for i, ln := range lns {
		strs = append(strs, Str(ln).TrimInternal(cutset, nkeep))
		if i < len(lns)-1 {
			strs = append(strs, "\n")
		}
	}
	if s.V()[s.L()-1] == '\n' {
		return sJ(strs, "") + "\n"
	}
	return sJ(strs, "")
}

// TrimAllInternal :
func (s Str) TrimAllInternal(cutset string) (r string) {
	cuts := []rune(cutset)
	for _, c := range cuts {
		r = s.TrimInternal(c, 0)
		s = Str(r)
	}
	return
}

// KeyValueMap :
func (s Str) KeyValueMap(delimiter, assign, terminator rune) (r map[string]string) {
	r = make(map[string]string)
	str := s.RmBlankNear(string(assign))
	if pt := sI(str, string(terminator)); pt > 0 {
		str = str[:pt]
	}
	for _, kv := range sFF(str, func(c rune) bool { return c == delimiter }) {
		if sCtn(kv, string(assign)) {
			kvpair := sSpl(kv, string(assign))
			r[kvpair[0]] = Str(kvpair[1]).RmQuotes()
		}
	}
	return
}

// KeyValuePair : (if assign mark cannot be found, k is empty, v is original string)
func (s Str) KeyValuePair(assign string, terminatorK, terminatorV rune, rmQuotes, trimBlank bool) (k, v string) {
	str := s.RmBlankNNear(1, assign)
	if p := sI(str, assign); p >= 0 {
		k, v = str[:p], str[p+len(assign):]
		if pk := sLI(k, string(terminatorK)); pk >= 0 {
			k = str[pk+1 : p]
		}
		if pv := sI(v, string(terminatorV)); pv >= 0 {
			v = str[p+len(assign) : p+len(assign)+pv]
		}
	} else {
		return "", s.V()
	}
	if trimBlank {
		k, v = sT(k, " \t"), sT(v, " \t")
	}
	if rmQuotes {
		k, v = Str(k).RmQuotes(), Str(v).RmQuotes()
	}
	return
}

// looseSearch2Strs :
func (s Str) looseSearch2Strs(first, second string, ignore ...rune) (bool, int, int) {
	fPosList, sPosList := s.Indices(first), s.Indices(second)
	fEndPosList := []int{}
	for _, fp := range fPosList {
		fEndPosList = append(fEndPosList, fp+len(first))
	}
	for i, fpe := range fEndPosList {
		for j, sp := range sPosList {
			if fpe < sp {
				if Str(s.V()[fpe+1 : sp]).IsMadeFrom(ignore...) {
					return true, fPosList[i], sPosList[j]
				}
			}
		}
	}
	return false, -1, -1
}

// LooseSearchStrs : last param is ignored runes string
func (s Str) LooseSearchStrs(aims ...string) (ok bool, findpos int) {
	PC(len(aims) < 3, fEf("At least 3 params, the last is ignored runes string"))
	ignored, fst, scd, offsets := []rune(aims[len(aims)-1]), -1, -1, []int{}

AGAIN:
	prevFst, prevScd := -1, -1
	for i := 0; i < len(aims)-2; i++ {
		if ok, fst, scd = s.looseSearch2Strs(aims[i], aims[i+1], ignored...); ok {
			if fst != prevScd && prevScd != -1 {
				s = s[prevFst+1:]
				offsets = append(offsets, prevFst+1)
				goto AGAIN
			}
			if i == 0 {
				offsum := 0
				for _, offset := range offsets {
					offsum += offset
				}
				findpos = fst + offsum
			}
			prevFst, prevScd = fst, scd
		} else {
			return false, -1
		}
	}
	return
}

// looseSearch2Chars :
func (s Str) looseSearch2Chars(first, second rune, ignore ...rune) (bool, int, int) {
	for p, c := range s.V() {
		if c == first && p < s.L()-1 {
			if pe := sI(s.V()[p+1:], string(second)); pe >= 0 {
				if Str(s.V()[p+1 : p+1+pe]).IsMadeFrom(ignore...) {
					return true, p, p + pe + 1
				}
			}
		}
	}
	return false, -1, -1
}

// LooseSearchChars :
func (s Str) LooseSearchChars(aim string, ignore ...rune) (ok bool, findpos int) {
	if len(aim) == 1 {
		if p := sI(s.V(), string(aim[0])); p >= 0 {
			return true, p
		}
		return false, -1
	}

	fst, scd, offsets := -1, -1, []int{}
AGAIN:
	prevFst, prevScd := -1, -1
	for i := 0; i < len(aim)-1; i++ {
		if ok, fst, scd = s.looseSearch2Chars(rune(aim[i]), rune(aim[i+1]), ignore...); ok {
			if fst != prevScd && prevScd != -1 {
				s = s[prevFst+1:]
				offsets = append(offsets, prevFst+1)
				goto AGAIN
			}
			if i == 0 {
				offsum := 0
				for _, offset := range offsets {
					offsum += offset
				}
				findpos = fst + offsum
			}
			prevFst, prevScd = fst, scd
		} else {
			return false, -1
		}
	}
	return
}

// Indices :
func (s Str) Indices(aim string) (posList []int) {
	str := s.V()
AGAIN:
	if p := sI(str, aim); p >= 0 {
		pos := p
		if len(posList) > 0 {
			pos = p + posList[len(posList)-1] + len(aim)
		}
		posList = append(posList, pos)
		str = str[p+len(aim):]
		goto AGAIN
	}
	return
}

// IsXMLSegSimple :
func (s Str) IsXMLSegSimple() bool {
	c := s.BracketPairCount(BAngle)
	if c == 0 {
		return false
	}
	tagsStr, _, _ := s.BracketsPos(BAngle, 1, 1)
	tageStr, _, _ := s.BracketsPos(BAngle, 1, c)
	tage := tageStr[2 : len(tageStr)-1]
	tags := tagsStr[1 : 1+len(tage)]
	return tags == tage &&
		(tagsStr[len(tags)+1] == ' ' || tagsStr[len(tags)+1] == '>') &&
		tagsStr[0] == '<' && tageStr[:2] == "</"
}

// IsJSON :
func (s Str) IsJSON() bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s.V()), &js) == nil
}

// IsUUID :
func (s Str) IsUUID() bool {
	_, e := uuid.Parse(s.V())
	return e == nil
}

// FieldsSeqContain :
func (s Str) FieldsSeqContain(str, sep string) bool {
	sArr0, sArr1 := sSpl(s.V(), sep), sSpl(str, sep)
	return Strs(sArr0).ToG().SeqContain(Strs(sArr1).ToG())
}

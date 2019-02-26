package util

// JSONChildValue : s has "{ }" wrapper. idx from 1 and only one value. If child is array, idx is ignored.
func (s Str) JSONChildValue(child string, idx ...int) (content string, pos int) {

	PC(!Str(s.MkBrackets(BCurly)).IsJSON(), fEf("Invalid JSON String"))

	json, child := s.V(), Str(child).MkQuotes(QDouble)+":"
AGAIN:
	if pos = sI(json, child); pos > 0 { // *** General Found, including nested ***
		above := json[:pos]
		if sCnt(above, "{")-sCnt(above, "}") == 1 { // *** FOUND ( Object OR Value ) ***

			if ok, pchk := s[pos:].LooseSearchChars(":{", ' ', '\t', '\n'); ok && pchk == len(child)-1 { //         *** (Object) ***
				content, lb, _ := s[pos:].BracketsPos(BCurly, 1, 1)
				return content, pos + lb
			} else if ok, pchk := s[pos:].LooseSearchChars(":\"", ' ', '\t', '\n'); ok && pchk == len(child)-1 { // *** (Value) ***
				pos += len(child)
				content, lq, _ := s[pos:].QuotesPos(QDouble, 1)
				return content, pos + lq
			} else if ok, pchk := s[pos:].LooseSearchChars(":[", ' ', '\t', '\n'); ok && pchk == len(child)-1 { //  *** (Array) (SAME type in array) ***
				i := 1
				if len(idx) > 0 {
					i = idx[0]					
				}

				content, _, _ = s[pos:].BracketsPos(BBox, 1, 1)
				content = sT(Str(content).RmBrackets(), " \t\n\r")
				if Str(content).TrimAllInternal(" \t\r\n") == "" { //          *** empty element ***
					return "", 0
				}

				nCurlyPair := s[pos:].BracketPairCount(BCurly)

				if nCurlyPair == 0 { //                                        *** All plain values ***
					ss := sSpl(content, ",")
					i = TerOp(i > len(ss), len(ss), i).(int)
					i = TerOp(i < 1, 1, i).(int)
					v, posTrue := sT(ss[i-1], " \t\n\r"), 0
					for _, p := range s[pos:].Indices(v) {
						if sCnt(s.V()[pos:pos+p], ",") == i-1 {
							posTrue = pos + p
							break
						}
					}
					return v, posTrue
				}

				for j := 1; j <= nCurlyPair; j++ {
					idx := fSf("#%d", j)
					contentObj, _, _ := Str(content).BracketsPos(BCurly, 1, j)
					content = sRep(content, contentObj, idx, 1)
				}

				ss := sSpl(content, ",")
				i = TerOp(i > len(ss), len(ss), i).(int)
				i = TerOp(i < 1, 1, i).(int)
				v, posTrue := sT(ss[i-1], " \t\n\r"), 0

				if posList := s[pos:].Indices(v); len(posList) > 0 { //        *** for plain value ***
					for _, p := range posList {
						if sCnt(s.V()[pos:pos+p], ",") == i-1 {
							posTrue = pos + p
							break
						}
					}
					return v, posTrue
				}

				//                                                             *** for #Obj value ***
				iObj := 0
				for iEachEle, eachEle := range ss {
					if sHP(sT(eachEle, " \t\n\r"), "#") {
						iObj++
					}
					if iEachEle+1 == i {
						break
					}
				}
				content, lb, _ := s[pos:].BracketsPos(BCurly, 1, iObj)
				return content, pos + lb
			}
		}

		// *** FAKE FOUND ***
		json = json[:pos] + "\"*" + json[pos+2:]
		goto AGAIN
	}
	return "", 0
}

// JSONXPathValue : s has "{ }" wrapper.  idx from 1 and only one value
func (s Str) JSONXPathValue(xpath, del string, idx ...int) (content string, posStart, posEnd int) {
	posEach, _ := 0, ""
	for _, seg := range sSpl(xpath, del) {
		s = TerOp(content != "", Str(content), s).(Str)
		content, posEach = s.JSONChildValue(seg, idx...)
		if content == "" {
			return
		}
		if posEach == -1 {
			posStart, posEach = -1, -1
			return
		}
		posStart += posEach
	}
	posEnd = posStart + len(content) - 1
	return
}

// JSONBuild :
func (s Str) JSONBuild(xpath, del string, idx int, property, value string) (string, bool) {
	if sT(s.V(), " \t") == "" {
		s = Str("{}")
	}

	property = Str(property).MkQuotes(QDouble) + ": "
	if value[0] != '{' && value[0] != '[' {
		value = Str(value).MkQuotes(QDouble)
	}

	if s == "{}" {
		s = Str(fSf(`{%s %s}`, property, value))
		return s.V(), s.IsJSON()
	}

	if _, start, end := s.JSONXPathValue(xpath, del, idx); start != -1 {
		left, right := s.V()[:end], s.V()[end:]
		if !sHS(left, "{") {
			left += ","
		}
		json := left + " " + property + value + " " + right
		return json, Str(json).IsJSON()
	}
	return "", false
}

// JSONRoot : The first json child
func (s Str) JSONRoot() string {
	if s.IsJSON() {
		str := s.RmBrackets()
		if p := sI(str, ":"); p > 0 {
			str = str[:p]
			str = sT(str, " \t\n\r")
			return Str(str).RmQuotes()
		}
	}
	return ""
}

// JSONChildren :
func (s Str) JSONChildren(xpath, del string) (children []string) {
	content, _, _ := s.JSONXPathValue(xpath, del)
	// fPln(content)

	posList := []int{}
	for _, p := range Str(content).Indices(`":`) {
		if Str(content).BracketDepth(BCurly, p) == 1 {
			posList = append(posList, p)
		}
	}
	// fPln(posList)

	for _, pe := range posList {
		str := content[:pe]
		if ps := sLI(str, "\""); ps >= 0 {
			children = append(children, str[ps+1:])

			// *** search "[" for array, add prefix "[]" to array child ***
			if sHP(sTL(content[pe+2:], " \t\n\r"), "[") {
				children[len(children)-1] = "[]" + children[len(children)-1]
			}
		}
	}
	return
}

// JSONFamilyTree :
func (s Str) JSONFamilyTree(xpath, del string, mapFT *map[string][]string) {
	PC(mapFT == nil, fEf("FamilyTree map is not inited"))
	children := s.JSONChildren(xpath, del)
	// fPln(xpath, children)
	if len(children) > 0 {
		(*mapFT)[xpath] = children
		for _, child := range (*mapFT)[xpath] {
			if sHP(child, "[]") {
				child = child[2:]
			}
			s.JSONFamilyTree(xpath+del+child, del, mapFT)
		}
	}
}

// JSONArrInfo :
func (s Str) JSONArrInfo(xpath, del, id string) map[string]struct {
	Count int
	ID    string
} {
	mapFT := &map[string][]string{}
	s.JSONFamilyTree(xpath, del, mapFT)

	mapAC := map[string]struct {
		Count int
		ID    string
	}{}

	for k, v := range *mapFT {
		if sHP(v[0], "[]") {
			content, _, _ := s.JSONXPathValue(k, ".")
			content = Str(content).RmBrackets()
			n := Str(content).BracketPairCount(BCurly)
			if n == 0 {
				bbox, _, _ := Str(content).BracketsPos(BBox, 1, 1)
				n = sCnt(bbox, ",") + 1
				if sT(bbox, " \t\n\r[]") == "" {
					n = 0
				}
			}
			mapAC[k] = struct {
				Count int
				ID    string
			}{Count: n, ID: id}
		}
	}
	return mapAC
}

// ******************************************************************************

// GQLBuild :
func (s Str) GQLBuild(typename, field, fieldtype string) (gql string) {
	if ok, pos := s.LooseSearchStrs("type", typename, "{", " \t"); ok {		
		_, _, r := s[pos:].BracketsPos(BCurly, 1, 1)
		gql = s.V()[:pos+r]
		tail := s.V()[pos+r+1:]
		add := fSf("\t%s: %s\n}", field, fieldtype)
		gql += add + tail
	} else {
		s += Str(fSf("\n\ntype %s {\n\t%s: %s\n}", typename, field, fieldtype))
		gql = s.V()
	}
	return gql
}

package util

import "encoding/json"

// IsJSON :
func (s Str) IsJSON() bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s.V()), &js) == nil
}

// JSONChildValue : s has "{ }" wrapper. idx from 1 and only one value.
func (s Str) JSONChildValue(child string, idx ...int) (content string, pos int) {

	if !sHP(s.V(), "[") && !sHS(s.V(), "]") {
		PC(!Str(s.MkBrackets(BCurly)).IsJSON(), fEf("Invalid JSON String"))
	}

	if child == "" { //                    *** empty child return whole json string ***
		return s.MkBrackets(BCurly), 0
	}

	json, child := s.V(), Str(child).MkQuotes(QDouble)+":"
AGAIN:
	if pos = sI(json, child); pos > 0 { // *** General Found, including nested ***
		above := json[:pos]
		if sCnt(above, "{")-sCnt(above, "}") == 1 { // *** FOUND ( Object OR Value ) ***

			if ok, pchk := s[pos:].LooseSearchChars(":{", " \t\n\r"); ok && pchk == len(child)-1 { //         *** (Object) ***
				content, lb, _ := s[pos:].BracketsPos(BCurly, 1, 1)
				return content, pos + lb
			} else if ok, pchk := s[pos:].LooseSearchChars(":\"", " \t\n\r"); ok && pchk == len(child)-1 { // *** (Value) ***
				pos += len(child)
				content, lq, _ := s[pos:].QuotesPos(QDouble, 1)
				return content, pos + lq
			} else if ok, pchk := s[pos:].LooseSearchChars(":[", " \t\n\r"); ok && pchk == len(child)-1 { //  *** (Array) (SAME type in array) ***

				content, lb, rb := s[pos:].BracketsPos(BBox, 1, 1)
				if len(idx) == 0 { //                                  *** no idx, return all array '[ *** ]' ***
					return content, pos + lb
				}

				content = sT(Str(content).RmBrackets(), " \t\n\r")
				if Str(content).TrimAllInternal(" \t\r\n") == "" { //  *** empty array '[]' ***
					return "[]", 0
				}

				i := 1
				if len(idx) > 0 {
					i = idx[0]
				}

				segment := s[pos : pos+rb+1] //                        *** specific part what we want to deal with ***

				for j := 1; j <= segment.BracketPairCount(BCurly); j++ {
					contentObj, _, _ := Str(content).BracketsPos(BCurly, 1, j)
					content = sRep(content, contentObj, fSf("##%d", j), 1)
				}

				ss := sSpl(content, ",")
				i = TerOp(i > len(ss), len(ss), i).(int)
				i = TerOp(i < 1, 1, i).(int)
				v := sT(ss[i-1], " \t\n\r")

				for _, p := range segment.Indices(v) {
					if sCnt(segment.V()[:p+1], ",") == i-1 {
						if !sHP(v, "##") { //        			*** plain value, position ***
							return v, pos + p
						}
					}
				}

				//                                              *** ##Obj value ***
				iObj := 0
				for iEachEle, eachEle := range ss {
					if sHP(sT(eachEle, " \t\n\r"), "##") {
						iObj++
					}
					if iEachEle == i-1 {
						break
					}
				}
				content, lb, _ = segment.BracketsPos(BCurly, 1, iObj)
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
	PC(!s.IsJSON(), fEf("Invalid JSON"))
	str := s.RmBrackets()
	if p := sI(str, ":"); p > 0 {
		str = str[:p]
		str = sT(str, " \t\n\r")
	}
	return Str(str).RmQuotes()
}

// JSONRootEx : if No root JSON, add "root", return the modified JSON. if root JSON, same as JSONRoot
func (s Str) JSONRootEx() (root string, ext bool, extJSON string) {
	PC(!s.IsJSON(), fEf("Invalid JSON"))
	if children := s.JSONChildren("", "."); len(children) > 1 {
		root, ext = "root", true
		prefix := "{\n\t\"root\": "
		suffix := "\n}"
		extJSON = prefix + s.V() + suffix
	} else {
		root, ext, extJSON = s.JSONRoot(), false, ""
	}
	return
}

// JSONChildren :
func (s Str) JSONChildren(xpath, del string) (children []string) {
	content, _, _ := s.JSONXPathValue(xpath, del)
	posList := []int{}
	for _, p := range Str(content).Indices(`":`) {
		if Str(content).BracketDepth(BCurly, p) == 1 {
			posList = append(posList, p)
		}
	}

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

	if children := s.JSONChildren(xpath, del); len(children) > 0 {
		// fPln(xpath, children)
		(*mapFT)[xpath] = children
		for _, child := range (*mapFT)[xpath] {
			if sHP(child, "[]") {
				child = child[2:]
			}
			s.JSONFamilyTree(xpath+del+child, del, mapFT) // *** delimiter in key ***
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
		for _, e := range v {
			if sHP(e, "[]") {
				e = e[2:]
				content, _, _ := s.JSONXPathValue(k+del+e, del) // *** get all array '[ *** ]' ***
				n := Str(content).BracketPairCount(BCurly)
				if n == 0 {
					bbox, _, _ := Str(content).BracketsPos(BBox, 1, 1)
					n = sCnt(bbox, ",") + 1
					if sT(bbox, " \t\n\r[]") == "" {
						n = 0
					}
				}
				mapAC[k+del+e] = struct {
					Count int
					ID    string
				}{Count: n, ID: id}
			}
		}
	}
	return mapAC
}

// ******************************************************************************

// GQLBuild : ignore identical field
func (s Str) GQLBuild(typename, field, fieldtype string) (gql string) {
	if ok, pos := s.LooseSearchStrs("type", typename, "{", " \t"); ok {
		content, _, r := s[pos:].BracketsPos(BCurly, 1, 1)
		if sCtn(content, field+":") {
			return s.V()
		}
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
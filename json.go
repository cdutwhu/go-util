package util

import "encoding/json"

// IsJSON :                                                                                                        &
func (s Str) IsJSON() bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s.V()), &js) == nil
}

// IsJSONRootArray : Array Info only be valid on one-type element
func (s Str) IsJSONRootArray() (rootarray bool, tStr string, n int) {
	if s.IsJSON() {
		s = s.T(BLANK)
		if s.C(0) == '[' && s.C(LAST) == ']' {
			rootarray = true

			if s.TrimAllLMR(BLANK).L() == 2 {
				tStr, n = JSONTypeDesc[JT_NULL], 0
				return
			}

			if ok, start, _ := s.LooseSearchStrs("[", "{", BLANK); ok && start == 0 {
				tStr, n = JSONTypeDesc[JT_OBJ], s.RmBrackets().BracketPairCount(BCurly)
			} else if ok, start, _ := s.LooseSearchStrs("[", "[", BLANK); ok && start == 0 {
				tStr, n = JSONTypeDesc[JT_ARR], s.RmBrackets().BracketPairCount(BBox)
			} else if ok, start, _ := s.LooseSearchStrs("[", "\"", BLANK); ok && start == 0 {
				tStr, n = JSONTypeDesc[JT_STR], s.RmBrackets().QuotePairCount(QDouble)
			} else {
				tStr, n = JSONTypeDesc[JT_NUM], sCnt(s.V(), ",")+1
			}
		}
	}
	return
}

// JSONRoot : The first json child                                                                                 &
func (s Str) JSONRoot() string {
	PC(!s.IsJSON(), fEf("Invalid JSON"))
	str := s.RmBrackets()
	if p := str.Idx(":"); p > 0 {
		str = str.S(0, p)
		str = str.T(BLANK)
	}
	return str.RmQuotes().V()
}

// JSONRootEx : if No root JSON, add "root", return the modified JSON. if root JSON, same as JSONRoot              &
func (s Str) JSONRootEx(rootExt string) (root string, ext bool, extJSON string) {
	PC(!s.IsJSON(), fEf("Invalid JSON"))
	if children := s.JSONChildren("", "."); len(children) > 1 {
		root, ext = rootExt, true
		prefix := "{\n\t\"" + rootExt + "\": "
		suffix := "\n}"
		extJSON = prefix + s.V() + suffix
	} else {
		root, ext, extJSON = s.JSONRoot(), false, ""
	}
	return
}

// JSONChildArrCnt : if this child is [{},{},...], return array count                                              ?
func (s Str) JSONChildArrCnt(child string) int {
	child = Str(child).MkQuotes(QDouble).V() + ":"
	Lc := Str(child).L()
AGAIN:
	L := s.L()
	if pos := s.Idx(child); pos > 0 { // *** General Found, including above nested ***
		above := s.S(0, pos, L).V()
		sBelow := s.S(pos, ALL, L)
		if sCnt(above, "{")-sCnt(above, "}") == 1 { // *** TRUELY FOUND ( Object OR Value ) ***
			if ok, start, _ := sBelow.LooseSearchStrs("\":", "[", BLANK); ok && start == Lc-2 {
				content, _, _ := sBelow.BracketsPos(BBox, 1, 1)
				content = content.RmBrackets().T(BLANK)
				switch content.C(0) {
				case '{':
					return content.BracketPairCount(BCurly)
				case '[':
					return content.BracketPairCount(BBox)
				case '"':
					return content.QuotePairCount(QDouble)
				case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
					panic("number count haven't been implemented")
				}
			} else {
				return -1
			}
		}
		// *** FAKE FOUND ***
		s = s.SegRep(pos, pos+2, "\"*")
		goto AGAIN
	}
	return -1
}

// JSONChildValueSimple : only return the value content                                                            ?
func (s Str) JSONChildValueSimple(child string) string {
	child = Str(child).MkQuotes(QDouble).V() + ":"
	Lc := Str(child).L()
AGAIN:
	L := s.L()
	if pos := s.Idx(child); pos > 0 { // *** General Found, including above nested ***
		above := s.S(0, pos, L).V()
		sBelow := s.S(pos, ALL, L)
		if sCnt(above, "{")-sCnt(above, "}") == 1 { // *** TRUELY FOUND ( Object OR Value ) ***
			if ok, start, _ := sBelow.LooseSearchStrs("\":", "{", BLANK); ok && start == Lc-2 { //         *** object ***
				str, _, _ := sBelow.BracketsPos(BCurly, 1, 1)
				return str.V()
			} else if ok, start, _ := sBelow.LooseSearchStrs("\":", "[", BLANK); ok && start == Lc-2 { //  *** array ***
				str, _, _ := sBelow.BracketsPos(BBox, 1, 1)
				return str.V()
			} else if ok, start, _ := sBelow.LooseSearchStrs("\":", "\"", BLANK); ok && start == Lc-2 { // *** string ***
				str, _, _ := sBelow.QuotesPos(QDouble, 1)
				return str.V()
			} else { //                                                                     *** maybe number ... ***
				panic("number value haven't been implemented")
			}
		}

		// *** FAKE FOUND ***
		s = s.SegRep(pos, pos+2, "\"*")
		goto AGAIN
	}
	return ""
}

// JSONChildValue : s has "{}" wrapper. idx from 1 and it's one value.                                             &
// if idx not given, and if child is array, get whole array content.
// if child is array, return the array's count
func (s Str) JSONChildValue(child string, idx ...int) (content string, pos int, nArr int) {
	nArr = -1

	if !s.HP("[") && !s.HS("]") {
		PC(!s.MkBrackets(BCurly).T(BLANK).IsJSON(), fEf("Invalid JSON String"))
	}

	if child == "" { //                    *** empty child return whole json string ***
		return s.MkBrackets(BCurly).V(), 0, 0
	}

	child = Str(child).MkQuotes(QDouble).V() + ":"
	Lc := Str(child).L()

AGAIN:
	L := s.L()
	if pos = s.Idx(child); pos > 0 { // *** General Found, including above nested ***

		above := s.S(0, pos, L).V()
		sBelow := s.S(pos, ALL, L)

		if sCnt(above, "{")-sCnt(above, "}") == 1 { // *** TRUELY FOUND ( Object OR Value ) ***

			if ok, pchk, _ := sBelow.LooseSearchStrs(":", "{", BLANK); ok && pchk == Lc-1 { //         *** (Object) ***

				content, lb, _ := sBelow.BracketsPos(BCurly, 1, 1)
				return content.V(), pos + lb, 0

			} else if ok, pchk, _ := sBelow.LooseSearchStrs(":", "\"", BLANK); ok && pchk == Lc-1 { // *** (Value string) ***

				content, lq, _ := sBelow.QuotesPos(QDouble, 2)
				return content.V(), pos + lq, 0

			} else if ok, pchk, _ := sBelow.LooseSearchAny2Strs([]string{":"}, DigStrArr, BLANK); ok && pchk == Lc-1 { // *** (Value digit) ***

				_, content := sBelow.KeyValuePair(":", " ", BLANK+",}", true, true)
				return content.V(), pos + sBelow.Idx(content.V()), 0

			} else if ok, pchk, _ := sBelow.LooseSearchStrs(":", "[", BLANK); ok && pchk == Lc-1 { //  *** (Array) ***

				content, lb, rb := sBelow.BracketsPos(BBox, 1, 1)
				sContent := content.RmBrackets().T(BLANK)

				if sContent.TrimAllInternal(BLANK) == "" { //          *** empty array '[]' ***
					return "[]", 0, 0
				}

				i := 1
				if len(idx) > 0 {
					i = idx[0]
				}

				nPairBCurly := sContent.BracketPairCount(BCurly)
				for j := 1; j <= nPairBCurly; j++ {
					objStr, _, _ := sContent.BracketsPos(BCurly, 1, 1)
					sContent = Str(sRep(sContent.V(), objStr.V(), fSf("#O%d", j), 1)) //     *** #O modified content ***
				}
				nPairBBox := sContent.BracketPairCount(BBox)
				for j := 1; j <= nPairBBox; j++ {
					subArrStr, _, _ := sContent.BracketsPos(BBox, 1, 1)
					sContent = Str(sRep(sContent.V(), subArrStr.V(), fSf("#A%d", j), 1)) //  *** #A modified content ***
				}
				nPairQuote, mapS := sContent.QuotePairCount(QDouble), map[string]string{}
				for j := 1; j <= nPairQuote; j++ {
					sQuotes, _, _ := sContent.QuotesPos(QDouble, 1)
					symbol, K := fSf("#S%d", j), ""
					for k, v := range mapS {
						if v == sQuotes.V() { //                                    *** same string use same #S
							K = k
							break
						}
					}
					symbol = TerOp(K != "", K, symbol).(string)
					sContent = Str(sRep(sContent.V(), sQuotes.V(), symbol, 1)) //   *** #S modified content ***
					mapS[symbol] = sQuotes.V()
				}

				ss := sSpl(sContent.V(), ",")
				nArr = len(ss)
				if len(idx) == 0 { //                                 *** no idx provided, return all array '[ *** ]' ***
					return content.V(), pos + lb, nArr
				}

				i = TerOp(i > nArr, nArr, i).(int)
				i = TerOp(i < 1, 1, i).(int)
				v := Str(ss[i-1]).T(BLANK)

				sSeg := s.S(pos, pos+rb+1, L) //                      *** original array line ***
				search := TerOp(v.HP("#S"), mapS[v.V()], v.V()).(string)

				for pIdx, p := range sContent.Indices(v.V()) {
					if sCnt(sContent.S(0, p+1).V(), ",") == i-1 { //  ***
						if !v.HP("#O") && !v.HP("#A") { //            *** plain value, position ***

							segIdx := 0
							for _, pSegPos := range sSeg.Indices(search) {
								if sSeg.BracketDepth(BBox, pSegPos) == 1 && sSeg.BracketDepth(BCurly, pSegPos) == 0 {
									if segIdx == pIdx {
										return search, pos + pSegPos, nArr
									}
									segIdx++
								}
							}
						}
					}
				}

				//                                                     *** #Obj & #Arr value ***
				iObj, iArr, iType := 0, 0, ""
				for iEle, eachEle := range ss {
					ele := Str(eachEle).T(BLANK)
					if ele.HP("#O") {
						iObj, iType = iObj+1, "#O"
					}
					if ele.HP("#A") {
						iArr, iType = iArr+1, "#A"
					}
					if iEle == i-1 {
						break
					}
				}

				content, lb, _ = sSeg.BracketsPos(BCurly, 1, iObj)
				if iType == "#A" {
					content, lb, _ = sSeg.BracketsPos(BBox, 2, iArr)
				}

				return content.V(), pos + lb, nArr
			}
		}

		// *** FAKE FOUND ***
		s = s.SegRep(pos, pos+2, "\"*")
		goto AGAIN
	}
	return "", 0, -1
}

// JSONXPathValue : s has "{}" wrapper. indices is from the 1st path-seg's array index.                            &
// if it's not array, use 1; if it is 0 and array, return the whole array
func (s Str) JSONXPathValue(xpath, del string, indices ...int) (content string, pStart, pEnd int, nArr int) {
	segs := sSpl(xpath, del)
	PC(len(segs) != len(indices), fEf("path & seg's index count not match"))
	for i := 0; i < len(indices)-1; i++ {
		PC(indices[i] <= 0, fEf("Only Last index can be 0 to get the whole array content"))
	}

	pEach := 0
	for i, seg := range segs {
		s = TerOp(content != "", Str(content), s).(Str)
		if indices[i] != 0 {
			content, pEach, nArr = s.JSONChildValue(seg, indices[i])
		} else {
			content, pEach, nArr = s.JSONChildValue(seg)
		}
		if content == "" {
			return
		}
		if pEach == -1 {
			pStart, pEach = -1, -1
			return
		}
		pStart += pEach
	}
	pEnd = pStart + Str(content).L() - 1
	return
}

// JSONChildren : if xpath is "", get top level's all children name                                                 &
func (s Str) JSONChildren(xpath, del string, indices ...int) (children []string) {

	indices = TerOp(xpath == "", []int{1}, indices).([]int)
	content, _, _, _ := s.JSONXPathValue(xpath, del, indices...)

	if ok, pos, _ := Str(content).LooseSearchStrs("[", "{", BLANK); ok && pos == 0 {
		content, _, _, _ = s.JSONXPathValue(xpath, del, 0)
	}

	sContent := Str(content)
	posList := []int{}
	for _, p := range sContent.Indices(`":`) {
		if sContent.BracketDepth(BCurly, p) == 1 {
			posList = append(posList, p)
		}
	}
	// fPln(posList) // DEBUG

	for _, pe := range posList {
		str := sContent.S(0, pe)
		if ps := str.LIdx("\""); ps >= 0 {
			children = append(children, str.S(ps+1, ALL).V())

			// *** search "[" for array, add prefix "[]" to array child ***
			if sContent.S(pe+2, ALL).TL(BLANK).HP("[") {
				children[len(children)-1] = "[]" + children[len(children)-1]
			}
		}
	}
	return
}

// JSONFamilyTree : Use One Sample to get FamilyTree, DO NOT use long array data                                   &
func (s Str) JSONFamilyTree(xpath, del string, mapFT *map[string][]string) {
	PC(mapFT == nil, fEf("FamilyTree return map is not initialised !"))
	// fPln(xpath) // DEBUG

	indices := []int{}
	for i := 0; i < len(sSpl(xpath, del)); i++ {
		indices = append(indices, 1)
	}

	if children := s.JSONChildren(xpath, del, indices...); len(children) > 0 {
		//fPln(xpath, children) // DEBUG
		(*mapFT)[xpath] = children
		for _, child := range (*mapFT)[xpath] {
			if Str(child).HP("[]") {
				child = Str(child).S(2, ALL).V()
			}
			nextPath := Str(xpath + del + child).TL(".").V()
			s.JSONFamilyTree(nextPath, del, mapFT) //                  *** delimiter in key ***
		}
	}
}

// JSONArrInfo : Only deal with SAME TYPE element array                                                            ?
func (s Str) JSONArrInfo(xpath, del, id string, mapFT *map[string][]string) (*map[string][]string, *map[string]struct {
	Count int
	ID    string
}) {
	if mapFT == nil {
		mapFT = &map[string][]string{}
		s.JSONFamilyTree(xpath, del, mapFT)
	}

	// for k, v := range *mapFT {
	// 	fPln(k, " : ", v)
	// }
	// return nil, nil
	// fPln(" ------------------------------------------------- ")

	mapA := map[string]bool{}
	for attr, children := range *mapFT {
		for _, child := range children {
			sChild := Str(child)
			if sChild.HP("[]") {
				sChild = sChild.S(2, ALL)
				path := Str(attr + del + sChild.V()).T(del).V()
				mapA[path] = true
			}
		}
	}

	for k, v := range mapA {
		fPln(k, " : ", v)
	}
	//return nil, nil

	mapAC := &map[string]struct {
		Count int
		ID    string
	}{}

	for k := range mapA {
		if len(sSpl(k, del)) == 1 {
			if n := s.JSONChildArrCnt(k); n >= 0 {
				(*mapAC)[k] = struct {
					Count int
					ID    string
				}{Count: n, ID: id}
			}
		}
	}

	// for k := range mapA {
	// 	ss := sSpl(k, del)
	// 	if len(ss) == 2 {
	// 		s1, s2, s12, s12ns := ss[0], ss[1], sJ(ss, del), []string{}
	// 		if cntid1, ok := (*mapAC)[s1]; ok { //                                ** get upper level's count
	// 			for i := 1; i <= cntid1.Count; i++ {
	// 				s12ns = append(s12ns, fSf("%s#%d%s%s", s1, i, del, s2))
	// 			}
	// 			for i, ns := range s12ns {
	// 				idx := Str(sSpl(sSpl(ns, del)[0], "#")[1]).ToInt()
	// 				_, _, _, n := s.JSONXPathValue(s12, del, []int{idx, 0}...) // ** get this level's count
	// 				(*mapAC)[s12ns[i]] = struct {
	// 					Count int
	// 					ID    string
	// 				}{Count: n, ID: id}
	// 			}
	// 		} else {
	// 			if _, _, _, n := s.JSONXPathValue(s12, del, []int{1, 0}...); n >= 0 {
	// 				(*mapAC)[s12] = struct {
	// 					Count int
	// 					ID    string
	// 				}{Count: n, ID: id}
	// 			}
	// 		}
	// 	}
	// }

	// for k := range mapA {
	// 	ss := sSpl(k, del)
	// 	if len(ss) == 3 {
	// 		s1, s1nArr := ss[0], []string{}
	// 		if n1, ok := (*mapAC)[s1]; ok {
	// 			for i := 1; i <= n1.Count; i++ {
	// 				s1nArr = append(s1nArr, fSf("%s#%d", s1, i))
	// 			}
	// 		} else {

	// 		}

	// 		s2, s1ns2nArr := ss[1], []string{}
	// 		for _, s1n := range s1nArr {
	// 			s1ns2 := s1n + del + s2
	// 			if n12, ok := (*mapAC)[s1ns2]; ok {
	// 				for i := 1; i <= n12.Count; i++ {
	// 					s1ns2nArr = append(s1ns2nArr, fSf("%s#%d", s1ns2, i))
	// 				}
	// 			} else {

	// 			}
	// 		}

	// 		s3 := ss[2]
	// 		for _, sn := range s1ns2nArr {
	// 			sArr, indices := []string{}, []int{}
	// 			for _, sn := range sSpl(sn, del) {
	// 				strnum := sSpl(sn, "#")
	// 				sArr = append(sArr, strnum[0])
	// 				indices = append(indices, Str(strnum[1]).ToInt())
	// 			}
	// 			path := sJ(sArr, del) + del + s3
	// 			indices = append(indices, 0)
	// 			if _, _, _, n := s.JSONXPathValue(path, del, indices...); n >= 0 { // ** get this level's count
	// 				(*mapAC)[sn+del+s3] = struct {
	// 					Count int
	// 					ID    string
	// 				}{Count: n, ID: id}
	// 			}
	// 		}
	// 	}
	// }

	// ****************************************************************************************

	//for k := range mapA {
	//ss := sSpl(k, del)
	//lss := len(ss)

	// n := -1
	// if lss == 1 {

	// 	indices := []int{0}
	// 	_, _, _, n = s.JSONXPathValue(k, del, indices...)
	// 	(*mapAC)[k] = struct {
	// 		Count int
	// 		ID    string
	// 	}{Count: n, ID: id}

	// } else if lss == 2 {

	// 	s1, s2, s12, s12ns := ss[0], ss[1], sJ(ss, del), []string{}
	// 	_, _, _, n = s.JSONXPathValue(s1, del, []int{0}...)
	// 	for i := 1; i <= n; i++ {
	// 		s12ns = append(s12ns, fSf("%s#%d%s%s", s1, i, del, s2))
	// 	}
	// 	for i, ns := range s12ns {
	// 		idx := Str(sSpl(sSpl(ns, del)[0], "#")[1]).ToInt()
	// 		_, _, _, n = s.JSONXPathValue(s12, del, []int{idx, 0}...)
	// 		(*mapAC)[s12ns[i]] = struct {
	// 			Count int
	// 			ID    string
	// 		}{Count: n, ID: id}
	// 	}

	// } else if lss == 3 {

	// }

	/******************************************/

	// if lss < 3 {

	// 	indices := TerOp(len(ss) == 1, []int{0}, []int{1, 0}).([]int)
	// 	_, _, _, n := s.JSONXPathValue(k, del, indices...)
	// 	(*mapAC)[k] = struct {
	// 		Count int
	// 		ID    string
	// 	}{Count: n, ID: id}

	// } else if lss == 3 {

	// 	s12, s3, s123, s123ns := sJ(ss[:2], del), ss[2], sJ(ss, del), []string{}
	// 	_, _, _, n := s.JSONXPathValue(s12, del, []int{1, 0}...)
	// 	for i := 1; i <= n; i++ {
	// 		s123ns = append(s123ns, fSf("%s#%d%s%s", s12, i, del, s3))
	// 	}
	// 	// fPln(s123ns)

	// 	for i, ns := range s123ns {
	// 		idx := Str(sSpl(sSpl(ns, del)[1], "#")[1]).ToInt()
	// 		_, _, _, n = s.JSONXPathValue(s123, del, []int{1, idx, 0}...)
	// 		(*mapAC)[s123ns[i]] = struct {
	// 			Count int
	// 			ID    string
	// 		}{Count: n, ID: id}
	// 	}

	// } else if lss == 4 {

	// 	// s123, s4, s1234, s1234ns := sJ(ss[:3], del), ss[3], sJ(ss, del), []string{}
	// 	// _, _, _, n := s.JSONXPathValue(s12, del, []int{1, 0}...)

	// } else {

	// 	panic("haven't implemented this level nested array function")

	// }
	//}

	return mapFT, mapAC
}

// JSONBuild : NOT support mixed (atomic & object) types in one array                                              &
func (s Str) JSONBuild(xpath, del, property, value string, indices ...int) (string, bool) {
	if s.T(BLANK) == "" {
		s = Str("{}")
	}

	PC(len(indices) != len(sSpl(xpath, del)), fEf("indices count must be xpath seg's count"))

	property = Str(property).MkQuotes(QDouble).V() + ": "
	sValue := Str(value)
	if sValue.C(0) != '{' && sValue.C(0) != '[' {
		value = sValue.MkQuotes(QDouble).V()
	}

	if s == "{}" {
		s = Str(fSf(`{ %s%s}`, property, value))
		return s.V(), s.IsJSON()
	}

	if sub, start, end, _ := s.JSONXPathValue(xpath, del, indices...); start != -1 {

		for _, p0 := range Str(sub).Indices(property) { //                               ** incoming p-v's property already exists **
			sub02p0 := Str(sub).S(0, p0).V()
			if sCnt(sub02p0, "{")-sCnt(sub02p0, "}") == 1 { //                           ** 1 level child property **

				Subp02end := Str(sub).S(p0, ALL)
				if p1 := Subp02end.Idx(property + "["); p1 == 0 { //                     ** already array format, 3rd, 4th... coming **
					box, _, _ := Str(sub).BracketsPos(BBox, 1, 1)
					inBox := box.RmBrackets().TrimAllLMR(BLANK).V()
					ss := append(sSpl(inBox, ","), value)
					newBox := Str(sJ(ss, ", ")).MkBrackets(BBox).V()
					sub = sRep(sub, box.V(), newBox, 1)
				} else { //                                                              ** only one exists, the 2nd coming, change to array format **
					k, v := Subp02end.KeyValuePair(":", " ", " ", false, false)
					if !v.HP("{") { //                                                   ** atomic value **
						v = v.TR(",")
						sub = sRep(sub, property+v.V(), property+"["+v.V()+", "+value+"]", 1)
					} else { //                                                          ** object value **
						sub = sRep(sub, k.V()+":", k.V()+": [", 1)
						sub = Str(sub).S(0, ALL-1).T(BLANK).V() + ", " + "{}" + " ] " + "}"
					}
				}

				left, right := s.S(0, start).V(), s.S(end+1, ALL).V()
				json := left + sub + right
				return json, Str(json).IsJSON()
			}
		}

		// **********************************************

		left, right := s.S(0, end).T(BLANK).V(), s.S(end, ALL).T(BLANK).V()
		left = TerOp(!Str(left).HS("{"), left+",", left).(string)
		json := left + " " + Str(property+value).T(BLANK).V() + " " + right
		return json, Str(json).IsJSON()
	}
	return "", false
}

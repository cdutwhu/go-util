package util

import (
	"io/ioutil"
	"testing"
)

func TestJSONChild(t *testing.T) {

	jsonSample := Str(` [ "a" ] `)
	fPln(jsonSample.IsJSON())
	fPln(jsonSample.IsJSONRootArray())

	jsonBytes, _ := ioutil.ReadFile("json_test.json") // only LF at end of line
	s := Str(jsonBytes)
	s.SetEnC()

	if !s.IsJSON() {
		t.Errorf("JSON Format Error\n")
	}
	if s.JSONRoot() != "test最最" {
		t.Errorf("JSONRoot() Error\n")
	}

	if root, ext, json := s.JSONRootEx("fakeRoot"); !ext {
		t.Errorf("JSONRootEx() Error <%s> <%v> <%s>\n", root, ext, json)
	}

	fPln(s.JSONChildValue("Nam最e"))

	for i := 1; i <= 11; i++ {

		v, p, nArr := s.JSONChildValue("Nam最e", i)

		if nArr != 11 {
			t.Errorf("JSONChildValue() Error <%d>\n", nArr)
		}

		switch i {
		case 1:
			if v != "23" || p != 42 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 2:
			if v != "45" || p != 46 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 3:
			if v != "23" || p != 50 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 4:
			if v != `"ab,最c"` || p != 54 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 5:
			if v != `{"test": "ab,最c", "test1": "AB最C"}` || p != 63 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 6:
			if v != `[ 2, "ab最c", 3]` || p != 99 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 7:
			if v != `"ab,最c"` || p != 117 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 8:
			if v != `{"p2":  "v2最"}` || p != 126 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 9:
			if v != `[ 22, 33]` || p != 142 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 10:
			if v != `",ab最c"` || p != 153 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		case 11:
			if v != `"def最"` || p != 162 || nArr != 11 {
				t.Errorf("JSONChildValue() Error <%d> <%s> <%d> <%d>\n", i, v, p, nArr)
			}
		}
	}

	{
		content, start, end, nArr := s.JSONXPathValue("actor.member.innerArr", ".", 1, 1, 0)
		if content != `[1, 2, 3]` || start != 393 || end != 401 || nArr != 3 {
			t.Errorf("JSONXPathValue() Error <%s> <%d> <%d> <%d>\n", content, start, end, nArr)
		}
		content, start, end, nArr = s.JSONXPathValue("actor.member.innerArr", ".", 1, 2, 2)
		if content != `5` || start != 837 || end != 837 || nArr != 4 {
			t.Errorf("JSONXPathValue() Error <%s> <%d> <%d> <%d>\n", content, start, end, nArr)
		}
		content, start, end, nArr = s.JSONXPathValue("actor.member.account.homePage", ".", 1, 3, 1, 1)
		if content != `"http://www.example3.com"` || start != 1382 || end != 1406 || nArr != 0 {
			t.Errorf("JSONXPathValue() Error <%s> <%d> <%d> <%d>\n", content, start, end, nArr)
		}
		content, start, end, nArr = s.JSONXPathValue("actor.member.mbox_sha1sum", ".", 1, 2, 1)
		if content != `"ebd31e95054c018b1072最7ccffd2ef2ec3a016ee9222"` || start != 1134 || end != 1179 || nArr != 0 {
			t.Errorf("JSONXPathValue() Error <%s> <%d> <%d> <%d>\n", content, start, end, nArr)
		}
	}

	{
		children := s.JSONChildren("", ".")
		if len(children) != 6 {
			t.Errorf("JSONChildren() Error <%d>\n", len(children))
		}
		for i, child := range children {
			switch i {
			case 0:
				if child != "test最最" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 1:
				if child != "[]Nam最e" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 2:
				if child != "test" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 3:
				if child != "test1" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 4:
				if child != "actor" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 5:
				if child != "test2" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			}
		}

		children = s.JSONChildren("actor", ".", 1)
		if len(children) != 5 {
			t.Errorf("JSONChildren() Error <%d>\n", len(children))
		}
		for i, child := range children {
			switch i {
			case 0:
				if child != "name" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 1:
				if child != "mbox" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 2:
				if child != "[]simple" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 3:
				if child != "[]member" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			case 4:
				if child != "objectType" {
					t.Errorf("JSONChildren() Error <%d> <%s>\n", i, child)
				}
			}
		}

		children = s.JSONChildren("actor.name", ".", 1, 1)
		if len(children) != 0 {
			t.Errorf("actor.name's nChildren <%d>\n", len(children))
		}
		children = s.JSONChildren("actor.member", ".", 1, 1)
		if len(children) != 6 {
			t.Errorf("actor.member's nChildren <%d> <%v>\n", len(children), children)
		}
	}
}

func TestJSONArrInfo(t *testing.T) {
	jsonBytesTemp, _ := ioutil.ReadFile("xapi.json")
	sTemp := Str(jsonBytesTemp)
	fPln(sTemp.SetEnC())

	{
		mapFT := &map[string][]string{}
		sTemp.JSONFamilyTree("", " ~ ", mapFT)
		for k, v := range *mapFT {
			fPln(k, "::", v)
		}

		jsonBytes, _ := ioutil.ReadFile("xapi.1.json")
		s := Str(jsonBytes)
		fPln(s.SetEnC())

		// mapFT = &map[string][]string{}
		// root, _, newJSON := s.JSONRootEx("fakeRoot")
		// Str(newJSON).JSONFamilyTree(root, ".", mapFT)
		// fPln(*mapFT)

		mapFT, mapAI := s.JSONArrInfo("", " ~ ", "535e966a-931e-430f-a809-d90401147864", mapFT)
		// for k, v := range *mapFT {
		// 	fPln(k, "::", v)
		// }
		fPln("----------------------------------------------------------------------")
		if mapAI != nil {
			for k, v := range *mapAI {
				fPln(k, "::", v)
			}
		}
	}
}

func TestJSONMake(t *testing.T) {
	json, ok := Str("").JSONBuild("", ".", "StaffPersonal", "{}", 1)
	// //json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "-RefId", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "LocalId", "946379881", 1)
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "LocalId", "946379882", 1)
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "LocalIdTest", "tttttttt", 1)
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "StateProvinceId", "C2345681", 1)
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "OtherIdList", "{}", 1)            //                                ***
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList", ".", "OtherId", "{}", 1, 1) //                     *** 1
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", "-Type", "0004", 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", "#content", "333333333", 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "PersonInfo", "{}", 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", "Name", "{}", 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Name", ".", "-Type", "LGL", 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", "OtherNames", "{}", 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames", ".", "Name", "[{},{}]", 1, 1, 1) //      ***
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames.Name", ".", "-Type", "AKA", 1, 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames.Name", ".", "-Type", "PRF", 1, 1, 1, 2) // ***
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", "Demographics", "{}", 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics", ".", "CountriesOfCitizenship", "{}", 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics.CountriesOfCitizenship", ".", "CountryOfCitizenship", "\"8104\"", 1, 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics.CountriesOfCitizenship", ".", "CountryOfCitizenship", "\"1101\"", 1, 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "LocalId", "946379883", 1)
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "LocalIdTest", "iiiiiiii", 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics.CountriesOfCitizenship", ".", "CountryOfCitizenship", "\"2202\"", 1, 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList", ".", "OtherId", "{}", 1, 1) //                     *** 2
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", "-Type", "0005", 1, 1, 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", "-Type1", "0008", 1, 1, 2)
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", "#content", "44444444", 1, 1, 2)
	// // json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalId", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", "LocalId", "test2", 1)
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", "#content", "44444444", 1, 1, 2)
	// // // fPln(Str(json).JSONXPath("StaffPersonal.PersonInfo.OtherNames.Name", ".", 1))

	fPln(json, ok)
	ioutil.WriteFile("temp.json", []byte(json), 0666)
}

func TestJSONObjectMerge(t *testing.T) {

	json1 := ``
	json2 := `{
		"TeachingGroup1": {
			"RefId": "F47C2C6D-BD49-40E6-A430-111111111111",
			"SchoolYear": "2018",
			"LocalId": "2018-English-8-1-B",
			"ShortName": "8B English 1",
			"LongName": "Year 8B English 1",
			"TimeTableSubjectRefId": "FD3E4B1F-0FC6-4607-BB95-78791ABA8997",
			"TeacherList": {
				"TeachingGroupTeacher": {
					"StaffPersonalRefId": "D4A3C1E3-3F6E-4B31-ABA6-26809DF5FD63",
					"StaffLocalId": "kafaj506",
					"Name": {
						"Type": "LGL",
						"FamilyName": "Knoll",
						"GivenName": "Ina"
					},
					"Association": "Class Teacher"
				}
			},
			"TeachingGroupPeriodList": {
				"TeachingGroupPeriod": [
					{
						"RoomNumber": "171",
						"DayId": "Fr",
						"PeriodId": "12:00:00"
					},
					{
						"RoomNumber": "166",
						"DayId": "Fr",
						"PeriodId": "15:00:00"
					}
				]
			}
		}
	}`

	rst := Str(json1).JSONObjectMerge(json2)
	ioutil.WriteFile("debug_temp.json", []byte(rst), 0666)
}

func TestJSONRoot(t *testing.T) {
	// jsonbytes, _ := ioutil.ReadFile("./test1.json")
	json := Str(`{
		"TeachingGroup": {
			"RefId": "F47C2C6D-BD49-40E6-A430-360333274DB2",
			"SchoolYear": "2018",
			"LocalId": "2018-English-8-1-B",
			"ShortName": "8B English 1",
			"LongName": "Year 8B English 1",
			"TimeTableSubjectRefId": "FD3E4B1F-0FC6-4607-BB95-78791ABA8997",
			"TeacherList": {
				"TeachingGroupTeacher": {
					"StaffPersonalRefId": "D4A3C1E3-3F6E-4B31-ABA6-26809DF5FD63",
					"StaffLocalId": "kafaj506",
					"Name": {
						"Type": "LGL",
						"FamilyName": "Knoll",
						"GivenName": "Ina"
					},
					"Association": "Class Teacher"
				}
			}
		}
	}`)
	fPln(json.JSONRoot())
	// root, ext, newJSON := Str(json).JSONRootEx("MyRoot")
	// fPln(root)
	// if ext {
	// 	json = newJSON
	// }

	// mapFT, mapAC := Str(json).JSONArrInfo("", " ~ ", "1234567890", nil)
	// if mapFT != nil && mapAC != nil {
	// 	for k, v := range *mapFT {
	// 		fPln(k, v)
	// 	}
	// 	fPln("<----------------------------------->")
	// 	for k, v := range *mapAC {
	// 		fPln(k, v)
	// 	}
	// }
}

func TestGQLBuild(t *testing.T) {
	s := Str(`
type StaffPersonal {
	-RefId: String
	LocalId: String
	StateProvinceId: String
	OtherIdList: OtherIdList
}
	
type OtherIdList {
	OtherId: OtherId
}

type OtherIdList1 {
	OtherId1: OtherId1
}

type OtherIdList2 {
	OtherId2: OtherId2
}
	
type OtherId {
	-Type: String
}`)

	s = Str(s.GQLBuild("OtherIdList2", "OtherId2", "String"))

	// s := Str("")
	// s = Str(s.GQLBuild("StaffPersonal", "RefId", "String"))
	// s = Str(s.GQLBuild("StaffPersonal", "LocalId", "String"))
	// s = Str(s.GQLBuild("Recent", "SchoolLocalId", "String"))
	// s = Str(s.GQLBuild("Recent", "LocalCampusId", "String"))
	// s = Str(s.GQLBuild("StaffPersonal", "StateProvinceId", "String"))
	// s = Str(s.GQLBuild("NAPLANClassListType", "ClassCode", "[String]"))
	// s = Str(s.GQLBuild("StaffPersonal", "OtherIdList", "OtherIdList"))

	fPln(s)
}

package util

import (
	"io/ioutil"
	"testing"
)

func TestJSONChild(t *testing.T) {
	// s := Str(`{
	// 	"data": {
	// 		"Name": [ 23, 45,   23,  {"p1":   "v1"},  "ab   c", {"p2":     "v2"}   ,  "def" ]
	// 	}
	// }`)

	// root, ext, newJSON := s.JSONRootEx()
	// fPln(root, ext)
	// fPln(newJSON)

	//children := s.JSONChildren("", ".")
	//fPln(children)

	s := Str(`{
		"actor": {
			"name": "Team PB",
			"mbox": "mailto:teampb@example.com",
			"simple": [ 1, 3, 4, 5 ],
			"member": [
				{
					"name": "Andrew Downes",
					"account": {
						"homePage": "http://www.example1.com",
						"name": "13936749111"
					},
					"openid": "http://toby.openid.example1.org/",
					"mbox_sha1sum": "ebd31e95054c018b10727ccffd2ef2ec3a016ee9111",
					"objectType": "Agent"
				},
				{
					"name": "Toby Nichols",
					"account": {
						"homePage": "http://www.example2.com",
						"name": "13936749222"
					},
					"openid": "http://toby.openid.example2.org/",
					"mbox_sha1sum": "ebd31e95054c018b10727ccffd2ef2ec3a016ee9222",
					"objectType": "Agent"
				},
				{
					"name": "Ena Hills",
					"account": {
						"homePage": "http://www.example3.com",
						"name": "13936749333"
					},
					"openid": "http://toby.openid.example3.org/",
					"mbox_sha1sum": "ebd31e95054c018b10727ccffd2ef2ec3a016ee9333",
					"objectType": "Agent"
				}
			],
			"objectType": "Group"
		}
	}`)
	//fPln(s.JSONChildValue("Name", 4))
	fPln(s.JSONXPathValue("actor.member", "."))
	fPln(s.JSONChildren("actor.member", "."))
}

func TestJSONMake(t *testing.T) {
	json, ok := Str("").JSONBuild("", "", 1, "StaffPersonal", "{}")
	//json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "-RefId", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalId", "946379881")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalId", "946379882")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalIdTest", "tttttttt")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "StateProvinceId", "C2345681")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "OtherIdList", "{}") // ***
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList", ".", 1, "OtherId", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", 1, "-Type", "0004")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", 1, "#content", "333333333")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "PersonInfo", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", 1, "Name", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Name", ".", 1, "-Type", "LGL")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", 1, "OtherNames", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames", ".", 1, "Name", "[{},{}]") // ***
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames.Name", ".", 1, "-Type", "AKA")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames.Name", ".", 2, "-Type", "PRF") // ***
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", 1, "Demographics", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics", ".", 1, "CountriesOfCitizenship", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics.CountriesOfCitizenship", ".", 1, "CountryOfCitizenship", "\"8104\"")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics.CountriesOfCitizenship", ".", 1, "CountryOfCitizenship", "\"1101\"")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalId", "946379883")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalIdTest", "iiiiiiii")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics.CountriesOfCitizenship", ".", 1, "CountryOfCitizenship", "\"2202\"")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList", ".", 1, "OtherId", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", 1, "-Type", "0005")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", 2, "-Type1", "0008")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", 2, "#content", "44444444")
	// json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalId", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalId", "test2")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", 2, "#content", "44444444")
	// // fPln(Str(json).JSONXPath("StaffPersonal.PersonInfo.OtherNames.Name", ".", 1))

	fPln(json, ok)
}

func TestJSONRoot(t *testing.T) {
	jsonbytes, _ := ioutil.ReadFile("./test1.json")
	json := string(jsonbytes)
	root, ext, newJSON := Str(json).JSONRootEx("MyRoot")
	fPln(root)
	if ext {
		json = newJSON
	}

	mapFT, mapAC := Str(json).JSONArrInfo(root, " ~ ", "1234567890")
	for k, v := range *mapFT {
		fPln(k, v)
	}
	for k, v := range *mapAC {
		fPln(k, v)
	}
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

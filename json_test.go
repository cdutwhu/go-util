package util

import (
	"io/ioutil"
	"testing"
)

func TestJSONChild(t *testing.T) {
	s := Str(`{ "data": {
		"Name": [ 23, 45, "abc" ]
	}}`)
	fPln(s.JSONChildValue("data", 3))
	// fPln(s.JSONXPathValue("data.Name", ".", 3))
}

func TestJSONMake(t *testing.T) {
	json, ok := Str("").JSONBuild("", "", 1, "StaffPersonal", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "-RefId", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "LocalId", "946379881")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "StateProvinceId", "C2345681")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "OtherIdList", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList", ".", 1, "OtherId", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", 1, "-Type", "0004")
	json, ok = Str(json).JSONBuild("StaffPersonal.OtherIdList.OtherId", ".", 1, "#content", "333333333")
	json, ok = Str(json).JSONBuild("StaffPersonal", ".", 1, "PersonInfo", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", 1, "Name", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Name", ".", 1, "-Type", "LGL")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", 1, "OtherNames", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames", ".", 1, "Name", "[{},{}]") // ***
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames.Name", ".", 1, "-Type", "AKA")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.OtherNames.Name", ".", 2, "-Type", "PRF")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo", ".", 1, "Demographics", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics", ".", 1, "CountriesOfCitizenship", "{}")
	json, ok = Str(json).JSONBuild("StaffPersonal.PersonInfo.Demographics.CountriesOfCitizenship", ".", 1, "CountryOfCitizenship", "[\"8104\", \"1101\"]") // ***
	// fPln(Str(json).JSONXPath("StaffPersonal.PersonInfo.OtherNames.Name", ".", 1))

	fPln(json, ok)
}

func TestJSONRoot(t *testing.T) {
	jsonbytes, _ := ioutil.ReadFile("./test.json")
	json := string(jsonbytes)
	root := Str(json).JSONRoot()
	fPln(root)

	fPln(Str(json).JSONChildren("StaffPersonal", "."))

	mapFT := &map[string][]string{}
	Str(json).JSONFamilyTree("StaffPersonal", ".", mapFT)

	mapAC := Str(json).JSONArrInfo("StaffPersonal", ".")
	for k, v := range mapAC {
		fPln(k, v)
	}

}

package util

import (
	"testing"
)

func TestIsMadeFrom(t *testing.T) {
	pln(Str("abc").IsMadeFrom('a', 'b'))
}

func TestLooseSearch(t *testing.T) {
	pln(Str("ab C C		C 	* d	 e *.fc  * c			d").LooseSearch("Cde.", ' ', '\t', '*'))
}

func TestTrimInternal(t *testing.T) {
	pln(Str(`*****abc*******abCCC**
***de***.**fc*c**d****`).TrimInternal('*', 1))
}

func TestBeCoveredInMapSIKeys(t *testing.T) {
	m := map[string]int{"abcd": 100, "def": 200}
	pln(Str("abc").BeCoveredInMapSIKeys(m))
}

func TestCoverAnyKeyInMapSI(t *testing.T) {
	m := map[string]int{"abc": 100, "deff": 200}
	pln(Str("def").CoverAnyKeyInMapSI(m))
}

func TestRemovePrefix(t *testing.T) {
	pln(Str("sif.abc").RemovePrefix("sif."))
}

func TestRemoveSuffix(t *testing.T) {
	pln(Str("sif.abc").RemoveSuffix("abc"))
}

func TestRemoveTailFromLast(t *testing.T) {
	pln(Str("a.sif.abc").RemoveTailFromLast("."))
}

func TestRemoveBlankBefore(t *testing.T) {
	pln(Str(`a            :  m   c		=   b  e 	 :  	 f`).RemoveBlankBefore("=", ":"))
}

func TestRemoveBlankNBefore(t *testing.T) {
	pln(Str(`a a            :  m   c		:   b  e 	 :  	 f`).RemoveBlankNBefore(2, ":"))
}

func TestRemoveBlankAfter(t *testing.T) {
	pln(Str(`a   :    t         c    =	b   e  		=		f`).RemoveBlankAfter("=", ":"))
}

func TestRemoveBlankNAfter(t *testing.T) {
	pln(Str(`a a            :  m   c		=   b  e 	 :  	 f`).RemoveBlankNAfter(0, ":"))
}

func TestRemoveBlankNNear(t *testing.T) {
	pln(Str(`a   :    t         c    =	b   e  		=		f`).RemoveBlankNNear(0, "="))
}

func TestRemoveBlankNear(t *testing.T) {
	pln(Str(`a   :    t         c    =	b   e  		=		f`).RemoveBlankNear(":", "="))
}

func TestKeyValueMap(t *testing.T) {
	pln(Str(`<abc a =	"dd"  c		= 	fff>>>>>`).KeyValueMap(' ', '=', '>'))
	pln(Str(`<abc a	 	: 	"dd"  c   =			fff>>>>>`).KeyValueMap(' ', ':', '|'))
}

func TestKeyValuePair(t *testing.T) {
	k, v := Str(`   <abc a =     "dd"  c	: 	= 	fff>>>>>`).KeyValuePair("= ", ' ', ' ', true, false)
	pln(k)
	pln(v)
}

func TestMakeQuotes(t *testing.T) {
	pln(Str("abc").MakeQuotes(QDouble))
}

func TestRemoveQuotes(t *testing.T) {
	pln(Str("'abc'").RemoveQuotes())
}

func TestMakeBrackets(t *testing.T) {
	pln(Str("abc").MakeBrackets(BCurly))
}

func TestRemoveBrackets(t *testing.T) {
	pln(Str("<abc>").RemoveBrackets())
}

func TestBracketsPos(t *testing.T) {
	s := Str(`	"actor": {
		"name": "Team PB",
		"mbox": "mailto:teampb@example.com",
		"member": [
			{
				"name": "Andrew Downes",
				"account": {
					"homePage": "http://www.example.com",
					"name": "13936749"
				},
				"objectType": "Agent"
			},
			{
				"name": "Toby Nichols",
				"openid": "http://toby.openid.example.org/",
				"objectType": "Agent"
			},
			{
				"name": "Ena Hills",
				"mbox_sha1sum": "ebd31e95054c018b10727ccffd2ef2ec3a016ee9",
				"objectType": "Agent"
			}
		],
		"objectType": "Group"
	},`)

	pln(s.BracketsPos(BCurly, 2, 3))
}

func TestBracketPairCount(t *testing.T) {
	s := Str(`{"member": [
		{
			"name": "Andrew Downes",
			"account": {
				"homePage": "http://www.example.com",
				"name": "13936749"
			},
			"objectType": "Agent"
		},
		{
			"name": "Toby Nichols",
			"openid": "http://toby.openid.example.org/",
			"objectType": "Agent"
		},
		{
			"name": "Ena Hills",
			"mbox_sha1sum": "ebd31e95054c018b10727ccffd2ef2ec3a016ee9",
			"objectType": "Agent"
		}
	]}`)

	pln(s.BracketPairCount(BCurly))
	pln(s.IsJSON())
}

func TestIsXML(t *testing.T) {
	s := Str(` <StaffPersonal RefId="D3E34F41-9D75-101A-8C3D-00AA001A1652">
    <LocalId>946379881</LocalId>
    <StateProvinceId>C2345681</StateProvinceId>
    <!--
  <ElectronicIdList><ElectronicId Type="01">206655</ElectronicId></ElectronicIdList>
  -->
    <OtherIdList>
        <OtherId Type="0004">333333333</OtherId>
    </OtherIdList>    
</StaffPersonal> `)

	pln(s.IsXMLSegSimple())
}

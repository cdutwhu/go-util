package util

import (
	"testing"
)

func TestIsMadeFrom(t *testing.T) {
	fPln(Str("abc").IsMadeFrom('a', 'b'))
}

func TestLooseSearch(t *testing.T) {
	fPln(Str("ab C C		C 	* d	 e *.fc  * c			d").LooseSearch("Cde.", ' ', '\t', '*'))
}

func TestTrimInternal(t *testing.T) {
	fPln(Str(`*****abc*******abCCC**
***de***.**fc*c**d****`).TrimInternal('*', 1))
}

func TestBeCoveredInMapSIKeys(t *testing.T) {
	m := map[string]int{"abcd": 100, "def": 200}
	fPln(Str("abc").BeCoveredInMapSIKeys(m))
}

func TestCoverAnyKeyInMapSI(t *testing.T) {
	m := map[string]int{"abc": 100, "deff": 200}
	fPln(Str("def").CoverAnyKeyInMapSI(m))
}

func TestRemovePrefix(t *testing.T) {
	fPln(Str("sif.abc").RemovePrefix("sif."))
}

func TestRemoveSuffix(t *testing.T) {
	fPln(Str("sif.abc").RemoveSuffix("abc"))
}

func TestRemoveTailFromLast(t *testing.T) {
	fPln(Str("a.sif.abc").RemoveTailFromLast("."))
}

func TestRemoveBlankBefore(t *testing.T) {
	fPln(Str(`a            :  m   c		=   b  e 	 :  	 f`).RemoveBlankBefore("=", ":"))
}

func TestRemoveBlankNBefore(t *testing.T) {
	fPln(Str(`a a            :  m   c		:   b  e 	 :  	 f`).RemoveBlankNBefore(2, ":"))
}

func TestRemoveBlankAfter(t *testing.T) {
	fPln(Str(`a   :    t         c    =	b   e  		=		f`).RemoveBlankAfter("=", ":"))
}

func TestRemoveBlankNAfter(t *testing.T) {
	fPln(Str(`a a            :  m   c		=   b  e 	 :  	 f`).RemoveBlankNAfter(0, ":"))
}

func TestRemoveBlankNNear(t *testing.T) {
	fPln(Str(`a   :    t         c    =	b   e  		=		f`).RemoveBlankNNear(0, "="))
}

func TestRemoveBlankNear(t *testing.T) {
	fPln(Str(`a   :    t         c    =	b   e  		=		f`).RemoveBlankNear(":", "="))
}

func TestKeyValueMap(t *testing.T) {
	fPln(Str(`<abc a =	"dd"  c		= 	fff>>>>>`).KeyValueMap(' ', '=', '>'))
	fPln(Str(`<abc a	 	: 	"dd"  c   =			fff>>>>>`).KeyValueMap(' ', ':', '|'))
}

func TestKeyValuePair(t *testing.T) {
	k, v := Str(`   <abc a =     "dd"  c	: 	= 	fff>>>>>`).KeyValuePair("= ", ' ', ' ', true, false)
	fPln(k)
	fPln(v)
}

func TestMakeQuotes(t *testing.T) {
	fPln(Str("abc").MakeQuotes(QDouble))
}

func TestRemoveQuotes(t *testing.T) {
	fPln(Str("'abc'").RemoveQuotes())
}

func TestMakeBrackets(t *testing.T) {
	fPln(Str("abc").MakeBrackets(BCurly))
}

func TestRemoveBrackets(t *testing.T) {
	fPln(Str("<abc>").RemoveBrackets())
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

	fPln(s.BracketsPos(BCurly, 2, 3))
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

	fPln(s.BracketPairCount(BCurly))
	fPln(s.IsJSON())
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

	fPln(s.IsXMLSegSimple())
}

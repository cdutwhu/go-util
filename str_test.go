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
***de***.**fc*c**d****`).TrimAllInternal("*.\n\r\t"))
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
	fPln(Str("sif.abc").RmPrefix("sif."))
}

func TestRemoveSuffix(t *testing.T) {
	fPln(Str("sif.abc").RmSuffix("abc"))
}

func TestRemoveTailFromLast(t *testing.T) {
	fPln(Str("a.sif.abc").RmTailFromLast("."))
}

func TestRemoveBlankBefore(t *testing.T) {
	fPln(Str(`a            :  m   c		=   b  e 	 :  	 f`).RmBlankBefore("=", ":"))
}

func TestRemoveBlankNBefore(t *testing.T) {
	fPln(Str(`a a            :  m   c		:   b  e 	 :  	 f`).RmBlankNBefore(2, ":"))
}

func TestRemoveBlankAfter(t *testing.T) {
	fPln(Str(`a   :    t         c    =	b   e  		=		f`).RmBlankAfter("=", ":"))
}

func TestRemoveBlankNAfter(t *testing.T) {
	fPln(Str(`a a            :  m   c		=   b  e 	 :  	 f`).RmBlankNAfter(0, ":"))
}

func TestRemoveBlankNNear(t *testing.T) {
	fPln(Str(`a   :    t         c    =	b   e  		=		f`).RmBlankNNear(0, "="))
}

func TestRemoveBlankNear(t *testing.T) {
	fPln(Str(`a   :    t         c    =	b   e  		=		f`).RmBlankNear(":", "="))
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
	fPln(Str("abc").MkQuotes(QDouble))
}

func TestRemoveQuotes(t *testing.T) {
	fPln(Str("'abc'").RmQuotes())
}

func TestMakeBrackets(t *testing.T) {
	fPln(Str("abc").MkBrackets(BCurly))
}

func TestRemoveBrackets(t *testing.T) {
	fPln(Str("<abc>").RmBrackets())
}

func TestIndices(t *testing.T) {
	fPln(Str(`abc77abc77abc77ccc77`).Indices("c77"))
}

func TestBracketsPos(t *testing.T) {
	s := Str(`
		"Name": [
				[ ]
		]
	}
`)

	// fPln(s.QuotesPos(QDouble, 1))
	fPln(s.BracketsPos(BCurly, 1, 1))
	fPln(s.BracketDepth(BBox, 23))
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

func TestFieldsSeqContain(t *testing.T) {
	s0 := Str("-RefId + LocalId + StateProvinceId + OtherIdList + PersonInfo + Title + EmploymentStatus + MostRecent")
	fPln(s0.FieldsSeqContain("-RefId + LocalId + StateProvinceId + OtherIdList + PersonInfo + Title + EmploymentStatus + MostRecent", " + "))
}

func TestIsUUID(t *testing.T) {
	fPln(Str("fbd3036f-0f1c-4e98-b71c-d4cd61213f93").IsUUID())
}

package util

import (
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {

	s := Str(`Strait ’ Islander 最 Peoples 1’2  最 connections最`)
	// s := Str(`Strait , Islander , Peoples 1,2  , connections,`)
	fPln(s.SetEnC())

	if s.L() != 47 {
		t.Errorf("s.L() error\n")
	}
	if s.SegRep(7, 8, "TT").V() != "Strait TT Islander 最 Peoples 1’2  最 connections最" {
		t.Errorf("s.SegRep() error\n")
	}
	if s.S(1, ALL-2).V() != "trait ’ Islander 最 Peoples 1’2  最 connection" {
		t.Errorf("s.S() error\n")
	}
	if s.C(LAST) != '最' {
		t.Errorf("s.C() error\n")
	}
	if s.Idx("Islander") != 9 {
		t.Errorf("s.Idx() error\n")
	}
	if s.LIdx("Peoples") != 20 {
		t.Errorf("s.LIdx() error\n")
	}
	if !s.HP("St") {
		t.Errorf("s.HP() error\n")
	}
	if !s.HS("最") {
		t.Errorf("s.HS() error\n")
	}
}

func TestToInt(t *testing.T) {
	s := Str(`100`)
	fPln(s.ToInt())
	fPln(s.ToInt64())
}

func TestDefValue(t *testing.T) {
	s := Str(``)
	fPln(s.DefValue("abc"))
}

func TestRepeat(t *testing.T) {
	s := Str(`abc`)
	fPln(s.Repeat(3, ", "))
}

func TestHasAny(t *testing.T) {
	s := Str(`abc  d 最`)
	fPln(s.HasAny('d', 'c'))
	fPln(s.HasAny('最'))
	fPln(s.IsMadeFrom('e', 'b', 'c', 'd', '最', ' '))
}

func TestInArr(t *testing.T) {
	s := Str(`abc  d 最`)
	fPln(s.InArr("abcd", "abc  d 最"))
}

// *****************************

func TestIsMadeFrom(t *testing.T) {
	fPln(Str("abc").IsMadeFrom('a', 'b'))
}

func TestLooseSearchChars(t *testing.T) {

	fPln(Str("ab C 最t de C d e C. 最t * d TTT 最 e最td #* e *#  #.	C 	* d	 e *# . 最 .fc  * c	d").LooseSearchStrs("最t", "d", "e", ".", " \t*#"))

	fPln(Str(`type StaffPersonal {
		RefId最: String
		LocalId: Stringt最’
		Recent: 123
	}

	 type *  Recent  ** {
		SchoolLocalId: String
	}`).LooseSearchStrs("type", "Recent", "{", " \t*"))
}

func TestLooseSearchAny2Strs(t *testing.T) {

	fPln(Str("ab C 最t : 最	3 de C d e C. 最t * d : 最 1 e最td #* e *#  #").LooseSearchAny2Strs([]string{":"}, []string{"1", "2"}, " \t最"))

}

func TestTrimInternal(t *testing.T) {

	s := Str(`*****abc****最最最    1’2  ***abCCC**
	***de**最最最    1’2  	*.**fc*c**d**最最最    1’2  **`)

	fPln(s.TrimAllInternal("*. \n\r\t"))
	fPln(s.TrimAllLMR("*. \n\r\t"))

	// ’   最

	s = Str(`  Strait 最 Islander Peoples    最最最    1’2       connections with land, 最 sea and animals of their place 最    `)
	fPln(s.TrimInternal('最', 0))
	fPln(s.TrimInternal(' ', 1))

	fPln(s.S(0, 34))
	fPln(s.V()[:34])

	fPln(strings.Index(s.V(), "1"), s.Idx("1"))
	fPln(strings.Index(s.V(), "2"), s.Idx("2"))
}

func TestBeCoveredInMapSIKeys(t *testing.T) {
	m := map[string]int{"def": 100, "abcd最": 200}
	m1 := map[string]string{"def": "100", "abcd": "200"}
	fPln(Str("最").BeCoveredInMapSIKeys(m))
	fPln(Str("abc").InMapSIKeys(m))
	fPln(Str("100").InMapSSValues(m1))
}

func TestCoverAnyKeyInMapSI(t *testing.T) {
	m := map[string]int{"df": 100, "e": 200}
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
	fPln(Str(`a     最       :  m   c	最		=   b  e 最	   	 :  	 f`).RmBlankBefore("=", ":"))
	fPln(len(`最`))
}

func TestRemoveBlankNBefore(t *testing.T) {
	s := Str(`a a     最       :  m 最  c		: 最  b  e 最	 :   e3 最	 f`)
	fPln(s.RmBlankNBefore(1, ":"))
	fPln(s.RmBlankNBefore(2, ":"))
	fPln(s.RmBlankNBefore(3, ":"))
}

func TestRemoveBlankAfter(t *testing.T) {
	fPln(Str(`a   :		   最  t         c    =   最	b   e  		=	 最	  f`).RmBlankAfter("=", ":"))
}

func TestRemoveBlankNAfter(t *testing.T) {
	s := Str(`a a     最       :  m 最  c		: 最  b  e 最	 :   e3 最	 f`)
	fPln(s.RmBlankNAfter(0, ":"))
	fPln(s.RmBlankNAfter(1, ":"))
	fPln(s.RmBlankNAfter(2, ":"))
	fPln(s.RmBlankNAfter(3, ":"))
}

func TestRemoveBlankNNear(t *testing.T) {
	s := Str(`a a     最       :  m 最  c		: 最  b  e 最	 :   e3 最	 f`)
	fPln(s.RmBlankNNear(0, ":"))
	fPln(s.RmBlankNNear(1, ":"))
	fPln(s.RmBlankNNear(2, ":"))
	fPln(s.RmBlankNNear(3, ":"))
}

func TestRemoveBlankNear(t *testing.T) {
	fPln(Str(`a 最  :    t     最    c    =	b   e  	最	=		f`).RmBlankNear(":", "="))
}

func TestKeyValueMap(t *testing.T) {
	fPln(Str(`<abc a =	"最dd"  c最	= 	fff>>>>>`).KeyValueMap(' ', '=', 'T'))
	fPln(Str(`<abc a最 	: 	"dd" 最 c   :			最fff>>>>>最`).KeyValueMap(' ', ':', 'T'))
}

func TestKeyValuePair(t *testing.T) {
	k, v := Str(` 最  <abc a 88     "dd"  c: 	= 最*     fff>>>>>`).KeyValuePair("= ", "", "*", true, false)
	fPln(k)
	fPln(v)
}

func TestMakeQuotes(t *testing.T) {
	fPln(Str(" abc").MkQuotes(QDouble))
	fPln(Str(" abc").MkQuotes(QSingle))
	fPln(Str(`"abc"`).MkQuotes(QDouble))
	fPln(Str(`"abc"`).MkQuotes(QSingle))
	fPln(Str(`abc`).MkPrefix("."))
}

func TestRemoveQuotes(t *testing.T) {
	fPln(Str("'abc'").RmQuotes())
	fPln(Str("\" abc'").RmQuotes())
}

func TestMakeBrackets(t *testing.T) {
	fPln(Str(" abc ").MkBrackets(BCurly))
	fPln(Str(" abc ").MkBrackets(BBox))
	fPln(Str("[ abc ]").MkBrackets(BBox))
}

func TestRemoveBrackets(t *testing.T) {
	fPln(Str("<abc>").RmBrackets())
	fPln(Str("[abc]").RmBrackets())
	fPln(Str("{abc]").RmBrackets())
}

func TestIndices(t *testing.T) {
	fPln(Str(`ab最c77ab最c77a最bc77c最cc77`).Indices("c最"))
	//fPln(Str(`abc77abc77abc77ccc7c7`).IndicesASCII("c7", ' ', ' '))
}

func TestBracketsPos(t *testing.T) {

	s1 := Str(` abc def " "  `)
	fPln(s1.QuotesPos(QDouble, 1))
	fPln(s1.QuotePairCount(QDouble))

	// 	s := Str(`
	// 		"Name": 最 [ 最
	// 				[ 最 ]   [ TTT ]
	// 		]
	// 	}
	// `)

	// 	fPln(s.QuotesPos(QDouble, 1))
	// 	fPln(s.BracketsPos(BBox, 1, 2))
	// 	fPln(s.BracketsPos(BBox, 2, 2))
	// 	fPln(s.BracketDepth(BBox, 20))
	// 	fPln(s.BracketDepth(BBox, 21))
	// 	fPln(s.BracketDepth(BBox, 25))
	// 	fPln(s.BracketDepth(BBox, 26))
}

func TestBracketPairCount(t *testing.T) {
	s := Str(`{"member": [
		{
			"name": "Andrew Downes 最",
			"account": {
				"homePage": "http://www.example.com",
				"name": "13936749"
			},
			"objectType": "Agent"
		},
		{
			"name": "Toby Nichols 最",
			"openid": "http://toby.openid.example.org/",
			"objectType": "Agent"
		},
		{
			"name": "Ena Hills 最",
			"mbox_sha1sum": "ebd31e95054c018b10727ccffd2ef2ec3a016ee9",
			"objectType": "Agent"
		}
	]
	}`)

	fPln(s.BracketPairCount(BCurly))
	fPln(s.IsJSON())
}

func TestIsXML(t *testing.T) {
	// 	s := Str(` <最StaffPersonal最 RefId="D3E34F41-9D75-101A-8C3D-00AA001A1652">
	//     <LocalId>946379881</LocalId>
	//     <StateProvinceId>C2345681</StateProvinceId>
	//     <!--
	//   <ElectronicIdList><ElectronicId Type="01">206655</ElectronicId></ElectronicIdList>
	//   -->
	//     <OtherIdList>
	//         <OtherId Type="0004">333333333</OtherId>
	//     </OtherIdList>
	// </最StaffPersonal最> `)

	// fPln(s.IsXMLSegSimple())
}

func TestFieldsSeqContain(t *testing.T) {
	s0 := Str("-RefId + LocalId + StateProvinceId + OtherIdList + PersonInfo + Title + EmploymentStatus + MostRecent")
	fPln(s0.FieldsSeqContain("-RefId + LocalId + StateProvinceId + OtherIdList + PersonInfo + Title + EmploymentStatus + MostRecent", " + "))
}

func TestIsUUID(t *testing.T) {
	fPln(Str("fbd3036f-0f1c-4e98-b71c-d4cd61213f93").IsUUID())
}

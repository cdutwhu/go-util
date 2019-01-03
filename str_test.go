package util

import (
	"fmt"
	"testing"
)

var pln = fmt.Println

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

func TestRemoveBlankAfter(t *testing.T) {
	pln(Str(`a   :    t         c    =	b   e  		=		f`).RemoveBlankAfter("=", ":"))
}

func TestRemoveBlankNear(t *testing.T) {
	pln(Str(`a   :    t         c    =	b   e  		=		f`).RemoveBlankNear(":", "="))
}

func TestKeyValueMap(t *testing.T) {
	pln(Str(`<abc a =	"dd"  c		= 	fff>>>>>`).KeyValueMap(' ', '=', '>'))
	pln(Str(`<abc a	 	: 	"dd"  c   =			fff>>>>>`).KeyValueMap(' ', ':', '|'))
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

	pln(s.BracketsPos(BCurly, 1, 1))
}

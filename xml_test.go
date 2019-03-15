package util

import (
	"io/ioutil"
	"testing"
)

func TestXML(t *testing.T) {
	sifbytes, _ := ioutil.ReadFile("./sif.xml")
	sif := Str(sifbytes)
	sif.SetEnC()

	tag, xml, l, r := sif.XMLSegPos(3, 1)
	fPln(tag)
	fPln(xml)
	fPln(l, r)

	fPln(sif.XMLSegsCount())

	tag, xml, l, r = sif.XMLSegPos(1, 726)
	fPln(tag)
	fPln(xml)
	fPln(l, r)

	fPln(Str(xml).XMLSegsCount())
}

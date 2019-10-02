package util

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	sCtn = strings.Contains
	sSpl = strings.Split
	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
	fSp  = fmt.Sprint

	pc = PanicOnCondition
	ph = PanicHandle
	pe = PanicOnError

	defLog = path.Join(os.TempDir(), "/GoErrorLog.txt")
)

const (
	NOFATAL = "#NOFATAL#"
)

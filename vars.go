package util

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	sCtn = strings.Contains

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
	fSp  = fmt.Sprint

	PC = PanicOnCondition
	PE = PanicOnError
	PH = PanicHandle

	defLog = path.Join(os.TempDir(), "/GoErrorLog.txt")
)

const (
	NONFATALMARK = "#NONFATAL#"
)

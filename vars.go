package util

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	sCtn = strings.Contains
	sCnt = strings.Count
	sJ   = strings.Join
	sSpl = strings.Split
	sFF  = strings.FieldsFunc
	sRep = strings.Replace

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
	fSp  = fmt.Sprint

	sc2Int   = strconv.ParseInt
	sc2Float = strconv.ParseFloat

	PC  = PanicOnCondition
	PE  = PanicOnError
	PE1 = PanicOnError1
	PH  = PanicHandle
	PHx = PanicHandleEx

	defLog = path.Join(os.TempDir(), "/GoErrorLog.txt")
)

const (
	NOFATALMARK = "#NOFATAL#"
)

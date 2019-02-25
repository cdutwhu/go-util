package util

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	sHP  = strings.HasPrefix
	sHS  = strings.HasSuffix
	sCtn = strings.Contains
	sCnt = strings.Count
	sI   = strings.Index
	sLI  = strings.LastIndex
	sJ   = strings.Join
	sSpl = strings.Split
	sT   = strings.Trim
	sTL  = strings.TrimLeft
	sTR  = strings.TrimRight
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

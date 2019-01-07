package util

import (
	"fmt"
	"strings"
)

var (
	sHP = strings.HasPrefix
	sHS = strings.HasSuffix
	sC  = strings.Contains
	sI  = strings.Index
	sLI = strings.LastIndex
	sJ  = strings.Join
	sS  = strings.Split
	sT  = strings.Trim
	sTL = strings.TrimLeft
	sTR = strings.TrimRight
	sFF = strings.FieldsFunc
	pln = fmt.Println
	pf  = fmt.Printf
)

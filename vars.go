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

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf

	PC   = PanicOnCondition
	PE   = PanicOnError
	PE1  = PanicOnError1
	PH   = PanicHandle
	PHEx = PanicHandleEx
)

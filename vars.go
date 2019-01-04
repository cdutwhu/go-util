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
	pln = fmt.Println
	pf  = fmt.Printf
)

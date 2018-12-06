package util

import (
	"strings"
)

// Str is string 'class'
type Str string

// V : get string value
func (s *Str) V() string {
	return string(*s)
}

// DefValue : if s is blank, assign it with input string value, otherwise keep its current value
func (s Str) DefValue(def string) string {
	if len(s) == 0 {
		return def
	}
	return s.V()
}

// InArr : check if at least a same value exists in string array
func (s Str) InArr(arr ...string) bool {
	for _, a := range arr {
		if s.V() == a {
			return true
		}
	}
	return false
}

// RemoveQuotations : Remove single or double Quotations from a string. If no quotations, do nothing
func (s Str) RemoveQuotations() string {
	if strings.HasPrefix(s.V(), "\"") && strings.HasSuffix(s.V(), "\"") {
		return s.V()[1 : len(s.V())-1]
	}
	if strings.HasPrefix(s.V(), "'") && strings.HasSuffix(s.V(), "'") {
		return s.V()[1 : len(s.V())-1]
	}
	return s.V()
}

// AllAreIdentical : check all the input strings are identical
func AllAreIdentical(arr ...string) bool {
	if len(arr) > 1 {
		for _, a := range arr {
			if arr[0] != a {
				return false
			}
		}
	}
	return true
}

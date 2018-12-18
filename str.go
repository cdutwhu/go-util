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

// MakePrefix :
func (s Str) MakePrefix(prefix string) string {
	if !strings.HasPrefix(s.V(), prefix) {
		return prefix + s.V()
	}
	return s.V()
}

// RemovePrefix :
func (s Str) RemovePrefix(prefix string) string {
	if strings.HasPrefix(s.V(), prefix) {
		return s.V()[len(prefix):len(s.V())]
	}
	return s.V()
}

// MakeSuffix :
func (s Str) MakeSuffix(suffix string) string {
	if !strings.HasSuffix(s.V(), suffix) {
		return s.V() + suffix
	}
	return s.V()
}

// RemoveSuffix :
func (s Str) RemoveSuffix(suffix string) string {
	if strings.HasSuffix(s.V(), suffix) {
		return s.V()[:len(s.V())-len(suffix)]
	}
	return s.V()
}

// RemoveTailFromLast :
func (s Str) RemoveTailFromLast(tail string) string {
	if i := strings.LastIndex(s.V(), tail); i >= 0 {
		return s.V()[:i]
	}
	return s.V()
}

// KeyValueMap :
func (s Str) KeyValueMap(delimiter, assign, terminator rune) (r map[string]string) {
	r = make(map[string]string)
	str := s.V()
	if pt := strings.Index(str, string(terminator)); pt > 0 {
		str = str[:pt]
	}
	for _, kv := range strings.FieldsFunc(str, func(c rune) bool { return c == delimiter }) {
		if strings.Contains(kv, string(assign)) {
			kvpair := strings.Split(kv, string(assign))
			r[kvpair[0]] = Str(kvpair[1]).RemoveQuotations()
		}
	}
	return
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

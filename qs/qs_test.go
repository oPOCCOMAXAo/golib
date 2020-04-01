package qs

import (
	"testing"
)

type testStringify struct {
	In       map[string]string
	Out      []string
	Computed string
}

func (t testStringify) Ok() bool {
	for _, v := range t.Out {
		if v == t.Computed {
			return true
		}
	}
	return false
}

func TestStringify(t *testing.T) {
	var testStringify = []testStringify{
		{
			In:  map[string]string{"a": "3", "b": "", "": ""},
			Out: []string{"a=3&b=&=", "a=3&=&b=", "=&a=3&b=", "=&b=&a=3", "b=&a=3&=", "b=&=&a=3"},
		},
		{
			In:  nil,
			Out: []string{""},
		},
	}
	for _, v := range testStringify {
		v.Computed = Stringify(v.In)
		if !v.Ok() {
			t.Errorf("Wanted one of %v, but received %s at input %#v\n", v.Out, v.Computed, v.In)
		}
	}
}

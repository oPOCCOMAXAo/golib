package qs

import (
	"testing"
)

type testStringify struct {
	In       map[string]string
	Out      string
	Computed string
}

func TestStringify(t *testing.T) {
	var testStringify = []testStringify{
		{
			In:  map[string]string{"a": "3", "b": "", "": ""},
			Out: "a=3&b=&=",
		},
		{
			In:  nil,
			Out: "",
		},
	}
	for _, v := range testStringify {
		v.Computed = Stringify(v.In)
		if v.Computed != v.Out {
			t.Errorf("Wanted %s, but received %s at input %#v\n", v.Out, v.Computed, v.In)
		}
	}
}

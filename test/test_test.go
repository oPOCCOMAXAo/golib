package test

import "testing"

type testStruct struct {
	s string
	i int
	f float64
}

func TestCheckValue(t *testing.T) {
	CheckValue(t, "bool", true, true)
	CheckValue(t, "bool", false, false)
	CheckValue(t, "int", 12345, 12345)
	CheckValue(t, "float", 123.45, 123.45)
	CheckValue(t, "string", "", "")
	CheckValue(t, "string", "123", "123")
	CheckValue(t, "object", testStruct{
		s: "test",
		i: 1234,
		f: 123.45,
	}, testStruct{
		s: "test",
		i: 1234,
		f: 123.45,
	})
}

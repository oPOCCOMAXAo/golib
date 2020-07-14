package test

import "testing"

func CheckValue(t *testing.T, name string, expected interface{}, got interface{}) {
	if expected != got {
		t.Errorf("%s: expected %v, but got %v", name, expected, got)
	}
}

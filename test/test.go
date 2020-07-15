package test

import (
	"reflect"
	"testing"
)

func CheckValue(t *testing.T, name string, expected interface{}, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s: expected %v, but got %v", name, expected, got)
	}
}

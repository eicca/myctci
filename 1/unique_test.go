package base

import (
	"testing"
)

var uniqueExamples = map[string]bool{
	"foo":      false,
	"bar":      true,
	"":         true,
	"ютф8":     true,
	"ютф8т":    false,
	"spaces  ": false,
}

func TestUniqueMap(t *testing.T) {
	testUnique(t, uniqueMap)
}

func TestUniqueSort(t *testing.T) {
	testUnique(t, uniqueSort)
}

func testUnique(t *testing.T, fn func(string) bool) {
	for in, out := range uniqueExamples {
		res := fn(in)
		if res != out {
			t.Errorf("unique returned %t for %s, expected %t", res, in, out)
		}
	}
}

package base

import (
	"testing"
)

func TestGetSet(t *testing.T) {
	a := NewArrayList(3)
	a.Set(0, 20)
	if a.Get(0) != 20 {
		t.Errorf("Expected 20 at index 0")
	}
}

func TestAppend(t *testing.T) {
	a := NewArrayList(0)
	a.Append(10)
	a.Append(20)
	a.Append(30)
	if a.Get(0) != 10 {
		t.Errorf("Expected 10 at index 0")
	}
	if a.Get(2) != 30 {
		t.Errorf("Expected 30 at index 2")
	}
}

func TestLen(t *testing.T) {
	a := NewArrayList(0)
	a.Append(1)
	a.Append(2)
	a.Append(3)
	if a.Len() != 3 {
		t.Error("Expected Len() to return 3")
	}
}

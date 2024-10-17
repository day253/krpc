package objects

import (
	"testing"
)

func TestString(t *testing.T) {
	type Student struct {
		Name string
	}
	s := Student{
		Name: "test",
	}
	if String(s) != "{\"Name\":\"test\"}" {
		t.Errorf("String(%v)", s)
	}
}

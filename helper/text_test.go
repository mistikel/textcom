package helper

import (
	"testing"
)

func TestUpperCase(t *testing.T) {
	expect := "HELLO WORLD"
	input := "Hello world"
	got := UpperCase(input)

	if got != expect {
		t.Errorf("Want: [%s] got [%s]", expect, got)
	}
}

func TestLowerCase(t *testing.T) {
	expect := "hello world"
	input := "Hello world"
	got := LowerCase(input)

	if got != expect {
		t.Errorf("Want: [%s] got [%s]", expect, got)
	}
}

func TestCamelCase(t *testing.T) {
	expect := "hElLo wOrLd"
	input := "Hello world"
	got := CamelCase(input)

	if got != expect {
		t.Errorf("Want: [%s] got [%s]", expect, got)
	}
}

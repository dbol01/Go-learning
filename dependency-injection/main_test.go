package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Dan")

	got := buffer.String()
	expected := "Hello Dan!"

	if got != expected {
		t.Errorf("got %q wanted %q", got, expected)
	}
}

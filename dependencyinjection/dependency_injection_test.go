package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Simon")

	got := buffer.String()
	want := "Hello, Simon"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

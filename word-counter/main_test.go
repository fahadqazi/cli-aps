package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	buf := bytes.NewBufferString("one two three")

	got := count(buf, true, false, false)
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountLines(t *testing.T) {
	buf := bytes.NewBufferString("one two three")

	got := count(buf, false, true, false)
	want := 1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountBytes(t *testing.T) {
	buf := bytes.NewBufferString("one two three")

	got := count(buf, false, false, true)
	want := 13

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

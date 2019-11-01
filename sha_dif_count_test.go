package main

import (
	"testing"
)

func TestGetHashDiff(t *testing.T) {
	c1 := "x"
	c2 := "X"

	_, _, got := getHashDiff(c1, c2)
	want := 31

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

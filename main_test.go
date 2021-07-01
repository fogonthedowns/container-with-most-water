package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	out := maxArea([]int{4, 3, 2, 1, 4})
	want := 16
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}
}

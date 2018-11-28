package main

import (
	"testing"
)

func TestIncrementCase1(t *testing.T) {
	if Increment(2) != 3 {
		t.Error("Expected increment of 2 to be 3")
	}
}
func TestIncrementCase2(t *testing.T) {
	if Increment(5) != 7 {
		t.Error("Expected increment of 5 to be 7")
	}
}

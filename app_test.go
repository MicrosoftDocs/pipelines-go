package main

import (
	"testing"
)

func TestIncrementCase0(t *testing.T) {
	tables := []struct {
		x int
		y int
	}{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
	}

	for _, table := range tables {
		res := Increment(table.x)
		if res != table.y {
			t.Errorf("Increment of (%d) was incorrect returned as (%d) instead of (%d)!", table.x, table.y, res)
		}
	}
}

func TestIncrementCase1(t *testing.T) {
	if Increment(2) != 3 {
		t.Error("Expected increment of 2 to be 3")
	}
}
func TestIncrementCase2(t *testing.T) {
	if Increment(5) != 6 {
		t.Error("Expected increment of 5 to be 6")
	}
}


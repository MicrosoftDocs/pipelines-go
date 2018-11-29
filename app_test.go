package main

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	tables := []struct {
		x int
		y int
	}{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 6},
		{5, 7},
	}

	for _, table := range tables {
		res := Increment(table.x)
		if res != table.y {
			t.Errorf("Increment of (%d) was incorrect returned as (%d) instead of (%d)!", table.x, table.y, res)
		}
	}
}

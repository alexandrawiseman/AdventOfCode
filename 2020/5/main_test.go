package main

import (
	"testing"
)

func TestConversionToSeat(t *testing.T) {
	row, col := convertLineIntoRowAndCol("FBFBBFFRLR")
	if row != 44 || col != 5 {
		t.Errorf("Got %v,%v, expected 44,5", row, col)
	}
}

func TestFindSeat(t *testing.T) {
	input := [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
		{false, false, false},
		{true, false, true},
		{false, false, false},
	}

	row, col := findEmptySeat(input)
	if row != 4 && col != 1 {
		t.Errorf("Got %v,%v, expected 4,1", row, col)
	}
}

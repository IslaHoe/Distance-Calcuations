package main

import (
	"math"
	"testing"
)

func TestConvertToRadians(t *testing.T) {
	tables := []struct {
		rad float64
		deg float64
	}{
		{1, 57.2958},
		{5, 286.479},
		{90, 5156.6},
		{180, 10313.2},
	}

	for _, table := range tables {
		total := convertToRadians(table.deg)
		if math.Round(total) != math.Round(table.rad) {
			t.Errorf("Conversion of (%b) was incorrect, got: %b, want: %b.", table.deg, total, table.rad)
		}
	}
}

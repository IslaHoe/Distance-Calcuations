package main

import (
	"fmt"
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
			t.Errorf("Conversion of (%v) was incorrect, got: %v, want: %v.", table.deg, total, table.rad)
		}
	}
}

func TestCalcuateDistance(t *testing.T) {
	coords := []struct {
		lat2 float64
		lng2 float64
		dist float64
	}{
		{52.833502, -8.522366, 161.35},
		{54.080556, -6.361944, 82.69},
		{52.228056, -7.915833, 166.44},
		{53.1302756, -6.2397222, 23.29},
	}

	for _, coord := range coords {
		total := calculateDistance(coord.lat2, coord.lng2)
		fmt.Println(math.Round(total))
		if math.Round(total) != math.Round(coord.dist) {
			t.Errorf("Distance is incorrect got: %v, want: %v.", total, coord.dist)
		}
	}
}

func Test_calculateDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lng1 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateDistance(tt.args.lat1, tt.args.lng1); got != tt.want {
				t.Errorf("calculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

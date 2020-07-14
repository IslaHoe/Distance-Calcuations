package main

import (
	"math"
	"testing"
)

func Test_ConvertToRadians(t *testing.T) {
	tests := []struct {
		rad float64
		deg float64
	}{
		{1, 57.2958},
		{5, 286.479},
		{90, 5156.6},
		{180, 10313.2},
	}

	for _, test := range tests {
		total := convertToRadians(test.deg)
		if math.Round(total) != math.Round(test.rad) {
			t.Errorf("Conversion of (%v) was incorrect, got: %v, want: %v.", test.deg, total, test.rad)
		}
	}
}

func Test_calculateDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lng1 float64
		lat2 float64
		lng2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				lat1: 53.339428,
				lng1: -6.257664,
				lat2: 52.986375,
				lng2: -6.043701,
			},
			want: math.Round(41.77),
		},
		{
			name: "test2",
			args: args{
				lat1: 53.339428,
				lng1: -6.257664,
				lat2: 53.1229599,
				lng2: -6.2705202,
			},
			want: math.Round(24.08),
		},
		{
			name: "test3",
			args: args{
				lat1: 53.339428,
				lng1: -6.2576,
				lat2: 55.033,
				lng2: -6.2705202,
			},
			want: math.Round(188.31),
		},
		{
			name: "test4",
			args: args{
				lat1: 52.240382,
				lng1: -6.257664,
				lat2: 52.240382,
				lng2: -6.972413,
			},
			want: math.Round(48.66),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.Round(calculateDistance(tt.args.lat1, tt.args.lng1, tt.args.lat2, tt.args.lng2)); got != tt.want {
				t.Errorf("calculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

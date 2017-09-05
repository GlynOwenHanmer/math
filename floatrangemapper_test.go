package math_test

import (
	"testing"
	"github.com/GlynOwenHanmer/math"
)

func TestFloatRangeMapper_Map(t *testing.T) {
	tests := []struct{
		math.FloatRangeMapper
		io map[float64]float64
	}{
		{
			FloatRangeMapper:math.FloatRangeMapper{
				Source:math.FloatRange{1,2},
				Target:math.FloatRange{2,4},
			},
			io:map[float64]float64{
				0.5:1,
				1:2,
				1.5:3,
				2:4,
				2.5:5,
			},
		},
		{
			FloatRangeMapper:math.FloatRangeMapper{
				Source:math.FloatRange{2,4},
				Target:math.FloatRange{2,-2},
			},
			io:map[float64]float64{
				1:4,
				2:2,
				3:0,
				4:-2,
				5:-4,
			},
		},
	}
	for _, test := range tests {
		for value, expected := range test.io {
			actual := test.FloatRangeMapper.Map(value)
			if actual != expected {
				t.Errorf("Expected %f, got %f", expected, actual)
			}
		}
	}
}
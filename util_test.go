package spatialID

import (
	"testing"
)

func TestCalculateArithmaticShift(t *testing.T) {
	data := []struct {
		index          int64
		shift          int64
		expectedOutput int64
	}{
		{index: 1, shift: 25, expectedOutput: 33554432},
		{index: 1, shift: 23, expectedOutput: 8388608},
		{index: 1, shift: 0, expectedOutput: 1},
		{index: 47, shift: 0, expectedOutput: 47},
		{index: 47, shift: -1, expectedOutput: 23},
		{index: 47, shift: -100, expectedOutput: 0},
		{index: -47, shift: 1, expectedOutput: -94},
		{index: -47, shift: 0, expectedOutput: -47},
		{index: -47, shift: -1, expectedOutput: -24},
		{index: -47, shift: -3, expectedOutput: -6},
		{index: 123456, shift: -1, expectedOutput: 61728},
	}

	for _, p := range data {
		result := CalculateArithmeticShift(p.index, p.shift)
		if result != p.expectedOutput {
			t.Log(t.Name())
			t.Errorf("calculateMinVerticalIndex(%v, %v) == %v, result: %v", p.index, p.shift, p.expectedOutput, result)
		}
	}
}

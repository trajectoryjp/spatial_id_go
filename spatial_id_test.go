package spatialID

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpatialIDOverlaps(t *testing.T) {
	testCases := []struct {
		id             SpatialID
		another        SpatialID
		expectedResult bool
	}{
		{
			id:             SpatialID{23, 0, 0, 0},
			another:        SpatialID{23, 0, 0, 0},
			expectedResult: true,
		},
		{
			id:             SpatialID{22, 0, 0, 0},
			another:        SpatialID{23, 1, 1, 1},
			expectedResult: true,
		},
		{
			id:             SpatialID{23, 0, 0, 0},
			another:        SpatialID{23, 1, 0, 0},
			expectedResult: false,
		},
		{
			id:             SpatialID{23, 0, 0, 0},
			another:        SpatialID{23, 0, 1, 0},
			expectedResult: false,
		},
		{
			id:             SpatialID{23, 0, 0, 0},
			another:        SpatialID{23, 0, 0, 1},
			expectedResult: false,
		},
		{
			id:             SpatialID{22, 1, 1, 1},
			another:        SpatialID{23, 1, 1, 1},
			expectedResult: false,
		},
	}

	for testCaseIndex, testCase := range testCases {
		for swapIndex := range 2 {
			actualResult := testCase.id.Overlaps(testCase.another)
			assert.Equal(
				t,
				testCase.expectedResult,
				actualResult,
				fmt.Sprintf("testCaseIndex: %d, swapIndex: %d", testCaseIndex, swapIndex),
			)
		}
		testCase.id, testCase.another = testCase.another, testCase.id
	}
}

func TestMergeSpatialIDs(t *testing.T) {
	testCases := []struct {
		ids            []*SpatialID
		expectedResult []*SpatialID
	}{
		{
			ids: []*SpatialID{
				{23, 1, 2, 4},
			},
			expectedResult: []*SpatialID{
				{23, 1, 2, 4},
			},
		},
		{
			ids: []*SpatialID{
				{23, 0, 0, 0},
				{23, 0, 0, 0},
				{24, 1, 1, 1},
			},
			expectedResult: []*SpatialID{
				{23, 0, 0, 0},
			},
		},
		{
			ids: []*SpatialID{
				{23, 0, 0, 0},
				{23, 0, 0, 1},
				{23, 0, 1, 0},
				{23, 0, 1, 1},
				{23, 1, 0, 0},
				{23, 1, 0, 1},
				{23, 1, 1, 0},
			},
			expectedResult: []*SpatialID{
				{23, 0, 0, 0},
				{23, 0, 0, 1},
				{23, 0, 1, 0},
				{23, 0, 1, 1},
				{23, 1, 0, 0},
				{23, 1, 0, 1},
				{23, 1, 1, 0},
			},
		},
		{
			ids: []*SpatialID{
				{23, 0, 0, 0},
				{23, 0, 0, 1},
				{23, 0, 1, 0},
				{23, 0, 1, 1},
				{23, 1, 0, 0},
				{23, 1, 0, 1},
				{23, 1, 1, 0},
				{23, 1, 1, 1},
				{23, 0, 1, 2},
			},
			expectedResult: []*SpatialID{
				{22, 0, 0, 0},
				{23, 0, 1, 2},
			},
		},
		{
			ids: []*SpatialID{
				{23, -1, 0, 0},
				{23, -1, 0, 1},
				{23, -1, 1, 0},
				{23, -1, 1, 1},
				{23, -2, 0, 0},
				{23, -2, 0, 1},
				{23, -2, 1, 0},
				{23, -2, 1, 1},
			},
			expectedResult: []*SpatialID{
				{22, -1, 0, 0},
			},
		},
		{
			ids: []*SpatialID{
				{23, 0, 0, 0},
				{24, 0, 0, 2},
				{24, 0, 0, 3},
				{24, 0, 1, 2},
				{24, 0, 1, 3},
				{24, 1, 0, 2},
				{24, 1, 0, 3},
				{24, 1, 1, 2},
				{24, 1, 1, 3},
				{23, 0, 1, 0},
				{23, 0, 1, 1},
				{23, 1, 0, 0},
				{23, 1, 0, 1},
				{23, 1, 1, 0},
				{23, 1, 1, 1},
			},
			expectedResult: []*SpatialID{
				{22, 0, 0, 0},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		if testCaseIndex != 3 {
			continue
		}
		actualResult := MergeSpatialIDs(testCase.ids)
		assert.ElementsMatch(
			t,
			testCase.expectedResult,
			actualResult,
			fmt.Sprintf("testCaseIndex: %d", testCaseIndex),
		)
	}
}

func BenchmarkMergeSpatialIDs(b *testing.B) {
	type Benchmark struct {
		name string
		ids  []*SpatialID
	}

	benchmarks := []*Benchmark{}
	for _, n := range []int64{16} {
		for _, z := range []int8{23} {
			ids := []*SpatialID{}
			for f := range n {
				for x := range n {
					for y := range n {
						id, _ := NewSpatialID(z, f, x, y)
						ids = append(ids, id)
					}
				}
			}
			benchmarks = append(
				benchmarks,
				&Benchmark{fmt.Sprintf("n=%d,z=%d", n, z), ids},
			)
		}
	}

	for _, benchmark := range benchmarks {
		b.Run(
			benchmark.name,
			func(b *testing.B) {
				MergeSpatialIDs(benchmark.ids)
			},
		)

		startTotalAlloc := getTotalAlloc()
		MergeSpatialIDs(benchmark.ids)
		fmt.Printf("BenchmarkMergeSpatialIDs/%s\t%d B\n", benchmark.name, getTotalAlloc()-startTotalAlloc)
	}
}

func getTotalAlloc() uint64 {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return mem.TotalAlloc
}

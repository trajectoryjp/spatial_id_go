package spatialID

import (
	"fmt"
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

func TestSummarizeSpatialIDs(t *testing.T) {
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
		actualResult := MergeSpatialIDs(testCase.ids)
		assert.ElementsMatch(
			t,
			testCase.expectedResult,
			actualResult,
			fmt.Sprintf("testCaseIndex: %d", testCaseIndex),
		)
	}
}

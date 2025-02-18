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

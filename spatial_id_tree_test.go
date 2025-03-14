package spatialID

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	minF = int64(-1)<<MaxZ + 1
	maxF = int64(1)<<MaxZ - 1
	maxX = int64(1)<<MaxZ - 1
	maxY = int64(1)<<MaxZ - 1
)

var (
	z0Tree   *SpatialIDTree = NewSpatialIDTree([]SpatialID{{0, 0, 0, 0}})
	maxZTree *SpatialIDTree = NewSpatialIDTree(
		[]SpatialID{
			{MaxZ, 0, 0, 0},
			{MaxZ, minF, maxX, maxY},
			{MaxZ, maxF, maxX, maxY},
		},
	)
	z23Tree *SpatialIDTree = NewSpatialIDTree(
		[]SpatialID{
			{23, 1, 1, 1},
		},
	)
)

func TestSpatialIDTreeOverlaps(t *testing.T) {
	testCases := []struct {
		tree           *SpatialIDTree
		ids            []SpatialID
		expectedResult bool
	}{
		{
			tree:           z0Tree,
			ids:            []SpatialID{{0, 0, 0, 0}},
			expectedResult: true,
		},
		{
			tree:           maxZTree,
			ids:            []SpatialID{{MaxZ, 0, 0, 0}},
			expectedResult: true,
		},
		{
			tree:           maxZTree,
			ids:            []SpatialID{{MaxZ, minF, maxX, maxY}},
			expectedResult: true,
		},
		{
			tree:           maxZTree,
			ids:            []SpatialID{{MaxZ, maxF, maxX, maxY}},
			expectedResult: true,
		},
		{
			tree:           z23Tree,
			ids:            []SpatialID{{22, 0, 0, 0}},
			expectedResult: true,
		},
		{
			tree:           z23Tree,
			ids:            []SpatialID{{24, 2, 2, 2}},
			expectedResult: true,
		},
		{
			tree:           z23Tree,
			ids:            []SpatialID{{24, 3, 3, 3}},
			expectedResult: true,
		},
		{
			tree:           z23Tree,
			ids:            []SpatialID{{23, 0, 0, 1}, {23, 0, 1, 0}, {23, 1, 0, 0}},
			expectedResult: false,
		},
		{
			tree:           z23Tree,
			ids:            []SpatialID{{22, 1, 1, 1}, {23, 0, 0, 0}, {23, 2, 2, 2}, {24, 0, 0, 0}, {24, 1, 1, 1}},
			expectedResult: false,
		},
	}

	for testCaseIndex, testCase := range testCases {
		actualResult := testCase.tree.Overlaps(testCase.ids)
		assert.Equal(
			t,
			testCase.expectedResult,
			actualResult,
			fmt.Sprintf("testCaseIndex: %d", testCaseIndex),
		)
	}
}

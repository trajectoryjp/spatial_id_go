package spatialID

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpatialIDDetector(t *testing.T) {
	testCases := []struct {
		spatialIDs       SpatialIDs
		targetSpatialIDs SpatialIDs
		expectedResult   bool
	}{
		// ズームレベル最小、インデックス最小値/最大値
		{
			spatialIDs:       SpatialIDs{{0, 0, 0, 0}},
			targetSpatialIDs: SpatialIDs{{0, 0, 0, 0}},
			expectedResult:   true,
		},
		// ズームレベル最大、インデックス最小値
		{
			spatialIDs:       SpatialIDs{{MaxZ, 0, 0, 0}},
			targetSpatialIDs: SpatialIDs{{MaxZ, 0, 0, 0}},
			expectedResult:   true,
		},
		// ズームレベル最大、インデックス最大値
		{
			spatialIDs:       SpatialIDs{{MaxZ, int64(1)<<MaxZ - 1, int64(1)<<MaxZ - 1, int64(1)<<MaxZ - 1}},
			targetSpatialIDs: SpatialIDs{{MaxZ, int64(1)<<MaxZ - 1, int64(1)<<MaxZ - 1, int64(1)<<MaxZ - 1}},
			expectedResult:   true,
		},
		// ズームレベル違い
		{
			spatialIDs:       SpatialIDs{{22, 0, 0, 0}},
			targetSpatialIDs: SpatialIDs{{23, 1, 1, 1}},
			expectedResult:   true,
		},
		{
			spatialIDs:       SpatialIDs{{24, 2, 2, 2}},
			targetSpatialIDs: SpatialIDs{{23, 1, 1, 1}},
			expectedResult:   true,
		},
		{
			spatialIDs:       SpatialIDs{{24, 3, 3, 3}},
			targetSpatialIDs: SpatialIDs{{23, 1, 1, 1}},
			expectedResult:   true,
		},
		{
			spatialIDs:       SpatialIDs{{22, 1, 1, 1}, {23, 0, 0, 0}, {23, 2, 2, 2}, {24, 0, 0, 0}, {24, 1, 1, 1}},
			targetSpatialIDs: SpatialIDs{{23, 1, 1, 1}},
			expectedResult:   false,
		},
	}

	for testCaseIndex, testCase := range testCases {
		for swapIndex := range 2 {
			for detectorIndex, newSpatialIDDetector := range []func(SpatialIDs) SpatialIDDetector{
				NewSpatialIDGreedyDetector,
				NewSpatialIDTreeDetector,
			} {
				spatialIDDetector := newSpatialIDDetector(testCase.spatialIDs)
				actualResult := spatialIDDetector.IsOverlap(testCase.targetSpatialIDs)
				assert.Equal(
					t,
					testCase.expectedResult,
					actualResult,
					fmt.Sprintf("testCaseIndex: %d, swapIndex: %d, detectorIndex: %d", testCaseIndex, swapIndex, detectorIndex),
				)
			}
			testCase.spatialIDs, testCase.targetSpatialIDs = testCase.targetSpatialIDs, testCase.spatialIDs
		}
	}
}

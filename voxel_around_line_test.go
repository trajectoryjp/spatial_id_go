package spatialID

import (
	"testing"

	"github.com/trajectoryjp/geodesy_go/coordinates"
)

func TestGetExtendedSpatialIdsWithinRadiusOfLine02_1(t *testing.T) {
	convexHull := []*coordinates.Geodetic{
		{
			139.788452,
			35.67093015,
			0,
		},
		{
			139.788452,
			35.670840,
			0,
		},
	}
	clearance := 0.0
	quadkeyZoomLevel := int8(23)
	altitudekeyZoomLevel := int8(23)

	expectedCount := 4

	geodeticBox, theError := NewGeodeticBoxFromConvexHull(convexHull, clearance)
	if theError != nil {
		t.Error(theError)
	}

	tileXYZBox, theError := NewTileXYZBoxFromGeodeticBox(*geodeticBox, quadkeyZoomLevel, altitudekeyZoomLevel)
	if theError != nil {
		t.Error(theError)
	}

	count := 0
	for _ = range tileXYZBox.AllCollisionWithConvexHull(convexHull, clearance) {
		count += 1
	}

	if count != expectedCount {
		t.Errorf("Expected %v voxels, but got %v", expectedCount, count)
	}
}

func TestGetExtendedSpatialIdsWithinRadiusOfLine02_2(t *testing.T) {
	convexHull := []*coordinates.Geodetic{
		{
			139.788452,
			35.67093015,
			0,
		},
		{
			139.788452,
			35.670840,
			0,
		},
	}
	clearance := 0.1
	quadkeyZoomLevel := int8(23)
	altitudekeyZoomLevel := int8(23)

	expectedCount := 54

	geodeticBox, theError := NewGeodeticBoxFromConvexHull(convexHull, clearance)
	if theError != nil {
		t.Error(theError)
	}

	tileXYZBox, theError := NewTileXYZBoxFromGeodeticBox(*geodeticBox, quadkeyZoomLevel, altitudekeyZoomLevel)
	if theError != nil {
		t.Error(theError)
	}

	count := 0
	for _ = range tileXYZBox.AllCollisionWithConvexHull(convexHull, clearance) {
		count += 1
	}

	if count != expectedCount {
		t.Errorf("Expected %v voxels, but got %v", expectedCount, count)
	}
}

package spatialID

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-gl/mathgl/mgl64"
	closest "github.com/trajectoryjp/closest_go"
	"github.com/trajectoryjp/geodesy_go/coordinates"
	geodesy "github.com/trajectoryjp/geodesy_go/coordinates"
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

	geodeticBox, error := NewGeodeticBoxFromConvexHull(convexHull, clearance)
	if error != nil {
		t.Error(error)
	}

	tileXYZBox, error := NewTileXYZBoxFromGeodeticBox(*geodeticBox, quadkeyZoomLevel, altitudekeyZoomLevel)
	if error != nil {
		t.Error(error)
	}

	count := 0
	tileXYZBox.ForCollisionWithConvexHull(convexHull, clearance, func(tile *TileXYZ) {
		count += 1
	})

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

	geodeticBox, error := NewGeodeticBoxFromConvexHull(convexHull, clearance)
	if error != nil {
		t.Error(error)
	}

	tileXYZBox, error := NewTileXYZBoxFromGeodeticBox(*geodeticBox, quadkeyZoomLevel, altitudekeyZoomLevel)
	if error != nil {
		t.Error(error)
	}

	count := 0
	tileXYZBox.ForCollisionWithConvexHull(convexHull, clearance, func(tile *TileXYZ) {
		count += 1
	})

	if count != expectedCount {
		t.Errorf("Expected %v voxels, but got %v", expectedCount, count)
	}
}

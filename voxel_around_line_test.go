package spatialID

import (
	"fmt"
	"image/color"
	"os"
	"testing"

	"github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/twpayne/go-kml/v3"
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
		t.Fatal(theError)
	}

	tileXYZBox, theError := NewTileXYZBoxFromGeodeticBox(*geodeticBox, quadkeyZoomLevel, altitudekeyZoomLevel)
	if theError != nil {
		t.Fatal(theError)
	}

	count := 0
	for range tileXYZBox.AllCollisionWithConvexHull(convexHull, clearance) {
		count += 1
	}

	if count != expectedCount {
		t.Errorf("Expected %v voxels, but got %v", expectedCount, count)
	}
}

func TestGetExtendedSpatialIdsWithinRadiusOfLine02_2(t *testing.T) {
	if testing.Short() {
		t.Skip("Failed by https://github.com/trajectoryjp/closest_go/issues/2")
	}

	convexHull := []*coordinates.Geodetic{
		{
			139.788452,
			35.67093015,
			0.0,
		},
		{
			139.788452,
			35.670840,
			0.0,
		},
	}
	clearance := 0.1
	quadkeyZoomLevel := int8(23)
	altitudekeyZoomLevel := int8(12)

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
	document := kml.Document(
		kml.Name(t.Name()),
		kml.Placemark(
			kml.LineString(
				kml.AltitudeMode(kml.AltitudeModeAbsolute),
				kml.Coordinates(
					kml.Coordinate{
						Lon: *convexHull[0].Longitude(),
						Lat: *convexHull[0].Latitude(),
						Alt: *convexHull[0].Altitude(),
					},
					kml.Coordinate{
						Lon: *convexHull[1].Longitude(),
						Lat: *convexHull[1].Latitude(),
						Alt: *convexHull[1].Altitude(),
					},
				),
			),
		),
	)
	for tileXYZ := range tileXYZBox.AllCollisionWithConvexHull(convexHull, clearance) {
		tileXYZBox, theError := NewTileXYZBox(tileXYZ, tileXYZ)
		if theError != nil {
			t.Fatal(theError)
		}

		geodeticBox := NewGeodeticBoxFromTileXYZBox(*tileXYZBox)

		document.Append(
			kml.Placemark(
				kml.Name(fmt.Sprint(tileXYZ)),
				kml.Style(
					kml.PolyStyle(
						kml.Color(color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}),
					),
				),
				NewMultiGeometryFromGeodeticBox(*geodeticBox),
			),
		)
		count += 1
	}

	if count != expectedCount {
		t.Errorf("Expected %v voxels, but got %v", expectedCount, count)
	}

	kmlFile := kml.KML(
		document,
	)

	file, error := os.Create(t.Name() + ".kml")
	if error != nil {
		t.Fatal(error)
	}
	defer file.Close()

	error = kmlFile.WriteIndent(file, "", "  ")
	if error != nil {
		t.Fatal(error)
	}
}

package transform

import (
	"log"
	"testing"
	"time"

	"github.com/go-gl/mathgl/mgl64"
	closest "github.com/trajectoryjp/closest_go"
	geodesy "github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/spatial_id_go/common/object"
	"github.com/trajectoryjp/spatial_id_go/shape"
)

func TestGetSpatialIdsWithinRadiusOfLine(t *testing.T) {

	// sumida river, 200m
	startPoint, error := object.NewPoint(139.788452, 35.670935, 100)
	if error != nil {
		t.Fatal(error)
	}
	endPoint, error := object.NewPoint(139.788074, 35.675711, 100)
	if error != nil {
		t.Fatal(error)
	}

	// measure distance
	startCartesian := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{startPoint.Lon(), startPoint.Lat(), startPoint.Alt()})
	endCartesian := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{endPoint.Lon(), endPoint.Lat(), endPoint.Alt()})

	measure := closest.Measure{}
	measure.ConvexHulls[0] = []*mgl64.Vec3{(*mgl64.Vec3)(&startCartesian)}
	measure.ConvexHulls[1] = []*mgl64.Vec3{(*mgl64.Vec3)(&endCartesian)}

	measure.MeasureDistance()

	log.Printf("Distance: %vm", measure.Distance)

	start := time.Now()
	spatialIds, error := GetSpatialIdsWithinRadiusOfLine(startPoint, endPoint, 5, 22, 22, false)
	elapsed := time.Since(start)
	if error != nil {
		t.Fatal(error)
	}

	log.Printf("\n%v Spatial IDs found in %v \n", len(spatialIds), elapsed)

}

func TestFitClearanceAroundExtendedSpatialID(t *testing.T) {

	point, error := object.NewPoint(139.788081, 35.672680, 100)
	if error != nil {
		t.Fatal(error)
	}
	points := []*object.Point{point}
	// 25, 25 is approx 1m box
	spatialId, error := shape.GetExtendedSpatialIdsOnPoints(points, 25, 25)
	if error != nil {
		t.Fatal(error)
	}

	if len(spatialId) != 1 {
		t.Fatalf("N. Id's != 1")
	}

	start := time.Now()
	hLayer, vLayer, error := FitClearanceAroundExtendedSpatialID(spatialId[0], 1000)
	end := time.Since(start)
	if error != nil {
		t.Fatal(error)
	}
	log.Printf("\nCalculation time: %v\n", end)
	log.Printf("\nHorizontal Layers: %v\tVertical Layers: %v", hLayer, vLayer)

}

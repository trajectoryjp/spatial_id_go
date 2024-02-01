package transform

import (
	"testing"
	"time"

	"github.com/go-gl/mathgl/mgl64"
	closest "github.com/trajectoryjp/closest_go"
	geodesy "github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/spatial_id_go/common/errors"
	"github.com/trajectoryjp/spatial_id_go/common/object"
	"github.com/trajectoryjp/spatial_id_go/shape"
)

func TestGetSpatialIdsWithinRadiusOfLine(t *testing.T) {

	// sumida river, 200m
	startPoint, error := object.NewPoint(139.788452, 35.670935, 100)
	if error != nil {
		t.Error(error)
	}
	endPoint, error := object.NewPoint(139.788074, 35.761311, 100)
	if error != nil {
		t.Error(error)
	}

	// measure distance
	startCartesian := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{startPoint.Lon(), startPoint.Lat(), startPoint.Alt()})
	endCartesian := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{endPoint.Lon(), endPoint.Lat(), endPoint.Alt()})

	measure := closest.Measure{}
	measure.ConvexHulls[0] = []*mgl64.Vec3{(*mgl64.Vec3)(&startCartesian)}
	measure.ConvexHulls[1] = []*mgl64.Vec3{(*mgl64.Vec3)(&endCartesian)}

	measure.MeasureDistance()

	t.Logf("Distance: %vm", measure.Distance)

	start := time.Now()
	spatialIds, error := GetSpatialIdsWithinRadiusOfLine(startPoint, endPoint, 5, 23, 23, false)
	elapsed := time.Since(start)
	if error != nil {
		t.Error(error)
	}

	t.Logf("\n%v Spatial IDs found in %v \n", len(spatialIds), elapsed)

}

func TestFitClearanceAroundExtendedSpatialID01(t *testing.T) {

	var clearance float64 = 0
	var expectedHLayer int64 = 1
	var expectedVLayer int64 = 1

	point, error := object.NewPoint(139.788081, 35.672680, 100)
	if error != nil {
		t.Error(error)
	}
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(t, point, clearance, 25, 25)
	if error != nil {
		t.Error(error)
	}

	if hLayer != expectedHLayer {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectedHLayer, hLayer)
	}
	if vLayer != expectedVLayer {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectedVLayer, vLayer)
	}

}

func TestFitClearanceAroundExtendedSpatialID02(t *testing.T) {

	var clearance float64 = 10
	var expectedHLayer int64 = 11
	var expectedVLayer int64 = 11

	point, error := object.NewPoint(139.788081, 35.672680, 100)
	if error != nil {
		t.Error(error)
	}
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(t, point, clearance, 25, 25)
	if error != nil {
		t.Error(error)
	}

	if hLayer != expectedHLayer {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectedHLayer, hLayer)
	}
	if vLayer != expectedVLayer {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectedVLayer, vLayer)
	}

}

func TestFitClearanceAroundExtendedSpatialID03(t *testing.T) {

	var clearance float64 = 10
	var expectedHLayer int64 = 3
	var expectedVLayer int64 = 3

	point, error := object.NewPoint(139.788081, 35.672680, 100)
	if error != nil {
		t.Error(error)
	}
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(t, point, clearance, 23, 23)
	if error != nil {
		t.Error(error)
	}

	if hLayer != expectedHLayer {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectedHLayer, hLayer)
	}
	if vLayer != expectedVLayer {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectedVLayer, vLayer)
	}

}

func TestFitClearanceAroundExtendedSpatialID04(t *testing.T) {

	var clearance float64 = 50
	var expectedHLayer int64 = 4
	var expectedVLayer int64 = 4

	point, error := object.NewPoint(139.788081, 35.672680, 100)
	if error != nil {
		t.Error(error)
	}
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(t, point, clearance, 21, 21)
	if error != nil {
		t.Error(error)
	}

	if hLayer != expectedHLayer {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectedHLayer, hLayer)
	}
	if vLayer != expectedVLayer {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectedVLayer, vLayer)
	}

}

func testFitClearanceAroundSpatialID(t *testing.T, point *object.Point, clearance float64, hZoom int64, vZoom int64) (hLayer int64, vlayer int64, error error) {

	points := []*object.Point{point}
	// 25, 25 is just over 1m box
	spatialId, error := shape.GetExtendedSpatialIdsOnPoints(points, hZoom, vZoom)
	if error != nil {
		return 0, 0, error
	}
	if len(spatialId) != 1 {
		return 0, 0, errors.NewSpatialIdError(errors.OtherErrorCode, "returned more than 1 spatialID. Please check point again")
	}

	hLayer, vLayer, error := FitClearanceAroundExtendedSpatialID(spatialId[0], clearance)
	if error != nil {
		return 0, 0, error
	}

	return hLayer, vLayer, nil

}

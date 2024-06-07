package transform

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-gl/mathgl/mgl64"
	closest "github.com/trajectoryjp/closest_go"
	geodesy "github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v4/common/object"
	"github.com/trajectoryjp/spatial_id_go/v4/shape"
)

// TestGetExtendedSpatialIdsWithinRadiusOfLine01 tests when skipsMeasurement=true and radius =0
// Expected value should return the exact same voxels as GetExtendedSpatialIdsOnLine
func TestGetExtendedSpatialIdsWithinRadiusOfLine01(t *testing.T) {

	var radius float64 = 0
	var hZoom int64 = 25
	var vZoom int64 = 25

	startPoint, error := object.NewPoint(139.788452, 35.67093015, 0)
	if error != nil {
		t.Error(error)
	}
	endPoint, error := object.NewPoint(139.788452, 35.670840, 0)
	if error != nil {
		t.Error(error)
	}

	idsOnLine, error := shape.GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)
	if error != nil {
		t.Error(error)
	}
	idsWithinRadiusOfLine, error := GetExtendedSpatialIdsWithinRadiusOfLine(startPoint, endPoint, radius, hZoom, vZoom, true)
	if error != nil {
		t.Error(error)
	}

	map1, map2 := make(map[string]string), make(map[string]string)
	for _, value := range idsOnLine {
		map1[value] = value
	}
	for _, value := range idsWithinRadiusOfLine {
		map2[value] = value
	}
	if !reflect.DeepEqual(map1, map2) {
		t.Errorf("期待値: %v 取得値: %v", idsOnLine, idsWithinRadiusOfLine)
	}
	t.Log("テスト終了")

}

// TestGetSpatialIdsWithinRadiusOfLine02 tests when skipsMeasurement=true and radius > 0 but radius is
// less than than the length of a voxcel.
// Expected value should return 54 voxcels:
// - GetExtendedSpatialIdsOnLine: 4 ids
// - FitClearanceAroundSpatialId: hlayer=1, vlayer=1
// - GetNspatialIdsAroundVoxcels: (count excludes ids on line) 50
// - GetSpatialIdsWithinRadiusOfLine: (count includes ids on line) 9 ids per layer * 6 layers = 54 ids expected
func TestGetExtendedSpatialIdsWithinRadiusOfLine02(t *testing.T) {

	var radius float64 = 0.1
	var hZoom int64 = 23
	var vZoom int64 = 23

	startPoint, error := object.NewPoint(139.788452, 35.67093015, 0)
	if error != nil {
		t.Error(error)
	}
	endPoint, error := object.NewPoint(139.788452, 35.670840, 0)
	if error != nil {
		t.Error(error)
	}

	idsOnLine, error := shape.GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)
	if error != nil {
		t.Error(error)
	}
	idsWithinRadiusOfLine, error := GetExtendedSpatialIdsWithinRadiusOfLine(startPoint, endPoint, radius, hZoom, vZoom, true)
	if error != nil {
		t.Error(error)
	}

	if len(idsOnLine) != 4 || len(idsWithinRadiusOfLine) != 54 {
		t.Fatalf("Expected values not returned. \nExpected Number of Extended Spatial Ids on Line: 4; returned: %v\nExpected number of extended Spatial Ids within radius of line: 54; returned: %v", len(idsOnLine), len(idsWithinRadiusOfLine))
	}

}

func TestGetExtendedSpatialIdsWithinRadiusOfLine10m_r0_horizontal(t *testing.T) {

	var radius float64 = 0
	var hZoom int64 = 25
	var vZoom int64 = 25

	startPoint, error := object.NewPoint(139.788452, 35.67093015, 0)
	if error != nil {
		t.Error(error)
	}
	endPoint, error := object.NewPoint(139.788452, 35.670840, 0)
	if error != nil {
		t.Error(error)
	}

	startCartesian := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{startPoint.Lon(), startPoint.Lat(), startPoint.Alt()})
	endCartesian := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{endPoint.Lon(), endPoint.Lat(), endPoint.Alt()})

	measure := closest.Measure{}
	measure.ConvexHulls[0] = []*mgl64.Vec3{(*mgl64.Vec3)(&startCartesian)}
	measure.ConvexHulls[1] = []*mgl64.Vec3{(*mgl64.Vec3)(&endCartesian)}

	measure.MeasureDistance()

	t.Logf("Distance: %vm", measure.Distance)

	idsOnLine, error := shape.GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)
	if error != nil {
		t.Error(error)
	}
	start := time.Now()
	idsWithinRadiusOfLine, error := GetExtendedSpatialIdsWithinRadiusOfLine(startPoint, endPoint, radius, hZoom, vZoom, false)
	elapsed := time.Since(start)
	if error != nil {
		t.Error(error)
	}

	map1, map2 := make(map[string]string), make(map[string]string)
	for _, value := range idsOnLine {
		map1[value] = value
	}
	for _, value := range idsWithinRadiusOfLine {
		map2[value] = value
	}
	if !reflect.DeepEqual(map1, map2) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("nIDs: 期待値: %v 取得値: %v \n空間ID - 期待値:%v, \n\n取得値: %v", len(idsOnLine), len(idsWithinRadiusOfLine), map1, map2)
	}
	t.Logf("IdsWithinRadiusOfLine calculation time: %v", elapsed)
	t.Log("テスト終了")
}

func TestFitClearanceAroundExtendedSpatialID01(t *testing.T) {

	var clearance float64 = 0
	var expectedHLayer int64 = 0
	var expectedVLayer int64 = 0

	point, error := object.NewPoint(139.788081, 35.672680, 100)
	if error != nil {
		t.Error(error)
	}
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(point, clearance, 25, 25)
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
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(point, clearance, 25, 25)
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
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(point, clearance, 23, 23)
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
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(point, clearance, 21, 21)
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

func TestFitClearanceAroundExtendedSpatialID05(t *testing.T) {

	var clearance float64 = 0.01
	var expectedHLayer int64 = 1
	var expectedVLayer int64 = 1

	point, error := object.NewPoint(139.788081, 35.672680, 100)
	if error != nil {
		t.Error(error)
	}
	hLayer, vLayer, error := testFitClearanceAroundSpatialID(point, clearance, 21, 21)
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

func testFitClearanceAroundSpatialID(point *object.Point, clearance float64, hZoom int64, vZoom int64) (hLayer int64, vlayer int64, error error) {

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

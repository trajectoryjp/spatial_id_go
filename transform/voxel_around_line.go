package transform

import (
	"fmt"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/v2/common"
	"github.com/trajectoryjp/spatial_id_go/v2/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v2/common/enum"
	"github.com/trajectoryjp/spatial_id_go/v2/common/object"
	"github.com/trajectoryjp/spatial_id_go/v2/operated"
	"github.com/trajectoryjp/spatial_id_go/v2/shape"

	"github.com/go-gl/mathgl/mgl64"

	closest "github.com/trajectoryjp/closest_go"
	geodesy "github.com/trajectoryjp/geodesy_go/coordinates"
)

// GetSpatialIdsWithinRadiusOfLine
// 直線と線からradiusの距離以内の拡張空間IDを取得する
//
// 引数：
//
//	start: 始点
//	end: 終点
//	radius: 半径。拡張空間IDはradius以内だと、戻り値のスライスを追加される。
//  hZoom: 垂直方向の精度レベル
//  vZoom: 水平方向の精度レベル
//
// 戻り値：
//
//	拡張空間IDスライス： []string
//  error: エラー

func GetSpatialIdsWithinRadiusOfLine(startPoint *object.Point, endPoint *object.Point, radius float64, hZoom int64, vZoom int64, skipsMeasurement bool) ([]string, error) {

	// 1. Return the Extended Spatial Ids on the line determined by startPoint and endPoint

	idsOnLine, error := shape.GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)
	if error != nil {
		return nil, error
	}
	// 2. Find the Spatial Ids that are not on the line but within the radius distance of the line

	// Variable Setup

	// the closest.measure object to determine IDs not on the path line but within the criterion distance
	var measure1 = closest.Measure{}
	// the slice of ExtendedSpatialIDs around the Ids on the route line. (returned result already cleaned for duplicates).
	var idsAroundVoxcels []string
	// the Extended Spatial IDs with the IDs on the route path line removed
	var idsAroundLine []string // fix this
	// the slice of Extended Spatial Ids from noLinePathIdsAroundLine found within the radius but not on the route line
	var idsToAdd []string
	// points is the [2]slice / list of startPoint and endPoint in geodesic format (lat/lon)
	var points = []*object.Point{startPoint, endPoint}
	// cartesianPoints is the list of points converted into cartesian (x, y)
	var cartesianPoints []geodesy.Geocentric
	// the unique list of spatialIds that are either on the route path line or within the radius distance
	var idsWithinCriterion []string

	// Convert points to cartesian and append to cartesianPoints
	for _, point := range points {

		cartesianPoint := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{
			point.Lon(),
			point.Lat(),
			point.Lat(),
		})

		cartesianPoints = append(cartesianPoints, cartesianPoint)

	}

	// Store the Start and End points in the 0-th Convex Hull of Measure 1.
	// measure1.ConvexHull[0] will represent the line from which we measure distances to
	// each SpatialID
	measure1.ConvexHulls[0] = []*mgl64.Vec3{
		(*mgl64.Vec3)(&cartesianPoints[0]),
		(*mgl64.Vec3)(&cartesianPoints[1]),
	}

	// create megaboxIds

	// Determine the number of layers around the spatialID to search.
	// All SpatialIds are virtually the same size, so use the first to measure
	hLayers, vLayers, error := FitClearanceAroundExtendedSpatialID(idsOnLine[0], radius)
	if error != nil {
		return nil, error
	}

	// Return the SpatialIDs within the box created by hLayers and vLayers
	idsAroundVoxcels, error = operated.GetNspatialIdsAroundVoxcels(idsOnLine, hLayers, vLayers)
	if error != nil {
		return nil, error
	}

	// Remove the spatial ids in the line from the ids around the line so that only the ids around the line remain
	idsAroundLine = common.Difference(idsAroundVoxcels, idsOnLine)

	// if skipsMeasurement=true, measure the distance between the route line and each id in idsAroundLine
	if !skipsMeasurement {

		for _, id := range idsAroundLine {

			// Get 8 vertexes of the SpatialID
			IdVertexes, error := shape.GetPointOnExtendedSpatialId(id, enum.Vertex)
			if error != nil {
				return nil, error
			}

			// idConvex is the list of vectors that the SpatialID's 8 vertexes
			var idConvex = []*mgl64.Vec3{}

			// convert to cartesian and append to idCovex
			for _, vertex := range IdVertexes {

				cartesianPoint := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{
					vertex.Lon(),
					vertex.Lat(),
					vertex.Lat(),
				})

				idConvex = append(idConvex, (*mgl64.Vec3)(&cartesianPoint))

			}

			// Put idConvex into measure's ConvexHulls[1]
			measure1.ConvexHulls[1] = idConvex

			// Measure the distance between the line (ConvexHull[0]) and the
			// SpatialIDs vertex vectors (ConvexHull[1])
			measure1.MeasureNonnegativeDistance()

			// dist is the closest distance between the line (ConvexHull[0]) and the vertexes of Spatial ID[i]
			var dist = measure1.Distance

			// Since MeasureNonnegativeDistance() was used the distance value in dist
			// will always be non-zero. If dist < radius, add the spatialID to idsToAdd
			if dist < (radius) {
				idsToAdd = append(idsToAdd, id)
			}

		}

		// combine idsToAdd + idsOnLine
		idsWithinCriterion = common.Unique(common.Union(idsToAdd, idsOnLine))

	} else {
		// if skipsMeasurement=false, return all ids in noLinePathIdsAroundLine and combine with idsOnLine
		idsWithinCriterion = common.Unique(common.Union(idsAroundLine, idsOnLine))
	}

	return idsWithinCriterion, nil

}

// FitClearanceAroundExtendedSpatialID
// ユーザーが設定する拡張空間IDがクリアランスを保つには、何番目の垂直方向拡張空間IDと何番目の水平方向の拡張空間IDまで離れる必要があるか
//
// 引数：
//
//	spatialID: 拡張空間ID
//	clearance: クリアランス

// 戻り値：
//
//		horizontalLayer: 垂直方向の層目
//	 verticalLayer: 垂直方向の層目
//	 error: エラー
func FitClearanceAroundExtendedSpatialID(spatialID string, clearance float64) (horizontalLayer int64, verticalLayer int64, error error) {

	// validate clearance: There are two special cases:
	// 1. The clearance must be non-negative (clearance must be >= 0)
	// 2. If the clearance is exactly 0, return 0 for both the horizontalLayer and verticalLayer. This is because
	// if clearance is 0, the only Spatial IDs that are on the route line should be used -- no surrounding Spatial IDs
	if clearance < 0 {
		return 0, 0, fmt.Errorf("\ninvalid clearance value. Clearance must be >= 0")
	}

	// validate and extract zoom info from spatialID
	idElements := strings.Split(spatialID, consts.SpatialIDDelimiter)

	if len(idElements) != 5 {
		return 0, 0, fmt.Errorf("\ninvalid ExtendedSpatialID format. SpatialID must be 'hZoom/x/y/vZoom/z' format")
	}

	// hZoom, _ := strconv.ParseInt(idElements[0], 10, 64)
	// vZoom, _ := strconv.ParseInt(idElements[3], 10, 64)

	// hLayer is the number of horizonal spatialID distances required to fit the clearance
	var hLayer int64
	// vLayyer is the number of vertical spatialID distances required to fit the clearance
	var vLayer int64

	var hUnits int64 = 1
	var vUnits int64 = 1

	// Begin horizonal fitting loop (determine hLayer)
	for {

		// OrigianlConvex is the list of vectors that the original SpatialID's 8 vertexes
		var OriginalConvex = []*mgl64.Vec3{}
		// ShiftedConvex is the list of vectors that the shifted SpatialID's 8 vertexes
		var ShiftedConvex = []*mgl64.Vec3{}

		// shift spatialID over n units
		shiftedID := operated.GetShiftingSpatialID(spatialID, hUnits, 0, 0)

		// measure closest distance between original and shifted IDs
		measure := closest.Measure{}

		// Original Point: get verticies, loop through verticies, convert them to cartesian, append to OriginalConvex
		OriginalVertexes, error := shape.GetPointOnExtendedSpatialId(spatialID, enum.Vertex)
		if error != nil {
			return 0, 0, error
		}

		for _, vertex := range OriginalVertexes {

			cartesianPoint := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{
				vertex.Lon(),
				vertex.Lat(),
				vertex.Lat(),
			})

			OriginalConvex = append(OriginalConvex, (*mgl64.Vec3)(&cartesianPoint))

		}

		measure.ConvexHulls[0] = OriginalConvex

		// Shifted Point: get verticies, loop through verticies, convert them to cartesian, append to OriginalConvex
		ShiftedVertexes, error := shape.GetPointOnExtendedSpatialId(shiftedID, enum.Vertex)
		if error != nil {
			return 0, 0, error
		}

		for _, vertex := range ShiftedVertexes {

			cartesianPoint := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{
				vertex.Lon(),
				vertex.Lat(),
				vertex.Lat(),
			})

			ShiftedConvex = append(ShiftedConvex, (*mgl64.Vec3)(&cartesianPoint))

		}

		measure.ConvexHulls[1] = ShiftedConvex

		// Measure Distance between Original and Shifted Convex
		measure.MeasureNonnegativeDistance()

		// if clearance is greater than the distance, continue loop. Otherwise, return hUnits-1 to hLaye
		if clearance > measure.Distance {
			hUnits = hUnits + 1
			continue
		}

		// the last value to fit the clearance gets returned to hLayer
		hLayer = hUnits - 1
		break

		// end horizontal fitting loop
	}

	// Begin vertical fitting loop (determine vLayer)
	for {

		// OrigianlConvex is the list of vectors that the original SpatialID's 8 vertexes
		var OriginalConvex = []*mgl64.Vec3{}
		// ShiftedConvex is the list of vectors that the shifted SpatialID's 8 vertexes
		var ShiftedConvex = []*mgl64.Vec3{}

		// shift spatialID over n units
		shiftedID := operated.GetShiftingSpatialID(spatialID, 0, vUnits, 0)

		// measure closest distance between original and shifted IDs
		measure := closest.Measure{}

		// Original Point: get verticies, loop through verticies, convert them to cartesian, append to OriginalConvex
		OriginalVertexes, error := shape.GetPointOnExtendedSpatialId(spatialID, enum.Vertex)
		if error != nil {
			return 0, 0, error
		}

		for _, vertex := range OriginalVertexes {

			cartesianPoint := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{
				vertex.Lon(),
				vertex.Lat(),
				vertex.Lat(),
			})

			OriginalConvex = append(OriginalConvex, (*mgl64.Vec3)(&cartesianPoint))

		}

		measure.ConvexHulls[0] = OriginalConvex

		// Shifted Point: get verticies, loop through verticies, convert them to cartesian, append to OriginalConvex
		ShiftedVertexes, error := shape.GetPointOnExtendedSpatialId(shiftedID, enum.Vertex)
		if error != nil {
			return 0, 0, error
		}

		for _, vertex := range ShiftedVertexes {

			cartesianPoint := geodesy.GeocentricFromGeodetic(geodesy.Geodetic{
				vertex.Lon(),
				vertex.Lat(),
				vertex.Lat(),
			})

			ShiftedConvex = append(ShiftedConvex, (*mgl64.Vec3)(&cartesianPoint))

		}

		measure.ConvexHulls[1] = ShiftedConvex

		// Measure Distance between Original and Shifted Convex
		measure.MeasureNonnegativeDistance()

		// if clearance is greater than the distance, continue loop. Otherwise, return vUnits-1 to hLayer
		if clearance > measure.Distance {
			vUnits = vUnits + 1
			continue
		}

		// the last value to fit the clearance gets returned to hLayer
		vLayer = vUnits - 1
		break

		// end horizontal fitting loop
	}

	return hLayer, vLayer, nil
}

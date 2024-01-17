package transform

import (
	"github.com/trajectoryjp/spatial_id_go/common"
	"github.com/trajectoryjp/spatial_id_go/common/enum"
	"github.com/trajectoryjp/spatial_id_go/common/object"
	"github.com/trajectoryjp/spatial_id_go/operated"
	"github.com/trajectoryjp/spatial_id_go/shape"

	"github.com/go-gl/mathgl/mgl64"

	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/closest.git/v3"
)

// GetSpatialIdsWithinRadiusOfLine

//
//
// 引数：
//
//	start:
//	end:
//	radius
//  hZoom
//  vZoom
//
// 戻り値：
//
//	拡張空間IDスライス： []string

func GetSpatialIdsWithinRadiusOfLine(startPoint *object.Point, endPoint *object.Point, radius float64, hZoom int64, vZoom int64) ([]string, error) {

	// 1. Return the Extended Spatial Ids on the line determined by startPoint and endPoint

	idsOnLine, error := shape.GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)
	if error != nil {
		return nil, error
	}

	// 2. Find the Spatial Ids that are not on the line but within the radius distance of the line

	// the closest.measure object to determine IDs not on the path line but within the criterion distance
	var measure1 = closest.Measure{}
	// the distances from each SpatialID in uniqueMegaBoxIDs to the route path line
	var measureDistances1 []float64
	// the slice of Spatial IDs created by combining the result of all adjacent 26 SpatialIDs from all spatialIDs on the route line
	var megaBoxIds []string
	// the slice of unique Spatial IDs (non-duplicates) from megaBoxIDs
	var uniqueMegaBoxIds []string
	// the Spatial IDs in uniqueMegaBoxIds with the IDs on the route path line removed
	var noRoutePathMegaBoxIds []string
	// the slice of Spatial Ids from noRoutePathMegaBoxIds found within the radius but not on the route line
	var idsToAdd []string

	// put start and end point in convex
	points := []*object.Point{startPoint, endPoint}
	startEndVectors := make([]mgl64.Vec3, len(points))

	for i, v := range points {
		startEndVectors[i] = mgl64.Vec3(geodesy.Geodetic{
			v.Lon(),
			v.Lat(),
			float64(v.Alt()),
		}.ConvertToCartesian())
	}

	// put line from startPoint - endPoint in measure1.Convexes[0]
	measure1.Convexes[0] = []*mgl64.Vec3{
		&startEndVectors[0],
		&startEndVectors[1],
	}

	// loop through all spatialIds to add to measure1.convexes[1], measuredistance
	for _, id := range idsOnLine {

		// find the 26 surrounding spatialids ajacent
		adjacent26ids := operated.Get26spatialIdsAroundVoxel(id)

		// add to megaBoxIds
		megaBoxIds = append(megaBoxIds, adjacent26ids...)

	}

	// make unique list of spatial ids
	uniqueMegaBoxIds = common.Unique(megaBoxIds)

	// subtract the spatial ids in the line
	noRoutePathMegaBoxIds = common.Difference(uniqueMegaBoxIds, idsOnLine)

	// loop through the spatialids not included in the route path to determine their
	// distance from the path line.
	for _, id := range noRoutePathMegaBoxIds {

		// get 8 vertexes of spatialid
		IdVertexes, error := shape.GetPointOnExtendedSpatialId(id, enum.Vertex)
		if error != nil {
			return nil, error
		}

		idVectors := make([]mgl64.Vec3, len(IdVertexes))
		idConvex := []*mgl64.Vec3{}

		for i, v := range IdVertexes {
			idVectors[i] = mgl64.Vec3(geodesy.Geodetic{
				v.Lon(),
				v.Lat(),
				float64(v.Alt()),
			}.ConvertToCartesian())

			idConvex = append(idConvex, (*mgl64.Vec3)(&idVectors[i]))
		}

		// put these vertexes into convex
		measure1.Convexes[1] = idConvex

		// measure distance and append to list of distances
		measure1.MeasurePlusDistance()

		var dist = measure1.Distance

		measureDistances1 = append(measureDistances1, dist)

		// since spatialIDs on the flight path line have been removed (noRoutePathMegaBoxIds)
		// the distance should always been > 0. If dist < radius, add the spatialID to idsToAdd
		if dist < (radius) {
			idsToAdd = append(idsToAdd, id)
		}

	} // end noRoutePathMegaBoxIds/distance to line loop

	// combine idsToAdd + idsOnLine
	// the unique list of spatialIds that are either on the route path line or within the radius distance
	var idsWithinCriterion []string
	idsWithinCriterion = common.Unique(common.Union(idsToAdd, idsOnLine))

	return idsWithinCriterion, nil

}

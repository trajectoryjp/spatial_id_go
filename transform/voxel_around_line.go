package transform

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/trajectoryjp/spatial_id_go/common"
	"github.com/trajectoryjp/spatial_id_go/common/consts"
	"github.com/trajectoryjp/spatial_id_go/common/enum"
	"github.com/trajectoryjp/spatial_id_go/common/object"
	"github.com/trajectoryjp/spatial_id_go/operated"
	"github.com/trajectoryjp/spatial_id_go/shape"

	"github.com/go-gl/mathgl/mgl64"

	closest "github.com/trajectoryjp/closest_go"
	geodesy "github.com/trajectoryjp/geodesy_go/coordinates"
)

// type Geocentric mgl64.Vec3
// type Geodetic mgl64.Vec3

// var GeocentricReferenceSystem = wgs84.GeocentricReferenceSystem{}
// var GeodeticReferenceSystem = wgs84.LonLat()

// func GeocentricFromGeodetic(geodetic Geodetic) (geocentric Geocentric) {
// 	geocentric[0], geocentric[1], geocentric[2] = GeodeticReferenceSystem.To(GeocentricReferenceSystem)(geodetic[0], geodetic[1], geodetic[2])
// 	return
// }

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
	// the distances from each SpatialID in uniqueMegaBoxIDs to the route path line
	var measureDistances1 []float64
	// the slice of Spatial IDs created by combining the result of all adjacent 26 SpatialIDs from all spatialIDs on the route line
	var megaBoxIds []string
	// the slice of unique Spatial IDs (non-duplicates) from megaBoxIDs
	var uniqueMegaBoxIds []string
	// the Spatial IDs in uniqueMegaBoxIds with the IDs on the route path line removed
	var noLinePathMegaBoxIds []string
	// the slice of Spatial Ids from noLinePathMegaBoxIds found within the radius but not on the route line
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

	startGetids := time.Now()
	// Loop through all idsOnLine to create megaBoxIds
	for _, id := range idsOnLine {

		// Determine the number of layers around the spatialID to search.
		hLayers, vLayers, error := FitClearanceAroundExtendedSpatialID(id, radius)
		if error != nil {
			return nil, error
		}

		// Return the SpatialIDs within the box created by hLayers and vLayers
		adjacentIds, error := operated.GetNspatialIdsAroundVoxcel(id, hLayers, vLayers)

		// add to megaBoxIds
		megaBoxIds = append(megaBoxIds, adjacentIds...)

	}
	endGetids := time.Since(startGetids)

	// Make unique list of spatial ids
	startUnique := time.Now()
	uniqueMegaBoxIds = common.Unique(megaBoxIds)

	// Subtract the spatial ids in the line
	noLinePathMegaBoxIds = common.Difference(uniqueMegaBoxIds, idsOnLine)
	endUnique := time.Since(startUnique)
	// loop through the spatialids not included in the line path (noLinePathMegaBoxIds)
	// to determine their distance from the path line.
	startMeasure := time.Now()
	if !skipsMeasurement {

		for _, id := range noLinePathMegaBoxIds {

			// Get 8 vertexes of the SpatialID
			IdVertexes, error := shape.GetPointOnExtendedSpatialId(id, enum.Vertex)
			if error != nil {
				return nil, error
			}

			//idVectors := make([]mgl64.Vec3, len(IdVertexes))

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

				// idVectors[i] = mgl64.Vec3(geodesy.Geodetic{
				// 	v.Lon(),
				// 	v.Lat(),
				// 	float64(v.Alt()),
				// }.ConvertToCartesian())

				// idConvex = append(idConvex, (*mgl64.Vec3)(&idVectors[i]))
			}

			// Put idConvex into measure's ConvexHulls[1]
			measure1.ConvexHulls[1] = idConvex

			// Measure the distance between the line (ConvexHull[0]) and the
			// SpatialIDs vertex vectors (ConvexHull[1])
			measure1.MeasureNonnegativeDistance()

			// dist is the closest distance between the line (ConvexHull[0]) and the vertexes of Spatial ID[i]
			var dist = measure1.Distance

			measureDistances1 = append(measureDistances1, dist)

			// Since MeasureNonnegativeDistance() was used the distance value in dist
			// will always be non-zero. If dist < radius, add the spatialID to idsToAdd
			if dist < (radius) {
				idsToAdd = append(idsToAdd, id)
			}

		} // end noLinePathMegaBoxIds/ distance to line loop

		// combine idsToAdd + idsOnLine
		idsWithinCriterion = common.Unique(common.Union(idsToAdd, idsOnLine))

	} else {
		idsWithinCriterion = megaBoxIds
	}
	endMeasure := time.Since(startMeasure)

	log.Printf("\nGetNids: %v\nUnique/Subtract: %v\nMeasure: %v\n", endGetids, endUnique, endMeasure)

	Nadd := float64(len(idsToAdd))
	NTotal := float64(len(uniqueMegaBoxIds))
	fraction := Nadd / NTotal
	pct := mgl64.Round(fraction, 5) * 100
	log.Printf("\nIds from GetIdsWithinRadiusOfLine: %v\nIds measured and added from above: %v\nPercent within radius: %v pct",
		len(uniqueMegaBoxIds),
		len(idsToAdd),
		pct)

	return idsWithinCriterion, nil

}

func GetSpatialIdsWithinRadiusOfLine_vector(startPoint *object.Point, endPoint *object.Point, radius float64, hZoom int64, vZoom int64, skipsMeasurement bool) ([]string, error) {

	// 1. Return the Extended Spatial Ids on the line determined by startPoint and endPoint

	idsOnLine, error := shape.GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)
	if error != nil {
		return nil, error
	}

	// 2. Find the Spatial Ids that are not on the line but within the radius distance of the line

	// Variable Setup

	// the closest.measure object to determine IDs not on the path line but within the criterion distance
	var measure1 = closest.Measure{}
	// the distances from each SpatialID in uniqueMegaBoxIDs to the route path line
	var measureDistances1 []float64
	// the slice of Spatial IDs created by combining the result of all adjacent 26 SpatialIDs from all spatialIDs on the route line
	var megaBoxIds []string
	// the slice of unique Spatial IDs (non-duplicates) from megaBoxIDs
	var uniqueMegaBoxIds []string
	// the Spatial IDs in uniqueMegaBoxIds with the IDs on the route path line removed
	var noLinePathMegaBoxIds []string
	// the slice of Spatial Ids from noLinePathMegaBoxIds found within the radius but not on the route line
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

	startGetids := time.Now()

	// create megaboxIds

	// Determine the number of layers around the spatialID to search.
	// All SpatialIds are virtually the same size, so use the first to measure
	hLayers, vLayers, error := FitClearanceAroundExtendedSpatialID(idsOnLine[0], radius)
	if error != nil {
		return nil, error
	}

	// Return the SpatialIDs within the box created by hLayers and vLayers
	megaBoxIds, error = operated.GetNspatialIdsAroundVoxcels(idsOnLine, hLayers, vLayers)
	if error != nil {
		return nil, error
	}

	endGetids := time.Since(startGetids)

	// Make unique list of spatial ids
	startUnique := time.Now()
	uniqueMegaBoxIds = common.Unique(megaBoxIds)

	// Subtract the spatial ids in the line
	noLinePathMegaBoxIds = common.Difference(uniqueMegaBoxIds, idsOnLine)
	endUnique := time.Since(startUnique)
	// loop through the spatialids not included in the line path (noLinePathMegaBoxIds)
	// to determine their distance from the path line.
	startMeasure := time.Now()
	if !skipsMeasurement {

		for _, id := range noLinePathMegaBoxIds {

			// Get 8 vertexes of the SpatialID
			IdVertexes, error := shape.GetPointOnExtendedSpatialId(id, enum.Vertex)
			if error != nil {
				return nil, error
			}

			//idVectors := make([]mgl64.Vec3, len(IdVertexes))

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

				// idVectors[i] = mgl64.Vec3(geodesy.Geodetic{
				// 	v.Lon(),
				// 	v.Lat(),
				// 	float64(v.Alt()),
				// }.ConvertToCartesian())

				// idConvex = append(idConvex, (*mgl64.Vec3)(&idVectors[i]))
			}

			// Put idConvex into measure's ConvexHulls[1]
			measure1.ConvexHulls[1] = idConvex

			// Measure the distance between the line (ConvexHull[0]) and the
			// SpatialIDs vertex vectors (ConvexHull[1])
			measure1.MeasureNonnegativeDistance()

			// dist is the closest distance between the line (ConvexHull[0]) and the vertexes of Spatial ID[i]
			var dist = measure1.Distance

			measureDistances1 = append(measureDistances1, dist)

			// Since MeasureNonnegativeDistance() was used the distance value in dist
			// will always be non-zero. If dist < radius, add the spatialID to idsToAdd
			if dist < (radius) {
				idsToAdd = append(idsToAdd, id)
			}

		} // end noLinePathMegaBoxIds/ distance to line loop

		// combine idsToAdd + idsOnLine
		idsWithinCriterion = common.Unique(common.Union(idsToAdd, idsOnLine))

	} else {
		idsWithinCriterion = megaBoxIds
	}
	endMeasure := time.Since(startMeasure)

	log.Printf("\nGetNids: %v\nUnique/Subtract: %v\nMeasure: %v\n", endGetids, endUnique, endMeasure)

	Nadd := float64(len(idsToAdd))
	NTotal := float64(len(uniqueMegaBoxIds))
	fraction := Nadd / NTotal
	pct := mgl64.Round(fraction, 5) * 100
	log.Printf("\nIds from GetIdsWithinRadiusOfLine: %v\nIds measured and added from above: %v\nPercent within radius: %v pct",
		len(uniqueMegaBoxIds),
		len(idsToAdd),
		pct)

	return idsWithinCriterion, nil

}

func FitClearanceAroundExtendedSpatialID(spatialID string, clearance float64) (horizontalLayer int64, verticalLayer int64, error error) {

	// validate clearance
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

	var hUnits int64 = 2
	var vUnits int64 = 2

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

package spatialID

import (
	"math"

	mathematics "github.com/HarutakaMatsumoto/mathematics_go"
	"github.com/HarutakaMatsumoto/mathematics_go/geometry/rectangular/solid"
	"github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
)

type GeodeticBox struct {
	Min coordinates.Geodetic
	Max coordinates.Geodetic
}

func NewGeodeticBoxFromConvexHull(convexHull []*coordinates.Geodetic, clearance float64) (*GeodeticBox, error) {
	if len(convexHull) == 0 {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "The convex hull is empty.")
	}

	geodeticMin := *convexHull[0]
	geodeticMax := geodeticMin
	for _, vertex := range convexHull {
		if *vertex.Longitude() < *geodeticMin.Longitude() {
			*geodeticMin.Longitude() = *vertex.Longitude()
		} else if *vertex.Longitude() > *geodeticMax.Longitude() {
			*geodeticMax.Longitude() = *vertex.Longitude()
		}

		if *vertex.Latitude() < *geodeticMin.Latitude() {
			*geodeticMin.Latitude() = *vertex.Latitude()
		} else if *vertex.Latitude() > *geodeticMax.Latitude() {
			*geodeticMax.Latitude() = *vertex.Latitude()
		}

		if *vertex.Altitude() < *geodeticMin.Altitude() {
			*geodeticMin.Altitude() = *vertex.Altitude()
		} else if *vertex.Altitude() > *geodeticMax.Altitude() {
			*geodeticMax.Altitude() = *vertex.Altitude()
		}
	}

	localFromGeocentric := geodeticMin.GenerateLocalFromGeocentric()
	geocentricFromLocal := geodeticMin.GenerateGeocentricFromLocal()

	geocentricMin := coordinates.GeocentricFromGeodetic(geodeticMin)
	geocentricMax := coordinates.GeocentricFromGeodetic(geodeticMax)
	localMin := localFromGeocentric(geocentricMin)
	localMax := localFromGeocentric(geocentricMax)
	
	localMin[0] -= clearance
	localMin[1] -= clearance
	localMin[2] -= clearance
	localMax[0] += clearance
	localMax[1] += clearance
	localMax[2] += clearance

	geocentricMin = geocentricFromLocal(localMin)
	geocentricMax = geocentricFromLocal(localMax)

	return &GeodeticBox{
		Min: coordinates.GeodeticFromGeocentric(geocentricMin),
		Max: coordinates.GeodeticFromGeocentric(geocentricMax),
	}, nil
}

func NewGeodeticBoxFromSpatialIDBox(spatialIDBox SpatialIDBox) *GeodeticBox {
	box := &GeodeticBox{}

	max := float64(int(1) << spatialIDBox.GetMin().GetZ())

	*box.Min.Longitude() = 360.0 * float64(spatialIDBox.GetMin().GetX()) / max - 180.0
	*box.Max.Longitude() = 360.0 * float64(spatialIDBox.GetMax().GetX()+1) / max - 180.0

	*box.Min.Latitude() = mathematics.RadianPerDegree * math.Atan(math.Sinh(math.Pi * (1.0 - 2.0*float64(spatialIDBox.GetMin().GetY())/max)))
	*box.Max.Latitude() = mathematics.RadianPerDegree * math.Atan(math.Sinh(math.Pi * (1.0 - 2.0*float64(spatialIDBox.GetMax().GetY()+1)/max)))

	altitudeResolution := float64(int(1) << (SpatialIDZBaseExponent - spatialIDBox.GetMin().GetZ()))

	*box.Min.Altitude() = float64(spatialIDBox.GetMin().GetZ()) * altitudeResolution
	*box.Max.Altitude() = float64(spatialIDBox.GetMax().GetZ()+1) * altitudeResolution

	return box
}

func NewGeodeticBoxFromTileXYZBox(TileXYZBox TileXYZBox) *GeodeticBox {
	box := &GeodeticBox{}

	quadMax := float64(int(1) << TileXYZBox.GetMin().GetQuadkeyZoomLevel())

	*box.Min.Longitude() = 360.0 * float64(TileXYZBox.GetMin().GetX()) / quadMax - 180.0
	*box.Max.Longitude() = 360.0 * float64(TileXYZBox.GetMax().GetX()+1) / quadMax - 180.0

	*box.Min.Latitude() = mathematics.RadianPerDegree * math.Atan(math.Sinh(math.Pi * (1.0 - 2.0*float64(TileXYZBox.GetMin().GetY())/quadMax)))
	*box.Max.Latitude() = mathematics.RadianPerDegree * math.Atan(math.Sinh(math.Pi * (1.0 - 2.0*float64(TileXYZBox.GetMax().GetY()+1)/quadMax)))

	altitudeResolution := float64(int(1) << (TileXYZZBaseExponent - TileXYZBox.GetMin().GetAltitudekeyZoomLevel()))

	*box.Min.Altitude() = float64(TileXYZBox.GetMin().GetAltitudekeyZoomLevel()) * altitudeResolution
	*box.Max.Altitude() = float64(TileXYZBox.GetMax().GetAltitudekeyZoomLevel()+1) * altitudeResolution

	return box
}

func (box GeodeticBox) GetVertices() []*coordinates.Geodetic {
	vertices := make([]*coordinates.Geodetic, len(solid.VertexIntervals))
	for i, interval := range solid.VertexIntervals {
		vertices[i] = &coordinates.Geodetic{}

		if interval[0] == 1.0 {
			*vertices[i].Longitude() = *box.Max.Longitude()
		} else {
			*vertices[i].Longitude() = *box.Min.Longitude()
		}

		if interval[1] == 1.0 {
			*vertices[i].Latitude() = *box.Max.Latitude()
		} else {
			*vertices[i].Latitude() = *box.Min.Latitude()
		}

		if interval[2] == 1.0 {
			*vertices[i].Altitude() = *box.Max.Altitude()
		} else {
			*vertices[i].Altitude() = *box.Min.Altitude()
		}
	}

	return vertices
}

func (box GeodeticBox) GetCenter() coordinates.Geodetic {
	center := coordinates.Geodetic{}

	*center.Longitude() = (*box.Min.Longitude() + *box.Max.Longitude()) / 2.0
	*center.Latitude() = (*box.Min.Latitude() + *box.Max.Latitude()) / 2.0
	*center.Altitude() = (*box.Min.Altitude() + *box.Max.Altitude()) / 2.0

	return center
}

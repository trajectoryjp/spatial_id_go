package spatialID

import (
	"math"

	mathematics "github.com/HarutakaMatsumoto/mathematics_go"
	"github.com/HarutakaMatsumoto/mathematics_go/geometry/rectangular/solid"
	"github.com/trajectoryjp/geodesy_go/coordinates"
)

type GeodeticBox struct {
	Min coordinates.Geodetic
	Max coordinates.Geodetic
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

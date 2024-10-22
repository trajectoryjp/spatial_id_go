package spatialID

import (
	"math"

	"github.com/HarutakaMatsumoto/mathematics_go/geometry/rectangular/solid"
	"github.com/go-gl/mathgl/mgl64"
	closest "github.com/trajectoryjp/closest_go"
	"github.com/trajectoryjp/geodesy_go/coordinates"
)

type SpatialIDBox struct {
	min SpatialID
	max SpatialID
}

func NewSpatialIDBox(min SpatialID, max SpatialID) (*SpatialIDBox, error) {
	delta := max.GetZ() - min.GetZ()
	if delta < 0 {
		newMax, error := max.NewMaxChild(-delta)
		if error != nil {
			return nil, error
		}

		max = *newMax
	} else if delta > 0 {
		newMin, error := min.NewMinChild(delta)
		if error != nil {
			return nil, error
		}

		min = *newMin
	}

	return &SpatialIDBox{
		min: min,
		max: max,
	}, nil
}

func (box *SpatialIDBox) AddZ(delta int8) error {
	if delta < 0 {
		newMin, error := box.min.NewParent(-delta)
		if error != nil {
			return error
		}

		box.min = *newMin

		newMax, error := box.max.NewParent(-delta)
		if error != nil {
			return error
		}

		box.max = *newMax
	} else if delta > 0 {
		newMin, error := box.min.NewMinChild(delta)
		if error != nil {
			return error
		}

		box.min = *newMin

		newMax, error := box.max.NewMaxChild(delta)
		if error != nil {
			return error
		}

		box.max = *newMax
	}

	return nil
}

func (box SpatialIDBox) GetMin() SpatialID {
	return box.min
}

func (box SpatialIDBox) GetMax() SpatialID {
	return box.max
}

func ForSpatialIDsCollidedWithConvexHull(zoomLevel int8, convexHull []*coordinates.Geodetic, clearance float64, function func(id *SpatialID)) error {
	if len(convexHull) == 0 {
		return nil
	}

	geocentricMin := coordinates.GeocentricFromGeodetic(*convexHull[0])
	geocentricMax := geocentricMin
	geocentricConvexHull := make([]*mgl64.Vec3, len(convexHull))
	for _, vertex := range convexHull {
		geocentric := coordinates.GeocentricFromGeodetic(*vertex)
		geocentricConvexHull = append(geocentricConvexHull, (*mgl64.Vec3)(&geocentric))

		if *geocentric.X() < *geocentricMin.X() {
			*geocentricMin.X() = *geocentric.X()
		} else if *geocentric.X() > *geocentricMax.X() {
			*geocentricMax.X() = *geocentric.X()
		}

		if *geocentric.Y() < *geocentricMin.Y() {
			*geocentricMin.Y() = *geocentric.Y()
		} else if *geocentric.Y() > *geocentricMax.Y() {
			*geocentricMax.Y() = *geocentric.Y()
		}

		if *geocentric.Z() < *geocentricMin.Z() {
			*geocentricMin.Z() = *geocentric.Z()
		} else if *geocentric.Z() > *geocentricMax.Z() {
			*geocentricMax.Z() = *geocentric.Z()
		}
	}

	*geocentricMin.X() -= clearance
	*geocentricMin.Y() -= clearance
	*geocentricMin.Z() -= clearance
	*geocentricMax.X() += clearance
	*geocentricMax.Y() += clearance
	*geocentricMax.Z() += clearance

	minSpatialID, error := NewSpatialIDFromGeodetic(coordinates.GeodeticFromGeocentric(geocentricMin), zoomLevel)
	if error != nil {
		return error
	}
	maxSpatialID, error := NewSpatialIDFromGeodetic(coordinates.GeodeticFromGeocentric(geocentricMax), zoomLevel)
	if error != nil {
		return error
	}
	spatialIDBox, error := NewSpatialIDBox(*minSpatialID, *maxSpatialID)
	if error != nil {
		return error
	}

	measure := closest.Measure{
		ConvexHulls: [2][]*mgl64.Vec3{
			geocentricConvexHull,
			make([]*mgl64.Vec3, len(solid.VertexIntervals)),
		},
	}

	oldDistance := math.Inf(1)
	spatialIDBox.ForXYF(func(id *SpatialID) {
		spatialIDBox, _ := NewSpatialIDBox(*id, *id)
		geodeticBox := NewGeodeticBoxFromSpatialIDBox(*spatialIDBox)

		for i, vertex := range geodeticBox.GetVertices() {
			geocentric := coordinates.GeocentricFromGeodetic(*vertex)
			measure.ConvexHulls[1][i] = (*mgl64.Vec3)(&geocentric)
		}

		measure.MeasureNonnegativeDistance()

		if measure.Distance > clearance {
			if measure.Distance > oldDistance {
				id.SetF(spatialIDBox.GetMax().GetF())
				oldDistance = math.MaxFloat64
			} else {
				deltaAltitude := *geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude()
				newF := int64(measure.Distance / deltaAltitude) + id.GetF()
				if newF >= spatialIDBox.GetMax().GetF() {
					id.SetF(spatialIDBox.GetMax().GetF())
				} else {
					id.SetF(newF)
				}
			}
		}

		function(id)

		if id.GetF() == spatialIDBox.GetMax().GetF() {
			oldDistance = math.Inf(1)
		} else {
			oldDistance = measure.Distance
		}
	})

	return nil
}

func (box SpatialIDBox) ForXYF(function func(id *SpatialID)) {
	current := SpatialID{}

	for current.SetX(box.GetMin().GetX()); ; current.SetX(current.GetX() + 1) {
		for current.SetY(box.GetMin().GetY()); ; current.SetY(current.GetY() + 1) {
			for current.SetF(box.GetMin().GetF()); ; current.SetF(current.GetF() + 1) {
				function(&current)

				if current.GetF() == box.GetMax().GetF() {
					break
				}
			}

			if current.GetY() == box.GetMax().GetY() {
				break
			}
		}

		if current.GetX() == box.GetMax().GetX() {
			break
		}
	}
}

func NewSpatialIDBoxFromTileXYZBox(tileXYZBox TileXYZBox) (*SpatialIDBox, error) {
	deltaQuad := tileXYZBox.GetMin().GetQuadkeyZoomLevel() - SpatialIDZBaseExponent
	deltaAltitude := tileXYZBox.GetMin().GetAltitudekeyZoomLevel() - TileXYZZBaseExponent
	tileXYZBox.AddZoomLevel(-deltaQuad, -deltaAltitude)

	baseMinID, error := NewSpatialID(
		tileXYZBox.GetMin().GetQuadkeyZoomLevel(),
		tileXYZBox.GetMin().GetZ()-TileXYZZBaseOffset+SpatialIDZBaseOffset,
		tileXYZBox.GetMin().GetX(),
		tileXYZBox.GetMin().GetY(),
	)
	if error != nil {
		return nil, error
	}

	baseMaxID, error := NewSpatialID(
		tileXYZBox.GetMax().GetQuadkeyZoomLevel(),
		tileXYZBox.GetMax().GetZ()-TileXYZZBaseOffset+SpatialIDZBaseOffset,
		tileXYZBox.GetMax().GetX(),
		tileXYZBox.GetMax().GetY(),
	)
	if error != nil {
		return nil, error
	}

	box, error := NewSpatialIDBox(*baseMinID, *baseMaxID)
	if error != nil {
		return nil, error
	}

	delta := deltaQuad
	if deltaAltitude < delta {
		delta = deltaAltitude
	}

	error = box.AddZ(delta)
	if error != nil {
		return nil, error
	}

	return box, nil
}

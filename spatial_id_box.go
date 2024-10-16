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

func (box SpatialIDBox) ForCollisionWithConvexHull(convexHull []*coordinates.Geodetic, clearance float64, function func(SpatialID) SpatialID) {
	measure := closest.Measure{
		ConvexHulls: [2][]*mgl64.Vec3{
			make([]*mgl64.Vec3, len(convexHull)),
			make([]*mgl64.Vec3, len(solid.VertexIntervals)),
		},
	}

	for i, vertex := range convexHull {
		measure.ConvexHulls[0][i] = (*mgl64.Vec3)(vertex)
	}

	oldDistance := math.Inf(1)
	box.ForXYF(func(id SpatialID) SpatialID {
		spatialIDBox, _ := NewSpatialIDBox(id, id)
		geodeticBox := NewGeodeticBoxFromSpatialIDBox(*spatialIDBox)

		for i, vertex := range geodeticBox.GetVertices() {
			measure.ConvexHulls[1][i] = (*mgl64.Vec3)(vertex)
		}

		measure.MeasureNonnegativeDistance()

		geocentric0 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[0]))
		geocentric1 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[1]))
		distance := mgl64.Vec3(geocentric0).Sub(mgl64.Vec3(geocentric1)).Len() // TODO: Embed

		if distance > clearance {
			if distance > oldDistance {
				id.SetF(box.GetMax().GetF() - 1) // TODO: Consider
				oldDistance = math.MaxFloat64
			}

			return id
		}

		id = function(id)

		if id.GetF() == box.GetMax().GetF() {
			oldDistance = math.Inf(1)
		} else {
			oldDistance = distance
		}

		return id
	})
}

func (box SpatialIDBox) ForXYF(function func(SpatialID) SpatialID) {
	current := SpatialID{}

	for current.SetX(box.GetMin().GetX()); current.GetX() != box.GetMax().GetX(); current.SetX(current.GetX() + 1) {
		for current.SetY(box.GetMin().GetY()); current.GetY() != box.GetMax().GetY(); current.SetY(current.GetY() + 1) {
			for current.SetF(box.GetMin().GetF()); current.GetF() != box.GetMax().GetF(); current.SetF(current.GetF() + 1) {
				current = function(current)
			}

			current = function(current)
		}

		for current.SetF(box.GetMin().GetF()); current.GetF() != box.GetMax().GetF(); current.SetF(current.GetF() + 1) {
			current = function(current)
		}

		current = function(current)
	}

	for current.SetY(box.GetMin().GetY()); current.GetY() != box.GetMax().GetY(); current.SetY(current.GetY() + 1) {
		for current.SetF(box.GetMin().GetF()); current.GetF() != box.GetMax().GetF(); current.SetF(current.GetF() + 1) {
			current = function(current)
		}

		current = function(current)
	}

	for current.SetF(box.GetMin().GetF()); current.GetF() != box.GetMax().GetF(); current.SetF(current.GetF() + 1) {
		current = function(current)
	}

	current = function(current)
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

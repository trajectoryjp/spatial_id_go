package spatialID

import (
	"math"

	"github.com/HarutakaMatsumoto/mathematics_go/geometry/rectangular/solid"
	"github.com/go-gl/mathgl/mgl64"
	closest "github.com/trajectoryjp/closest_go"
	"github.com/trajectoryjp/geodesy_go/coordinates"
)

type TileXYZBox struct {
	min TileXYZ
	max TileXYZ
}

func NewTileXYZBox(min TileXYZ, max TileXYZ) (*TileXYZBox, error) {
	quadDelta := max.GetQuadkeyZoomLevel() - min.GetQuadkeyZoomLevel()
	altitudeDelta := max.GetAltitudekeyZoomLevel() - min.GetAltitudekeyZoomLevel()
	if quadDelta < 0 || altitudeDelta < 0 {
		newMax, error := max.NewMaxChild(-quadDelta, -altitudeDelta)
		if error != nil {
			return nil, error
		}

		max = *newMax
	} else if quadDelta > 0 || altitudeDelta > 0 {
		newMin, error := min.NewMinChild(quadDelta, altitudeDelta)
		if error != nil {
			return nil, error
		}

		min = *newMin
	}

	return &TileXYZBox{
		min: min,
		max: max,
	}, nil
}

func (box *TileXYZBox) AddZoomLevel(quadDelta, altitudeDelta int8) error {
	if quadDelta < 0 || altitudeDelta < 0 {
		newMin, error := box.min.NewParent(-quadDelta, -altitudeDelta)
		if error != nil {
			return error
		}

		box.min = *newMin

		newMax, error := box.max.NewParent(-quadDelta, -altitudeDelta)
		if error != nil {
			return error
		}

		box.max = *newMax
	} else if quadDelta > 0 || altitudeDelta > 0 {
		newMin, error := box.min.NewMinChild(quadDelta, altitudeDelta)
		if error != nil {
			return error
		}

		box.min = *newMin

		newMax, error := box.max.NewMaxChild(quadDelta, altitudeDelta)
		if error != nil {
			return error
		}

		box.max = *newMax
	}

	return nil
}

func (box TileXYZBox) GetMin() TileXYZ {
	return box.min
}

func (box TileXYZBox) GetMax() TileXYZ {
	return box.max
}

func (box TileXYZBox) ForCollisionWithConvexHull(convexHull []*coordinates.Geodetic, clearance float64, function func(TileXYZ) TileXYZ) {
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
	box.ForXYZ(func(tile TileXYZ) TileXYZ {
		tileXYZBox, _ := NewTileXYZBox(tile, tile)
		geodeticBox := NewGeodeticBoxFromTileXYZBox(*tileXYZBox)

		for i, vertex := range geodeticBox.GetVertices() {
			measure.ConvexHulls[1][i] = (*mgl64.Vec3)(vertex)
		}

		measure.MeasureNonnegativeDistance()

		if measure.Distance > clearance {
			if measure.Distance > oldDistance {
				tile.SetZ(box.GetMax().GetZ() - 1) // TODO: Consider
				oldDistance = math.MaxFloat64
			}

			return tile
		}

		tile = function(tile)

		if tile.GetZ() == box.GetMax().GetZ() {
			oldDistance = math.Inf(1)
		} else {
			oldDistance = measure.Distance
		}

		return tile
	})
}

func (box TileXYZBox) ForXYZ(function func(TileXYZ) TileXYZ) {
	current := TileXYZ{}

	for current.SetX(box.GetMin().GetX()); current.GetX() != box.GetMax().GetX(); current.SetX(current.GetX() + 1) {
		for current.SetY(box.GetMin().GetY()); current.GetY() != box.GetMax().GetY(); current.SetY(current.GetY() + 1) {
			for current.SetZ(box.GetMin().GetZ()); current.GetZ() != box.GetMax().GetZ(); current.SetZ(current.GetZ() + 1) {
				current = function(current)
			}

			current = function(current)
		}

		for current.SetZ(box.GetMin().GetZ()); current.GetZ() != box.GetMax().GetZ(); current.SetZ(current.GetZ() + 1) {
			current = function(current)
		}

		current = function(current)
	}

	for current.SetY(box.GetMin().GetY()); current.GetY() != box.GetMax().GetY(); current.SetY(current.GetY() + 1) {
		for current.SetZ(box.GetMin().GetZ()); current.GetZ() != box.GetMax().GetZ(); current.SetZ(current.GetZ() + 1) {
			current = function(current)
		}

		current = function(current)
	}

	for current.SetZ(box.GetMin().GetZ()); current.GetZ() != box.GetMax().GetZ(); current.SetZ(current.GetZ() + 1) {
		current = function(current)
	}

	current = function(current)
}

func NewTileXYZBoxFromSpatialIDBox(spatialIDBox SpatialIDBox) (*TileXYZBox, error) {
	deltaQuad := spatialIDBox.GetMin().GetZ() - SpatialIDZBaseExponent
	deltaAltitude := spatialIDBox.GetMax().GetZ() - TileXYZZBaseExponent
	spatialIDBox.AddZ(-deltaQuad)

	baseMinTile, error := NewTileXYZ(
		spatialIDBox.GetMin().GetZ(),
		spatialIDBox.GetMin().GetZ(),
		spatialIDBox.GetMin().GetX(),
		spatialIDBox.GetMin().GetY(),
		spatialIDBox.GetMin().GetF()-SpatialIDZBaseOffset+TileXYZZBaseOffset,
	)
	if error != nil {
		return nil, error
	}

	baseMaxTile, error := NewTileXYZ(
		spatialIDBox.GetMax().GetZ(),
		spatialIDBox.GetMin().GetZ(),
		spatialIDBox.GetMax().GetX(),
		spatialIDBox.GetMax().GetY(),
		spatialIDBox.GetMax().GetF()-SpatialIDZBaseOffset+TileXYZZBaseOffset,
	)
	if error != nil {
		return nil, error
	}

	box, error := NewTileXYZBox(*baseMinTile, *baseMaxTile)
	if error != nil {
		return nil, error
	}

	error = box.AddZoomLevel(deltaQuad, deltaAltitude)
	if error != nil {
		return nil, error
	}

	return box, nil
}

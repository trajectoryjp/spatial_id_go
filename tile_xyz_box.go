package spatialID

import (
	"iter"
	"math"

	"github.com/HarutakaMatsumoto/mathematics_go/geometry/rectangular/solid"
	"github.com/go-gl/mathgl/mgl64"
	closest "github.com/trajectoryjp/closest_go"
	"github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
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

	// Pass x
	if min.GetY() > max.GetY() {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	if min.GetZ() > max.GetZ() {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	return &TileXYZBox{
		min: min,
		max: max,
	}, nil
}

func NewTileXYZBoxFromGeodeticBox(geodeticBox GeodeticBox, quadkeyZoomLevel int8, altitudekeyZoomLevel int8) (*TileXYZBox, error) {
	minTile, error := NewTileXYZFromGeodetic(geodeticBox.Min, quadkeyZoomLevel, altitudekeyZoomLevel)
	if error != nil {
		return nil, error
	}

	maxTile, error := NewTileXYZFromGeodetic(geodeticBox.Max, quadkeyZoomLevel, altitudekeyZoomLevel)
	if error != nil {
		return nil, error
	}

	return NewTileXYZBox(*minTile, *maxTile)
}

func (box *TileXYZBox) AddZoomLevel(quadDelta, altitudeDelta int8) error {
	if quadDelta < 0 || altitudeDelta < 0 { // TODO: 入れ違いになった時落ちる
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

func (box TileXYZBox) IsCollidedWith(another TileXYZBox) bool {
	another.AddZoomLevel(box.GetMin().GetQuadkeyZoomLevel()-another.GetMin().GetQuadkeyZoomLevel(), box.GetMin().GetAltitudekeyZoomLevel()-another.GetMin().GetAltitudekeyZoomLevel())

	if box.GetMin().GetX() > another.GetMax().GetX() || box.GetMax().GetX() < another.GetMin().GetX() {
		return false
	}
	if box.GetMin().GetY() > another.GetMax().GetY() || box.GetMax().GetY() < another.GetMin().GetY() {
		return false
	}
	if box.GetMin().GetZ() > another.GetMax().GetZ() || box.GetMax().GetZ() < another.GetMin().GetZ() {
		return false
	}

	return true
}

func (box TileXYZBox) AllCollisionWithConvexHull(convexHull []*coordinates.Geodetic, clearance float64) iter.Seq[TileXYZ] {
	measure := closest.Measure{
		ConvexHulls: [2][]*mgl64.Vec3{
			make([]*mgl64.Vec3, len(convexHull)),
			make([]*mgl64.Vec3, len(solid.VertexIntervals)),
		},
	}

	for i, vertex := range convexHull {
		measure.ConvexHulls[0][i] = (*mgl64.Vec3)(vertex)
	}

	return iter.Seq[TileXYZ](func(yield func(tile TileXYZ) bool) {
		current := box.GetMin()
		for ; ; current.SetX(current.GetX() + 1) {
		yLoop:
			for current.SetY(box.GetMin().GetY()); ; current.SetY(current.GetY() + 1) {
				bottom := current
				oldDistance := math.Inf(1)
				for bottom.SetZ(box.GetMin().GetZ()); ; bottom.SetZ(bottom.GetZ() + 1) {
					tileXYZBox, _ := NewTileXYZBox(bottom, bottom)
					geodeticBox := NewGeodeticBoxFromTileXYZBox(*tileXYZBox)

					for i, vertex := range geodeticBox.GetVertices() {
						measure.ConvexHulls[1][i] = (*mgl64.Vec3)(vertex)
					}

					measure.MeasureNonnegativeDistance()

					geocentric0 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[0]))
					geocentric1 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[1]))
					distance := mgl64.Vec3(geocentric0).Sub(mgl64.Vec3(geocentric1)).Len() // TODO: Embed

					if distance > clearance {
						if distance > oldDistance {
							continue yLoop
						} else {
							deltaAltitude := *geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude()
							newZ := int64(distance/deltaAltitude) + bottom.GetZ()
							if newZ >= tileXYZBox.GetMax().GetZ() {
								continue yLoop
							}

							bottom.SetZ(newZ)
							continue
						}
					}

					break
				}

				top := current
				oldDistance = math.Inf(1)
				for top.SetZ(box.GetMax().GetZ()); ; top.SetZ(top.GetZ() - 1) {
					tileXYZBox, _ := NewTileXYZBox(top, top)
					geodeticBox := NewGeodeticBoxFromTileXYZBox(*tileXYZBox)

					for i, vertex := range geodeticBox.GetVertices() {
						measure.ConvexHulls[1][i] = (*mgl64.Vec3)(vertex)
					}

					measure.MeasureNonnegativeDistance()

					geocentric0 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[0]))
					geocentric1 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[1]))
					distance := mgl64.Vec3(geocentric0).Sub(mgl64.Vec3(geocentric1)).Len() // TODO: Embed

					if distance > clearance {
						if distance > oldDistance {
							continue yLoop
						} else {
							deltaAltitude := *geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude()
							newZ := -int64(distance/deltaAltitude) + top.GetZ()
							if newZ <= tileXYZBox.GetMin().GetZ() {
								continue yLoop
							}

							top.SetZ(newZ)
							continue
						}
					}

					break
				}

				for currentZ := bottom; currentZ.GetZ() <= top.GetZ(); currentZ.SetZ(currentZ.GetZ() + 1) {
					if !yield(currentZ) {
						return
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
	})
}

func (box TileXYZBox) AllXYZ() iter.Seq[TileXYZ] {
	current := box.GetMin()

	return iter.Seq[TileXYZ](func(yield func(tile TileXYZ) bool) {
		for ; ; current.SetX(current.GetX() + 1) {
			for current.SetY(box.GetMin().GetY()); ; current.SetY(current.GetY() + 1) {
				for current.SetZ(box.GetMin().GetZ()); ; current.SetZ(current.GetZ() + 1) {
					if !yield(current) {
						return
					}

					if current.GetZ() == box.GetMax().GetZ() {
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
	})
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

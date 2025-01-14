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

	if min.GetF() > max.GetF() {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	// Pass x
	if min.GetY() > max.GetY() {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	return &SpatialIDBox{
		min: min,
		max: max,
	}, nil
}

func NewSpatialIDBoxFromGeodeticBox(geodeticBox GeodeticBox, zoomLevel int8) (*SpatialIDBox, error) {
	minSpatialID, error := NewSpatialIDFromGeodetic(geodeticBox.Min, zoomLevel)
	if error != nil {
		return nil, error
	}

	maxSpatialID, error := NewSpatialIDFromGeodetic(geodeticBox.Max, zoomLevel)
	if error != nil {
		return nil, error
	}

	return NewSpatialIDBox(*minSpatialID, *maxSpatialID)
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

func (box SpatialIDBox) IsCollidedWith(another SpatialIDBox) bool {
	another.AddZ(box.GetMin().GetZ() - another.GetMin().GetZ())

	if box.GetMin().GetF() > another.GetMax().GetF() || box.GetMax().GetF() < another.GetMin().GetF() {
		return false
	}
	if box.GetMin().GetX() > another.GetMax().GetX() || box.GetMax().GetX() < another.GetMin().GetX() {
		return false
	}
	if box.GetMin().GetY() > another.GetMax().GetY() || box.GetMax().GetY() < another.GetMin().GetY() {
		return false
	}

	return true
}

func (box SpatialIDBox) AllCollisionWithConvexHull(convexHull []*coordinates.Geodetic, clearance float64) iter.Seq[SpatialID] {
	measure := closest.Measure{
		ConvexHulls: [2][]*mgl64.Vec3{
			make([]*mgl64.Vec3, len(convexHull)),
			make([]*mgl64.Vec3, len(solid.VertexIntervals)),
		},
	}

	for i, vertex := range convexHull {
		measure.ConvexHulls[0][i] = (*mgl64.Vec3)(vertex)
	}

	return iter.Seq[SpatialID](func(yield func(id SpatialID) bool) {
		current := box.GetMin()
		for ; ; current.SetX(current.GetX() + 1) {
		yLoop:
			for current.SetY(box.GetMin().GetY()); ; current.SetY(current.GetY() + 1) {
				bottom := current
				oldDistance := math.Inf(1)
				for bottom.SetF(box.GetMin().GetF()); ; bottom.SetF(bottom.GetF() + 1) {
					spatialIDBox, _ := NewSpatialIDBox(bottom, bottom)
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
							continue yLoop
						} else {
							deltaAltitude := *geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude()
							newF := int64(distance/deltaAltitude) + bottom.GetF()
							if newF >= spatialIDBox.GetMax().GetF() {
								continue yLoop
							}

							bottom.SetF(newF)
							continue
						}
					}

					break
				}

				top := current
				oldDistance = math.Inf(1)
				for top.SetF(box.GetMax().GetF()); ; top.SetF(top.GetF() - 1) {
					spatialIDBox, _ := NewSpatialIDBox(top, top)
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
							continue yLoop
						} else {
							deltaAltitude := *geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude()
							newF := -int64(distance/deltaAltitude) + top.GetF()
							if newF <= spatialIDBox.GetMin().GetF() {
								continue yLoop
							}

							top.SetF(newF)
							continue
						}
					}

					break
				}

				for currentF := bottom; currentF.GetF() <= top.GetF(); currentF.SetF(currentF.GetF() + 1) {
					if !yield(currentF) {
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

func (box SpatialIDBox) AllZXYF() iter.Seq[SpatialID] {
	generations := [MaxZ]SpatialIDBox{}

	ancestor := box
	error := ancestor.AddZ(-1)
	for ; error != nil; error = ancestor.AddZ(-1) {
		generations[ancestor.GetMin().GetZ()] = ancestor
	}

	for ; error != nil; error = box.AddZ(1) {
		generations[box.GetMin().GetZ()] = box
	}

	return iter.Seq[SpatialID](func(yield func(id SpatialID) bool) {
		for _, generation := range generations {
			for id := range generation.AllXYF() {
				if !yield(id) {
					return
				}
			}
		}
	})
}

func (box SpatialIDBox) AllXYF() iter.Seq[SpatialID] {
	return iter.Seq[SpatialID](func(yield func(id SpatialID) bool) {
		current := box.GetMin()
		for ; ; current.SetX(current.GetX() + 1) {
			for current.SetY(box.GetMin().GetY()); ; current.SetY(current.GetY() + 1) {
				for current.SetF(box.GetMin().GetF()); ; current.SetF(current.GetF() + 1) {
					if !yield(current) {
						return
					}

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
	})
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

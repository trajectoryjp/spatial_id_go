package spatialID

import (
	"iter"
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

	if min.GetF() > max.GetF() {
		return nil, NewSpatialIdError(InputValueErrorCode, "")
	}
	// Pass x
	if min.GetY() > max.GetY() {
		return nil, NewSpatialIdError(InputValueErrorCode, "")
	}

	return &SpatialIDBox{
		min: min,
		max: max,
	}, nil
}

func NewSpatialIDBoxFromGeodeticBox(geodeticBox GeodeticBox, zoomLevel int8) (*SpatialIDBox, error) {
	minSpatialID, error := NewSpatialIDFromGeodetic(
		coordinates.Geodetic{
			*geodeticBox.Min.Longitude(),
			*geodeticBox.Max.Latitude(),
			*geodeticBox.Min.Altitude(),
		},
		zoomLevel,
	)
	if error != nil {
		return nil, error
	}

	maxSpatialID, error := NewSpatialIDFromGeodetic(
		coordinates.Geodetic{
			*geodeticBox.Max.Longitude(),
			*geodeticBox.Min.Latitude(),
			*geodeticBox.Max.Altitude(),
		},
		zoomLevel,
	)
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

// TODO: Test
func (box SpatialIDBox) Contains(another SpatialIDBox) bool {
	deltaZ := box.GetMin().GetZ() - another.GetMin().GetZ()
	if deltaZ > 0 {
		another.AddZ(deltaZ)
	} else if deltaZ < 0 {
		box.AddZ(-deltaZ)
	}

	if box.GetMin().GetF() > another.GetMin().GetF() || box.GetMax().GetF() < another.GetMax().GetF() {
		return false
	}
	if box.GetMin().GetX() > another.GetMin().GetX() || box.GetMax().GetX() < another.GetMax().GetX() {
		return false
	}
	if box.GetMin().GetY() > another.GetMin().GetY() || box.GetMax().GetY() < another.GetMax().GetY() {
		return false
	}

	return true
}

// TODO: Test
func (box SpatialIDBox) Overlaps(another SpatialIDBox) bool {
	deltaZ := box.GetMin().GetZ() - another.GetMin().GetZ()
	if deltaZ > 0 {
		another.AddZ(deltaZ)
	} else if deltaZ < 0 {
		box.AddZ(-deltaZ)
	}

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
			for current.SetY(box.GetMin().GetY()); ; current.SetY(current.GetY() + 1) {
				bottom := current
				top := current
				oldDistance := math.Inf(1)
				for bottom.SetF(box.GetMin().GetF()); ; bottom.SetF(bottom.GetF() + 1) {
					spatialIDBox, _ := NewSpatialIDBox(bottom, bottom)
					geodeticBox := NewGeodeticBoxFromSpatialIDBox(*spatialIDBox)

					for i, vertex := range geodeticBox.GetVertices() {
						measure.ConvexHulls[1][i] = (*mgl64.Vec3)(vertex)
					}

					measure.MeasureNonnegativeDistance()

					distance := measure.Distance
					if distance > 0.0 {
						geocentric0 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[0]))
						geocentric1 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[1]))
						distance = mgl64.Vec3(geocentric0).Sub(mgl64.Vec3(geocentric1)).Len() // TODO: Embed
					}

					if distance > clearance {
						if distance > oldDistance {
							goto yEnd
						} else {
							deltaAltitude := *geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude()
							newF := int64(distance/deltaAltitude) + bottom.GetF()
							if newF >= spatialIDBox.GetMax().GetF() {
								goto yEnd
							}

							bottom.SetF(newF)
							continue
						}
					}

					break
				}

				oldDistance = math.Inf(1)
				for top.SetF(box.GetMax().GetF()); ; top.SetF(top.GetF() - 1) {
					spatialIDBox, _ := NewSpatialIDBox(top, top)
					geodeticBox := NewGeodeticBoxFromSpatialIDBox(*spatialIDBox)

					for i, vertex := range geodeticBox.GetVertices() {
						measure.ConvexHulls[1][i] = (*mgl64.Vec3)(vertex)
					}

					measure.MeasureNonnegativeDistance()

					distance := measure.Distance
					if distance > 0.0 {
						geocentric0 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[0]))
						geocentric1 := coordinates.GeocentricFromGeodetic(coordinates.Geodetic(measure.Points[1]))
						distance = mgl64.Vec3(geocentric0).Sub(mgl64.Vec3(geocentric1)).Len() // TODO: Embed
					}

					if distance > clearance {
						if distance > oldDistance {
							goto yEnd
						} else {
							deltaAltitude := *geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude()
							newF := -int64(distance/deltaAltitude) + top.GetF()
							if newF <= spatialIDBox.GetMin().GetF() {
								goto yEnd
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

			yEnd:
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
	deltaBaseExponent := SpatialIDZBaseExponent - TileXYZZBaseExponent
	deltaBaseZ := MaxZ - SpatialIDZBaseExponent
	deltaOffset := (SpatialIDZBaseOffset - TileXYZZBaseOffset) << deltaBaseZ

	deltaQuad := tileXYZBox.GetMin().GetQuadkeyZoomLevel() - MaxQuadkeyZoomLevel
	deltaAltitude := tileXYZBox.GetMin().GetAltitudekeyZoomLevel() - (MaxAltitudekeyZoomLevel - deltaBaseExponent)

	error := tileXYZBox.AddZoomLevel(-deltaQuad, -deltaAltitude)
	if error != nil {
		return nil, error
	}

	baseMinID, error := NewSpatialID(
		tileXYZBox.GetMin().GetQuadkeyZoomLevel(),
		tileXYZBox.GetMin().GetZ()+deltaOffset,
		tileXYZBox.GetMin().GetX(),
		tileXYZBox.GetMin().GetY(),
	)
	if error != nil {
		return nil, error
	}

	baseMaxID, error := NewSpatialID(
		tileXYZBox.GetMax().GetQuadkeyZoomLevel(),
		tileXYZBox.GetMax().GetZ()+deltaOffset,
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

	maxDelta := deltaQuad
	if deltaAltitude > maxDelta {
		maxDelta = deltaAltitude
	}

	error = box.AddZ(maxDelta)
	if error != nil {
		return nil, error
	}

	return box, nil
}

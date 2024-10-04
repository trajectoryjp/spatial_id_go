package spatialID

import tileXYZ "github.com/trajectoryjp/spatial_id_go/v4/tile_xyz"

type Box struct {
	min SpatialID
	max SpatialID
}

func NewBox(min SpatialID, max SpatialID) (*Box, error) {
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

	return &Box{
		min: min,
		max: max,
	}, nil
}

func (box *Box) AddZ(delta int8) error {
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

func (box Box) GetMin() SpatialID {
	return box.min
}

func (box Box) GetMax() SpatialID {
	return box.max
}

func (box Box) ForXYF(function func(SpatialID) SpatialID) {
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

func NewBoxFromTileXYZBox(tileXYZBox tileXYZ.Box) (*Box, error) {
	deltaZ := spatialIDBox.GetMin().GetZ() - spatialID.ZBaseExponent
	spatialIDBox.AddZ(-deltaZ)

	baseMinTile, error := NewTileXYZ(
		spatialIDBox.GetMin().GetZ(),
		ZBaseExponent,
		spatialIDBox.GetMin().GetX(),
		spatialIDBox.GetMin().GetY(),
		spatialIDBox.GetMin().GetF() >> (spatialID.ZBaseOffset - ZBaseOffset),
	)
	if error != nil {
		return nil, error
	}

	baseMaxTile, error := NewTileXYZ(
		spatialIDBox.GetMax().GetZ(),
		ZBaseExponent,
		spatialIDBox.GetMax().GetX(),
		spatialIDBox.GetMax().GetY(),
		spatialIDBox.GetMax().GetF() >> (spatialID.ZBaseOffset - ZBaseOffset),
	)
	if error != nil {
		return nil, error
	}

	box, error := NewBox(*baseMinTile, *baseMaxTile)
	if error != nil {
		return nil, error
	}

	error = box.AddZoomLevel(deltaZ, deltaZ)
	if error != nil {
		return nil, error
	}

	return box, nil
}

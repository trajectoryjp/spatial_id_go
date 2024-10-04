package tileXYZ

import spatialID "github.com/trajectoryjp/spatial_id_go/v4"

type Box struct {
	min TileXYZ
	max TileXYZ
}

func NewBox(min TileXYZ, max TileXYZ) (*Box, error) {
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

	return &Box{
		min: min,
		max: max,
	}, nil
}

func (box *Box) AddZoomLevel(quadDelta, altitudeDelta int8) error {
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

func (box Box) GetMin() TileXYZ {
	return box.min
}

func (box Box) GetMax() TileXYZ {
	return box.max
}

func (box Box) ForXYZ(function func(TileXYZ) TileXYZ) {
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

func NewBoxFromSpatialIDBox(spatialIDBox spatialID.Box) (*Box, error) {
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

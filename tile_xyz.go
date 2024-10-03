package spatialID

import (
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
)

const MaxQuadkeyZoomLevel = 35
const MaxAltitudekeyZoomLevel = 35

type TileXYZ struct {
	quadkeyZoomLevel int8
	altitudekeyZoomLevel int8
	x                   int64
	y                   int64
	z                   int64
}

func NewTileXYZ(
	quadkeyZoomLevel int8,
	altitudekeyZoomLevel int8,
	x                   int64,
	y                   int64,
	z                   int64,
) (*TileXYZ, error) {
	tile := &TileXYZ{}

	if quadkeyZoomLevel < 0 || MaxQuadkeyZoomLevel < quadkeyZoomLevel {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	tile.quadkeyZoomLevel = quadkeyZoomLevel

	if altitudekeyZoomLevel < 0 || MaxAltitudekeyZoomLevel < altitudekeyZoomLevel {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	tile.altitudekeyZoomLevel = altitudekeyZoomLevel

	tile.SetX(x)
	tile.SetY(y)
	tile.SetZ(z)

	return tile, nil
}

func (tile *TileXYZ) SetX(x int64) {
	tile.x = x%(1 << tile.GetQuadkeyZoomLevel())
	if tile.x < 0 {
		tile.x += 1 << tile.GetQuadkeyZoomLevel()
	}
}

func (tile *TileXYZ) SetY(y int64) {
	tile.y = y%(1 << tile.GetQuadkeyZoomLevel())
	if tile.y < 0 {
		tile.y += 1 << tile.GetQuadkeyZoomLevel()
	}
}

func (tile *TileXYZ) SetZ(z int64) {
	tile.z = z%(1 << tile.GetAltitudekeyZoomLevel())
	if tile.z < 0 {
		tile.z += 1 << tile.GetAltitudekeyZoomLevel()
	}
}

func (tile TileXYZ) GetQuadkeyZoomLevel() int8 {
	return tile.quadkeyZoomLevel
}

func (tile TileXYZ) GetAltitudekeyZoomLevel() int8 {
	return tile.altitudekeyZoomLevel
}

func (tile TileXYZ) GetX() int64 {
	return tile.x
}

func (tile TileXYZ) GetY() int64 {
	return tile.y
}

func (tile TileXYZ) GetZ() int64 {
	return tile.z
}

func (tile TileXYZ) NewParent(quadNumber, altitudeNumber int8) (*TileXYZ, error) {
	return NewTileXYZ(
		tile.GetQuadkeyZoomLevel()-quadNumber,
		tile.GetAltitudekeyZoomLevel()-altitudeNumber,
		tile.GetX() >> quadNumber,
		tile.GetY() >> quadNumber,
		tile.GetZ() >> altitudeNumber,
	)
}

func (tile TileXYZ) NewMinChild(quadNumber, altitudeNumber int8) (*TileXYZ, error) {
	return NewTileXYZ(
		tile.GetQuadkeyZoomLevel()+quadNumber,
		tile.GetAltitudekeyZoomLevel()+altitudeNumber,
		tile.GetX() << quadNumber,
		tile.GetY() << quadNumber,
		tile.GetZ() << altitudeNumber,
	)
}

func (tile TileXYZ) NewMaxChild(quadNumber, altitudeNumber int8) (*TileXYZ, error) {
	return NewTileXYZ(
		tile.GetQuadkeyZoomLevel()+quadNumber,
		tile.GetAltitudekeyZoomLevel()+altitudeNumber,
		(tile.GetX() << quadNumber) + (1 << quadNumber) - 1,
		(tile.GetY() << quadNumber) + (1 << quadNumber) - 1,
		(tile.GetZ() << altitudeNumber) + (1 << altitudeNumber) - 1,
	)
}

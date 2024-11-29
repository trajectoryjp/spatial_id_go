package spatialID

import (
	"math"

	mathematics "github.com/HarutakaMatsumoto/mathematics_go"
	"github.com/trajectoryjp/geodesy_go/coordinates"
)

const MaxQuadkeyZoomLevel = 35
const MaxAltitudekeyZoomLevel = 35
var TileXYZZBaseExponent int8 = 14
var TileXYZZBaseOffset int64 = 512

type TileXYZ struct {
	quadkeyZoomLevel int8
	altitudekeyZoomLevel int8
	x                   int64
	y                   int64
	z                   int64
}

func NewTileXYZFromGeodetic(geodetic coordinates.Geodetic, quadkeyZoomLevel int8, altitudeZoomLevel int8) (*TileXYZ, error) {
	quadMax := float64(int(1) << quadkeyZoomLevel)

	// 経度方向のインデックスの計算
	x := int64(math.Floor(quadMax * math.Mod(*geodetic.Longitude() + 180.0, 360.0)))

	radianLatitude := mathematics.RadianPerDegree * *geodetic.Latitude()

	y := int64(math.Floor(quadMax * (1.0 - math.Log(math.Tan(radianLatitude)+(1.0/math.Cos(radianLatitude)))/math.Pi) / 2.0))

	// 高さ全体の精度あたりの垂直方向の精度
	altitudeResolution := float64(CalculateArithmeticShift(1, int64(TileXYZZBaseExponent-altitudeZoomLevel)))

	// 垂直方向の位置を計算する
	z := int64(math.Floor(*geodetic.Altitude() / altitudeResolution))+TileXYZZBaseOffset

	return NewTileXYZ(quadkeyZoomLevel, altitudeZoomLevel, x, y, z)
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
		return nil, NewSpatialIdError(InputValueErrorCode, "")
	}
	tile.quadkeyZoomLevel = quadkeyZoomLevel

	if altitudekeyZoomLevel < 0 || MaxAltitudekeyZoomLevel < altitudekeyZoomLevel {
		return nil, NewSpatialIdError(InputValueErrorCode, "")
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
	max := int64(1 << tile.GetQuadkeyZoomLevel()) -  1
	min := int64(0)

	if y > max {
		tile.y = max
	} else if y < min {
		tile.y = min
	} else {
		tile.y = y
	}
}

func (tile *TileXYZ) SetZ(z int64) {
	max := int64(1 << tile.GetAltitudekeyZoomLevel()) - 1
	min := int64(0)

	if z > max {
		tile.z = max
	} else if z < min {
		tile.z = min
	} else {
		tile.z = z
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

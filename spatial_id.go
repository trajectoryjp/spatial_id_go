package spatialID

import (
	"math"
	"strconv"
	"strings"

	mathematics "github.com/HarutakaMatsumoto/mathematics_go"
	"github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/spatial_id_go/v4/common"
	"github.com/trajectoryjp/spatial_id_go/v4/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
)

const MaxZ = 35
const SpatialIDZBaseExponent int8 = 25
const SpatialIDZBaseOffset int64 = 0

const delimiter = "/"

// SpatialID 空間IDクラス
type SpatialID struct {
	z int8 // 精度
	f                   int64 // 高さID
	x                   int64 // 経度ID
	y                   int64 // 緯度ID
}

func NewSpatialIDFromString(string string) (*SpatialID, error) {
	// 空間IDを区切り文字で分割
	attributes := strings.Split(string, consts.SpatialIDDelimiter)

	if len(attributes) != 4 {
		// 区切り文字数がフォーマットに従っていない場合エラーインスタンスを返却
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 整数型変換後のフィールド値格納用
	int64s := []int64{}
	bitLengths := []int{8, 64, 64, 64}
	for i, attribute := range attributes {
		// フィールド値をint64型に返却
		int64_, error := strconv.ParseInt(attribute, 10, bitLengths[i])

		if error != nil {
			// int64変換時にエラーが発生した場合エラーインスタンスを返却
			return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}

		int64s = append(int64s, int64_)
	}

	return NewSpatialID(int8(int64s[0]), int64s[1], int64s[2], int64s[3])
}

func NewSpatialIDFromGeodetic(geodetic coordinates.Geodetic, z int8) (*SpatialID, error) {
	max := float64(int(1) << z)

	// 経度方向のインデックスの計算
	x := math.Floor(max * math.Mod(*geodetic.Longitude() + 180.0, 360.0))

	radianLatitude := mathematics.RadianPerDegree * *geodetic.Latitude()

	y := math.Floor(max * (1.0 - math.Log(math.Tan(radianLatitude)+(1.0/math.Cos(radianLatitude)))/math.Pi) / 2.0)

	// 高さ全体の精度あたりの垂直方向の精度
	altitudeResolution := float64(common.CalculateArithmeticShift(1, int64(SpatialIDZBaseExponent-z)))

	// 垂直方向の位置を計算する
	f := math.Floor(*geodetic.Altitude() / altitudeResolution)

	return NewSpatialID(z, int64(f), int64(x), int64(y))
}

func NewSpatialID(
	z int8,
	f int64,
	x int64,
	y int64,
) (*SpatialID, error) {
	id := &SpatialID{}

	if z < 0 || MaxZ < z {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	id.z = z

	id.SetF(f)
	id.SetX(x)
	id.SetY(y)

	return id, nil
}

func (id *SpatialID) SetX(x int64) {
	id.x = x%(1 << id.GetZ())
	if id.x < 0 {
		id.x += 1 << id.GetZ()
	}
}

func (id *SpatialID) SetY(y int64) {
	id.y = y%(1 << id.GetZ())
	if id.y < 0 {
		id.y += 1 << id.GetZ()
	}
}

func (id *SpatialID) SetF(f int64) {
	max := int64(1 << id.GetZ())
	min := -max
	max -= 1

	if f > max {
		id.f = max
	} else if f < min {
		id.f = min
	} else {
		id.f = f
	}
}

func (id SpatialID) GetZ() int8 {
	return id.z
}

func (id SpatialID) GetF() int64 {
	return id.f
}

func (id SpatialID) GetX() int64 {
	return id.x
}

func (id SpatialID) GetY() int64 {
	return id.y
}

func (id SpatialID) String() string {
	return strconv.FormatInt(int64(id.GetZ()), 10) + delimiter +
	strconv.FormatInt(id.GetF(), 10) + delimiter +
	strconv.FormatInt(id.GetX(), 10) + delimiter +
	strconv.FormatInt(id.GetY(), 10)
}

func (id SpatialID) NewParent(number int8) (*SpatialID, error) {
	return NewSpatialID(
		id.GetZ()-number,
		id.GetF() >> number,
		id.GetX() >> number,
		id.GetY() >> number,
	)
}

func (id SpatialID) NewMinChild(number int8) (*SpatialID, error) {
	return NewSpatialID(
		id.GetZ()+number,
		id.GetF() << number,
		id.GetX() << number,
		id.GetY() << number,
	)
}

func (id SpatialID) NewMaxChild(number int8) (*SpatialID, error) {
	return NewSpatialID(
		id.GetZ()+number,
		(id.GetF() << number) + (1 << number) - 1,
		(id.GetX() << number) + (1 << number) - 1,
		(id.GetY() << number) + (1 << number) - 1,
	)
}

package spatialID

import (
	"math"
	"strconv"
	"strings"

	mathematics "github.com/HarutakaMatsumoto/mathematics_go"
	"github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/multidimensional-radix-tree/src/tree"
	"github.com/trajectoryjp/spatial_id_go/v4/common"
	"github.com/trajectoryjp/spatial_id_go/v4/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
)

const MaxZ = 35
const SpatialIDZBaseExponent int8 = 25
const SpatialIDZBaseOffset int64 = 0
const SpatialIDMaxNumberOfChildren = 8

const delimiter = "/"

var SpatialIDZoomSetTable = tree.Create3DTable()

// SpatialID 空間IDクラス
type SpatialID struct {
	z int8  // 精度
	f int64 // 高さID
	x int64 // 経度ID
	y int64 // 緯度ID
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
	x := math.Floor(max * math.Mod(*geodetic.Longitude()+180.0, 360.0))

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
	limit := int64(1 << id.GetZ())

	id.x = x % limit
	if id.x < 0 {
		id.x += limit
	}
}

func (id *SpatialID) SetY(y int64) {
	limit := int64(1 << id.GetZ())
	max := limit - 1
	min := int64(0)

	if y > max {
		id.y = max
	} else if y < min {
		id.y = min
	} else {
		id.y = y
	}
}

func (id *SpatialID) SetF(f int64) {
	limit := int64(1 << id.GetZ())
	max := limit - 1
	min := -limit

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
		id.GetF()>>number,
		id.GetX()>>number,
		id.GetY()>>number,
	)
}

func (id SpatialID) NewMinChild(number int8) (*SpatialID, error) {
	return NewSpatialID(
		id.GetZ()+number,
		id.GetF()<<number,
		id.GetX()<<number,
		id.GetY()<<number,
	)
}

func (id SpatialID) NewMaxChild(number int8) (*SpatialID, error) {
	return NewSpatialID(
		id.GetZ()+number,
		(id.GetF()+1)<<number-1,
		(id.GetX()+1)<<number-1,
		(id.GetY()+1)<<number-1,
	)
}

func (id SpatialID) Contains(another SpatialID) bool {
	return id.GetZ() <= another.GetZ() && id.Overlaps(another)
}

func (id SpatialID) Overlaps(another SpatialID) bool {
	deltaZ := id.GetZ() - another.GetZ()
	if deltaZ < 0 {
		anotherParent, _ := another.NewParent(-deltaZ)
		another = *anotherParent
	} else if deltaZ > 0 {
		idParent, _ := id.NewParent(deltaZ)
		id = *idParent
	}

	return id == another
}

func MergeSpatialIDs(ids []*SpatialID) []*SpatialID {
	positiveSpatialIDs := []*SpatialID{}
	negativeSpatialIDs := []*SpatialID{}
	for _, id := range ids {
		if id.GetF() >= 0 {
			positiveSpatialIDs = append(positiveSpatialIDs, id)
		} else {
			id.SetF(^id.GetF())
			negativeSpatialIDs = append(negativeSpatialIDs, id)
		}
	}

	positiveSpatialIDs = mergeSpatialIDs(positiveSpatialIDs, 0)
	negativeSpatialIDs = mergeSpatialIDs(negativeSpatialIDs, 0)

	resultIDs := positiveSpatialIDs
	for _, id := range negativeSpatialIDs {
		id.SetF(^id.GetF())
		resultIDs = append(resultIDs, id)
	}
	return resultIDs
}

// All f index must be greater than or equal to 0.
func mergeSpatialIDs(ids []*SpatialID, z int8) []*SpatialID {
	if len(ids) == 0 {
		return []*SpatialID{}
	}

	idsTable := [SpatialIDMaxNumberOfChildren][]*SpatialID{}
	for _, id := range ids {
		if id.GetZ() == z {
			return []*SpatialID{id}
		}

		parent, _ := id.NewParent(id.GetZ() - (z + 1))
		i := (parent.GetF()&1)<<2 | (parent.GetX()&1)<<1 | (parent.GetY() & 1)
		idsTable[i] = append(idsTable[i], id)
	}

	resultIDs := []*SpatialID{}
	for _, ids := range idsTable {
		resultIDs = append(resultIDs, mergeSpatialIDs(ids, z+1)...)
	}

	if len(resultIDs) == SpatialIDMaxNumberOfChildren {
		id, _ := resultIDs[0].NewParent(1)
		return []*SpatialID{id}
	}
	return resultIDs
}

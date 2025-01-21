package spatialID

import (
	"math"
	// "slices"
	"strconv"
	"strings"

	mathematics "github.com/HarutakaMatsumoto/mathematics_go"
	"github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/multidimensional-radix-tree/src/tree"
)

const MaxZ = 35
const SpatialIDZBaseExponent int8 = 25
const SpatialIDZBaseOffset int64 = 0

const delimiter = "/"

var SpatialIDZoomSetTable = tree.Create3DTable()

// func MergeSpatialIDs (spatialIDs []*SpatialID) []*SpatialID {
// 	type element struct {
// 		id SpatialID
// 		isExprored bool
// 	}
// 	compare := func(a, b *element) int {
// 		return CompareSpatialIDs(&a.id, &b.id)
// 	}

// 	zooms := [MaxZ + 1][]*element{}
// 	for _, spatialID := range spatialIDs {
// 		zooms[spatialID.GetZ()] = append(zooms[spatialID.GetZ()], &element{
// 			id: *spatialID,
// 			isExprored: false,
// 		})
// 	}

// 	for _, zoom := range zooms {
// 		slices.SortFunc(zoom, compare)
// 	}

// 	for i := MaxZ; i >= 0; i -= 1 {
// 		zoom := zooms[i]
// 		for j := len(zoom) - 1; j >= 0; j -= 1 {
// 			// {0, 0, 0}
// 			element0 := zoom[j]
// 			if element0.isExprored {
// 				continue
// 			}
// 			element0.isExprored = true

// 			if element0.id.GetF() % 2 == 0 {
// 				continue
// 			}

// 			// {1, 0, 0}
// 			j -= 1
// 			element1 := zoom[j]
// 			if element1.isExprored {
// 				continue
// 			}
// 			element1.isExprored = true

// 			for element1.id.GetF() != element0.id.GetF() {
// 				j -= 1
// 				element1 = zoom[j]
// 				if element1.isExprored {
// 					continue
// 				}
// 				element1.isExprored = true
// 			}
// 			if element1.id.GetF() != element0.id.GetF() + 1 {
// 				continue
// 			}
// 			if element1.id.GetX() != element0.id.GetX() {
// 				continue
// 			}
// 			if element1.id.GetY() != element0.id.GetY() {
// 				continue
// 			}

// 			// {0, 1, 0}
// 			slices.BinarySearchFunc(zoom, &element{
// 				id
// 			})
// 		}

// 	}

// 	// 空間IDのマージ
// 	mergedSpatialIDs := []*SpatialID{}
// 	for _, spatialID := range spatialIDs {
// 		mergedSpatialIDs = append(mergedSpatialIDs, spatialID)
// 	}

// 	return mergedSpatialIDs
// }

func CompareSpatialIDs (a, b *SpatialID) int {
	if a.GetZ() < b.GetZ() {
		return -1
	} else if a.GetZ() > b.GetZ() {
		return 1
	}

	if a.GetF() < b.GetF() {
		return -1
	} else if a.GetF() > b.GetF() {
		return 1
	}

	if a.GetX() < b.GetX() {
		return -1
	} else if a.GetX() > b.GetX() {
		return 1
	}

	if a.GetY() < b.GetY() {
		return -1
	} else if a.GetY() > b.GetY() {
		return 1
	}

	return 0
}

// SpatialID 空間IDクラス
type SpatialID struct {
	z int8 // 精度
	f                   int64 // 高さID
	x                   int64 // 経度ID
	y                   int64 // 緯度ID
}

func NewSpatialIDFromString(string string) (*SpatialID, error) {
	// 空間IDを区切り文字で分割
	attributes := strings.Split(string, delimiter)

	if len(attributes) != 4 {
		// 区切り文字数がフォーマットに従っていない場合エラーインスタンスを返却
		return nil, NewSpatialIdError(InputValueErrorCode, "")
	}

	// 整数型変換後のフィールド値格納用
	int64s := []int64{}
	bitLengths := []int{8, 64, 64, 64}
	for i, attribute := range attributes {
		// フィールド値をint64型に返却
		int64_, error := strconv.ParseInt(attribute, 10, bitLengths[i])

		if error != nil {
			// int64変換時にエラーが発生した場合エラーインスタンスを返却
			return nil, NewSpatialIdError(InputValueErrorCode, "")
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
	altitudeResolution := float64(CalculateArithmeticShift(1, int64(SpatialIDZBaseExponent-z)))

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
		return nil, NewSpatialIdError(InputValueErrorCode, "")
	}
	id.z = z

	id.SetF(f)
	id.SetX(x)
	id.SetY(y)

	return id, nil
}

func (id *SpatialID) SetX(x int64) {
	id.x = x%(1 << id.GetZ()) // TODO: 統合
	if id.x < 0 {
		id.x += 1 << id.GetZ()
	}
}

func (id *SpatialID) SetY(y int64) {
	max := int64(1 << id.GetZ()) -  1
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
		(id.GetF() << number) + (1 << number) - 1, // TODO: 統合
		(id.GetX() << number) + (1 << number) - 1,
		(id.GetY() << number) + (1 << number) - 1,
	)
}

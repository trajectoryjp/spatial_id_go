package object

import (
	"strconv"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/v4/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
)

const MaxHorizontalZoomLevel = 35
const MaxVerticalZoomLevel = 35

const delimiter = "/"

// ExtendedSpatialID 拡張空間IDクラス
type ExtendedSpatialID struct {
	horizontalZoomLevel int8 // 水平方向精度
	x                   int64 // 経度ID
	y                   int64 // 緯度ID
	verticalZoomLevel   int8 // 垂直方向精度
	z                   int64 // 高さID
}

func NewExtendedSpatialIDFromString(string string) (*ExtendedSpatialID, error) {
	// 拡張空間IDを区切り文字で分割
	attributes := strings.Split(string, consts.SpatialIDDelimiter)

	if len(attributes) != 5 {
		// 区切り文字数がフォーマットに従っていない場合エラーインスタンスを返却
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 整数型変換後のフィールド値格納用
	int64s := []int64{}
	bitLengths := []int{8, 64, 64, 8, 64}
	for i, attribute := range attributes {
		// フィールド値をint64型に返却
		int64_, error := strconv.ParseInt(attribute, 10, bitLengths[i])

		if error != nil {
			// int64変換時にエラーが発生した場合エラーインスタンスを返却
			return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}

		int64s = append(int64s, int64_)
	}

	return NewExtendedSpatialID(int8(int64s[0]), int64s[1], int64s[2], int8(int64s[3]), int64s[4])
}

func NewExtendedSpatialID(
	horizontalZoomLevel int8,
	x int64,
	y int64,
	verticalZoomLevel int8,
	z int64,
) (*ExtendedSpatialID, error) {
	id := &ExtendedSpatialID{}

	if horizontalZoomLevel < 0 || MaxHorizontalZoomLevel < horizontalZoomLevel {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	id.horizontalZoomLevel = horizontalZoomLevel

	if verticalZoomLevel < 0 || MaxVerticalZoomLevel < verticalZoomLevel {
		return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	id.verticalZoomLevel = verticalZoomLevel

	id.SetX(x)
	id.SetY(y)
	id.SetZ(z)

	return id, nil
}

// func (id *ExtendedSpatialID) SetHorizontalZoomLevel(horizontalZoomLevel int8) error {
// 	if horizontalZoomLevel < 0 || MaxHorizontalZoomLevel < horizontalZoomLevel {
// 		return errors.NewSpatialIdError(errors.InputValueErrorCode, "")
// 	}

// 	id.horizontalZoomLevel = horizontalZoomLevel

// 	id.SetX(id.GetX())
// 	id.SetY(id.GetY())
// 	return nil
// }

func (id *ExtendedSpatialID) SetX(x int64) {
	id.x = x%(1 << id.GetHorizontalZoomLevel())
	if id.x < 0 {
		id.x += 1 << id.GetHorizontalZoomLevel()
	}
}

// SetY 緯度ID設定関数
//
// ExtendedSpatialIDオブジェクトのyを引数の入力値に設定する。
//
// 引数：
//
//	y：緯度ID
func (id *ExtendedSpatialID) SetY(y int64) {
	id.y = y%(1 << id.GetHorizontalZoomLevel())
	if id.y < 0 {
		id.y += 1 << id.GetHorizontalZoomLevel()
	}
}

// func (id *ExtendedSpatialID) SetVerticalZoomLevel(verticalZoomLevel int8) error {
// 	if verticalZoomLevel < 0 || MaxVerticalZoomLevel < verticalZoomLevel {
// 		return errors.NewSpatialIdError(errors.InputValueErrorCode, "")
// 	}

// 	id.verticalZoomLevel = verticalZoomLevel
// 	return nil
// }

// SetZ 高さID設定関数
//
// ExtendedSpatialIDオブジェクトのzを引数の入力値に設定する。
//
// 引数：
//
//	z：高さID
func (id *ExtendedSpatialID) SetZ(z int64) {
	max := int64(1 << id.GetVerticalZoomLevel())
	min := -max
	max -= 1

	if z > max {
		id.z = max
	} else if z < min {
		id.z = min
	} else {
		id.z = z
	}
}

func (id ExtendedSpatialID) GetHorizontalZoomLevel() int8 {
	return id.horizontalZoomLevel
}

// X 経度ID取得関数
//
// ExtendedSpatialIDオブジェクトの経度IDを返却する。
//
// 戻り値：
//
//	経度ID
func (id ExtendedSpatialID) GetX() int64 {
	return id.x
}

// Y 緯度ID取得関数
//
// ExtendedSpatialIDオブジェクトの緯度IDを返却する。
//
// 戻り値：
//
//	緯度ID
func (id ExtendedSpatialID) GetY() int64 {
	return id.y
}


func (id ExtendedSpatialID) GetVerticalZoomLevel() int8 {
	return id.verticalZoomLevel
}

// Z 高さID取得関数
//
// ExtendedSpatialIDオブジェクトの高さIDを返却する。
//
// 戻り値：
//
//	高さID
func (id ExtendedSpatialID) GetZ() int64 {
	return id.z
}

func (id ExtendedSpatialID) String() string {
	return strconv.FormatInt(int64(id.GetHorizontalZoomLevel()), 10) + delimiter +
	strconv.FormatInt(id.GetX(), 10) + delimiter +
	strconv.FormatInt(id.GetY(), 10) + delimiter +
	strconv.FormatInt(int64(id.GetVerticalZoomLevel()), 10) + delimiter +
	strconv.FormatInt(id.GetZ(), 10)
}

func (id ExtendedSpatialID) NewParent(horizonalNumber, verticalNumber int8) (*ExtendedSpatialID, error) {
	return NewExtendedSpatialID(
		id.GetHorizontalZoomLevel()-horizonalNumber,
		id.GetX() >> horizonalNumber,
		id.GetY() >> horizonalNumber,
		id.GetVerticalZoomLevel()-verticalNumber,
		id.GetZ() >> verticalNumber,
	)
}

func (id ExtendedSpatialID) NewMinChild(horizonalNumber, verticalNumber int8) (*ExtendedSpatialID, error) {
	return NewExtendedSpatialID(
		id.GetHorizontalZoomLevel()+horizonalNumber,
		id.GetX() << horizonalNumber,
		id.GetY() << horizonalNumber,
		id.GetVerticalZoomLevel()+verticalNumber,
		id.GetZ() << verticalNumber,
	)
}

func (id ExtendedSpatialID) NewMaxChild(horizonalNumber, verticalNumber int8) (*ExtendedSpatialID, error) {
	return NewExtendedSpatialID(
		id.GetHorizontalZoomLevel()+horizonalNumber,
		(id.GetX() << horizonalNumber) + (1 << horizonalNumber) - 1,
		(id.GetY() << horizonalNumber) + (1 << horizonalNumber) - 1,
		id.GetVerticalZoomLevel()+verticalNumber,
		(id.GetZ() << verticalNumber) + (1 << verticalNumber) - 1,
	)
}

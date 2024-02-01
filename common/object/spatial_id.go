package object

import (
	"math"
	"strconv"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/common/consts"
	"github.com/trajectoryjp/spatial_id_go/common/errors"
)

// ExtendedSpatialID 拡張空間IDクラス
type ExtendedSpatialID struct {
	hZoom int64 // 水平方向精度
	x     int64 // 経度ID
	y     int64 // 緯度ID
	vZoom int64 // 垂直方向精度
	z     int64 // 高さID
}

// NewExtendedSpatialID 拡張空間ID初期化関数
//
// 拡張空間IDのフォーマットに違反した引数を入力した場合エラーインスタンスが返却される。
//
// 引数：
//
//	extendedSpatialID：拡張空間ID
//
// 戻り値：
//
//	初期化したExtendedSpatialIDオブジェクト
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 引数にExtendedSpatialIDのフォーマットに違反した値が入力された場合
func NewExtendedSpatialID(extendedSpatialID string) (*ExtendedSpatialID, error) {
	s := &ExtendedSpatialID{}

	err := s.ResetExtendedSpatialID(extendedSpatialID)

	return s, err
}

// ResetExtendedSpatialID 拡張空間ID再設定関数
//
// 拡張空間IDのフォーマットに違反した引数を入力した場合エラーインスタンスが返却される。
//
// 引数：
//
//	extendedSpatialID：拡張空間ID
//
// 戻り値：
//
//	フィールド値を再設定したExtendedSpatialIDオブジェクト
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 引数にExtendedSpatialIDのフォーマットに違反した値が入力された場合
func (s *ExtendedSpatialID) ResetExtendedSpatialID(extendedSpatialID string) error {
	// 拡張空間IDを区切り文字で分割
	attr := strings.Split(extendedSpatialID, consts.SpatialIDDelimiter)

	if len(attr) != 5 {
		// 区切り文字数がフォーマットに従っていない場合エラーインスタンスを返却
		return errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 整数型変換後のフィールド値格納用
	convAttr := []int64{}

	for _, v := range attr {
		// フィールド値をint64型に返却
		i, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			// int64変換時にエラーが発生した場合エラーインスタンスを返却
			return errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}

		convAttr = append(convAttr, i)
	}

	// 経度、緯度、高さIDを設定
	s.SetX(convAttr[1])
	s.SetY(convAttr[2])
	s.SetZ(convAttr[4])

	// 水平方向精度、高さ方向精度を設定
	s.SetZoom(convAttr[0], convAttr[3])

	return nil
}

// SetX 位置設定関数
//
// ExtendedSpatialIDオブジェクトのxを引数の入力値に設定する。
//
// 引数：
//
//	x：経度ID
func (s *ExtendedSpatialID) SetX(x int64) {

	hZoom := s.HZoom()

	// インデックスの最大値を取得
	maxIndex := int64(math.Pow(2, float64(hZoom)) - 1)

	// 新インデックスを保存する
	newXIndex := x

	// 新インデックスが存在しているかチェックする。
	// x方向インデックスのチェック
	if newXIndex > maxIndex || newXIndex < 0 {
		// インデックスが負の場合は精度-2^精度%abs(index)が
		// インデックスの範囲を超えている場合はn周分を無視する
		for newXIndex < 0 {
			newXIndex += int64(math.Pow(2, float64(hZoom)))
		}
		newXIndex = int64(math.Mod(float64(newXIndex), math.Pow(2, float64(hZoom))))
	}

	s.x = newXIndex
}

// SetY 緯度ID設定関数
//
// ExtendedSpatialIDオブジェクトのyを引数の入力値に設定する。
//
// 引数：
//
//	y：緯度ID
func (s *ExtendedSpatialID) SetY(y int64) {

	hZoom := s.HZoom()

	// インデックスの最大値を取得
	maxIndex := int64(math.Pow(2, float64(hZoom)) - 1)

	// 新インデックスを計算する
	newYIndex := y

	// 新インデックスが存在しているかチェックする。
	// y方向インデックスのチェック
	if newYIndex > maxIndex || newYIndex < 0 {
		// インデックスが負の場合は精度-2^精度%abs(index)が
		// インデックスの範囲を超えている場合はn周分を無視する
		for newYIndex < 0 {
			newYIndex += int64(math.Pow(2, float64(hZoom)))
		}
		newYIndex = int64(math.Mod(float64(newYIndex), math.Pow(2, float64(hZoom))))
	}
	s.y = newYIndex
}

// SetZ 高さID設定関数
//
// ExtendedSpatialIDオブジェクトのzを引数の入力値に設定する。
//
// 引数：
//
//	z：高さID
func (s *ExtendedSpatialID) SetZ(z int64) {
	s.z = z
}

// SetZoom 精度設定関数
//
// ExtendedSpatialIDオブジェクトのhZoomを引数の入力値に設定する。
//
// 引数：
//
//	hZoom：水平方向精度
//	vZoom：垂直方向精度
func (s *ExtendedSpatialID) SetZoom(hZoom, vZoom int64) {
	s.hZoom = hZoom
	s.vZoom = vZoom
}

// X 経度ID取得関数
//
// ExtendedSpatialIDオブジェクトの経度IDを返却する。
//
// 戻り値：
//
//	経度ID
func (s ExtendedSpatialID) X() int64 {
	return s.x
}

// Y 緯度ID取得関数
//
// ExtendedSpatialIDオブジェクトの緯度IDを返却する。
//
// 戻り値：
//
//	緯度ID
func (s ExtendedSpatialID) Y() int64 {
	return s.y
}

// Z 高さID取得関数
//
// ExtendedSpatialIDオブジェクトの高さIDを返却する。
//
// 戻り値：
//
//	高さID
func (s ExtendedSpatialID) Z() int64 {
	return s.z
}

// HZoom 水平精度取得関数
//
// ExtendedSpatialIDオブジェクトのhZoomを返却する。
//
// 戻り値：
//
//	水平方向精度
func (s ExtendedSpatialID) HZoom() int64 {
	return s.hZoom
}

// VZoom 垂直精度取得関数
//
// ExtendedSpatialIDオブジェクトのvZoomを返却する。
//
// 戻り値：
//
//	垂直方向精度
func (s ExtendedSpatialID) VZoom() int64 {
	return s.vZoom
}

// FieldParams 拡張空間ID成分取得関数
//
// 拡張空間IDインスタンスの各フィールドに設定されたパラメータを返却する
//
// 戻り値：
//
//	拡張空間IDに含まれる水平、垂直精度、X, Y, Z成分を格納した以下の配列
//	 [水平精度, X成分, Y成分, 垂直精度, Z成分]
func (s *ExtendedSpatialID) FieldParams() []int64 {
	return []int64{s.hZoom, s.x, s.y, s.vZoom, s.z}
}

// ID 空間ID文字列返却関数
//
// 文字列に連結した空間IDを返却
//
// 戻り値：
//
//	拡張空間ID文字列
func (s ExtendedSpatialID) ID() string {
	var arr = []string{
		strconv.FormatInt(s.hZoom, 10),
		strconv.FormatInt(s.x, 10),
		strconv.FormatInt(s.y, 10),
		strconv.FormatInt(s.vZoom, 10),
		strconv.FormatInt(s.z, 10),
	}

	return strings.Join(arr, "/")
}

// Higher 最適化後拡張空間ID化関数
//
// 拡張空間IDインスタンスの最適化後となる拡張空間IDを返却する
//
// 引数：
//
//	hDiff：水平方向精度差分
//	vDiff：垂直方向精度差分
//
// 戻り値：
//
//	最適化後拡張空間ID化した拡張空間IDインスタンス
func (s ExtendedSpatialID) Higher(hDiff, vDiff int64) *ExtendedSpatialID {
	// 最適化後拡張空間ID精度
	var hZoom = s.hZoom - hDiff
	var vZoom = s.vZoom - vDiff

	var hDiv = int64(math.Pow(2, float64(hDiff)))
	var vDiv = int64(math.Pow(2, float64(vDiff)))
	var x = s.x / hDiv
	var y = s.y / hDiv
	var z = s.z / vDiv

	return &ExtendedSpatialID{
		hZoom: hZoom,
		vZoom: vZoom,
		x:     x,
		y:     y,
		z:     z,
	}
}

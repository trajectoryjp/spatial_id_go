// 拡張空間IDパッケージ
package object

import (
	"fmt"
	"github.com/trajectoryjp/spatial_id_go/v4/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
)

// FromExtendedSpatialIDToQuadkeyAndAltitudekey 拡張空間IDから変換したquadkeyとaltitudekeyの組み合わせを管理する構造体
type FromExtendedSpatialIDToQuadkeyAndAltitudekey struct {
	// quadkeyの精度
	quadkeyZoom int64
	// innerIDList [[quadkey, altitudekey]...]
	innerIDList [][2]int64
	// altitudekeyの精度
	altitudekeyZoom int64
	// altitudekeyの高さが1mとなるズームレベル
	zBaseExponent int64
	// ズームレベルがzBaseExponentで高度0mにおけるaltitudekey
	zBaseOffset int64
}

// NewFromExtendedSpatialIDToQuadkeyAndVerticalID FromExtendedSpatialIDToQuadkeyAndVerticalID初期化関数
//
// input 引数：
//
//	quadkeyZoom： quadkeyの精度
//	innerIDList： [[quadkey, altitudekey]...]
//	altitudekeyZoom： altitudekeyの精度
//	zBaseExponent： altitudekeyの高さが1mとなるズームレベル
//	zBaseOffset： ズームレベルがzBaseExponentで高度0mにおけるaltitudekey
//
// output 戻り値：
//
//	初期化したFromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクト
func NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(quadkeyZoom int64, innerIDList [][2]int64, altitudekeyZoom int64, zBaseExponent int64, zBaseOffset int64) *FromExtendedSpatialIDToQuadkeyAndAltitudekey {
	a := &FromExtendedSpatialIDToQuadkeyAndAltitudekey{}
	a.SetQuadkeyZoom(quadkeyZoom)
	a.SetInnerIDList(innerIDList)
	a.SetAltitudekeyZoom(altitudekeyZoom)
	a.SetZBaseExponent(zBaseExponent)
	a.SetZBaseOffset(zBaseOffset)
	return a
}

// SetQuadkeyZoom quadkeyの精度設定関数
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトのquadkeyZoomを引数の入力値に設定する。
//
// input 引数：
//
//	quadkeyZoom：quadkeyの精度
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) SetQuadkeyZoom(quadkeyZoom int64) {
	a.quadkeyZoom = quadkeyZoom
}

// SetInnerIDList innerIDList設定関数
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトのinnerIDListを引数の入力値に設定する。
//
// input 引数：
//
//	innerIDList：innerIDListのスライス
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) SetInnerIDList(innerIDList [][2]int64) {
	a.innerIDList = innerIDList
}

// SetAltitudeZoom altitudekeyの精度設定関数
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトのaltitudeKeyZoomを引数の入力値に設定する。
//
// input 引数：
//
//	altitudekeyZoom：altitudekeyの精度
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) SetAltitudekeyZoom(altitudekeyZoom int64) {
	a.altitudekeyZoom = altitudekeyZoom
}

// SetZBaseExponent
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトのzBaseExponentを引数の入力値に設定する。
//
// input 引数：
//
//	zBaseExponent： altitudekeyの高さが1mとなるズームレベル
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) SetZBaseExponent(zBaseExponent int64) {
	a.zBaseExponent = zBaseExponent
}

// SetZBaseOffset
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトのzBaseOffsetを引数の入力値に設定する。
//
// input 引数：
//
//	zBaseOffset： ズームレベルがzBaseExponentで高度0mにおけるaltitudekey
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) SetZBaseOffset(zBaseOffset int64) {
	a.zBaseOffset = zBaseOffset
}

// QuadkeyZoom quadkeyZoom設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているquadkeyZoomの値を取得する。
//
// output 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているquadkeyZoomの値
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) QuadkeyZoom() int64 {
	return a.quadkeyZoom
}

// InnerIDList innerIDList設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているinnerIDListの値を取得する。
//
// output 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているinnerIDListの値
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) InnerIDList() [][2]int64 {
	return a.innerIDList
}

// AltitudekeyZoom altitudekeyZoom設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているaltitudekeyZoomの値を取得する。
//
// output 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているaltitudekeyZoomの値
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) AltitudekeyZoom() int64 {
	return a.altitudekeyZoom
}

// ZBaseExponent zBaseExponent設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているzBaseExponentの値を取得する。
//
// output 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているzBaseExponentの値
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) ZBaseExponent() int64 {
	return a.zBaseExponent
}

// ZBaseOffset zBaseOffset設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているzBaseOffsetの値を取得する。
//
// output 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndAltitudekeyオブジェクトに設定されているzBaseOffsetの値
func (a *FromExtendedSpatialIDToQuadkeyAndAltitudekey) ZBaseOffset() int64 {
	return a.zBaseOffset
}

// FromExtendedSpatialIDToQuadkeyAndVerticalID 拡張空間IDから変換したquadkeyと高さのIDの組み合わせを管理する構造体
type FromExtendedSpatialIDToQuadkeyAndVerticalID struct {
	// quadkeyの精度
	quadkeyZoom int64
	// innerIDList [[quadkey,vIndex]...]
	innerIDList [][2]int64
	// 高さ方向の精度
	vZoom int64
	// 最高高度
	maxHeight float64
	// 最低高度
	minHeight float64
}

// NewFromExtendedSpatialIDToQuadkeyAndVerticalID FromExtendedSpatialIDToQuadkeyAndVerticalID初期化関数
//
// 引数：
//
//	quadkeyZoom： 水平精度
//	quadkey： quadkeyのスライス
//	vZoom： 垂直精度
//	verticalID： 高さインデックスのスライス
//	maxHeight : 最高高度
//	minHeight : 最低高度
//
// 戻り値：
//
//	初期化したFromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクト
func NewFromExtendedSpatialIDToQuadkeyAndVerticalID(quadkeyZoom int64, innerIDList [][2]int64, vZoom int64, maxHeight float64, minHeight float64) *FromExtendedSpatialIDToQuadkeyAndVerticalID {
	s := &FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	// quadkeyの精度
	s.SetQuadkeyZoom(quadkeyZoom)
	// 内部形式IDの配列
	s.SetInnerIDList(innerIDList)
	// 高さ方向の精度
	s.SetVerticalZoom(vZoom)
	// 最高高度
	s.SetMaxHeight(maxHeight)
	// 最低高度
	s.SetMinHeight(minHeight)
	return s
}

// SetQuadkeyZoom 水平精度設定関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトのquadkeyZoomを引数の入力値に設定する。
//
// 引数：
//
//	quadkeyZoom：quadkeyの精度
func (s *FromExtendedSpatialIDToQuadkeyAndVerticalID) SetQuadkeyZoom(quadkeyZoom int64) {
	// quadkeyを引数入力値に変更
	s.quadkeyZoom = quadkeyZoom
}

// SetInnerIDList innerIDList設定関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトのinnerIDListを引数の入力値に設定する。
//
// 引数：
//
//	innerIDList：innerIDListのスライス
func (s *FromExtendedSpatialIDToQuadkeyAndVerticalID) SetInnerIDList(innerIDList [][2]int64) {
	// innerIDListを引数入力値に変更
	s.innerIDList = innerIDList
}

// SetVerticalZoom 垂直精度設定関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトのVerticalZoomを引数の入力値に設定する。
//
// 引数：
//
//	vZoom：垂直精度
func (s *FromExtendedSpatialIDToQuadkeyAndVerticalID) SetVerticalZoom(vZoom int64) {
	// vZoomを引数入力値に変更
	s.vZoom = vZoom
}

// SetMaxHeight 最高高度設定関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトのMaxHeightを引数の入力値に設定する。
//
// 引数：
//
//	maxHeight：最高高度
func (s *FromExtendedSpatialIDToQuadkeyAndVerticalID) SetMaxHeight(maxHeight float64) {
	// maxHeightを引数入力値に変更
	s.maxHeight = maxHeight
}

// SetMinHeight 最低高度設定関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトのMinHeightを引数の入力値に設定する。
//
// 引数：
//
//	minHeight：最低高度
func (s *FromExtendedSpatialIDToQuadkeyAndVerticalID) SetMinHeight(minHeight float64) {
	// minHeightを引数入力値に変更
	s.minHeight = minHeight
}

// QuadkeyZoom QuadkeyZoom設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているquadkeyZoomの値を取得する。
//
// 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているquadkeyZoomの値
func (s FromExtendedSpatialIDToQuadkeyAndVerticalID) QuadkeyZoom() int64 {
	return s.quadkeyZoom
}

// InnerIDList InnerIDList設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているInnerIDListの値を取得する。
//
// 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているInnerIDListの値
func (s FromExtendedSpatialIDToQuadkeyAndVerticalID) InnerIDList() [][2]int64 {
	return s.innerIDList
}

// VerticalZoom VerticalZoom設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているvZoomの値を取得する。
//
// 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているvZoomの値
func (s FromExtendedSpatialIDToQuadkeyAndVerticalID) VerticalZoom() int64 {
	return s.vZoom
}

// MaxHeight MaxHeight設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているmaxHeightの値を取得する。
//
// 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているmaxHeightの値
func (s FromExtendedSpatialIDToQuadkeyAndVerticalID) MaxHeight() float64 {
	return s.maxHeight
}

// MinHeight MinHeight設定値取得関数
//
// FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているminHeightの値を取得する。
//
// 戻り値：
//
//	FromExtendedSpatialIDToQuadkeyAndVerticalIDオブジェクトに設定されているminHeightの値
func (s FromExtendedSpatialIDToQuadkeyAndVerticalID) MinHeight() float64 {
	return s.minHeight
}

// QuadkeyAndVerticalID quadkeyと高さ方向のIDの組み合わせを管理する構造体
type QuadkeyAndVerticalID struct {
	// 水平精度
	quadkeyZoom int64
	// quadkey
	quadkey int64
	// 垂直精度
	vZoom int64
	// 高さ方向のID
	vIndex int64
	// 最高高度
	maxHeight float64
	// 最低高度
	minHeight float64
}

// NewQuadkeyAndVerticalID QuadkeyAndVerticalID初期化関数
//
// quadkeyに数値以外が含まれていた場合エラーとなる。
//
// 引数：
//
//	quadkeyZoom： quadkeyZoom
//	quadkey： quadkey
//	vZoom： 垂直精度
//	vIndex： 高さ方向のインデックス
//	maxHeight : 最高高度
//	minHeight : 最低高度
//
// 戻り値：
//
//	初期化したQuadkeyAndVerticalIDオブジェクト
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 quadkey不正：quadkey文字列に数値以外が含まれていた場合
func NewQuadkeyAndVerticalID(quadkeyZoom int64, quadkey int64, vZoom int64, vIndex int64, maxHeight float64, minHeight float64) *QuadkeyAndVerticalID {
	qh := &QuadkeyAndVerticalID{}

	// 水平精度設定
	qh.SetQuadkeyZoom(quadkeyZoom)

	// quadkey設定
	qh.SetQuadkey(quadkey)

	// 垂直精度設定
	qh.SetVZoom(vZoom)

	// 高さのインデックス設定
	qh.SetVIndex(vIndex)

	// 最高高度設定
	qh.SetMaxHeight(maxHeight)
	// 最低高度設定
	qh.SetMinHeight(minHeight)

	return qh
}

// SetQuadkeyZoom quadkeyZoom設定関数
//
// QuadkeyAndVerticalIDオブジェクトのquadkeyZoomを引数の入力値に設定する。
// 入力値の文字列に数字以外の文字が含まれていた場合エラーとなる。
//
// 引数：
//
//	quadkeyZoom：quadkeyZoomの数値
func (qh *QuadkeyAndVerticalID) SetQuadkeyZoom(quadkeyZoom int64) {

	// quadkeyを引数入力値に変更
	qh.quadkeyZoom = quadkeyZoom

}

// SetQuadkey quadkey設定関数
//
// QuadkeyAndVerticalIDオブジェクトのquadkeyを引数の入力値に設定する。
// 入力値の文字列に数字以外の文字が含まれていた場合エラーとなる。
//
// 引数：
//
//	quadkey：quadkeyの数値
func (qh *QuadkeyAndVerticalID) SetQuadkey(quadkey int64) {

	// quadkeyを引数入力値に変更
	qh.quadkey = quadkey

}

// SetVZoom 垂直精度設定関数
//
// QuadkeyAndVerticalIDオブジェクトのvZoomを引数の入力値に設定する。
//
// 引数：
//
//	VerticalID：高さのID
func (qh *QuadkeyAndVerticalID) SetVZoom(vZoom int64) {

	qh.vZoom = vZoom
}

// SetVIndex 高さのインデックス設定関数
//
// QuadkeyAndVerticalIDオブジェクトのvIndexを引数の入力値に設定する。
//
// 引数：
//
//	VerticalID：高さのID
func (qh *QuadkeyAndVerticalID) SetVIndex(vIndex int64) {

	qh.vIndex = vIndex
}

// SetMaxHeight 最高高度設定関数
//
// QuadkeyAndVerticalIDオブジェクトのmaxHeightを引数の入力値に設定する。
//
// 引数：
//
//	maxHeight：最高高度
func (qh *QuadkeyAndVerticalID) SetMaxHeight(maxHeight float64) {
	qh.maxHeight = maxHeight
}

// SetMinHeight 最低高度設定関数
//
// QuadkeyAndVerticalIDオブジェクトのminHeightを引数の入力値に設定する。
//
// 引数：
//
//	minHeight：最高高度
func (qh *QuadkeyAndVerticalID) SetMinHeight(minHeight float64) {
	qh.minHeight = minHeight
}

// QuadkeyZoom quadkeyZoom設定値取得関数
//
// QuadkeyAndVerticalIDオブジェクトに設定されているquadkeyZoomの値を取得する。
//
// 戻り値：
//
//	QuadkeyAndVerticalIDオブジェクトに設定されているquadkeyZoomの値
func (qh QuadkeyAndVerticalID) QuadkeyZoom() int64 {
	return qh.quadkeyZoom
}

// Quadkey quadkey設定値取得関数
//
// QuadkeyAndVerticalIDオブジェクトに設定されているquadkeyの値を取得する。
//
// 戻り値：
//
//	QuadkeyAndVerticalIDオブジェクトに設定されているquadkeyの値
func (qh QuadkeyAndVerticalID) Quadkey() int64 {
	return qh.quadkey
}

// VZoom VZoom設定値取得関数
//
// QuadkeyAndVerticalIDオブジェクトに設定されているvZoomの値を取得する。
//
// 戻り値：
//
//	QuadkeyAndVerticalIDオブジェクトに設定されているvZoomの値
func (qh QuadkeyAndVerticalID) VZoom() int64 {
	return qh.vZoom
}

// VIndex VIndex設定値取得関数
//
// QuadkeyAndVerticalIDオブジェクトに設定されているvIndexの値を取得する。
//
// 戻り値：
//
//	QuadkeyAndVerticalIDオブジェクトに設定されているvIndexの値
func (qh QuadkeyAndVerticalID) VIndex() int64 {
	return qh.vIndex
}

// MaxHeight 最高高度設定値取得関数
//
// QuadkeyAndVerticalIDオブジェクトに設定されているmaxHeightの値を取得する。
//
// 戻り値：
//
//	QuadkeyAndVerticalIDオブジェクトに設定されているmaxHeightの値
func (qh QuadkeyAndVerticalID) MaxHeight() float64 {
	return qh.maxHeight
}

// MinHeight 最低高度設定値取得関数
//
// QuadkeyAndVerticalIDオブジェクトに設定されているminHeightの値を取得する。
//
// 戻り値：
//
//	QuadkeyAndVerticalIDオブジェクトに設定されているminHeightの値
func (qh QuadkeyAndVerticalID) MinHeight() float64 {
	return qh.minHeight
}

// TileXYZ 水平方向TileKey(x,y)と垂直方向TileKey(z)の組み合わせを管理する構造体
type TileXYZ struct {
	// 水平精度 0-(consts.MaxTileXYZZoom)の非負整数
	hZoom uint16
	// x 水平方向key x軸
	x int64
	// y 水平方向key y軸
	y int64
	// 垂直精度 0-(consts.MaxTileXYZZoom)の非負整数
	vZoom uint16
	// 垂直方向key
	z int64
}

// NewTileXYZ TileXYZ初期化関数
//
// quadkeyに数値以外が含まれていた場合エラーとなる。
//
// 引数：
//
//	hZoom： 水平ズームレベル
//	x： 水平方向xインデックス
//	y： 水平方向yインデックス
//	vZoom： 垂直ズームレベル
//	z： 高さ方向のインデックス
//
// 戻り値：
//
//	初期化したTileXYZオブジェクト
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 ズームレベル不正：hZoomまたはvZoomに0-(consts.MaxTileXYZZoom)以外の数値が含まれていた場合
func NewTileXYZ(hZoom uint16, x int64, y int64, vZoom uint16, z int64) (*TileXYZ, error) {
	tile := &TileXYZ{}

	// 水平精度設定
	err := tile.SetHZoom(hZoom)
	if err != nil {
		return nil, err
	}

	// X設定
	tile.SetX(x)
	// Y設定
	tile.SetY(y)

	// 垂直精度設定
	err = tile.SetVZoom(vZoom)
	if err != nil {
		return nil, err
	}

	// 高さのインデックス設定
	tile.SetZ(z)

	return tile, nil
}

// SetHZoom 水平精度設定関数
//
// TileXYZオブジェクトのhZoomを引数の入力値に設定する。
//
// input 引数：
//
//	hZoom：水平精度
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 ズームレベル不正：hZoomに0-(consts.MaxTileXYZZoom)以外の数値が含まれていた場合
func (a *TileXYZ) SetHZoom(hZoom uint16) error {
	if !(hZoom <= consts.MaxTileXYZZoom) {
		return errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("hZoom must be in 0-%v, but got %v", consts.MaxTileXYZZoom, hZoom))
	}
	a.hZoom = hZoom
	return nil
}

// SetX xインデックス設定関数
//
// TileXYZオブジェクトのxを引数の入力値に設定する。
//
// input 引数：
//
//	x：xインデックス値
func (a *TileXYZ) SetX(x int64) {
	a.x = x
}

// SetY yインデックス設定関数
//
// TileXYZオブジェクトのyを引数の入力値に設定する。
//
// input 引数：
//
//	y：yインデックス値
func (a *TileXYZ) SetY(y int64) {
	a.y = y
}

// SetVZoom 垂直精度設定関数
//
// TileXYZオブジェクトのvZoomを引数の入力値に設定する。
//
// input 引数：
//
//	vZoom：垂直精度
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 ズームレベル不正：はvZoomに0-35以外の数値が含まれていた場合
func (a *TileXYZ) SetVZoom(vZoom uint16) error {
	if !(vZoom <= consts.MaxTileXYZZoom) {
		return errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("vZoom must be in 0-%v, but got %v", consts.MaxTileXYZZoom, vZoom))
	}
	a.vZoom = vZoom
	return nil
}

// SetZ zインデックス設定関数
//
// TileXYZオブジェクトのzを引数の入力値に設定する。
//
// input 引数：
//
//	z：zインデックス値
func (a *TileXYZ) SetZ(z int64) {
	a.z = z
}

// HZoom HZoom設定値取得関数
//
// TileXYZオブジェクトに設定されているhZoomの値を取得する。
//
// output 戻り値：
//
//	TileXYZオブジェクトに設定されているHZoomの値
func (a *TileXYZ) HZoom() uint16 {
	return a.hZoom
}

// X x設定値取得関数
//
// TileXYZオブジェクトに設定されているxの値を取得する。
//
// output 戻り値：
//
//	TileXYZオブジェクトに設定されているxの値
func (a *TileXYZ) X() int64 {
	return a.x
}

// Y y設定値取得関数
//
// TileXYZオブジェクトに設定されているyの値を取得する。
//
// output 戻り値：
//
//	TileXYZオブジェクトに設定されているyの値
func (a *TileXYZ) Y() int64 {
	return a.y
}

// VZoom vZoom設定値取得関数
//
// TileXYZオブジェクトに設定されているvZoomの値を取得する。
//
// output 戻り値：
//
//	TileXYZオブジェクトに設定されているvZoomの値
func (a *TileXYZ) VZoom() uint16 {
	return a.vZoom
}

// Z z設定値取得関数
//
// TileXYZオブジェクトに設定されているzの値を取得する。
//
// output 戻り値：
//
//	TileXYZオブジェクトに設定されているzの値
func (a *TileXYZ) Z() int64 {
	return a.z
}

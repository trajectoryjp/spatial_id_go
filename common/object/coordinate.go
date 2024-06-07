// Package object ライブラリ内で共通的に使用するオブジェクトを管理するパッケージ。
package object

import (
	"math"

	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
)

// Point 地理座標用の構造体
type Point struct {
	lon float64 // 経度
	lat float64 // 緯度
	alt float64 // 高さ
}

// NewPoint Point初期化関数
//
// 経度入力値の絶対値が180を超える場合エラーとなる。
//
// また、緯度入力値の絶対値が85.0511287798を超える場合エラーとなる。
// 入力値が小数点11桁以上の場合、10桁となるよう切り捨てる
//
// 引数：
//
//	lon：経度
//	lat：緯度
//	alt：高さ
//
// 戻り値：
//
//	初期化したPointオブジェクト
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 経度入力値超過：絶対値が180を超える経度が入力された場合
//	 緯度入力値超過：絶対値が85.0511287798を超える緯度が入力された場合
func NewPoint(lon float64, lat float64, alt float64) (*Point, error) {
	s := &Point{}

	// 経度設定
	err := s.SetLon(lon)

	if err != nil {
		return s, err
	}

	// 緯度設定
	err = s.SetLat(lat)

	if err != nil {
		return s, err
	}

	// 高さ設定
	s.SetAlt(alt)

	return s, err
}

// SetLon 経度設定関数
//
// PointオブジェクトのLonを引数の入力値に設定する。
// 入力値の絶対値が180を超える場合エラーとなる。
//
// 引数：
//
//	lon：経度
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 経度入力値超過：絶対値が180を超える経度が入力された場合
func (p *Point) SetLon(lon float64) error {
	if math.Abs(lon) > 180 {
		// 経度入力値が制限値を超える場合、エラーインスタンスを返却
		return errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// 経度を引数入力値に変更
	p.lon = lon

	return nil
}

// SetLat 緯度設定関数
//
// PointオブジェクトのLatを引数の入力値に設定する。
// 入力値の絶対値が85.0511287798を超える場合エラーとなる。
// 入力値が小数点11桁以上の場合、10桁となるよう切り捨てる
//
// 引数：
//
//	lat：緯度
//
// 戻り値(エラー)：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 緯度入力値超過：絶対値が85.0511287798を超える緯度が入力された場合
func (p *Point) SetLat(lat float64) error {
	// 小数点11位以下は切り捨てる
	if lat > 0 {
		lat = math.Floor(lat*math.Pow(10, 10.0)) / math.Pow(10, 10.0)
	} else {
		lat = math.Ceil(lat*math.Pow(10, 10.0)) / math.Pow(10, 10.0)
	}

	if math.Abs(lat) > 85.0511287798 {
		// 緯度入力値が制限値を超える場合、エラーインスタンスを返却
		return errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	p.lat = lat
	return nil
}

// SetAlt 高さ設定関数
//
// Pointオブジェクトのaltを引数の入力値に設定する。
//
// 引数：
//
//	alt：高さ
func (p *Point) SetAlt(alt float64) {
	p.alt = alt
}

// Lon 経度設定値取得関数
//
// Pointオブジェクトに設定されているlonの値を取得する。
//
// 戻り値：
//
//	Pointオブジェクトに設定されているlonの値
func (p Point) Lon() float64 {
	return p.lon
}

// Lat 緯度設定値取得関数
//
// Pointオブジェクトに設定されているlatの値を取得する。
//
// 戻り値：
//
//	Pointオブジェクトに設定されているlatの値
func (p Point) Lat() float64 {
	return p.lat
}

// Alt 高さ設定値取得関数
//
// Pointオブジェクトに設定されているaltの値を取得する。
//
// 戻り値：
//
//	Pointオブジェクトに設定されているaltの値
func (p Point) Alt() float64 {
	return p.alt
}

// ProjectedPoint 投影座標用の構造体
type ProjectedPoint struct {
	X   float64 // X座標
	Y   float64 // Y座標
	Alt float64 // 高さ
}

// VerticalPoint 垂直方向位置用の構造体
type VerticalPoint struct {
	Alt        float64 // 高さ
	Resolution float64 // 分解能
}

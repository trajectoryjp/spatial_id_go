// Package spatial 空間座標操作パッケージ
package spatial

import (
	"github.com/trajectoryjp/spatial_id_go/v2/common"
	"github.com/trajectoryjp/spatial_id_go/v2/common/errors"
)

// UniqueAppend 点のユニークを保持した配列追加
//
// 点のスライスへユニークを保持して点を追加
//
// 引数：
//
//	points ：点のスライス
//	addPoint: 追加する点
//	epsilon： ユニーク判定時の小数点誤差
//
// 戻り値：
//
//	点のスライス
func UniqueAppend(points []*Point3, addPoint *Point3, epsilon float64) []*Point3 {
	// 追加点のユニークであるかの判定
	isUnique := true

	for _, point := range points {
		// スライス内に既に同一点がある場合
		if point.IsClose(*addPoint, epsilon) {
			isUnique = false
			break
		}

	}

	// 追加点のユニークである場合はスライスに追加
	if isUnique {
		points = append(points, addPoint)
	}

	return points
}

// MaxPoint 最大点取得関数
//
// 引数に入力されたPoint3スライスの最大値を返却する
//
// 引数：
//
//	points： Point3スライス
//	vec： Point3の比較値を出すためのベクトル
//
// 戻り値：
//
//	Point3スライスの最大値
//
// 戻り値（エラー）：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力不正：入力スライスが空の場合。
func MaxPoint(points []*Point3, vec Vector3) (*Point3, error) {
	// スライスが空の場合はエラー
	if len(points) == 0 {
		return new(Point3), errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	max := points[0]
	maxValue := Vector3(*max).Dot(vec)
	for _, point := range points {
		value := Vector3(*point).Dot(vec)
		if maxValue < value {
			max = point
			maxValue = value
		}
	}

	return max, nil
}

// MinPoint 最小点取得関数
//
// 引数に入力されたPoint3スライスの最小値を返却する
//
// 引数：
//
//	points： Point3スライス
//	vec： Point3の比較値を出すためのベクトル
//
// 戻り値：
//
//	Point3スライスの最小値
//
// 戻り値（エラー）：
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力不正：入力スライスが空の場合。
func MinPoint(points []*Point3, vec Vector3) (*Point3, error) {
	// スライスが空の場合はエラー
	if len(points) == 0 {
		return new(Point3), errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	min := points[0]
	minValue := Vector3(*min).Dot(vec)
	for _, point := range points {
		value := Vector3(*point).Dot(vec)
		if minValue > value {
			min = point
			minValue = value
		}
	}

	return min, nil
}

// Point3 点座標構造体
type Point3 struct {
	X, Y, Z float64 // X・Y・Z座標
}

// IsClose 点が同一であるかの確認
//
// 2点が浮動小数点誤差epsilonだけ許容して一致しているかを返却
//
// 引数：
//
//	q ：同一確認する点
//
// 戻り値：
//
//	true：2点が同一 false：2点が同一でない
func (p Point3) IsClose(q Point3, epsilon float64) bool {
	return common.AlmostEqual(p.X, q.X, epsilon) && common.AlmostEqual(p.Y, q.Y, epsilon) && common.AlmostEqual(p.Z, q.Z, epsilon)
}

// Translate 点の平行移動
//
// 点をベクトルに沿って平行移動する
//
// 引数：
//
//	a ：平行移動するベクトル
//
// 戻り値：
//
//	平行移動した点
func (p Point3) Translate(a Vector3) Point3 {
	return Point3(Vector3(p).Add(a))
}

// DistancePoint 2点間の距離
//
// 2点間の距離を返却する
//
// 引数：
//
//	q ：間の距離を求める点
//
// 戻り値：
//
//	2点間の距離
func (p Point3) DistancePoint(q Point3) float64 {
	return NewVectorFromPoints(p, q).Norm()
}

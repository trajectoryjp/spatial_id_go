// Package spatial 空間座標操作パッケージ
package spatial

import (
	"math"
	"gonum.org/v1/gonum/spatial/r3"
)

// Vector3 ベクトル構造体
type Vector3 r3.Vec

// NewVectorFromPoints 点からのベクトル作成
//
// 2点からベクトル初期化する
//
// 引数：
//  p ：ベクトルの始点
//  q ：ベクトルの終点
//
// 戻り値：
//  始点から終点へのベクトル
func NewVectorFromPoints(p, q Point3) Vector3 {
	return Vector3(r3.Sub(r3.Vec(q), r3.Vec(p)))
}

// Add ベクトルの加算
//
// ベクトルを成分ごとに加算し、a + b を返却する
//
// 引数：
//  b ：加算するベクトル
//
// 戻り値：
//  加算したベクトル
func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3(r3.Add(r3.Vec(a), r3.Vec(b)))
}

// Sub ベクトルの減算
//
// ベクトルを成分ごとに減算し、a - b を返却する
//
// 引数：
//  b ：減算するベクトル
//
// 戻り値：
//  減算したベクトル
func (a Vector3) Sub(b Vector3) Vector3 {
	return Vector3(r3.Sub(r3.Vec(a), r3.Vec(b)))
}

// Scale ベクトルのスカラー倍
//
// ベクトルを成分ごとにスカラー倍し、f * a を返却する
//
// 引数：
//  f ：乗算するスカラー値
//
// 戻り値：
//  スカラー倍したベクトル
func (a Vector3) Scale(f float64) Vector3 {
	return Vector3(r3.Scale(f, r3.Vec(a)))
}

// Dot ベクトルの内積
//
// ベクトルの内積を算出し、a・b を返却する
//
// 引数：
//  b ：内積として掛けるベクトル
//
// 戻り値：
//  ベクトル内積
func (a Vector3) Dot(b Vector3) float64 {
	return r3.Dot(r3.Vec(a), r3.Vec(b))
}

// Cross ベクトルの外積
//
// ベクトルの外積を算出し、a×b を返却する
//
// 引数：
//  b ：外積として掛けるベクトル
//
// 戻り値：
//  ベクトル外積
func (a Vector3) Cross(b Vector3) Vector3 {
	return Vector3(r3.Cross(r3.Vec(a), r3.Vec(b)))
}

// Norm ベクトルのユークリッドノルム
//
// ベクトルのユークリッドノルムを算出し、|a| を返却する
//
// 戻り値：
//  ベクトルのユークリッドノルム
func (a Vector3) Norm() float64 {
	return r3.Norm(r3.Vec(a))
}

// L1Norm ベクトルのL1ノルム
//
// ベクトルのL1ノルムを算出し、返却する
//
// 戻り値：
//  ベクトルのL1ノルム
func (a Vector3) L1Norm() float64 {
	return math.Abs(a.X) + math.Abs(a.Y) + math.Abs(a.Z)
}

// Unit ベクトルの正規化
//
// ベクトルを正規化し、単位ベクトル a/|a| を返却する
//
// 戻り値：
//  ベクトルの単位ベクトル
func (a Vector3) Unit() Vector3 {
	return Vector3(r3.Unit(r3.Vec(a)))
}

// Cos ベクトル間の余弦(cos)
//
// ベクトル間の余弦(cos)を返却する
//
// 引数：
//  b ：間の角度を求めるベクトル
//
// 戻り値：
//  ベクトル間の余弦(cos)
func (a Vector3) Cos(b Vector3) float64 {
	return r3.Cos(r3.Vec(a), r3.Vec(b))
}

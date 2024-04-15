// Package spatial 空間座標操作パッケージ
package spatial

import (
	"math"

	"github.com/trajectoryjp/spatial_id_go/v3/common/consts"
)

// Quat 四元数の構造体
type Quat struct {
	W, X, Y, Z float64
}

// 四元数の計算結果確認用メソッド
// 実装では使用しないためコメントアウト
// // Mul 四元数の積
// //
// // 四元数の積を算出して返却する。
// //
// // 引数：
// //  b： 掛け合わせる四元数
// //
// // 戻り値：
// //  四元数の積
// func (a Quat) Mul(b Quat) Quat {
// 	return Quat{
// 		(b.W * a.W) - (b.X * a.X) - (b.Y * a.Y) - (b.Z * a.Z),
// 		(b.X * a.W) + (b.W * a.X) - (b.Z * a.Y) + (b.Y * a.Z),
// 		(b.Y * a.W) + (b.Z * a.X) + (b.W * a.Y) - (b.X * a.Z),
// 		(b.Z * a.W) - (b.Y * a.X) + (b.X * a.Y) + (b.W * a.Z),
// 	}
// }

// // Conjugate 四元数の共役
// //
// // 共役四元数を算出して返却する。
// //
// // 戻り値：
// //  共役四元数
// func (a Quat) Conjugate() Quat {
// 	return Quat{
// 		a.W,
// 		-a.X,
// 		-a.Y,
// 		-a.Z,
// 	}
// }

// // TransformVec3 四元数でのベクトル変換
// //
// // 入力のベクトルを四元数で変換した値を返却する。
// //
// // 引数：
// //  v： 入力ベクトル
// //
// // 戻り値：
// //  変換後のベクトル
// func (a Quat) TransformVec3(v Vector3) Vector3 {
// 	vecQuat := Quat{0.0, v.X, v.Y, v.Z}
// 	conjugate := a.Conjugate()
// 	vecQuat = conjugate.Mul(vecQuat).Mul(a)
// 	return Vector3{vecQuat.X, vecQuat.Y, vecQuat.Z}
// }

// RotateBetweenVector 2ベクトル間の四元数算出処理
//
// 2ベクトル間の四元数を算出して返却する。
//
// 引数：
//
//	start： 開始ベクトル
//	end： 終了ベクトル
//
// 戻り値：
//
//	2ベクトル間の四元数
func RotateBetweenVector(start, end Vector3) Quat {
	// ベクトルを正規化
	startUnit := start.Unit()
	endUnit := end.Unit()
	// ベクトル間のcosを算出
	cos := startUnit.Cos(endUnit)

	// 法線ベクトルを回転軸ベクトルとする
	axis := startUnit.Cross(endUnit)

	// ベクトルが正反対方向を向いている場合
	if cos+1 < consts.Minima {
		// 垂直になるベクトルを回転軸ベクトルとする
		axis = startUnit.Cross(Vector3{0.0, 0.0, 1.0})
		if axis.Norm() < consts.Minima {
			axis = start.Cross(Vector3{1.0, 0.0, 0.0})
		}

		return QuatFromAxisAngle(axis, math.Pi)
	}

	s := math.Sqrt(2 * (1 + cos))
	inv := 1 / s
	return Quat{
		s * 0.5,
		axis.X * inv,
		axis.Y * inv,
		axis.Z * inv,
	}
}

// QuatFromAxisAngle 回転ベクトルの四元数算出処理
//
// 回転ベクトルの四元数を算出して返却する。
//
// 引数：
//
//	axis： 回転軸ベクトル
//	angle： 回転角度(ラジアン)
//
// 戻り値：
//
//	回転ベクトルの四元数
func QuatFromAxisAngle(axis Vector3, angle float64) Quat {
	// 回転軸ベクトルを正規化
	axis = axis.Unit()
	// 半角のsinを算出
	sinHalfAngle := math.Sin(angle * 0.5)
	return Quat{
		math.Cos(angle * 0.5),
		axis.X * sinHalfAngle,
		axis.Y * sinHalfAngle,
		axis.Z * sinHalfAngle,
	}
}

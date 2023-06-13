// Package physics 物理オブジェクト操作パッケージ
package physics

import (
	"github.com/azul3d/engine/native/ode"
	"github.com/trajectoryjp/spatial_id_go/src/spatial-id/common/spatial"
)

// CapsulePhysics カプセル用の物理オブジェクト構造体
type CapsulePhysics struct {
	BasePhysics // 基底物理オブジェクト構造体の埋め込み
}

// NewCapsulePhysics カプセル用の物理オブジェクト構造体コンストラクタ
//
// カプセル用の物理オブジェクト構造体の初期化を行う
//
// 引数：
//
//	radius：カプセルの半径
//	start ：カプセルの始点
//	end   ：カプセルの終点
//
// 戻り値：
//
//	カプセル用の物理オブジェクト構造体
func NewCapsulePhysics(radius float64, start spatial.Point3, end spatial.Point3) *CapsulePhysics {

	// カプセルの軸
	axis := spatial.NewVectorFromPoints(start, end)
	// カプセルの軸の高さ
	length := axis.Norm()
	// 重心
	center := spatial.NewLineFromPoints(start, end).ToPoint(0.5)

	// 基底物理オブジェクト構造体
	b := NewBasePhysics()
	// カプセル(剛体)
	body := b.world.NewBody()
	// 座標設定
	body.SetPosition(ode.NewVector3(center.X, center.Y, center.Z))
	// 傾き設定
	quat := spatial.RotateBetweenVector(
		spatial.Vector3{0.0, 0.0, 1.0},
		axis,
	)
	body.SetQuaternion(ode.NewQuaternion(quat.W, quat.X, quat.Y, quat.Z))

	// カプセル(ジオメトリ)
	capsule := b.space.NewCapsule(radius, length)
	capsule.SetBody(body)

	b.geom = capsule

	return &CapsulePhysics{*b}
}

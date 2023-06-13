// Package physics 物理オブジェクト操作パッケージ
package physics

import (
	"github.com/azul3d/engine/native/ode"
	"github.com/trajectoryjp/spatial_id_go/src/spatial-id/common/spatial"
)

// CylinderPhysics 円柱用の物理オブジェクト構造体
type CylinderPhysics struct {
	BasePhysics // 基底物理オブジェクト構造体の埋め込み
}

// NewCylinderPhysics 円柱用の物理オブジェクト構造体コンストラクタ
//
// 円柱用の物理オブジェクト構造体の初期化を行う
//
// 引数：
//
//	radius：円柱の半径
//	start ：円柱の始点
//	end   ：円柱の終点
//
// 戻り値：
//
//	円柱用の物理オブジェクト構造体
func NewCylinderPhysics(radius float64, start spatial.Point3, end spatial.Point3) *CylinderPhysics {

	axis := spatial.NewVectorFromPoints(start, end)
	length := axis.Norm()
	// 重心
	center := spatial.NewLineFromPoints(start, end).ToPoint(0.5)

	// 基底物理オブジェクト構造体
	b := NewBasePhysics()
	// 円柱(剛体)
	body := b.world.NewBody()
	// 座標設定
	body.SetPosition(ode.NewVector3(center.X, center.Y, center.Z))
	// 傾き設定
	quat := spatial.RotateBetweenVector(
		spatial.Vector3{0.0, 0.0, 1.0},
		axis,
	)
	body.SetQuaternion(ode.NewQuaternion(quat.W, quat.X, quat.Y, quat.Z))

	// 円柱(ジオメトリ)
	cylinder := b.space.NewCylinder(radius, length)
	cylinder.SetBody(body)

	b.geom = cylinder

	return &CylinderPhysics{*b}
}

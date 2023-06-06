// Package physics 物理オブジェクト操作パッケージ
package physics

import (
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common/spatial"
	"github.com/azul3d/engine/native/ode"
)

// SpherePhysics 球用の物理オブジェクト構造体
type SpherePhysics struct {
	BasePhysics // 基底物理オブジェクト構造体の埋め込み
}

// NewSpherePhysics 球用の物理オブジェクト構造体コンストラクタ
//
// 球用の物理オブジェクト構造体の初期化を行う
//
// 引数：
//
//	radius：球の半径
//	center：球の中心
//
// 戻り値：
//
//	球用の物理オブジェクト構造体
func NewSpherePhysics(radius float64, center spatial.Point3) *SpherePhysics {

	// 基底物理オブジェクト構造体
	b := NewBasePhysics()
	// 球(剛体)
	body := b.world.NewBody()
	// 座標設定
	body.SetPosition(ode.NewVector3(center.X, center.Y, center.Z))
	// 球(ジオメトリ)
	sphere := b.space.NewSphere(radius)
	sphere.SetBody(body)

	b.geom = sphere

	return &SpherePhysics{*b}
}

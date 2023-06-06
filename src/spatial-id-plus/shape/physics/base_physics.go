// Package physics 物理オブジェクト操作パッケージ
package physics

import (
	"git-codecommit.ap-northeast-1.amazonaws.com/v1/repos/drone_usecase.git/library_go/src/spatial-id/common/spatial"
	"github.com/azul3d/engine/native/ode"
)

// Physics 物理オブジェクトインターフェース
type Physics interface {
	// ボクセルオブジェクト追加処理
	//AddCollideVoxel(index int, center spatial.Point3, lens spatial.Vector3)
	IsCollideVoxel(center spatial.Point3, lens spatial.Vector3) bool
	Space() ode.Space
}

// BasePhysics 基底物理オブジェクト構造体
type BasePhysics struct {
	world ode.World // 動力学演算の対象を含む空間(ワールド)
	space ode.Space // 衝突検出の対象を含む空間(スペース)
	geom  ode.Geom  // 衝突検出の対象を物理オブジェクト
}

// NewBasePhysics 基底物理オブジェクト構造体コンストラクタ
//
// 基底物理オブジェクト構造体の初期化を行う
//
// 戻り値：
//
//	基底物理オブジェクト構造体
func NewBasePhysics() *BasePhysics {
	// ODE初期化
	ode.Init(0, ode.AllAFlag)

	b := BasePhysics{}

	// ワールド
	b.world = ode.NewWorld()

	// スペース
	b.space = ode.NilSpace().NewHashSpace()

	return &b
}

// Space 衝突検出の空間取得
//
// 衝突検出の空間を返却する
//
// 戻り値：
//
//	衝突検出の空間(スペース)
func (b BasePhysics) Space() ode.Space {
	return b.space
}

// IsCollideVoxel ボクセルオブジェクト衝突判定処理
//
// 物理オブジェクトとボクセルオブジェクトの衝突判定処理
//
// 引数：
//
//	center: ボクセル中心
//	lens: ボクセルの対角線ベクトル
//
// 戻り値：
//
//	衝突判定結果
func (b BasePhysics) IsCollideVoxel(center spatial.Point3, lens spatial.Vector3) bool {
	// ボクセル(剛体)
	body := b.world.NewBody()
	// 座標設定
	body.SetPosition(ode.NewVector3(center.X, center.Y, center.Z))
	// ボクセル(ジオメトリ)
	voxel := b.space.NewBox(ode.NewVector3(lens.X, lens.Y, lens.Z))
	voxel.SetBody(body)

	// 物理オブジェクトが未定義の場合
	if b.geom == nil {
		return false
	}

	collide := b.geom.Collide(voxel, 1, 1)

	return len(collide) == 1
}

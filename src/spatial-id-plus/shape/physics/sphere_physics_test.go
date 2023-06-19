// Package physics 物理オブジェクト操作パッケージ
package physics

import (
	"reflect"
	"github.com/trajectoryjp/spatial_id_go/src/spatial-id/common/spatial"
	"testing"

	"github.com/azul3d/engine/native/ode"
)

// TestNewSpherePhysics01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//  - 球の半径： 4.0
//  - 球の中心： (1,2,3)
//
// + 確認内容
//   - 戻り値の型が球用の物理オブジェクト構造体であること
//   - space内に物理オブジェクトが一つ格納されていること
//   - 格納されている物理オブジェクトがsphereであること
//   - 物理オブジェクトの座標と半径が設定した値であること
func TestNewSpherePhysics01(t *testing.T) {
	//入力値
	radius := 4.0
	center := spatial.Point3{-19567.87924100512, 19567.78714071522, -32768}

	// 基底物理オブジェクト構造体
	expectB := NewBasePhysics()
	// 球(剛体)
	expectBody := expectB.world.NewBody()
	// 座標設定
	expectBody.SetPosition(ode.NewVector3(center.X, center.Y, center.Z))
	// 球(ジオメトリ)
	sphere := expectB.space.NewSphere(radius)
	sphere.SetBody(expectBody)

	expectB.geom = sphere

	// 期待値
	expectP := &SpherePhysics{*expectB}

	// テスト対象呼び出し
	resultP := NewSpherePhysics(radius, center)

	// スペース取得
	space := resultP.Space()

	geom := new(ode.Geom)
	// space内の物理オブジェクト数
	geomNum := space.NumGeoms(*geom)
	// 格納されている物理オブジェクト
	geomObject := space.Geom(0)
	// 格納されている物理オブジェクトの型
	resultSphere := geomObject.(ode.Sphere)
	// 物理オブジェクトの座標
	resultPos := resultSphere.Position()
	// 物理オブジェクトの半径
	resultRadius := resultSphere.Radius()

	// 戻り値と期待値の型の比較
	if !reflect.DeepEqual(reflect.TypeOf(expectP), reflect.TypeOf(resultP)) {
		t.Errorf("球用の物理オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(expectP), reflect.TypeOf(resultP))
	}
	// space内に物理オブジェクトが一つ格納されていることを確認
	if geomNum != 1 {
		t.Errorf("spaceの要素数 - 期待値：1, 取得値：%v", geomNum)
	}
	// 格納されている物理オブジェクトがSphereであること
	if !reflect.DeepEqual(reflect.TypeOf(sphere), reflect.TypeOf(geomObject)) {
		t.Errorf("球用の物理オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(sphere), reflect.TypeOf(geomObject))
	}
	// 物理オブジェクトの座標の比較
	if !reflect.DeepEqual(sphere.Position(), resultPos) {
		t.Errorf("座標 - 期待値：%v, 取得値：%v", sphere.Quaternion(), resultPos)
	}
	// 物理オブジェクトの半径の比較
	if !reflect.DeepEqual(radius, resultRadius) {
		t.Errorf("半径 - 期待値：%v, 取得値：%v", radius, resultRadius)
	}

	t.Log("テスト終了")
}

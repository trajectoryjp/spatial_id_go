package physics

import (
	"reflect"
	"spatial-id/common/spatial"
	"testing"

	"github.com/azul3d/engine/native/ode"
)

// TestNewCylinderPhysics01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - 半径: 4
//   - 始点: (1,2,3)
//   - 終点: (5,6,7)
//
// + 確認内容
//   - 戻り値の型がCylinderPhysicsであること
//   - space内に物理オブジェクトが一つ格納されていること
//   - 格納されている物理オブジェクトがCylinderであること
//   - 物理オブジェクトの傾き、座標、半径、長さが設定した値であること
func TestNewCylinderPhysics01(t *testing.T) {
	//入力値
	radius := 4.0
	start := spatial.Point3{1, 2, 3}
	end := spatial.Point3{5, 6, 7}

	//　戻り値として期待される円柱用物理オブジェクトの型のポインタ
	expectP := &CylinderPhysics{}

	// 比較用処理
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
	cylinder.SetData(-1)

	// テスト対象呼び出し
	resultP := NewCylinderPhysics(radius, start, end)

	// スペース取得
	space := resultP.Space()

	geom := new(ode.Geom)
	// space内の物理オブジェクト数
	geomNum := space.NumGeoms(*geom)
	// 格納されている物理オブジェクト
	geomObject := space.Geom(0)
	// 格納されている物理オブジェクトの型
	resultCylinder := geomObject.(ode.Cylinder)
	// 物理オブジェクトの傾き
	resultQuat := resultCylinder.Quaternion()
	// 物理オブジェクトの座標
	resultPos := resultCylinder.Position()
	// 物理オブジェクトの半径、長さ
	resultRadius, resultLength := resultCylinder.Params()

	// 戻り値と期待値の型の比較
	if !reflect.DeepEqual(reflect.TypeOf(expectP), reflect.TypeOf(resultP)) {
		t.Errorf("円柱用の物理オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(expectP), reflect.TypeOf(resultP))
	}
	// space内に物理オブジェクトが一つ格納されていることを確認
	if geomNum != 1 {
		t.Errorf("spaceの要素数 - 期待値：1, 取得値：%v", geomNum)
	}
	// 格納されている物理オブジェクトがCylinderであること
	if !reflect.DeepEqual(reflect.TypeOf(cylinder), reflect.TypeOf(geomObject)) {
		t.Errorf("円柱用の物理オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(cylinder), reflect.TypeOf(geomObject))
	}
	// 物理オブジェクトの傾きの比較
	if !reflect.DeepEqual(cylinder.Quaternion(), resultQuat) {
		t.Errorf("傾き - 期待値：%v, 取得値：%v", cylinder.Quaternion(), resultQuat)
	}
	// 物理オブジェクトの座標の比較
	if !reflect.DeepEqual(cylinder.Position(), resultPos) {
		t.Errorf("座標 - 期待値：%v, 取得値：%v", cylinder.Quaternion(), resultPos)
	}
	// 物理オブジェクトの半径の比較
	if !reflect.DeepEqual(radius, resultRadius) {
		t.Errorf("半径 - 期待値：%v, 取得値：%v", radius, resultRadius)
	}
	// 物理オブジェクトの長さの比較
	if !reflect.DeepEqual(length, resultLength) {
		t.Errorf("長さ - 期待値：%v, 取得値：%v", length, resultLength)
	}

	t.Log("テスト終了")
}

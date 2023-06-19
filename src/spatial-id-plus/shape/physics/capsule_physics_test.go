// Package physics 物理オブジェクト操作パッケージ
package physics

import (
	"reflect"
	"github.com/trajectoryjp/spatial_id_go/src/spatial-id/common/spatial"
	"testing"

	"github.com/azul3d/engine/native/ode"
)

// TestNewCapsulePhysics01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//  カプセルの半径： 2.0
//  カプセルの始点： (3,5,7)
//  カプセルの終点： (-2,-5,9)
//
// + 確認内容
//   - 戻り値の型がカプセル用の物理オブジェクト構造体であること
//   - space内に物理オブジェクトが一つ格納されていること
//   - 格納されている物理オブジェクトがcapsuleであること
//   - 物理オブジェクトの傾き、座標、半径、長さが設定した値であること
func TestNewCapsulePhysics01(t *testing.T) {
	//入力値
	radius := 2.0
	start := spatial.Point3{3, 5, 7}
	end := spatial.Point3{-2, -5, 9}

	// 比較用処理
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

	expectVal := &CapsulePhysics{*b}

	// テスト対象呼び出し
	resultVal := NewCapsulePhysics(radius, start, end)

	// スペース取得
	space := resultVal.Space()

	geom := new(ode.Geom)
	// space内の物理オブジェクト数
	geomNum := space.NumGeoms(*geom)
	// 格納されている物理オブジェクト
	geomObject := space.Geom(0)
	// 格納されている物理オブジェクトの型
	resultCapsule := geomObject.(ode.Capsule)
	// 物理オブジェクトの傾き
	resultQuat := resultCapsule.Quaternion()
	// 物理オブジェクトの座標
	resultPos := resultCapsule.Position()
	// 物理オブジェクトの半径、長さ
	resultRadius, resultLength := resultCapsule.Params()

	// 戻り値と期待値の型の比較
	if !reflect.DeepEqual(reflect.TypeOf(expectVal), reflect.TypeOf(resultVal)) {
		t.Errorf("カプセル用の物理オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(expectVal), reflect.TypeOf(resultVal))
	}
	// space内に物理オブジェクトが一つ格納されていることを確認
	if geomNum != 1 {
		t.Errorf("spaceの要素数 - 期待値：1, 取得値：%v", geomNum)
	}
	// 格納されている物理オブジェクトがCapsuleであること
	if !reflect.DeepEqual(reflect.TypeOf(capsule), reflect.TypeOf(geomObject)) {
		t.Errorf("カプセル用の物理オブジェクトの型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(expectVal), reflect.TypeOf(geomObject))
	}
	// 物理オブジェクトの傾きの比較
	if !reflect.DeepEqual(capsule.Quaternion(), resultQuat) {
		t.Errorf("傾き - 期待値：%v, 取得値：%v", capsule.Quaternion(), resultQuat)
	}
	// 物理オブジェクトの座標の比較
	if !reflect.DeepEqual(capsule.Position(), resultPos) {
		t.Errorf("座標 - 期待値：%v, 取得値：%v", capsule.Quaternion(), resultPos)
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

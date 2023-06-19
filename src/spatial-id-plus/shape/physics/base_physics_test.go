package physics

import (
	"reflect"
	"github.com/trajectoryjp/spatial_id_go/src/spatial-id/common/spatial"
	"testing"
)

// TestNewBasePhysics01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - なし
//
// + 確認内容
//   - 戻り値の型がBasePhysicsであること
func TestNewBasePhysics01(t *testing.T) {

	//　戻り値として期待される円柱用物理オブジェクトの型のポインタ
	expectP := new(BasePhysics)

	// テスト対象呼び出し
	resultP := NewBasePhysics()

	// 戻り値と期待値の型の比較
	if !reflect.DeepEqual(reflect.TypeOf(expectP), reflect.TypeOf(resultP)) {
		t.Errorf("戻り値の型 - 期待値：%v, 取得値：%v",
			reflect.TypeOf(expectP), reflect.TypeOf(resultP))
	}
	t.Log("テスト終了")
}

// TestSpace01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - 始点： (1,2,3)
//   - 終点： (5,6,7)
//
// + 確認内容
//   - 戻り値の真偽値がfalseであること
func TestSpace01(t *testing.T) {
	// 基底物理オブジェクト構造体
	b := NewBasePhysics()

	// 期待値
	expectVal := b.space

	// テスト対象呼び出し
	resultVal := b.Space()

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		t.Errorf("衝突判定 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsCollideVoxel01 正常系動作確認(衝突あり)
//
// 試験詳細：
// + 試験データ
//   - ボクセルの中心： (-19567.87924100512, 19567.78714071522, -32768)
//   - ボクセルの対角線ベクトル： (39135.758482, 39135.758471, 65536)
//
// + 確認内容
//   - 戻り値の真偽値がtrueであること
func TestIsCollideVoxel01(t *testing.T) {
	//入力値
	center := spatial.Point3{-19567.87924100512, 19567.78714071522, -32768}
	lens := spatial.Vector3{39135.758482, 39135.758471, 65536}
	radius := 4.0

	// 期待値
	expectVal := true
	// 基底物理オブジェクト構造体
	resultB := NewBasePhysics()

	// 球(剛体)
	resultBody := resultB.world.NewBody()
	// 球(ジオメトリ)
	sphere := resultB.space.NewSphere(radius)
	sphere.SetBody(resultBody)
	resultB.geom = sphere

	// テスト対象呼び出し
	resultVal := resultB.IsCollideVoxel(center, lens)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		t.Errorf("衝突判定 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsCollideVoxel02 正常系動作確認(衝突なし)
//
// 試験詳細：
// + 試験データ
//   - ボクセルの中心： (19567.879241005117, 58703.36144845117, 32768})
//   - ボクセルの対角線ベクトル： (39135.758482, 39135.758471, 65536)
//
// + 確認内容
//   - 戻り値の真偽値がfalseであること
func TestIsCollideVoxel02(t *testing.T) {
	//入力値
	center := spatial.Point3{19567.879241005117, 58703.36144845117, 32768}
	lens := spatial.Vector3{39135.758482, 39135.758471, 65536}
	radius := 4.0

	// 期待値
	expectVal := false

	// 基底物理オブジェクト構造体
	resultB := NewBasePhysics()

	// 球(剛体)
	resultBody := resultB.world.NewBody()
	// 球(ジオメトリ)
	sphere := resultB.space.NewSphere(radius)
	sphere.SetBody(resultBody)
	resultB.geom = sphere

	// テスト対象呼び出し
	resultVal := resultB.IsCollideVoxel(center, lens)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		t.Errorf("衝突判定 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsCollideVoxel03 異常系動作確認(物理オブジェクトが未定義)
//
// 試験詳細：
// + 試験データ
//   - ボクセルの中心： (-19567.87924100512, 19567.78714071522, -32768)
//   - ボクセルの対角線ベクトル： (39135.758482, 39135.758471, 65536)
//
// + 確認内容
//   - 戻り値の真偽値がfalseであること
func TestIsCollideVoxel03(t *testing.T) {
	//入力値
	center := spatial.Point3{-19567.87924100512, 19567.78714071522, -32768}
	lens := spatial.Vector3{39135.758482, 39135.758471, 65536}

	// 基底物理オブジェクト構造体
	resultB := NewBasePhysics()

	// 期待値
	expectVal := false

	// テスト対象呼び出し
	resultVal := resultB.IsCollideVoxel(center, lens)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		t.Errorf("戻り値 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

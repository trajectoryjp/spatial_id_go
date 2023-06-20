package spatial

import (
	"testing"

	"github.com/trajectoryjp/spatial_id_go/common"
	"github.com/trajectoryjp/spatial_id_go/common/consts"
)

// TestQuat 四元数の構造体(試験用)
type TestQuat struct {
	W, X, Y, Z float64
}

// Mul 四元数の積
//
// 四元数の積を算出して返却する。
//
// 引数：
//
//	b： 掛け合わせる四元数
//
// 戻り値：
//
//	四元数の積
func (a TestQuat) Mul(b TestQuat) TestQuat {
	return TestQuat{
		(b.W * a.W) - (b.X * a.X) - (b.Y * a.Y) - (b.Z * a.Z),
		(b.X * a.W) + (b.W * a.X) - (b.Z * a.Y) + (b.Y * a.Z),
		(b.Y * a.W) + (b.Z * a.X) + (b.W * a.Y) - (b.X * a.Z),
		(b.Z * a.W) - (b.Y * a.X) + (b.X * a.Y) + (b.W * a.Z),
	}
}

// Conjugate 四元数の共役
//
// 共役四元数を算出して返却する。
//
// 戻り値：
//
//	共役四元数
func (a TestQuat) Conjugate() TestQuat {
	return TestQuat{
		a.W,
		-a.X,
		-a.Y,
		-a.Z,
	}
}

// TransformVec3 四元数でのベクトル変換
//
// 入力のベクトルを四元数で変換した値を返却する。
//
// 引数：
//
//	v： 入力ベクトル
//
// 戻り値：
//
//	変換後のベクトル
func (a TestQuat) TransformVec3(v Vector3) Vector3 {
	vecTestQuat := TestQuat{0.0, v.X, v.Y, v.Z}
	conjugate := a.Conjugate()
	vecTestQuat = conjugate.Mul(vecTestQuat).Mul(a)
	return Vector3{vecTestQuat.X, vecTestQuat.Y, vecTestQuat.Z}
}

// TestRotateBetweenVector01 正常系動作確認(ベクトルが正反対を向いていない)
//
// 試験詳細：
//   - 試験データ
//     開始ベクトル： (1,2,3)
//     終了ベクトル： (4,5,6)
//
// + 確認内容
//   - 入力に応じた2ベクトル間の四元数が返ってくること
//   - 四元数を元に開始ベクトルを回転させて、終了ベクトルと近似すること
func TestRotateBetweenVector01(t *testing.T) {
	start := Vector3{1, 2, 3}
	end := Vector3{4, 5, 6}

	// テスト対象呼び出し
	resultVal := RotateBetweenVector(start, end)

	// 検算値作成
	testQuat := TestQuat{
		resultVal.W,
		resultVal.X,
		resultVal.Y,
		resultVal.Z,
	}
	expectVal := testQuat.TransformVec3(start.Unit())

	// 検算結果を比較
	if !common.AlmostEqual(expectVal.X, end.Unit().X, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [X]  期待値%v, 取得値%v", expectVal.X, end.Unit().X)

	}
	if !common.AlmostEqual(expectVal.Y, end.Unit().Y, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [Y]  期待値%v, 取得値%v", expectVal.Y, end.Unit().Y)

	}
	if !common.AlmostEqual(expectVal.Z, end.Unit().Z, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [Z]  期待値%v, 取得値%v", expectVal.Z, end.Unit().Z)

	}

	t.Log("テスト終了")
}

// TestRotateBetweenVector02 正常系動作確認(ベクトルが正反対を向いている)
//
// 試験詳細：
//   - 試験データ
//     開始ベクトル： (1,2,3)
//     終了ベクトル： (-1,-2,-3)
//
// + 確認内容
//   - 入力に応じた2ベクトル間の四元数が返ってくること
//   - 四元数を元に開始ベクトルを回転させて、終了ベクトルと近似すること
func TestRotateBetweenVector02(t *testing.T) {
	start := Vector3{1, 2, 3}
	end := Vector3{-1, -2, -3}

	// テスト対象呼び出し
	resultVal := RotateBetweenVector(start, end)

	// 検算値作成
	testQuat := TestQuat{
		resultVal.W,
		resultVal.X,
		resultVal.Y,
		resultVal.Z,
	}
	expectVal := testQuat.TransformVec3(start.Unit())

	// 検算結果を比較
	if !common.AlmostEqual(expectVal.X, end.Unit().X, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [X]  期待値%v, 取得値%v", expectVal.X, end.Unit().X)

	}
	if !common.AlmostEqual(expectVal.Y, end.Unit().Y, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [Y]  期待値%v, 取得値%v", expectVal.Y, end.Unit().Y)

	}
	if !common.AlmostEqual(expectVal.Z, end.Unit().Z, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [Z]  期待値%v, 取得値%v", expectVal.Z, end.Unit().Z)

	}

	t.Log("テスト終了")
}

// TestRotateBetweenVector03 正常系動作確認
// (ベクトルが正反対を向いていて、かつ回転軸ベクトルののユークリッドノルムが小さい)
//
// 試験詳細：
//   - 試験データ
//     開始ベクトル： (0,0,0.00000000000000000001)
//     終了ベクトル： (0,0,-0.00000000000000000001)
//
// + 確認内容
//   - 入力に応じた2ベクトル間の四元数が返ってくること
//   - 四元数を元に開始ベクトルを回転させて、終了ベクトルと近似すること
func TestRotateBetweenVector03(t *testing.T) {
	start := Vector3{0, 0, 0.00000000000000000001}
	end := Vector3{0, 0, -0.00000000000000000001}

	// テスト対象呼び出し
	resultVal := RotateBetweenVector(start, end)

	// 検算値作成
	testQuat := TestQuat{
		resultVal.W,
		resultVal.X,
		resultVal.Y,
		resultVal.Z,
	}
	expectVal := testQuat.TransformVec3(start.Unit())

	// 検算結果を比較
	if !common.AlmostEqual(expectVal.X, end.Unit().X, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [X]  期待値%v, 取得値%v", expectVal.X, end.Unit().X)

	}
	if !common.AlmostEqual(expectVal.Y, end.Unit().Y, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [Y]  期待値%v, 取得値%v", expectVal.Y, end.Unit().Y)

	}
	if !common.AlmostEqual(expectVal.Z, end.Unit().Z, consts.Minima) {
		t.Errorf("4元数で回転しなおしたベクトル比較 - [Z]  期待値%v, 取得値%v", expectVal.Z, end.Unit().Z)

	}

	t.Log("テスト終了")
}

// TestQuatFromAxisAngle01 正常系動作確認
//
// 試験詳細：
//   - 試験データ
//     回転軸ベクトル：
//     回転角度(ラジアン)： math.Pi
//
// + 確認内容
//   - 入力に応じた回転ベクトルの四元数が返ってくること
func TestQuatFromAxisAngle01(t *testing.T) {
	// 下記テストケース同時消化
	// - TestRotateBetweenVector02
	// - TestRotateBetweenVector03
}

package spatial

import (
	"reflect"
	"testing"
)

// TestNewMatrix301 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     行列(0,0) = 1.1, 行列(0,1) = 2.2, 行列(0,2) = 3.2
//     行列(1,0) = 4.9, 行列(1,1) = 5.0, 行列(1,2) = 6
//     行列(2,0) = 7.1, 行列(2,1) = 8.55, 行列(2,2) = 9.123
//
// + 確認内容
//   - 入力に対応した3×3行列の構造体が取得できること
func TestNewMatrix301(t *testing.T) {
	//入力値
	var mm00 float64 = 1.1
	var mm01 float64 = 2.2
	var mm02 float64 = 3.2

	var mm10 float64 = 4.9
	var mm11 float64 = 5.0
	var mm12 float64 = 6

	var mm20 float64 = 7.1
	var mm21 float64 = 8.55
	var mm22 float64 = 9.123

	resultVal := NewMatrix3(mm00, mm01, mm02, mm10, mm11, mm12, mm20, mm21, mm22)

	//期待値
	expectVal := Matrix3{{mm00, mm01, mm02}, {mm10, mm11, mm12}, {mm20, mm21, mm22}}

	// 戻り値の3×3行列の構造体と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("3×3行列の構造体 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestNewUnitMatrix301 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (入力なし)
//
// + 確認内容
//   - 3×3単位行列の構造体が取得できること
func TestNewUnitMatrix301(t *testing.T) {

	resultVal := NewUnitMatrix3()
	//期待値
	expectVal := Matrix3{{1.0, 0.0, 0.0}, {0.0, 1.0, 0.0}, {0.0, 0.0, 1.0}}

	// 戻り値の3×3行列の構造体と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("3×3単位行列の構造体 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMul01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//
//   - パターン1：
//     行列A(0,0) = 1.0, 行列A(0,1) = 2.0, 行列A(0,2) = 3.0
//     行列A(1,0) = 4.0, 行列A(1,1) = 5.0, 行列A(1,2) = 6.0
//     行列A(2,0) = 7.0, 行列A(2,1) = 8.0, 行列A(2,2) = 9.0
//
//     行列B(0,0) = 1.0, 行列B(0,1) = 2.0, 行列B(0,2) = 3.0
//     行列B(1,0) = 4.0, 行列B(1,1) = 5.0, 行列B(1,2) = 6.0
//     行列B(2,0) = 7.0, 行列B(2,1) = 8.0, 行列B(2,2) = 9.0
//
// + 確認内容
//   - 入力に対応した3×3行列の行列積が取得できること
func TestMul01(t *testing.T) {
	//入力値
	var mA00 float64 = 1.0
	var mA01 float64 = 2.0
	var mA02 float64 = 3.0

	var mA10 float64 = 4.0
	var mA11 float64 = 5.0
	var mA12 float64 = 6.0

	var mA20 float64 = 7.0
	var mA21 float64 = 8.0
	var mA22 float64 = 9.0

	var mB00 float64 = 1.0
	var mB01 float64 = 2.0
	var mB02 float64 = 3.0

	var mB10 float64 = 4.0
	var mB11 float64 = 5.0
	var mB12 float64 = 6.0

	var mB20 float64 = 7.0
	var mB21 float64 = 8.0
	var mB22 float64 = 9.0

	A := NewMatrix3(mA00, mA01, mA02, mA10, mA11, mA12, mA20, mA21, mA22)
	B := NewMatrix3(mB00, mB01, mB02, mB10, mB11, mB12, mB20, mB21, mB22)

	resultVal := A.Mul(B)
	//期待値
	expectVal := Matrix3{{30.0, 36.0, 42.0}, {66.0, 81.0, 96.0}, {102.0, 126.0, 150.0}}

	// 戻り値の3×3行列の行列積と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("3×3行列の行列積 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMulVec01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     行列(0,0) = 1.0, 行列(0,1) = 2.0, 行列(0,2) = 3.0
//     行列(1,0) = 4.0, 行列(1,1) = 5.0, 行列(1,2) = 6.0
//     行列(2,0) = 7.0, 行列(2,1) = 8.0, 行列(2,2) = 9.0
//
// + 確認内容
//   - 入力に対応した3×3行列の構造体が取得できること
func TestMulVec01(t *testing.T) {
	//入力値
	var mm00 float64 = 1.0
	var mm01 float64 = 2.0
	var mm02 float64 = 3.0

	var mm10 float64 = 4.0
	var mm11 float64 = 5.0
	var mm12 float64 = 6.0

	var mm20 float64 = 7.0
	var mm21 float64 = 8.0
	var mm22 float64 = 9.0

	A := NewMatrix3(mm00, mm01, mm02, mm10, mm11, mm12, mm20, mm21, mm22)

	v := Vector3{1, 2, 3}

	resultVal := A.MulVec(v)

	//期待値
	expectVal := Vector3{14, 32, 50}

	// 戻り値のベクトルと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("ベクトル - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

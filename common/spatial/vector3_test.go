package spatial

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

// TestNewVectorFromPoints01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点座標：(1,0,6), 終点座標：(1,4,9))
//
// + 確認内容
//   - 始点から終点へのベクトルが取得できること
func TestNewVectorFromPoints01(t *testing.T) {
	//入力値
	p := Point3{1, 0, 6}
	q := Point3{1, 4, 9}
	resultVal := NewVectorFromPoints(p, q)

	//期待値
	expectVal := Vector3{0, 4, 3}

	// 始点から終点へのベクトルと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("ベクトル - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestAdd01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (加算前のベクトル：(1,2,3), 加算するベクトル：(5,6,7))
//
// + 確認内容
//   - 加算したベクトルが取得できること
func TestAdd01(t *testing.T) {
	//入力値
	p := Vector3{1, 2, 3}
	q := Vector3{5, 6, 7}
	resultVal := p.Add(q)

	//期待値
	expectVal := Vector3{p.X + q.X, p.Y + q.Y, p.Z + q.Z}

	// 戻り値のベクトルと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("ベクトル - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestSub01 正常系動作確認
//
// 試験詳細：
//
// + 試験データ
//   - パターン1：
//     (減算前のベクトル：(5,6,7), 減算するベクトル：(1,2,3))
//
// + 確認内容
//   - 減算したベクトルが取得できること
func TestSub01(t *testing.T) {
	//入力値
	p := Vector3{5, 6, 7}
	q := Vector3{1, 2, 3}
	resultVal := p.Sub(q)

	//期待値
	expectVal := Vector3{p.X - q.X, p.Y - q.Y, p.Z - q.Z}

	// 戻り値のベクトルと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("ベクトル - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestScale01 正常系動作確認
//
// 試験詳細：
//
// + 試験データ
//   - パターン1：
//     (元のベクトル：(1,2,3), 乗算するスカラー値：3.0)
//
// + 確認内容
//   - スカラー倍したベクトルが取得できること
func TestScale01(t *testing.T) {
	//入力値
	p := Vector3{1, 2, 3}
	q := 3.0
	resultVal := p.Scale(q)

	//期待値
	expectVal := Vector3{p.X * q, p.Y * q, p.Z * q}

	// 戻り値のベクトルと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("ベクトル - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestDot01 正常系動作確認
//
// 試験詳細：
//
// + 試験データ
//   - パターン1：
//     (ベクトル1：(4,8,12), ベクトル2：(2,2,2))
//
// + 確認内容
//   - ベクトルの内積を算出したa・bが取得できること
func TestDot01(t *testing.T) {
	//入力値
	p := Vector3{4, 8, 12}
	q := Vector3{2, 2, 2}
	resultVal := p.Dot(q)

	//期待値
	expectVal := p.X*q.X + p.Y*q.Y + p.Z*q.Z

	// 戻り値の内積と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("内積 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestCross01 正常系動作確認
//
// 試験詳細：
//
// + 試験データ
//   - パターン1：
//     (ベクトル1：(4,8,12), ベクトル2：(2,2,2))
//
// + 確認内容
//   - ベクトルの外積を算出したa×bが取得できること
func TestCross01(t *testing.T) {
	//入力値
	p := Vector3{4, 8, 12}
	q := Vector3{2, 2, 2}
	resultVal := p.Cross(q)

	//期待値
	expectVal := Vector3{p.Y*q.Z - p.Z*q.Y, p.Z*q.X - p.X*q.Z, p.X*q.Y - p.Y*q.X}

	// 戻り値の外積と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("外積 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestNorm01 正常系動作確認
//
// 試験詳細：
//
// + 試験データ
//   - パターン1：
//     (ベクトル1：(1,2,3))
//
// + 確認内容
//   - ベクトルのユークリッドノルムが取得できること
//
// 備考：
//   - 本入力値でr3パッケージとmathパッケージでは小数点16桁以降の算出結果に差分が出るが、
//     本試験では許容可能な誤差とする。
func TestNorm01(t *testing.T) {
	//入力値
	p := Vector3{1, 2, 3}
	resultVal := strconv.FormatFloat(p.Norm(), 'f', 14, 64)

	//期待値
	expectVal := strconv.FormatFloat(math.Sqrt((p.X*p.X + p.Y*p.Y + p.Z*p.Z)), 'f', 14, 64)

	// 戻り値のユークリッドノルムと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("ユークリッドノルム - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestL1Norm01 正常系動作確認
//
// 試験詳細：
//
// + 試験データ
//   - パターン1：
//     (ベクトル1：(1,-2,3))
//
// + 確認内容
//   - ベクトルのL1ノルムが取得できること
func TestL1Norm01(t *testing.T) {
	//入力値
	p := Vector3{1, -2, 3}
	resultVal := p.L1Norm()

	//期待値
	expectVal := math.Abs(p.X) + math.Abs(p.Y) + math.Abs(p.Z)

	// 戻り値のL1ノルムと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("L1ノルム - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestUnit01 正常系動作確認
//
// 試験詳細：
//
// + 試験データ
//   - パターン1：
//     (ベクトル1：(2,3,5))
//
// + 確認内容
//   - ベクトルを正規化し、単位ベクトルが取得できること
func TestUnit01(t *testing.T) {
	//入力値
	p := Vector3{2, 3, 5}
	resultVal := p.Unit()

	//期待値
	un := math.Sqrt((p.X*p.X + p.Y*p.Y + p.Z*p.Z))
	expectVal := Vector3{p.X / un, p.Y / un, p.Z / un}

	// 戻り値の単位ベクトルと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("単位ベクトル - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestCos01 正常系動作確認
//
// 試験詳細：
//
// + 試験データ
//   - パターン1：
//     (ベクトル1：(1,1,1), ベクトル2：(0,4,4))
//
// + 確認内容
//   - ベクトル間の余弦(cos)を取得できること
func TestCos01(t *testing.T) {
	//入力値
	p := Vector3{0, 1, 0}
	q := Vector3{0, 4, 4}
	resultVal := p.Cos(q)

	//期待値
	pv := math.Sqrt((p.X*p.X + p.Y*p.Y + p.Z*p.Z))
	qv := math.Sqrt((q.X*q.X + q.Y*q.Y + q.Z*q.Z))
	pqv := pv * qv
	pq := p.X*q.X + p.Y*q.Y + p.Z*q.Z
	expectVal := pq / pqv

	// 戻り値のコサインと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("コサイン - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

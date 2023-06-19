package spatial

import (
	"math"
	"reflect"
	"testing"
)

// TestUniqueAppend01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     点のスライス：{1.0, 2.0, 3.0},
//     追加する点：{4.0, 5.0, 6.0}(追加点が点のスライスに対し一意(ユニーク)である),
//     ユニーク判定時の小数点誤差：0.01
//
// + 確認内容
//   - 入力のスライスに、追加点が追加されたスライスが返却されること
func TestUniqueAppend01(t *testing.T) {
	// データ定義
	p := Point3{1.0, 2.0, 3.0}
	points := []*Point3{&p}

	addPoint := Point3{4.0, 5.0, 6.0}
	epsilon := 0.01

	// 取得値
	resultVal := UniqueAppend(points, &addPoint, float64(epsilon))

	//期待値
	expectVal := []*Point3{&p, &addPoint}
	// 戻り値の点のスライスと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		for index, exp := range expectVal {
			t.Errorf("スライス[%d] - 期待値：%v, 取得値：%v", index+1, exp, resultVal[index])
		}

		t.Errorf("スライス - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestUniqueAppend02 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン2：
//     点のスライス：{1.0, 2.0, 3.0},
//     追加する点：{1.0, 2.0, 3.0}(追加点が点のスライスに対し一意(ユニーク)でない),
//     ユニーク判定時の小数点誤差：0.01
//
// + 確認内容
//   - 入力のスライスに、追加点が追加されていないスライスが返却されること
func TestUniqueAppend02(t *testing.T) {
	// データ定義
	p := Point3{1.0, 2.0, 3.0}
	points := []*Point3{&p}
	addPoint := Point3{1.0, 2.0, 3.0}
	epsilon := 0.01

	// 取得値
	resultVal := UniqueAppend(points, &addPoint, float64(epsilon))

	// 期待値
	expectVal := points

	// 結果の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestUniqueAppend03 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン3：
//     点のスライス：{(未入力)}},（追加される前の配列が空の場合）
//     追加する点：{1.0, 2.0, 3.0}(追加点が点のスライスに対し一意(ユニーク)である),
//     ユニーク判定時の小数点誤差：0.01
//
// + 確認内容
//   - 入力のスライスに、追加点が追加されたスライスが返却されること
func TestUniqueAppend03(t *testing.T) {
	// データ定義
	points := []*Point3{}
	addPoint := Point3{1.0, 2.0, 3.0}
	epsilon := 0.01

	// 取得値
	resultVal := UniqueAppend(points, &addPoint, float64(epsilon))

	// 期待値
	expectVal := []*Point3{&addPoint}

	// 結果の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestUniqueAppend04 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン4：
//     点のスライス1：{1.0, 2.0, 3.0},
//     点のスライス2：{2.0, 3.0, 4.0},
//     点のスライス3：{3.0, 4.0, 5.0},
//     追加する点：{4.0, 5.0, 6.0}(追加点が点のスライスに対し一意(ユニーク)であり、追加される前の配列が複数の要素を含む場合
//     ユニーク判定時の小数点誤差：0.01
//
// + 確認内容
//   - 入力のスライスに、追加点が追加されたスライスが返却されること
func TestUniqueAppend04(t *testing.T) {
	// データ定義
	p01 := Point3{1.0, 2.0, 3.0}
	p02 := Point3{2.0, 3.0, 4.0}
	p03 := Point3{3.0, 4.0, 5.0}
	points := []*Point3{&p01, &p02, &p03}
	addPoint := Point3{4.0, 5.0, 6.0}

	epsilon := 0.01

	// 取得値
	resultVal := UniqueAppend(points, &addPoint, float64(epsilon))

	// 期待値
	expectVal := []*Point3{&p01, &p02, &p03, &addPoint}

	// 結果の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMaxPoint01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     Point3スライス1：{9, 9, 9},(Point3スライスの1番目の要素が最大値だった場合)
//     Point3スライス2：{1, 2, 3},
//     Point3スライス3：{2, 3, 4},
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - 入力値であるスライスの1番目の要素が返却されること
func TestMaxPoint01(t *testing.T) {
	// データ定義
	p01 := Point3{9, 9, 9}
	p02 := Point3{1, 2, 3}
	p03 := Point3{2, 3, 4}
	points := []*Point3{&p01, &p02, &p03}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MaxPoint(points, vec)

	// 期待値
	expectVal := &p01

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMaxPoint02 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン2：
//     Point3スライス1：{0, 0, 0},
//     Point3スライス2：{9, 8, 7},
//     Point3スライス3：{9, 9, 9},(Point3スライスの2番目以降の要素が最大値だった場合)
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - 入力値であるスライスの2番目以降の要素が返却されること
func TestMaxPoint02(t *testing.T) {
	// データ定義
	p01 := Point3{0, 0, 0}
	p02 := Point3{9, 8, 7}
	p03 := Point3{9, 9, 9}
	points := []*Point3{&p01, &p02, &p03}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MaxPoint(points, vec)

	// 期待値
	expectVal := &p03

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMaxPoint03 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン3：
//     Point3スライス1：{(未入力)}    (Point3スライスが空だった場合)
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - 入力チェックエラーとなること
func TestMaxPoint03(t *testing.T) {
	// データ定義
	points := []*Point3{}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MaxPoint(points, vec)

	// 期待値
	expectVal := new(Point3)
	expectErr := "InputValueError,入力チェックエラー"

	// 結果の比較
	if !reflect.DeepEqual(resultErr.Error(), expectErr) {
		t.Errorf("エラー - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 取得値：%v", resultVal)
	}
	t.Log("テスト終了")
}

// TestMaxPoint04 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン4：
//     Point3スライス1：{4, 5, 6},
//     Point3スライス2：{6, 4, 5},（境界値試験±0、スライスに同じ値のベクトルが入っている場合）
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - 入力したスライスの番号の若い最大値を示す要素が返却されること
func TestMaxPoint04(t *testing.T) {
	// データ定義 p01とp02は最大値が同じ値となる
	p01 := Point3{4, 5, 6}
	p02 := Point3{6, 4, 5}
	points := []*Point3{&p01, &p02}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MaxPoint(points, vec)

	// 期待値
	expectVal := &p01

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMaxPoint05 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン5：
//     Point3スライス1：{1, 1, 3},
//     Point3スライス2：{3, 2, 1},（境界値試験+1、スライスの1番目の要素より2番目の要素の最大値が1大きい）
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - スライスの最大値を示す要素が返却されること
func TestMaxPoint05(t *testing.T) {
	// データ定義
	p01 := Point3{1, 1, 3}
	p02 := Point3{3, 2, 1}
	points := []*Point3{&p01, &p02}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MaxPoint(points, vec)

	// 期待値
	expectVal := &p02

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMaxPoint06 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン6：
//     Point3スライス1：{1, 2, 3},
//     Point3スライス2：{3, 1, 1},（境界値試験-1、スライスの1番目の要素より2番目の要素の最大値が1小さい）
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - スライスの最大値を示す要素が返却されること
func TestMaxPoint06(t *testing.T) {
	// データ定義
	p01 := Point3{1, 2, 3}
	p02 := Point3{3, 1, 1}
	points := []*Point3{&p01, &p02}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MaxPoint(points, vec)

	// 期待値
	expectVal := &p01

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMinPoint01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     Point3スライス1：{1, 2, 3},(Point3スライスの1番目の要素が最小値だった場合)
//     Point3スライス2：{9, 9, 9},
//     Point3スライス3：{2, 3, 4},
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - 入力値であるスライスの1番目の要素が返却されること
func TestMinPoint01(t *testing.T) {
	// データ定義
	p01 := Point3{1, 2, 3}
	p02 := Point3{9, 9, 9}
	p03 := Point3{2, 3, 4}
	points := []*Point3{&p01, &p02, &p03}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MinPoint(points, vec)

	// 期待値
	expectVal := &p01

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMinPoint02 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン2：
//     Point3スライス1：{9, 9, 9},
//     Point3スライス2：{8, 9, 3},
//     Point3スライス3：{5, 3, 4},(Point3スライスの2番目以降の要素が最小値だった場合)
//     Point3スライス4：{9, 2, 9},
//     Point3スライス5：{4, 5, 6},
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - 入力値であるスライスの2番目以降の要素が返却されること
func TestMinPoint02(t *testing.T) {
	// データ定義
	p01 := Point3{9, 9, 9}
	p02 := Point3{8, 9, 3}
	p03 := Point3{5, 3, 4}
	p04 := Point3{9, 2, 9}
	p05 := Point3{4, 5, 6}
	points := []*Point3{&p01, &p02, &p03, &p04, &p05}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MinPoint(points, vec)

	// 期待値
	expectVal := &p03

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMinPoint03 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン3：
//     Point3スライス1：{(未入力)},(Point3スライスが空だった場合)
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - スライス空入力により入力チェックエラーとなること
func TestMinPoint03(t *testing.T) {
	// データ定義
	points := []*Point3{}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MinPoint(points, vec)

	// 期待値
	expectVal := new(Point3)
	expectErr := "InputValueError,入力チェックエラー"

	// 結果の比較
	if !reflect.DeepEqual(resultErr.Error(), expectErr) {
		t.Errorf("エラー - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 取得値：%v", resultVal)
	}
	t.Log("テスト終了")
}

// TestMinPoint04 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン4：
//     Point3スライス1：{4, 5, 6},（境界値試験±0、スライスに同じ値のベクトルが入っている場合）
//     Point3スライス2：{6, 4, 5},
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - 入力したスライスの番号の若い最小値を示す要素が返却されること
func TestMinPoint04(t *testing.T) {
	// データ定義 p01とp02は最大値が同じ値となる
	p01 := Point3{4, 5, 6}
	p02 := Point3{6, 4, 5}
	points := []*Point3{&p01, &p02}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MinPoint(points, vec)

	// 期待値
	expectVal := &p01

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMinPoint05 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン5：
//     Point3スライス1：{2, 2, 1},
//     Point3スライス2：{1, 2, 3},（境界値試験-1、スライスの1番目の要素より2番目の要素の最大値が1大きい）
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - スライスの最小値を示す要素が返却されること
func TestMinPoint05(t *testing.T) {
	// データ定義
	p01 := Point3{2, 2, 1}
	p02 := Point3{1, 2, 3}
	points := []*Point3{&p01, &p02}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MinPoint(points, vec)

	// 期待値
	expectVal := &p01

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestMinPoint06 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン6：
//     Point3スライス1：{3, 1, 2},
//     Point3スライス2：{1, 2, 2},（境界値試験+1、スライスの1番目の要素より2番目の要素の最大値が1小さい）
//     Point3の比較値を出すためのベクトル：{1, 1, 1}
//
// + 確認内容
//   - スライスの最小値を示す要素が返却されること
func TestMinPoint06(t *testing.T) {
	// データ定義
	p01 := Point3{3, 1, 2}
	p02 := Point3{1, 2, 2}
	points := []*Point3{&p01, &p02}
	vec := Vector3{1, 1, 1}

	// 取得値
	resultVal, resultErr := MinPoint(points, vec)

	// 期待値
	expectVal := &p02

	// 結果の比較
	if resultErr != nil {
		t.Errorf("エラー - 期待値：%v, 取得値：%v", nil, resultErr)
	}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("点のユニークを保持した配列追加 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsClose01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     点1座標：{1.0, 1.0, 1.0},
//     点2座標：{1.0, 1.0, 1.0},
//     浮動小数点誤差：0.0
//     2点がx, y, z座標において全て一致している
//
// + 確認内容
//   - trueが返却されること
func TestIsClose01(t *testing.T) {
	// データ定義
	p01 := Point3{1.0, 1.0, 1.0}
	p02 := Point3{1.0, 1.0, 1.0}
	epsilon := 0.0

	// 取得値
	resultVal := p01.IsClose(p02, epsilon)

	// 期待値
	expectVal := true

	// 結果の比較
	if resultVal != expectVal {
		t.Errorf("同一点である - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsClose02 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン2：
//     点1座標：{1.0, 1.0, 1.0},
//     点2座標：{3.0, 1.0, 1.0},
//     浮動小数点誤差：0.0
//     2点がx, y, z座標において、x座標の値のみ一致していない
//
// + 確認内容
//   - falseが返却されること
func TestIsClose02(t *testing.T) {
	// データ定義
	p01 := Point3{1.0, 1.0, 1.0}
	p02 := Point3{3.0, 1.0, 1.0}
	epsilon := 0.0

	// 取得値
	resultVal := p01.IsClose(p02, epsilon)

	// 期待値
	expectVal := false

	// 結果の比較
	if resultVal != expectVal {
		t.Errorf("同一点でない - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsClose03 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン3：
//     点1座標：{1.0, 1.0, 1.0},
//     点2座標：{1.0, 5.0, 1.0},
//     浮動小数点誤差：0.0
//     2点がx, y, z座標において、y座標の値のみ一致していない
//
// + 確認内容
//   - falseが返却されること
func TestIsClose03(t *testing.T) {
	// データ定義
	p01 := Point3{1.0, 1.0, 1.0}
	p02 := Point3{1.0, 5.0, 1.0}
	epsilon := 0.0

	// 取得値
	resultVal := p01.IsClose(p02, epsilon)

	// 期待値
	expectVal := false

	// 結果の比較
	if resultVal != expectVal {
		t.Errorf("同一点でない - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsClose04 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン4：
//     点1座標：{1.0, 1.0, 1.0},
//     点2座標：{1.0, 1.0, 4.0},
//     浮動小数点誤差：0.0
//     2点がx, y, z座標において、z座標の値のみ一致していない
//
// + 確認内容
//   - falseが返却されること
func TestIsClose04(t *testing.T) {
	// データ定義
	p01 := Point3{1.0, 1.0, 1.0}
	p02 := Point3{1.0, 1.0, 4.0}
	epsilon := 0.0

	// 取得値
	resultVal := p01.IsClose(p02, epsilon)

	// 期待値
	expectVal := false

	// 結果の比較
	if resultVal != expectVal {
		t.Errorf("同一点でない - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsClose05 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン5：
//     点1座標：{1.0, 1.0, 1.0},
//     点2座標：{1.1, 1.1, 1.1},
//     浮動小数点誤差：0.2
//     2点がx, y, z座標において、浮動小数点誤差epsilonを許容して一致している
//
// + 確認内容
//   - trueが返却されること
func TestIsClose05(t *testing.T) {
	// データ定義
	p01 := Point3{1.0, 1.0, 1.0}
	p02 := Point3{1.1, 1.1, 1.1}
	epsilon := 0.2

	// 取得値
	resultVal := p01.IsClose(p02, epsilon)

	// 期待値
	expectVal := true

	// 結果の比較
	if resultVal != expectVal {
		t.Errorf("同一点でない - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestIsClose06 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン6：
//     点1座標：{1.0, 1.0, 1.0},
//     点2座標：{1.2, 1.2, 1.2},
//     浮動小数点誤差：0.1
//     2点がx, y, z座標において、浮動小数点誤差epsilonを許容して一致していない
//
// + 確認内容
//   - falseが返却されること
func TestIsClose06(t *testing.T) {
	// データ定義
	p01 := Point3{1.0, 1.0, 1.0}
	p02 := Point3{1.2, 1.2, 1.2}
	epsilon := 0.1

	// 取得値
	resultVal := p01.IsClose(p02, epsilon)

	// 期待値
	expectVal := false

	// 結果の比較
	if resultVal != expectVal {
		t.Errorf("同一点でない - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestTranslate01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     座標点(1, 1, 1)、座標ベクトル(2, 3, 2)
//
// + 確認内容
//   - 平行移動した点が正しく返却されること
func TestTranslate01(t *testing.T) {
	// データ定義
	p := Point3{1, 1, 1}
	vec := Vector3{2, 2, 2}

	// 取得値
	resultVal := p.Translate(vec)

	// 期待値
	expectVal := Point3{3, 3, 3}

	// 結果の比較
	if resultVal != expectVal {
		t.Errorf("点の平行移動 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestDistancePoint01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     座標点01(1, 1, 0)、座標点02(2, 2, 0)
//
// + 確認内容
//   - 2点間の距離が正しく返却されること
func TestDistancePoint01(t *testing.T) {
	// データ定義
	p01 := Point3{1, 1, 0}
	p02 := Point3{2, 2, 0}

	// 取得値
	resultVal := p01.DistancePoint(p02)

	// 期待値
	expectVal := math.Sqrt(2)

	// 結果の比較
	if resultVal != expectVal {
		t.Errorf("点の平行移動 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

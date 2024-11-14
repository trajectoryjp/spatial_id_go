package spatialID

import (
	"reflect"
	"testing"
)

// TestAlmostEquals01 同値確認関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 数値1：5.0 数値2：5.0 浮動小数点誤差：0.1
//
// + 確認内容
//   - 入力値に対応してtrueが取得できること
func TestAlmostEqual01(t *testing.T) {
	// テスト用入力パラメータ
	x := 5.0
	y := 5.0
	absTol := 0.1

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := AlmostEqual(x, y, absTol)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値が期待値と異なる場合Errorをログに出力
		t.Errorf("bool値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TTestAlmostEqual02 同値確認関数 境界値確認 誤差より2つの数値の差が大きい場合
// 試験詳細：
// + 試験データ
//   - 数値1:5.0 数値2:4.8 浮動小数点誤差:0.1
//
// + 確認内容
//   - 入力値に対応してfalseが取得できること
func TestAlmostEqual02(t *testing.T) {
	// テスト用入力パラメータ
	x := 5.0
	y := 4.8
	absTol := 0.1

	// 期待値
	expectVal := false

	// テスト対象呼び出し
	resultVal := AlmostEqual(x, y, absTol)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値が期待値と異なる場合Errorをログに出力
		t.Errorf("bool値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestAlmostEqual03 同値確認関数 境界値確認 誤差と2つの数値の差が等しい場合
// 試験詳細：
// + 試験データ
//   - 数値1:5.0 数値2:4.9 浮動小数点誤差:0.1
//
// + 確認内容
//   - 入力値に対応してtrueが取得できること
func TestAlmostEqual03(t *testing.T) {
	// テスト用入力パラメータ
	x := 5.0
	y := 4.9
	absTol := 0.1

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := AlmostEqual(x, y, absTol)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値が期待値と異なる場合Errorをログに出力
		t.Errorf("bool値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestAlmostEqual04 同値確認関数 境界値確認 誤差と2つの数値の差が小さい場合
// 試験詳細：
// + 試験データ
//   - 数値1:5.0 数値2:4.9 浮動小数点誤差:0.2
//
// + 確認内容
//   - 入力値に対応してtrueが取得できること
func TestAlmostEqual04(t *testing.T) {
	// テスト用入力パラメータ
	x := 5.0
	y := 4.9
	absTol := 0.2

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := AlmostEqual(x, y, absTol)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値が期待値と異なる場合Errorをログに出力
		t.Errorf("bool値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestAlmostEqual05 同値確認関数 比較時の浮動小数点誤差が、"absTol"の値と等しい
// 試験詳細：
// + 試験データ
//   - 数値1:1.0 数値2:1.1 浮動小数点誤差:1.0000000000000009e-1
//
// + 確認内容
//   - 入力値に対応してtrueが取得できること
func TestAlmostEqual05(t *testing.T) {
	// テスト用入力パラメータ
	x := 1.0
	y := 1.1
	absTol := 1.0000000000000009e-1

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := AlmostEqual(x, y, absTol)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値が期待値と異なる場合Errorをログに出力
		t.Errorf("bool値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestAlmostEqual06 同値確認関数 比較時の浮動小数点誤差が、"absTol"の値未満
// 試験詳細：
// + 試験データ
//   - 数値1:1.0 数値2:1.1 浮動小数点誤差:1.000000000000001e-1
//
// + 確認内容
//   - 入力値に対応してtrueが取得できること
func TestAlmostEqual06(t *testing.T) {
	// テスト用入力パラメータ
	x := 1.0
	y := 1.1
	absTol := 1.000000000000001e-1

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := AlmostEqual(x, y, absTol)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値が期待値と異なる場合Errorをログに出力
		t.Errorf("bool値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestAlmostEqual07 同値確認関数 比較時の浮動小数点誤差が、"absTol"の値を超える
// 試験詳細：
// + 試験データ
//   - 数値1:1.0 数値2:1.1 浮動小数点誤差:1.0000000000000008e-1
//
// + 確認内容
//   - 入力値に対応してtrueが取得できること
func TestAlmostEqual07(t *testing.T) {
	// テスト用入力パラメータ
	x := 1.0
	y := 1.1
	absTol := 1.0000000000000008e-1

	// 期待値
	expectVal := false

	// テスト対象呼び出し
	resultVal := AlmostEqual(x, y, absTol)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値が期待値と異なる場合Errorをログに出力
		t.Errorf("bool値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestMax01 最大値取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T := {1, 2, 2, 3, 2, 5, 1, 9, 8}
//
// + 確認内容
//   - 入力値のスライスから最大値(9)が取得できること
func TestMax01(t *testing.T) {
	// テスト用入力パラメータ
	T := []int{1, 2, 2, 3, 2, 5, 1, 9, 8}

	// 期待値
	expectVal := 9

	// テスト対象呼び出し
	resultVal, resultErr := Max(T)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最大値が期待値と異なる場合Errorをログに出力
		t.Errorf("最大値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMax02 最大値取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T := {2}
//
// + 確認内容
//   - 入力値のスライスから最大値(2)が取得できること
func TestMax02(t *testing.T) {
	// テスト用入力パラメータ
	T := []int{2}

	// 期待値
	expectVal := 2

	// テスト対象呼び出し
	resultVal, resultErr := Max(T)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最大値が期待値と異なる場合Errorをログに出力
		t.Errorf("最大値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMax03 最大値取得関数 引数に値が存在しない場合
// 試験詳細：
// + 試験データ
//   - 数値T := {} (値なし)
//
// + 確認内容
//   - 入力不備で入力チェックエラーとなること
func TestMax03(t *testing.T) {
	// テスト用入力パラメータ
	T := []int{}

	// 期待値
	expectVal := 0
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := Max(T)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最大値が期待値と異なる場合Errorをログに出力
		t.Errorf("最大値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestMin01 最小値取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T := {2, 2, 3, 1, 5, 2}
//
// + 確認内容
//   - 入力値のスライスから最小値(1)が取得できること
func TestMin01(t *testing.T) {
	// テスト用入力パラメータ
	T := []int{2, 2, 3, 1, 5, 2}

	// 期待値
	expectVal := 1

	// テスト対象呼び出し
	resultVal, resultErr := Min(T)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最小値が期待値と異なる場合Errorをログに出力
		t.Errorf("最小値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMin02 最小値取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T := {2}
//
// + 確認内容
//   - 入力値のスライスから最小値(2)が取得できること
func TestMin02(t *testing.T) {
	// テスト用入力パラメータ
	T := []int{2}

	// 期待値
	expectVal := 2

	// テスト対象呼び出し
	resultVal, resultErr := Min(T)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最小値が期待値と異なる場合Errorをログに出力
		t.Errorf("最小値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestMin03 最小値取得関数 引数に値が存在しない場合
// 試験詳細：
// + 試験データ
//   - 数値T := {} (値なし)
//
// + 確認内容
//   - 入力不備で入力チェックエラーとなること
func TestMin03(t *testing.T) {
	// テスト用入力パラメータ
	T := []int{}

	// 期待値
	expectVal := 0
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := Min(T)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最小値が期待値と異なる場合Errorをログに出力
		t.Errorf("最小値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestDegreeToRadian01 ラジアン取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T := 40.11
//
// + 確認内容
//   - 入力値からラジアンが取得できること
func TestDegreeToRadian01(t *testing.T) {
	// テスト用入力パラメータ
	T := 40.11

	// 期待値
	expectVal := 0.7000515629749255

	// テスト対象呼び出し
	resultVal := DegreeToRadian(T)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最小値が期待値と異なる場合Errorをログに出力
		t.Errorf("最小値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestRadianToDegree01 角度取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T := 0.7000515629749255
//
// + 確認内容
//   - 入力値から角度が取得できること
func TestRadianToDegree01(t *testing.T) {
	// テスト用入力パラメータ
	T := 0.7000515629749255

	// 期待値
	expectVal := 40.11

	// テスト対象呼び出し
	resultVal := RadianToDegree(T)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最小値が期待値と異なる場合Errorをログに出力
		t.Errorf("最小値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestUnion01 和集合取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T1 := {1, 2, 3, 5}, 数値T2 := {0, 3, 4, 8}
//
// + 確認内容
//   - 入力値のスライスから和集合スライスが取得できること
func TestUnion01(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{1, 2, 3, 5}
	T2 := []int{0, 3, 4, 8}

	// 期待値
	expectVal := []int{0, 1, 2, 3, 4, 5, 8}

	// テスト対象呼び出し
	resultVal := Union(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestUnion02 和集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {(空入力)}, 数値T2 := {0, 3, 4, 8}
//
// + 確認内容
//   - 数値T2の値を取得できること
func TestUnion02(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{}
	T2 := []int{0, 3, 4, 8}

	// 期待値
	expectVal := []int{0, 3, 4, 8}

	// テスト対象呼び出し
	resultVal := Union(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestUnion03 和集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {1, 2, 3, 5}, 数値T2 := {(空入力)}
//
// + 確認内容
//   - 数値T1の値を取得できること
func TestUnion03(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{1, 2, 3, 5}
	T2 := []int{}

	// 期待値
	expectVal := []int{1, 2, 3, 5}

	// テスト対象呼び出し
	resultVal := Union(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestUnion04 和集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {(空入力)}, 数値T2 := {(空入力)}
//
// + 確認内容
//   - 値が空のまま返すこと
func TestUnion04(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{}
	T2 := []int{}

	// 期待値
	expectVal := []int{}

	// テスト対象呼び出し
	resultVal := Union(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestDifference01 差集合取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T1 := {1, 2, 3, 5, 10}, 数値T2 := {0, 3, 4, 8, 10}
//
// + 確認内容
//   - 入力値のスライスから差集合スライスが取得できること
func TestDifference01(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{1, 2, 3, 5, 10}
	T2 := []int{0, 3, 4, 8, 10}

	// 期待値
	expectVal := []int{1, 2, 5}

	// テスト対象呼び出し
	resultVal := Difference(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestDifference02 差集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {(空入力)}, 数値T2 := {0, 3, 4, 8, 10}
//
// + 確認内容
//   - 値が空のまま返すこと
func TestDifference02(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{}
	T2 := []int{0, 3, 4, 8, 10}

	// 期待値
	expectVal := []int{}

	// テスト対象呼び出し
	resultVal := Difference(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestDifference03 差集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {1, 2, 3, 5, 10}, 数値T2 := {(空入力)}
//
// + 確認内容
//   - 数値T1の値を取得できること
func TestDifference03(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{1, 2, 3, 5, 10}
	T2 := []int{}

	// 期待値
	expectVal := []int{1, 2, 3, 5, 10}

	// テスト対象呼び出し
	resultVal := Difference(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestDifference04 差集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {(空入力)}, 数値T2 := {(空入力)}
//
// + 確認内容
//   - 値が空のまま返すこと
func TestDifference04(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{}
	T2 := []int{}

	// 期待値
	expectVal := []int{}

	// テスト対象呼び出し
	resultVal := Difference(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestIntersect01 積集合取得関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T1 := {1, 2, 3, 5, 10}, 数値T2 := {1, 3, 4, 8, 10}
//
// + 確認内容
//   - 入力値のスライスから積集合スライスが取得できること
func TestIntersect01(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{1, 2, 3, 5, 10}
	T2 := []int{1, 3, 4, 8, 10}

	// 期待値
	expectVal := []int{1, 3, 10}

	// テスト対象呼び出し
	resultVal := Intersect(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestIntersect02 積集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {(空入力)}, 数値T2 := {1, 3, 4, 8, 10}
//
// + 確認内容
//   - 値が空のまま返すこと
func TestIntersect02(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{}
	T2 := []int{1, 3, 4, 8, 10}

	// 期待値
	expectVal := []int{}

	// テスト対象呼び出し
	resultVal := Intersect(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestIntersect03 積集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {1, 2, 3, 5, 10}, 数値T2 := {(空入力)}
//
// + 確認内容
//   - 値が空のまま返すこと
func TestIntersect03(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{1, 2, 3, 5, 10}
	T2 := []int{}

	// 期待値
	expectVal := []int{}

	// テスト対象呼び出し
	resultVal := Intersect(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestIntersect04 積集合取得関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {(空入力)}, 数値T2 := {(空入力)}
//
// + 確認内容
//   - 値が空のまま返すこと
func TestIntersect04(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{}
	T2 := []int{}

	// 期待値
	expectVal := []int{}

	// テスト対象呼び出し
	resultVal := Intersect(T1, T2)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestUnique01 ユニーク化関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T := {1, 2, 3, 5, 10, 3, 12, 5, 5, 2}
//
// + 確認内容
//   - 入力値のスライスからユニークスライスが取得できること
func TestUnique01(t *testing.T) {
	// テスト用入力パラメータ
	T := []int{1, 2, 3, 5, 10, 3, 12, 5, 5, 2}

	// 期待値
	expectVal := []int{1, 2, 3, 5, 10, 12}

	// テスト対象呼び出し
	resultVal := Unique(T)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestUnique02 ユニーク化関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T := {(空入力)}
//
// + 確認内容
//   - 値が空のまま返すこと
func TestUnique02(t *testing.T) {
	// テスト用入力パラメータ
	T := []int{}

	// 期待値
	expectVal := []int{}

	// テスト対象呼び出し
	resultVal := Unique(T)

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値と期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("要素 - 期待値：%v, 取得値：%v", expectVal, resultVal)
			break
		}
	}

	t.Log("テスト終了")
}

// TestInclude01 内包関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T1 := {1, 2, 3, 5, 10}, 数値T2 := 1
//
// + 確認内容
//   - 入力値のスライスに対象が含まれる場合、trueを返すこと。
func TestInclude01(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{1, 2, 3, 5, 10}
	T2 := 1

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := Include(T1, T2)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最小値が期待値と異なる場合Errorをログに出力
		t.Errorf("期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestInclude02 内包関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値T1 := {1, 2, 3, 5, 10}, 数値T2 := 7
//
// + 確認内容
//   - 入力値のスライスに対象が含まれていない場合、falseを返すこと。
func TestInclude02(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{1, 2, 3, 5, 10}
	T2 := 7

	// 期待値
	expectVal := false

	// テスト対象呼び出し
	resultVal := Include(T1, T2)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最小値が期待値と異なる場合Errorをログに出力
		t.Errorf("期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestInclude03 内包関数 空入力時動作
// 試験詳細：
// + 試験データ
//   - 数値T1 := {(空入力)}, 数値T2 := 0
//
// + 確認内容
//   - 入力値のスライスに対象が含まれていない場合、falseを返すこと。
func TestInclude03(t *testing.T) {
	// テスト用入力パラメータ
	T1 := []int{}
	T2 := 0

	// 期待値
	expectVal := false

	// テスト対象呼び出し
	resultVal := Include(T1, T2)

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の最小値が期待値と異なる場合Errorをログに出力
		t.Errorf("期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestCombinations01 組み合わせ関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - 数値Tn := 5(組み合わせる数(位数)), 値Tn := 2(選択数), 数値A{1,2,3,4,5}
//
// + 確認内容
//   - 引数に入力された組み合わせである、数値Aの各要素2個を足した値の組み合わせを返却する
func TestCombinations01(t *testing.T) {
	// テスト用入力パラメータ
	var Tn int64 = 5
	var Tk int64 = 2
	A := []int64{1, 2, 3, 4, 5}

	// 期待値
	expectVal := []int64{3, 4, 5, 6, 5, 6, 7, 7, 8, 9}

	//返り値格納用変数
	resultVal := []int64{}
	// テスト対象呼び出し
	Combinations(Tn, Tk, func(pair []int64) {
		start := A[pair[0]]
		end := A[pair[1]]
		line := start + end
		resultVal = append(resultVal, line)
	})

	// 戻り値と期待値の比較

	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値が期待値と異なる場合Errorをログに出力
		t.Errorf("期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// intスライスの中に指定整数を含むか判定する
//
// 引数：
//
//	slice： intスライス
//	target： 検索整数
//
// 戻り値：
//
//	含む場合：true
//	含まない場合：false
func contains(slice []int, target int) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

func TestCalculateArithmaticShift(t *testing.T) {
	data := []struct {
		index          int64
		shift          int64
		expectedOutput int64
	}{
		{index: 1, shift: 25, expectedOutput: 33554432},
		{index: 1, shift: 23, expectedOutput: 8388608},
		{index: 1, shift: 0, expectedOutput: 1},
		{index: 47, shift: 0, expectedOutput: 47},
		{index: 47, shift: -1, expectedOutput: 23},
		{index: 47, shift: -100, expectedOutput: 0},
		{index: -47, shift: 1, expectedOutput: -94},
		{index: -47, shift: 0, expectedOutput: -47},
		{index: -47, shift: -1, expectedOutput: -24},
		{index: -47, shift: -3, expectedOutput: -6},
		{index: 123456, shift: -1, expectedOutput: 61728},
	}

	for _, p := range data {
		result := CalculateArithmeticShift(p.index, p.shift)
		if result != p.expectedOutput {
			t.Log(t.Name())
			t.Errorf("calculateMinVerticalIndex(%v, %v) == %v, result: %v", p.index, p.shift, p.expectedOutput, result)
		}
	}
}

package spatialID

import (
	"reflect"
	"testing"
)

// TestNewExtendedSpatialID01 拡張空間ID初期化関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestNewExtendedSpatialID01(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"18/232837/103222/0",
		&SpatialID{18, 232837, 103222, 0},
		nil,
	)
}

// TestNewExtendedSpatialID02 拡張空間ID初期化関数 異常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222"
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestNewExtendedSpatialID02(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"18/232837/103222",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID03 拡張空間ID再設定関数 入力された拡張空間IDの水平精度がしきい値を超過していた場合(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："36/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
// 備考：
//  対象関数内では精度の値で判定が変わらないため、値が返せていることを確認。
func TestResetExtendedSpatialID03(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"36/232837/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID04 拡張空間ID再設定関数 入力された拡張空間IDの水平精度がしきい値より小さい場合(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："-1/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
// 備考：
//  対象関数内では精度の値で判定が変わらないため、値が返せていることを確認。
func TestResetExtendedSpatialID04(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"-1/232837/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID05 拡張空間ID再設定関数 正常系動作確認(水平精度境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："35/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID05(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"35/232837/103222/0",
		&SpatialID{35, 232837, 103222, 25},
		nil,
	)
}

// TestResetExtendedSpatialID06 拡張空間ID再設定関数 正常系動作確認(水平精度境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："0/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID06(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"0/232837/103222/0",
		&SpatialID{0, 0, 0, 0},
		nil,
	)
}

// TestResetExtendedSpatialID07 拡張空間ID再設定関数 入力された拡張空間IDの垂直精度がしきい値を超過していた場合(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："25/232837/103222/36/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
// 備考：
//  対象関数内では精度の値で判定が変わらないため、値が返せていることを確認。
func TestResetExtendedSpatialID07(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"36/232837/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID08 拡張空間ID再設定関数 入力された拡張空間IDの垂直精度がしきい値より小さい場合(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："25/232837/103222/-1/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
// 備考：
//  対象関数内では精度の値で判定が変わらないため、値が返せていることを確認。
func TestResetExtendedSpatialID08(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"-1/232837/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

// TestResetExtendedSpatialID09 拡張空間ID再設定関数 正常系動作確認(垂直精度境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/35/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID09(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"35/232837/103222/0",
		&SpatialID{35, 232837, 103222, 0},
		nil,
	)
}

// TestResetExtendedSpatialID10 拡張空間ID再設定関数 正常系動作確認(垂直精度境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/0/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID10(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"0/0/0/-1",
		&SpatialID{0, 0, 0, -1},
		nil,
	)
}

// TestResetExtendedSpatialID11 拡張空間ID再設定関数 入力値に整数以外が存在した場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/2328A7/103222/0/0"
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestResetExtendedSpatialID11(t *testing.T) {
	testNewSpatialIDFromString(
		t,
		"18/2328A7/103222/0",
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
	)
}

func testNewSpatialIDFromString(
	t *testing.T,
	spatialIDString string,
	expectedSpatialID *SpatialID,
	expectedError error,
	) {
	spatialID, error := NewSpatialIDFromString(spatialIDString)

	// 始点から終点へのベクトルと期待値の比較
	if !reflect.DeepEqual(spatialID, expectedSpatialID) {
		t.Errorf("空間IDオブジェクト - 期待値：%v, 取得値：%v", expectedSpatialID, spatialID)
	}
	if !reflect.DeepEqual(error, expectedError) {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectedError.Error, error.Error())
	}
}

// TestSetX01 位置設定関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     格納する整数：1
//
// + 確認内容
//   - 入力値を拡張空間IDオブジェクトに格納できること
func TestSetX01(t *testing.T) {
	expected := &SpatialID{1, 0, 1, 0}

	result, error := NewSpatialID(1, 0, 0, 0)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result.SetX(1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("空間IDオブジェクト - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestSetY01 緯度ID設定関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     格納する整数：1
//
// + 確認内容
//   - 入力値を拡張空間IDオブジェクトに格納できること
func TestSetY01(t *testing.T) {
	expected := &SpatialID{1, 0, 0, 1}

	result, error := NewSpatialID(1, 0, 0, 0)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result.SetY(1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("空間IDオブジェクト - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestSetZ01 高さID設定関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     格納する整数：1
//
// + 確認内容
//   - 入力値を拡張空間IDオブジェクトに格納できること
func TestSetZ01(t *testing.T) {
	expected := &SpatialID{1, 1, 0, 0}

	result, error := NewSpatialID(1, 0, 0, 0)
	if error != nil {
		t.Errorf("空間IDオブジェクト - 期待値：nil, 取得値：%v", error)
	}

	result.SetF(1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("空間IDオブジェクト - 期待値：%v, 取得値：%v", expected, result)
	}
}

// TestX01 経度ID取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、経度の値を取得できること。
func TestX01(t *testing.T) {
	//入力値
	EXSpatialID := ExtendedSpatialID{1, 2, 3, 4, 5}
	resultVal := EXSpatialID.X()

	//期待値
	var expectVal int64 = 2

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("経度 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestY01 緯度ID取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、緯度の値を取得できること。
func TestY01(t *testing.T) {
	//入力値
	EXSpatialID := ExtendedSpatialID{1, 2, 3, 4, 5}
	resultVal := EXSpatialID.Y()

	//期待値
	var expectVal int64 = 3

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("緯度 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestZ01 高さID取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、高さの値を取得できること。
func TestZ01(t *testing.T) {
	//入力値
	EXSpatialID := ExtendedSpatialID{1, 2, 3, 4, 5}
	resultVal := EXSpatialID.Z()

	//期待値
	var expectVal int64 = 5

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("高さ - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestHZoom01 水平精度取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、水平精度の値を取得できること。
func TestHZoom01(t *testing.T) {
	//入力値
	EXSpatialID := ExtendedSpatialID{1, 2, 3, 4, 5}
	resultVal := EXSpatialID.HZoom()

	//期待値
	var expectVal int64 = 1

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("水平精度 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestVZoom01 垂直精度取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、垂直精度の値を取得できること。
func TestVZoom01(t *testing.T) {
	//入力値
	EXSpatialID := ExtendedSpatialID{1, 2, 3, 4, 5}
	resultVal := EXSpatialID.VZoom()

	//期待値
	var expectVal int64 = 4

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("垂直精度 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestFieldParams01 拡張空間ID成分取得関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDオブジェクト：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に含まれる水平、垂直精度、X, Y, Z成分をint64配列で格納できること
func TestFieldParams01(t *testing.T) {
	//入力値
	object := &ExtendedSpatialID{1, 2, 3, 4, 5}
	resultVal := object.FieldParams()

	//期待値
	expectVal := []int64{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間ID配列 - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestID01 空間ID文字列返却関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{1, 2, 3, 4, 5}
//
// + 確認内容
//   - 入力値に対して、拡張空間ID文字列を取得できること。
func TestID01(t *testing.T) {
	//入力値
	EXSpatialID := ExtendedSpatialID{1, 2, 3, 4, 5}
	resultVal := EXSpatialID.ID()

	//期待値
	expectVal := "1/2/3/4/5"

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestHigher01 最適化後拡張空間ID化関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間IDクラス：{20, 1024, 2048, 20, 1024}
//
// + 確認内容
//   - 入力値に対して、拡張空間IDを取得できること。
func TestHigher01(t *testing.T) {
	//入力値
	EXSpatialID := ExtendedSpatialID{20, 1024, 2048, 20, 1024}
	var hDiff int64 = 3
	var vDiff int64 = 5
	resultVal := EXSpatialID.Higher(hDiff, vDiff)

	//期待値
	expectVal := &ExtendedSpatialID{17, 128, 256, 15, 32}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("最適化後拡張空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

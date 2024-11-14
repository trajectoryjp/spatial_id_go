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
//     拡張空間ID："18/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestNewExtendedSpatialID01(t *testing.T) {
	//入力値
	id := "18/232837/103222/25/0"
	resultVal, resultErr := NewExtendedSpatialID(id)

	//期待値
	expectVal := &ExtendedSpatialID{18, 232837, 103222, 25, 0}

	// 始点から終点へのベクトルと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestNewExtendedSpatialID02 拡張空間ID初期化関数 異常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/25"
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestNewExtendedSpatialID02(t *testing.T) {
	//入力値
	id := "18/232837/103222/25"
	resultVal, resultErr := NewExtendedSpatialID(id)

	//期待値
	expectVal := &ExtendedSpatialID{0, 0, 0, 0, 0}
	expectErr := "InputValueError,入力チェックエラー"

	// 始点から終点へのベクトルと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestResetExtendedSpatialID01 拡張空間ID再設定関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/25/0"
//
// + 確認内容
//   - 入力値から初期化した拡張空間IDオブジェクトを取得できること
func TestResetExtendedSpatialID01(t *testing.T) {
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("18/232837/103222/25/0")

	//期待値
	expectVal := &ExtendedSpatialID{18, 232837, 103222, 25, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
}

// TestResetExtendedSpatialID02 拡張空間ID再設定関数 引数にExtendedSpatialIDのフォーマットに違反した値が入力された場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     拡張空間ID："18/232837/103222/25"
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestResetExtendedSpatialID02(t *testing.T) {
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("18/232837/103222/25")

	//期待値
	expectVal := &ExtendedSpatialID{0, 0, 0, 0, 0}
	expectErr := "InputValueError,入力チェックエラー"

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("36/232837/103222/25/0")

	//期待値
	expectVal := &ExtendedSpatialID{36, 232837, 103222, 25, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("-1/232837/103222/25/0")

	//期待値
	expectVal := &ExtendedSpatialID{-1, 232837, 103222, 25, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("35/232837/103222/25/0")

	//期待値
	expectVal := &ExtendedSpatialID{35, 232837, 103222, 25, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("0/232837/103222/25/0")

	//期待値
	expectVal := &ExtendedSpatialID{0, 232837, 103222, 25, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("25/232837/103222/36/0")

	expectVal := &ExtendedSpatialID{25, 232837, 103222, 36, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("25/232837/103222/-1/0")

	expectVal := &ExtendedSpatialID{25, 232837, 103222, -1, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("18/232837/103222/35/0")

	//期待値
	expectVal := &ExtendedSpatialID{18, 232837, 103222, 35, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("18/232837/103222/0/0")

	//期待値
	expectVal := &ExtendedSpatialID{18, 232837, 103222, 0, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%+v", resultErr)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{}
	resultErr := resultVal.ResetExtendedSpatialID("18/2328A7/103222/0/0")

	//期待値
	expectVal := &ExtendedSpatialID{0, 0, 0, 0, 0}
	expectErr := "InputValueError,入力チェックエラー"

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{0, 0, 0, 0, 0}
	var x int64 = 1
	resultVal.SetX(x)

	//期待値
	expectVal := &ExtendedSpatialID{0, 1, 0, 0, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{0, 0, 0, 0, 0}
	var y int64 = 1
	resultVal.SetY(y)

	//期待値
	expectVal := &ExtendedSpatialID{0, 0, 1, 0, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
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
	//入力値
	resultVal := &ExtendedSpatialID{0, 0, 0, 0, 0}
	var z int64 = 1
	resultVal.SetZ(z)

	//期待値
	expectVal := &ExtendedSpatialID{0, 0, 0, 0, 1}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestSetZoom01 精度設定関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     水平精度に格納する整数：1, 垂直精度に格納する整数：2
//
// + 確認内容
//   - 入力値を拡張空間IDオブジェクトに格納できること
func TestSetZoom01(t *testing.T) {
	//入力値
	resultVal := &ExtendedSpatialID{0, 0, 0, 0, 0}
	var hzoom int64 = 1
	var vzoom int64 = 2
	resultVal.SetZoom(hzoom, vzoom)

	//期待値
	expectVal := &ExtendedSpatialID{1, 0, 0, 2, 0}

	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("拡張空間IDオブジェクト - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}

	t.Log("テスト終了")
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

package spatialID

import (
	"reflect"
	"testing"
)

// TestNewChangeSpatialIdsZoom01 空間IDの精度変換関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の空間ID：{"25/0/29803148/13212522"}, 変換後の精度：15
//
// + 確認内容
//   - 入力値から精度変換後の全空間IDを取得できること
func TestChangeSpatialIdsZoom01(t *testing.T) {
	//入力値
	SpatialIds := []string{"25/0/29803148/13212522"}
	var zoom int64 = 15
	resultVal, resultErr := ChangeSpatialIdsZoom(SpatialIds, zoom)

	//期待値
	expectVal := []string{"15/0/29104/12902"}

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestNewChangeSpatialIdsZoom02 空間IDの精度変換関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の空間ID：{"14/0/1024/2048", "15/0/1024/2048", "16/0/1024/2048"},
//     変換後の精度：15
//
// + 確認内容
//   - 入力値から精度変換後の全空間IDを取得できること
func TestChangeSpatialIdsZoom02(t *testing.T) {
	//入力値
	SpatialIds := []string{"14/0/1024/2048", "15/0/1024/2048", "16/0/1024/2048"}
	var zoom int64 = 15
	resultVal, resultErr := ChangeSpatialIdsZoom(SpatialIds, zoom)

	//期待値
	expectVal := []string{"15/0/2048/4096", "15/1/2048/4096", "15/0/2049/4096", "15/1/2049/4096",
		"15/0/2048/4097", "15/1/2048/4097", "15/0/2049/4097", "15/1/2049/4097", //"14/0/1024/2048"変換後
		"15/0/1024/2048", //"15/0/1024/2048"変換後
		"15/0/512/1024"}  //"16/0/1024/2048"変換後

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestNewChangeSpatialIdsZoom03 空間IDの精度変換関数 空入力時動作確認
//
// 試験詳細：
//   - パターン1：
//     精度変換対象の空間ID：{(空入力)}, 変換後の精度：15
//
// + 確認内容
//   - 空配列を取得できること
func TestChangeSpatialIdsZoom03(t *testing.T) {
	//入力値
	SpatialIds := []string{}
	var zoom int64 = 15
	resultVal, resultErr := ChangeSpatialIdsZoom(SpatialIds, zoom)

	//期待値
	expectVal := []string{}

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestNewChangeSpatialIdsZoom04 空間IDの精度変換関数 精度閾値超過
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の空間ID：{"25/0/29803148/13212522"},
//     変換後の精度：36
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestChangeSpatialIdsZoom04(t *testing.T) {
	//入力値
	SpatialIds := []string{"25/0/29803148/13212522"}
	var zoom int64 = 36
	resultVal, resultErr := ChangeSpatialIdsZoom(SpatialIds, zoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestNewChangeSpatialIdsZoom05 空間IDの精度変換関数 空間IDフォーマット不正
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の空間ID：{"25/0/29803148/13212522/777"},
//     変換後の精度：25
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestChangeSpatialIdsZoom05(t *testing.T) {
	//入力値
	SpatialIds := []string{"25/0/29803148/13212522/777"}
	var zoom int64 = 25
	resultVal, resultErr := ChangeSpatialIdsZoom(SpatialIds, zoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom01 拡張空間IDの精度変換関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"25/29803148/13212522/25/0"}, 変換後の水平方向精度：15, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 入力値から全拡張空間IDを格納した配列を取得できること
func TestChangeExtendedSpatialIdsZoom01(t *testing.T) {
	//入力値
	SpatialIds := []string{"25/29803148/13212522/25/0"}
	var hzoom int64 = 15
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{"15/29104/12902/15/0"}

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom02 拡張空間IDの精度変換関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"14/1024/2048/14/0", "15/1024/2048/15/0", "16/1024/2048/16/0"},
//     変換後の水平方向精度：15, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 入力値から全拡張空間IDを格納した配列を取得できること
func TestChangeExtendedSpatialIdsZoom02(t *testing.T) {
	//入力値
	SpatialIds := []string{"14/1024/2048/14/0", "15/1024/2048/15/0", "16/1024/2048/16/0"}
	var hzoom int64 = 15
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{"15/2048/4096/15/0", "15/2048/4096/15/1", "15/2049/4096/15/0",
		"15/2049/4096/15/1", "15/2048/4097/15/0", "15/2048/4097/15/1", "15/2049/4097/15/0", "15/2049/4097/15/1",
		"15/1024/2048/15/0",
		"15/512/1024/15/0"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom03 拡張空間IDの精度変換関数 空入力時動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{(空入力)}, 変換後の水平方向精度：15, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 空配列を取得できること
func TestChangeExtendedSpatialIdsZoom03(t *testing.T) {
	//入力値
	SpatialIds := []string{}
	var hzoom int64 = 15
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{}

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom04 拡張空間IDの精度変換関数 水平精度閾値超過(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"25/29803148/13212522/25/0"}, 変換後の水平方向精度：36, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestChangeExtendedSpatialIdsZoom04(t *testing.T) {
	//入力値
	SpatialIds := []string{"25/29803148/13212522/25/0"}
	var hzoom int64 = 36
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom05 拡張空間IDの精度変換関数 水平精度閾値より小さい値(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"25/29803148/13212522/25/0"}, 変換後の水平方向精度：-1, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestChangeExtendedSpatialIdsZoom05(t *testing.T) {
	//入力値
	SpatialIds := []string{"25/29803148/13212522/25/0"}
	var hzoom int64 = -1
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom06 拡張空間IDの精度変換関数 正常動作確認(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"34/1024/1024/34/0"}, 変換後の水平方向精度：35, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 入力値から全拡張空間IDを格納した配列を取得できること
func TestChangeExtendedSpatialIdsZoom06(t *testing.T) {
	//入力値
	SpatialIds := []string{"34/1024/1024/34/0"}
	var hzoom int64 = 35
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{"35/2048/2048/15/0", "35/2049/2048/15/0", "35/2048/2049/15/0", "35/2049/2049/15/0"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom07 拡張空間IDの精度変換関数 正常動作確認(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"34/1024/1024/34/0"}, 変換後の水平方向精度：0, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 入力値から全拡張空間IDを格納した配列を取得できること
func TestChangeExtendedSpatialIdsZoom07(t *testing.T) {
	//入力値
	SpatialIds := []string{"34/1024/1024/34/0"}
	var hzoom int64 = 0
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{"0/0/0/15/0"}

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom08 拡張空間IDの精度変換関数 垂直精度閾値超過(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"25/29803148/13212522/25/0"}, 変換後の水平方向精度：15, 変換後の垂直方向精度：36
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestChangeExtendedSpatialIdsZoom08(t *testing.T) {
	//入力値
	SpatialIds := []string{"25/29803148/13212522/25/0"}
	var hzoom int64 = 15
	var vzoom int64 = 36
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom09 拡張空間IDの精度変換関数 垂直精度閾値より小さい値(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"25/29803148/13212522/25/0"}, 変換後の水平方向精度：15, 変換後の垂直方向精度：-1
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestChangeExtendedSpatialIdsZoom09(t *testing.T) {
	//入力値
	SpatialIds := []string{"25/29803148/13212522/25/0"}
	var hzoom int64 = 15
	var vzoom int64 = -1
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom10 拡張空間IDの精度変換関数 正常動作確認(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"34/1024/1024/34/0"}, 変換後の水平方向精度：15, 変換後の垂直方向精度：35
//
// + 確認内容
//   - 入力値から全拡張空間IDを格納した配列を取得できること
func TestChangeExtendedSpatialIdsZoom10(t *testing.T) {
	//入力値
	SpatialIds := []string{"34/1024/1024/34/0"}
	var hzoom int64 = 15
	var vzoom int64 = 35
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{"15/0/0/35/0", "15/0/0/35/1"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom11 拡張空間IDの精度変換関数 正常動作確認(境界値)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"34/1024/1024/34/0"}, 変換後の水平方向精度：15, 変換後の垂直方向精度：0
//
// + 確認内容
//   - 入力値から全拡張空間IDを格納した配列を取得できること
func TestChangeExtendedSpatialIdsZoom11(t *testing.T) {
	//入力値
	SpatialIds := []string{"34/1024/1024/34/0"}
	var hzoom int64 = 15
	var vzoom int64 = 0
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{"15/0/0/0/0"}

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom12 拡張空間IDの精度変換関数 区切り文字数がフォーマットに従っていない場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"34/1024/1024/34"}, 変換後の水平方向精度：15, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestChangeExtendedSpatialIdsZoom12(t *testing.T) {
	//入力値
	SpatialIds := []string{"34/1024/1024/34"}
	var hzoom int64 = 15
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestChangeExtendedSpatialIdsZoom13 拡張空間IDの精度変換関数 int64変換時にエラーが発生した場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間ID：{"34/1024A/1024/34/0"}, 変換後の水平方向精度：15, 変換後の垂直方向精度：15
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestChangeExtendedSpatialIdsZoom13(t *testing.T) {
	//入力値
	SpatialIds := []string{"34/1024A/1024/34/0"}
	var hzoom int64 = 15
	var vzoom int64 = 15
	resultVal, resultErr := ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)

	//期待値
	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestHorizontalZoom01  拡張空間IDの水平方向の精度変換関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間IDの水平方向精度：25, 拡張空間IDのxIndex成分：1024,
//     拡張空間IDのyIndex成分：1024, 変換後の水平方向精度：26
//
// + 確認内容
//   - 入力値から精度変換後の全拡張空間IDの水平方向成分を格納したスライスを取得できること
func TestHorizontalZoom01(t *testing.T) {
	//入力値
	var inputZoom int64 = 25
	var xIndex int64 = 1024
	var yIndex int64 = 1024
	var outputZoom int64 = 26
	resultVal := HorizontalZoom(inputZoom, xIndex, yIndex, outputZoom)

	//期待値
	expectVal := []string{"26/2048/2048", "26/2049/2048", "26/2049/2048", "26/2049/2049"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間IDの水平方向成分スライス - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間IDの水平方向成分スライス - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}

	t.Log("テスト終了")
}

// TestHorizontalZoom02  拡張空間IDの水平方向の精度変換関数 変換後の精度が低い場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間IDの水平方向精度：25, 拡張空間IDのxIndex成分：1024,
//     拡張空間IDのyIndex成分：1024, 変換後の水平方向精度：24
//
// + 確認内容
//   - 入力値から精度変換後の全拡張空間IDの水平方向成分を格納したスライスを取得できること
func TestHorizontalZoom02(t *testing.T) {
	//入力値
	var inputZoom int64 = 25
	var xIndex int64 = 1024
	var yIndex int64 = 1024
	var outputZoom int64 = 24
	resultVal := HorizontalZoom(inputZoom, xIndex, yIndex, outputZoom)

	//期待値
	expectVal := []string{"24/512/512"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間IDの水平方向成分スライス - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間IDの水平方向成分スライス - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}

	t.Log("テスト終了")
}

// TestHorizontalZoom03  拡張空間IDの水平方向の精度変換関数 変換後の精度が変換前と等しい場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間IDの水平方向精度：25, 拡張空間IDのxIndex成分：1024,
//     拡張空間IDのyIndex成分：1024, 変換後の水平方向精度：25
//
// + 確認内容
//   - 入力値から精度変換後の全拡張空間IDの水平方向成分を格納したスライスを取得できること
func TestHorizontalZoom03(t *testing.T) {
	//入力値
	var inputZoom int64 = 25
	var xIndex int64 = 1024
	var yIndex int64 = 1024
	var outputZoom int64 = 25
	resultVal := HorizontalZoom(inputZoom, xIndex, yIndex, outputZoom)

	//期待値
	expectVal := []string{"25/1024/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間IDの水平方向成分スライス - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間IDの水平方向成分スライス - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}

	t.Log("テスト終了")
}

// TestVerticalZoom01  拡張空間IDの垂直方向の精度変換関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間IDの垂直方向精度：25, 拡張空間IDのvIndex成分：1024,
//     変換後の垂直方向精度：26
//
// + 確認内容
//   - 入力値から精度変換後の全拡張空間IDの垂直方向成分を格納したスライスを取得できること
func TestVerticalZoom01(t *testing.T) {
	//入力値
	var inputZoom int64 = 25
	var vIndex int64 = 1024
	var outputZoom int64 = 26
	resultVal := VerticalZoom(inputZoom, vIndex, outputZoom)

	//期待値
	expectVal := []string{"26/2048", "26/2049"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間IDの垂直方向成分スライス - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間IDの垂直方向成分スライス - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}

	t.Log("テスト終了")
}

// TestVerticalZoom02  拡張空間IDの垂直方向の精度変換関数 変換後の垂直方向精度が低い場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間IDの垂直方向精度：25, 拡張空間IDのvIndex成分：1024,
//     変換後の垂直方向精度：24
//
// + 確認内容
//   - 入力値から精度変換後の全拡張空間IDの垂直方向成分を格納したスライスを取得できること
func TestVerticalZoom02(t *testing.T) {
	//入力値
	var inputZoom int64 = 25
	var vIndex int64 = 1024
	var outputZoom int64 = 24
	resultVal := VerticalZoom(inputZoom, vIndex, outputZoom)

	//期待値
	expectVal := []string{"24/512"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間IDの垂直方向成分スライス - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間IDの垂直方向成分スライス - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}

	t.Log("テスト終了")
}

// TestVerticalZoom03  拡張空間IDの垂直方向の精度変換関数 変換後の精度が変換前と等しい場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     精度変換対象の拡張空間IDの垂直方向精度：25, 拡張空間IDのvIndex成分：1024,
//     変換後の垂直方向精度：25
//
// + 確認内容
//   - 入力値から精度変換後の全拡張空間IDの垂直方向成分を格納したスライスを取得できること
func TestVerticalZoom03(t *testing.T) {
	//入力値
	var inputZoom int64 = 25
	var vIndex int64 = 1024
	var outputZoom int64 = 25
	resultVal := VerticalZoom(inputZoom, vIndex, outputZoom)

	//期待値
	expectVal := []string{"25/1024"}

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("拡張空間IDの垂直方向成分スライス - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("拡張空間IDの垂直方向成分スライス - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}

	t.Log("テスト終了")
}

// stringスライスの中に指定文字列を含むか判定する
//
// 引数：
//
//	slice： stringスライス
//	target： 検索文字列
//
// 戻り値：
//
//	含む場合：true
//	含まない場合：false
func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

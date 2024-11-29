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
	testForXYF(
		t,
		[]string{"15/0/29104/12902"},
		"25/0/29803148/13212522",
		-10,
	)
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
func TestChangeSpatialIdsZoom02_01(t *testing.T) {
	testForXYF(
		t,
		[]string{
			"15/0/2048/4096",
			"15/1/2048/4096",
			"15/0/2049/4096",
			"15/1/2049/4096",
			"15/0/2048/4097",
			"15/1/2048/4097",
			"15/0/2049/4097",
			"15/1/2049/4097",
		},
		"14/0/1024/2048",
		1,
	)
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
func TestChangeSpatialIdsZoom02_02(t *testing.T) {
	testForXYF(
		t,
		[]string{"15/0/1024/2048"},
		"15/0/1024/2048",
		0,
	)
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
	expectedError := NewSpatialIdError(InputValueErrorCode, "")

	spatialID, theError := NewSpatialIDFromString("25/0/29803148/13212522")
	if theError != nil {
		t.Error(theError)
	}

	spatialIDBox, theError := NewSpatialIDBox(*spatialID, *spatialID)
	if theError != nil {
		t.Error(theError)
	}

	theError = spatialIDBox.AddZ(11)
	if !reflect.DeepEqual(theError, expectedError) {
		t.Errorf("error - 期待値：%v, 取得値：%v", expectedError, theError)
	}
}

func testForXYF(
	t *testing.T,
	expected []string,
	spatialIDString string,
	differenceZ int8,
) {
	spatialID, theError := NewSpatialIDFromString(spatialIDString)
	if theError != nil {
		t.Error(theError)
	}

	spatialIDBox, theError := NewSpatialIDBox(*spatialID, *spatialID)
	if theError != nil {
		t.Error(theError)
	}

	theError = spatialIDBox.AddZ(differenceZ)
	if theError != nil {
		t.Error(theError)
	}

	i := 0
	theError = spatialIDBox.ForXYF(func(spatialID SpatialID) error {
		if i >= len(expected) {
			t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expected), i)
		}
		if !reflect.DeepEqual(spatialID.String(), expected[i]) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expected[i], spatialID.String())
		}
		i += 1
		return nil
	})

	if theError != nil {
		t.Error(theError)
	}
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
	testTileXYZBoxAddZoomLevel(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 29104,
				y: 12902,
				z: 0,
			},
		},
		nil,

		25, 25, 29803148, 13212522, 0,

		-10, -10,
	)
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
func TestChangeExtendedSpatialIdsZoom02_01(t *testing.T) {
	testTileXYZBoxAddZoomLevel(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 2048,
				y: 4096,
				z: 0,
			},
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 2048,
				y: 4096,
				z: 1,
			},
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 2049,
				y: 4096,
				z: 0,
			},
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 2049,
				y: 4096,
				z: 1,
			},
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 2048,
				y: 4097,
				z: 0,
			},
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 2048,
				y: 4097,
				z: 1,
			},
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 2049,
				y: 4097,
				z: 0,
			},
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 2049,
				y: 4097,
				z: 1,
			},
		},
		nil,

		14, 14, 1024, 2048, 0,

		1, 1,
	)
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
func TestChangeExtendedSpatialIdsZoom02_02(t *testing.T) {
	testTileXYZBoxAddZoomLevel(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 15,
				x: 1024,
				y: 2048,
				z: 0,
			},
		},
		nil,

		15, 15, 1024, 2048, 0,

		0, 0,
	)
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
	testTileXYZBoxAddZoomLevel(
		t,

		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		25, 25, 29803148, 13212522, 0,

		11, -10,
	)
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
	testTileXYZBoxAddZoomLevel(
		t,

		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		25, 25, 29803148, 13212522, 0,

		-26, -10,
	)
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
	testTileXYZBoxAddZoomLevel(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 35,
				altitudekeyZoomLevel: 15,
				x: 2048,
				y: 2048,
				z: 0,
			},
			{
				quadkeyZoomLevel: 35,
				altitudekeyZoomLevel: 15,
				x: 2049,
				y: 2048,
				z: 0,
			},
			{
				quadkeyZoomLevel: 35,
				altitudekeyZoomLevel: 15,
				x: 2048,
				y: 2049,
				z: 0,
			},
			{
				quadkeyZoomLevel: 35,
				altitudekeyZoomLevel: 15,
				x: 2049,
				y: 2049,
				z: 0,
			},
		},
		nil,

		34, 34, 1024, 1024, 0,

		1, -19,
	)
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
	testTileXYZBoxAddZoomLevel(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: 0,
				altitudekeyZoomLevel: 15,
				x: 0,
				y: 0,
				z: 0,
			},
		},
		nil,

		34, 34, 1024, 1024, 0,

		-34, -19,
	)
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

func testTileXYZBoxAddZoomLevel(
	t *testing.T,
	expected []*TileXYZ,
	expectedError error,
	quadkeyZoomLevel int8,
	altitudeZoomLevel int8,
	x int64,
	y int64,
	z int64,
	deltaQuadkeyZoomLevel int8,
	deltaAltitudekeyZoomLevel int8,
) {
	tileXYZ, theError := NewTileXYZ(quadkeyZoomLevel, altitudeZoomLevel, x, y, z)
	if theError != nil {
		t.Error(theError)
	}

	tileXYZBox, theError := NewTileXYZBox(*tileXYZ, *tileXYZ)
	if theError != nil {
		t.Error(theError)
	}

	theError = tileXYZBox.AddZoomLevel(deltaQuadkeyZoomLevel, deltaAltitudekeyZoomLevel)
	if !reflect.DeepEqual(theError, expectedError) {
		t.Errorf("error - 期待値：%v, 取得値：%v", expectedError, theError)
	}

	i := 0
	theError = tileXYZBox.ForXYZ(func(tile TileXYZ) error {
		if i >= len(expected) {
			t.Errorf("TileXYZ - 期待要素数：%v, 取得要素数：%v", len(expected), i)
		}
		if !reflect.DeepEqual(tile, *expected[i]) {
			t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", expected[i], tile)
		}
		i += 1
		return nil
	})

	if theError != nil {
		t.Error(theError)
	}
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

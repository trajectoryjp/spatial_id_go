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
		nil,
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
		nil,
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
		nil,
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
	testForXYF(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),
		"25/0/29803148/13212522",
		11,
	)
}

func testForXYF(
	t *testing.T,
	expected []string,
	expectedError error,
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
	if !reflect.DeepEqual(theError, expectedError) {
		t.Errorf("error - 期待値：%v, 取得値：%v", expectedError, theError)
	}

	i := 0
	for spatialID := range spatialIDBox.AllXYF() {
		if i >= len(expected) {
			t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expected), i)
		}
		if !reflect.DeepEqual(spatialID.String(), expected[i]) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expected[i], spatialID.String())
		}
		i += 1
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
	testTileXYZBoxAddZoomLevel(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		25, 25, 29803148, 13212522, 0,

		-10, 11,
	)
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
	testTileXYZBoxAddZoomLevel(
		t,
		nil,
		NewSpatialIdError(InputValueErrorCode, ""),

		25, 25, 29803148, 13212522, 0,

		-10, -26,
	)
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
	testTileXYZBoxAddZoomLevel(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 35,
				x: 0,
				y: 0,
				z: 0,
			},
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 35,
				x: 0,
				y: 0,
				z: 1,
			},
		},
		nil,

		34, 34, 1024, 1024, 0,

		-19, 1,
	)
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
	testTileXYZBoxAddZoomLevel(
		t,
		[]*TileXYZ{
			{
				quadkeyZoomLevel: 15,
				altitudekeyZoomLevel: 0,
				x: 0,
				y: 0,
				z: 0,
			},
		},
		nil,

		34, 34, 1024, 1024, 0,

		-19, -34,
	)
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
	for tile := range tileXYZBox.AllXYZ() {
		if i >= len(expected) {
			t.Errorf("TileXYZ - 期待要素数：%v, 取得要素数：%v", len(expected), i)
		}
		if !reflect.DeepEqual(tile, *expected[i]) {
			t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", expected[i], tile)
		}
		i += 1
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
	expectedMinChild := &TileXYZ{
		quadkeyZoomLevel: 26,
		altitudekeyZoomLevel: 0,
		x: 2048,
		y: 2048,
		z: 0,
	}
	expectedMaxChild := &TileXYZ{
		quadkeyZoomLevel: 26,
		altitudekeyZoomLevel: 0,
		x: 2049,
		y: 2049,
		z: 0,
	}

	tileXYZ, error := NewTileXYZ(25, 0, 1024, 1024, 0)
	if error != nil {
		t.Error(error)
	}

	minChild, error := tileXYZ.NewMinChild(1, 0)
	if error != nil {
		t.Error(error)
	}
	maxChild, error := tileXYZ.NewMaxChild(1, 0)
	if error != nil {
		t.Error(error)
	}

	if !reflect.DeepEqual(minChild, expectedMinChild) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", expectedMinChild, minChild)
	}
	if !reflect.DeepEqual(maxChild, expectedMaxChild) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", expectedMaxChild, maxChild)
	}
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
	expectedParent := &TileXYZ{
		quadkeyZoomLevel: 24,
		altitudekeyZoomLevel: 0,
		x: 512,
		y: 512,
		z: 0,
	}

	tileXYZ, error := NewTileXYZ(25, 0, 1024, 1024, 0)
	if error != nil {
		t.Error(error)
	}

	parent, error := tileXYZ.NewParent(1, 0)
	if error != nil {
		t.Error(error)
	}

	if !reflect.DeepEqual(parent, expectedParent) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", expectedParent, parent)
	}
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
	tileXYZ := &TileXYZ{
		quadkeyZoomLevel: 25,
		altitudekeyZoomLevel: 0,
		x: 1024,
		y: 1024,
		z: 0,
	}

	minChild, error := tileXYZ.NewMinChild(0, 0)
	if error != nil {
		t.Error(error)
	}
	maxChild, error := tileXYZ.NewMaxChild(0, 0)
	if error != nil {
		t.Error(error)
	}
	parent, error := tileXYZ.NewParent(0, 0)
	if error != nil {
		t.Error(error)
	}

	if !reflect.DeepEqual(minChild, tileXYZ) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", tileXYZ, minChild)
	}
	if !reflect.DeepEqual(maxChild, tileXYZ) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", tileXYZ, maxChild)
	}
	if !reflect.DeepEqual(parent, tileXYZ) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", tileXYZ, parent)
	}
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
	expectedMinChild := &TileXYZ{
		quadkeyZoomLevel: 0,
		altitudekeyZoomLevel: 26,
		x: 0,
		y: 0,
		z: 2048,
	}
	expectedMaxChild := &TileXYZ{
		quadkeyZoomLevel: 0,
		altitudekeyZoomLevel: 26,
		x: 0,
		y: 0,
		z: 2049,
	}

	tileXYZ, error := NewTileXYZ(0, 25, 0, 0, 1024)
	if error != nil {
		t.Error(error)
	}

	minChild, error := tileXYZ.NewMinChild(0, 1)
	if error != nil {
		t.Error(error)
	}
	maxChild, error := tileXYZ.NewMaxChild(0, 1)
	if error != nil {
		t.Error(error)
	}

	if !reflect.DeepEqual(minChild, expectedMinChild) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", expectedMinChild, minChild)
	}
	if !reflect.DeepEqual(maxChild, expectedMaxChild) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", expectedMaxChild, maxChild)
	}
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
	expectedParent := &TileXYZ{
		quadkeyZoomLevel: 0,
		altitudekeyZoomLevel: 24,
		x: 0,
		y: 0,
		z: 512,
	}

	tileXYZ, error := NewTileXYZ(0, 25, 0, 0, 1024)
	if error != nil {
		t.Error(error)
	}

	parent, error := tileXYZ.NewParent(0, 1)
	if error != nil {
		t.Error(error)
	}

	if !reflect.DeepEqual(parent, expectedParent) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", expectedParent, parent)
	}
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
	tileXYZ := &TileXYZ{
		quadkeyZoomLevel: 0,
		altitudekeyZoomLevel: 25,
		x: 0,
		y: 0,
		z: 1024,
	}

	minChild, error := tileXYZ.NewMinChild(0, 0)
	if error != nil {
		t.Error(error)
	}
	maxChild, error := tileXYZ.NewMaxChild(0, 0)
	if error != nil {
		t.Error(error)
	}
	parent, error := tileXYZ.NewParent(0, 0)
	if error != nil {
		t.Error(error)
	}

	if !reflect.DeepEqual(minChild, tileXYZ) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", tileXYZ, minChild)
	}
	if !reflect.DeepEqual(maxChild, tileXYZ) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", tileXYZ, maxChild)
	}
	if !reflect.DeepEqual(parent, tileXYZ) {
		t.Errorf("TileXYZ - 期待値：%v, 取得値：%v", tileXYZ, parent)
	}
}

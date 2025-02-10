package spatialID

import (
	"math"
	"reflect"
	"testing"

	"github.com/trajectoryjp/geodesy_go/coordinates"
)

// TestGetSpatialIdsOnPoints01 空間ID取得関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0)], 精度レベル：25)
//
// + 確認内容
//   - 入力座標群に対応した空間IDが取得できること
func TestGetSpatialIdsOnPoints01(t *testing.T) {
	testNewSpatialIDFromGeodetic(
		t,

		"25/0/29803148/13212522",
		nil,

		coordinates.Geodetic{
			139.753098,
			35.68537,
			0.0,
		},
		25,
	)
}

// TestGetSpatialIdsOnPoints02 空間ID取得関数 エラー確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0)], 精度：36)
//
// + 確認内容
//   - 精度の入力不備で入力チェックエラーとなること
func TestGetSpatialIdsOnPoints02(t *testing.T) {
	testNewSpatialIDFromGeodetic(
		t,

		"",
		NewSpatialIdError(InputValueErrorCode, ""),

		coordinates.Geodetic{
			139.753098,
			35.685371,
			0.0,
		},
		36,
	)
}

// TestGetSpatialIdsOnPoints03 空間ID取得関数 エラー確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0)], 精度：-1)
//
// + 確認内容
//   - 精度の入力不備で入力チェックエラーとなること
func TestGetSpatialIdsOnPoints03(t *testing.T) {
	testNewSpatialIDFromGeodetic(
		t,

		"",
		NewSpatialIdError(InputValueErrorCode, ""),

		coordinates.Geodetic{
			139.753098,
			35.685371,
			0.0,
		},
		-1,
	)
}

func testNewSpatialIDFromGeodetic(
	t *testing.T,

	expectedSpatialIDString string,
	expectedError error,

	geodetic coordinates.Geodetic,
	z int8,
) {
	spatialID, error := NewSpatialIDFromGeodetic(geodetic, z)
	if error != expectedError {
		t.Errorf("error - 期待値：%s, 取得値：%s", expectedError, error)
	}

	if !reflect.DeepEqual(spatialID.String(), expectedSpatialIDString) {
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectedSpatialIDString, spatialID.String())
	}
}

// TestGetExtendedSpatialIdsOnPoints01 拡張空間ID取得関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0)], 水平精度：18, 垂直精度：25)
//
// + 確認内容
//   - 入力座標群に対応した拡張空間IDが取得できること
func TestGetExtendedSpatialIdsOnPoints01(t *testing.T) {
	testNewTileXYZFromGeodetic(
		t,

		TileXYZ{
			quadkeyZoomLevel: 18,
			altitudekeyZoomLevel: 25,
			x: 232837,
			y: 103222,
			z: 0,
		},
		nil,

		coordinates.Geodetic{
			139.753098,
			35.685371,
			0.0,
		},
		18,
		25,
	)
}

// TestGetExtendedSpatialIdsOnPoints02 拡張空間ID取得関数 エラー確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0)], 水平精度：36, 垂直精度：25)
//
// + 確認内容
//   - 水平精度の入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdsOnPoints02(t *testing.T) {
	testNewTileXYZFromGeodetic(
		t,

		TileXYZ{},
		NewSpatialIdError(InputValueErrorCode, ""),

		coordinates.Geodetic{
			139.753098,
			35.685371,
			0.0,
		},
		36,
		25,
	)
}

// TestGetExtendedSpatialIdsOnPoints03 拡張空間ID取得関数 エラー確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0)], 水平精度：18, 垂直精度：-1)
//
// + 確認内容
//   - 垂直精度の入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdsOnPoints03(t *testing.T) {
	testNewTileXYZFromGeodetic(
		t,

		TileXYZ{},
		NewSpatialIdError(InputValueErrorCode, ""),

		coordinates.Geodetic{
			139.753098,
			35.685371,
			0.0,
		},
		18,
		-1,
	)
}

// TestGetExtendedSpatialIdsOnPoints04 拡張空間ID取得関数 エラー確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0), nil],  水平精度：18, 垂直精度：25)
//
// + 確認内容
//   - 地理座標群の入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdsOnPoints04(t *testing.T) {
	testNewTileXYZFromGeodetic(
		t,

		TileXYZ{},
		NewSpatialIdError(InputValueErrorCode, ""),
		
		coordinates.Geodetic{
			139.753098,
			35.685371,
			0.0,
		},
		18,
		25,
	)
}

func testNewTileXYZFromGeodetic(
	t *testing.T,

	expected TileXYZ,
	expectedError error,

	geodetic coordinates.Geodetic,
	quadkeyZoomLevel int8,
	altitudeZoomLevel int8,
) {
	tileXYZ, error := NewTileXYZFromGeodetic(geodetic, quadkeyZoomLevel, altitudeZoomLevel)
	if error != expectedError {
		t.Errorf("error - 期待値：%s, 取得値：%s", expectedError, error)
	}

	if !reflect.DeepEqual(tileXYZ, expected) {
		t.Errorf("タイルXYZ - 期待値：%v, 取得値：%v", expected, tileXYZ)
	}
}

// TestGetPointOnSpatialId01 空間IDの座標取得関数 空間IDの中心座標の取得確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID："18/32/232837/103222", option：enum.Center (空間IDの中心座標を取得))
//
// + 確認内容
//   - 入力値に対して、空間IDの中心点の座標が格納されたインスタンスのリストが適切に出力されること。
func TestGetPointOnSpatialId01(t *testing.T) {
	expected := coordinates.Geodetic{
		139.75364685058594,
		35.6857446882,
		4160.0,
	}

	spatialIDString := "18/32/232837/103222"
	spatialID, error := NewSpatialIDFromString(spatialIDString)
	if error != nil {
		t.Errorf("error - 期待値：%s", error)
	}

	spatialIDBox, error := NewSpatialIDBox(*spatialID, *spatialID)
	if error != nil {
		t.Errorf("error - 期待値：%s", error)
	}

	geodeticBox := NewGeodeticBoxFromSpatialIDBox(*spatialIDBox)
	centor := geodeticBox.GetCenter()

	if !reflect.DeepEqual(centor, expected) {
		t.Errorf("中心座標 - 期待値：%v, 取得値：%v", expected, centor)
	}
}

// TestGetPointOnSpatialId02 空間IDの座標取得関数 空間IDの頂点座標を取得
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID："18/32/0/0", option：enum.Vertex(空間IDの頂点座標を取得))
//
// + 確認内容
//   - 入力値に対して、空間IDの中心点の座標が格納されたインスタンスのリストが適切に出力されること。
func TestGetPointOnSpatialId02(t *testing.T) {
	expected := []*coordinates.Geodetic{
		// TODO: そもそもがおかしいので作り直す
	}

	spatialIDString := "3/32/0/0"
	spatialID, error := NewSpatialIDFromString(spatialIDString)
	if error != nil {
		t.Errorf("error - 期待値：%s", error)
	}

	spatialIDBox, error := NewSpatialIDBox(*spatialID, *spatialID)
	if error != nil {
		t.Errorf("error - 期待値：%s", error)
	}

	geodeticBox := NewGeodeticBoxFromSpatialIDBox(*spatialIDBox)
	vertices := geodeticBox.GetVertices()

	if !reflect.DeepEqual(vertices, expected) {
		t.Errorf("頂点座標 - 期待値：%v, 取得値：%v", expected, vertices)
	}
}

// TestGetPointOnExtendedSpatialId01 拡張空間IDの座標取得関数 拡張空間IDの中心座標の取得確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID："18/232837/103222/18/32", option：enum.Center (拡張空間IDの中心座標を取得))
//
// + 確認内容
//   - 入力値に対して、拡張空間IDの中心点の座標が格納されたインスタンスのリストが適切に出力されること。
func TestGetPointOnExtendedSpatialId01(t *testing.T) {
	expected := coordinates.Geodetic{
		139.75364685058594,
		35.6857446882,
		4160.0,
	}

	tileXYZ := TileXYZ{
		quadkeyZoomLevel: 18,
		altitudekeyZoomLevel: 7,
		x: 232837,
		y: 103222,
		z: 36,
	}

	tileXYZBox, error := NewTileXYZBox(tileXYZ, tileXYZ)
	if error != nil {
		t.Errorf("error - 期待値：%s", error)
	}

	geodeticBox := NewGeodeticBoxFromTileXYZBox(*tileXYZBox)
	centor := geodeticBox.GetCenter()

	if !reflect.DeepEqual(centor, expected) {
		t.Errorf("中心座標 - 期待値：%v, 取得値：%v", expected, centor)
	}
}

// TestGetPointOnExtendedSpatialId02 拡張空間IDの座標取得関数 拡張空間IDの頂点座標を取得
// 試験詳細：
// + 試験データ
//   - パターン：
//     (拡張空間ID："3/0/0/18/32", option：enum.Vertex(拡張空間IDの頂点座標を取得))
//
// + 確認内容
//   - 入力値に対して、拡張空間IDの中心点の座標が格納されたインスタンスのリストが適切に出力されること。
func TestGetPointOnExtendedSpatialId02(t *testing.T) {
	expected := []*coordinates.Geodetic{
		// TODO: そもそもがおかしいので作り直す
	}

	tileXYZ := TileXYZ{
		quadkeyZoomLevel: 3,
		altitudekeyZoomLevel: 7,
		x: 0,
		y: 0,
		z: 36,
	}

	tileXYZBox, error := NewTileXYZBox(tileXYZ, tileXYZ)
	if error != nil {
		t.Errorf("error - 期待値：%s", error)
	}

	geodeticBox := NewGeodeticBoxFromTileXYZBox(*tileXYZBox)
	vertices := geodeticBox.GetVertices()

	if !reflect.DeepEqual(vertices, expected) {
		t.Errorf("頂点座標 - 期待値：%v, 取得値：%v", expected, vertices)
	}
}

// TestGetAltitudeOnVerticalIndexAndZoom01 高さ及び分解能取得関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 高さの位置：56, 垂直方向精度：25
//
// + 確認内容
//   - 入力値に対して、適切な高さの位置と垂直方向精度が返却されること。
func TestGetAltitudeOnVerticalIndexAndZoom01(t *testing.T) {
	tileXYZ := TileXYZ{
		quadkeyZoomLevel: 3,
		altitudekeyZoomLevel: 14,
		x: 0,
		y: 0,
		z: 568,
	}
	
	// 期待値
	//高さ期待値
	expectedAltitude := float64(tileXYZ.GetZ()) * math.Pow(2, 25.0) / math.Pow(2, float64(tileXYZ.GetAltitudekeyZoomLevel()))

	//分解能期待値
	expectResolution := math.Pow(2, 25.0) / math.Pow(2, float64(tileXYZ.GetAltitudekeyZoomLevel()))

	tileXYZBox, error := NewTileXYZBox(tileXYZ, tileXYZ)
	if error != nil {
		t.Errorf("error - 期待値：%s", error)
	}

	geodeticBox := NewGeodeticBoxFromTileXYZBox(*tileXYZBox)

	// 高さと期待値の比較
	if !reflect.DeepEqual(*geodeticBox.Min.Altitude(), expectedAltitude) {
		// 高さが期待値と異なる場合Errorをログに出力
		t.Errorf("高さ - 期待値：%+v, 取得値：%+v", expectedAltitude, *geodeticBox.Min.Altitude())
	}
	// 分解能と期待値の比較
	if !reflect.DeepEqual(*geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude(), expectResolution) {
		// 分解能が期待値と異なる場合Errorをログに出力
		t.Errorf("分解能 - 期待値：%+v, 取得値：%+v", expectResolution, *geodeticBox.Max.Altitude() - *geodeticBox.Min.Altitude())
	}
}

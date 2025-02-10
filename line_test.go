package spatialID

import (
	"reflect"
	"testing"

	"github.com/trajectoryjp/geodesy_go/coordinates"
)

// TestGetSpatialIdsOnLine01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10, 終点：10/10/13/10, 精度レベル:10)
//
// + 確認内容
//   - 入力に対応した空間ID集合が取得できること
func TestGetSpatialIdsOnLine01(t *testing.T) {
	testGetSpatialIdsOnLine(
		t,

		[]string{"10/10/10/10", "10/10/11/10", "10/10/12/10", "10/10/13/10"},
		nil,

		"10/10/10/10",
		"10/10/13/10",

		10,
	)
}

// TestGetSpatialIdsOnLine02 精度閾値超過
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10, 終点：10/10/13/10, 精度レベル:36)
//
// + 確認内容
//   - エラーインスタンス（InputValueErrorCode）が返却されること
func TestGetSpatialIdsOnLine02(t *testing.T) {
	testGetSpatialIdsOnLine(
		t,

		[]string{},
		NewSpatialIdError(InputValueErrorCode, "入力チェックエラー"),

		"10/10/10/10",
		"10/10/13/10",

		36,
	)
}

func testGetSpatialIdsOnLine(
	t *testing.T,
	expectedSpatialIDStrings []string,
	expectedError error,
	startSpatialIDString string,
	endSpatialIDString string,
	z int8,
) {
	startSpatialID, error := NewSpatialIDFromString(startSpatialIDString)
	if error != nil {
		if error.Error() == expectedError.Error() {
			return
		}

		t.Fatal(error)
	}
	startSpatialIDBox, error := NewSpatialIDBox(*startSpatialID, *startSpatialID)
	if error != nil {
		if error.Error() == expectedError.Error() {
			return
		}

		t.Fatal(error)
	}
	startGeodeticBox := NewGeodeticBoxFromSpatialIDBox(*startSpatialIDBox)

	endSpatialID, error := NewSpatialIDFromString(endSpatialIDString)
	if error != nil {
		if error.Error() == expectedError.Error() {
			return
		}

		t.Fatal(error)
	}
	endSpatialIDBox, error := NewSpatialIDBox(*endSpatialID, *endSpatialID)
	if error != nil {
		if error.Error() == expectedError.Error() {
			return
		}

		t.Fatal(error)
	}
	endGeodeticBox := NewGeodeticBoxFromSpatialIDBox(*endSpatialIDBox)

	spatialIDBox, error := NewSpatialIDBox(*startSpatialID, *endSpatialID)
	if error != nil {
		if error.Error() == expectedError.Error() {
			return
		}

		t.Fatal(error)
	}
	spatialIDBox.AddZ(z - spatialIDBox.GetMin().GetZ())

	i := 0
	for spatialID := range spatialIDBox.AllCollisionWithConvexHull(
		[]*coordinates.Geodetic{
			{
				(*startGeodeticBox.Min.Longitude() + *endGeodeticBox.Min.Longitude()) / 2,
				(*startGeodeticBox.Min.Latitude() + *endGeodeticBox.Min.Latitude()) / 2,
				(*startGeodeticBox.Min.Altitude() + *endGeodeticBox.Min.Altitude()) / 2,
			},
			{
				(*startGeodeticBox.Max.Longitude() + *endGeodeticBox.Max.Longitude()) / 2,
				(*startGeodeticBox.Max.Latitude() + *endGeodeticBox.Max.Latitude()) / 2,
				(*startGeodeticBox.Max.Altitude() + *endGeodeticBox.Max.Altitude()) / 2,
			},
		},
		0.0,
	) {
		if i >= len(expectedSpatialIDStrings) {
			t.Fatalf("Too many spatial IDs: %v", i)
		}
		if spatialID.String() != expectedSpatialIDStrings[i] {
			t.Fatalf("Unexpected spatial ID: %v", spatialID.String())
		}

		i += 1
	}
	if i != len(expectedSpatialIDStrings) {
		t.Fatalf("Too few spatial IDs: %v", i)
	}
}

// TestGetExtendedSpatialIdsOnLine01 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10/10, 終点：10/13/10/10/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力に対応した空間ID集合が取得できること
func TestGetExtendedSpatialIdsOnLine01(t *testing.T) {
	testGetTileXYZsOnLine(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 10,
				x: 10,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 10,
				x: 11,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 10,
				x: 12,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 10,
				x: 13,
				y: 10,
				z: 10,
			},
		},
		nil,

		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 10,
			x: 10,
			y: 10,
			z: 10,
		},
		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 10,
			x: 13,
			y: 10,
			z: 10,
		},

		10, 10,
	)
}

// TestGetExtendedSpatialIdsOnLine04 精度閾値超過
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10/10, 終点：10/13/10/10/10, 水平方向の精度レベル:36, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - エラーインスタンス（InputValueErrorCode）が返却されること
func TestGetExtendedSpatialIdsOnLine04(t *testing.T) {
	testGetTileXYZsOnLine(
		t,

		[]*TileXYZ{},
		NewSpatialIdError(InputValueErrorCode, "入力チェックエラー"),

		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 10,
			x: 10,
			y: 10,
			z: 10,
		},
		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 10,
			x: 13,
			y: 10,
			z: 10,
		},

		36, 10,
	)
}

// TestGetExtendedSpatialIdsOnLine05 始点終点同値
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10/10, 終点：10/10/10/10/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力に対応した空間ID集合が取得できること
func TestGetExtendedSpatialIdsOnLine05(t *testing.T) {
	testGetTileXYZsOnLine(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 10,
				x: 10,
				y: 10,
				z: 10,
			},
		},
		nil,

		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 10,
			x: 10,
			y: 10,
			z: 10,
		},
		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 10,
			x: 10,
			y: 10,
			z: 10,
		},

		10, 10,
	)
}

// TestGetExtendedSpatialIdsOnLine06 中点取得時再設定閾値(水平精度31以上)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：31/10/10/10/10, 終点：31/13/10/10/10, 水平方向の精度レベル:31, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力に対応した空間ID集合が取得できること
func TestGetExtendedSpatialIdsOnLine06(t *testing.T) {
	testGetTileXYZsOnLine(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 31,
				altitudekeyZoomLevel: 10,
				x: 10,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 31,
				altitudekeyZoomLevel: 10,
				x: 11,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 31,
				altitudekeyZoomLevel: 10,
				x: 12,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 31,
				altitudekeyZoomLevel: 10,
				x: 13,
				y: 10,
				z: 10,
			},
		},
		nil,

		TileXYZ{
			quadkeyZoomLevel: 31,
			altitudekeyZoomLevel: 10,
			x: 10,
			y: 10,
			z: 10,
		},
		TileXYZ{
			quadkeyZoomLevel: 31,
			altitudekeyZoomLevel: 10,
			x: 13,
			y: 10,
			z: 10,
		},

		31, 10,
	)
}

// TestGetExtendedSpatialIdsOnLine07 中点取得時再設定閾値(水平精度30以下)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：30/10/10/10/10, 終点：30/13/10/10/10, 水平方向の精度レベル:30, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力に対応した空間ID集合が取得できること
func TestGetExtendedSpatialIdsOnLine07(t *testing.T) {
	testGetTileXYZsOnLine(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 30,
				altitudekeyZoomLevel: 10,
				x: 10,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 30,
				altitudekeyZoomLevel: 10,
				x: 11,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 30,
				altitudekeyZoomLevel: 10,
				x: 12,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 30,
				altitudekeyZoomLevel: 10,
				x: 13,
				y: 10,
				z: 10,
			},
		},
		nil,

		TileXYZ{
			quadkeyZoomLevel: 30,
			altitudekeyZoomLevel: 10,
			x: 10,
			y: 10,
			z: 10,
		},
		TileXYZ{
			quadkeyZoomLevel: 30,
			altitudekeyZoomLevel: 10,
			x: 13,
			y: 10,
			z: 10,
		},

		30, 10,
	)
}

// TestGetExtendedSpatialIdsOnLine08 中点取得時再設定閾値(垂直精度34以上)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/34/10, 終点：10/13/10/34/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:34)
//
// + 確認内容
//   - 入力に対応した空間ID集合が取得できること
func TestGetExtendedSpatialIdsOnLine08(t *testing.T) {
	testGetTileXYZsOnLine(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 34,
				x: 10,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 34,
				x: 11,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 34,
				x: 12,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 34,
				x: 13,
				y: 10,
				z: 10,
			},
		},
		nil,

		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 34,
			x: 10,
			y: 10,
			z: 10,
		},
		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 34,
			x: 13,
			y: 10,
			z: 10,
		},

		10, 34,
	)
}

// TestGetExtendedSpatialIdsOnLine09 中点取得時再設定閾値(垂直精度33以下)
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/33/10, 終点：10/13/10/33/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:33)
//
// + 確認内容
//   - 入力に対応した空間ID集合が取得できること
func TestGetExtendedSpatialIdsOnLine09(t *testing.T) {
	testGetTileXYZsOnLine(
		t,

		[]*TileXYZ{
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 33,
				x: 10,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 33,
				x: 11,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 33,
				x: 12,
				y: 10,
				z: 10,
			},
			{
				quadkeyZoomLevel: 10,
				altitudekeyZoomLevel: 33,
				x: 13,
				y: 10,
				z: 10,
			},
		},
		nil,

		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 33,
			x: 10,
			y: 10,
			z: 10,
		},
		TileXYZ{
			quadkeyZoomLevel: 10,
			altitudekeyZoomLevel: 33,
			x: 13,
			y: 10,
			z: 10,
		},

		10, 33,
	)
}

func testGetTileXYZsOnLine(
	t *testing.T,
	expected []*TileXYZ,
	expectedError error,
	start TileXYZ,
	end TileXYZ,
	quadkeyZoomLevel int8,
	altitudekeyZoomLevel int8,
) {
	startTileXYZBox, error := NewTileXYZBox(start, start)
	if error != nil {
		if error.Error() == expectedError.Error() {
			return
		}

		t.Fatal(error)
	}
	startGeodeticBox := NewGeodeticBoxFromTileXYZBox(*startTileXYZBox)

	endTileXYZBox, error := NewTileXYZBox(end, end)
	if error != nil {
		if error.Error() == expectedError.Error() {
			return
		}

		t.Fatal(error)
	}
	endGeodeticBox := NewGeodeticBoxFromTileXYZBox(*endTileXYZBox)

	tileXYZBox, error := NewTileXYZBox(start, end)
	if error != nil {
		if error.Error() == expectedError.Error() {
			return
		}

		t.Fatal(error)
	}
	tileXYZBox.AddZoomLevel(quadkeyZoomLevel - tileXYZBox.GetMin().GetQuadkeyZoomLevel(), altitudekeyZoomLevel - tileXYZBox.GetMin().GetAltitudekeyZoomLevel())

	i := 0
	for spatialID := range tileXYZBox.AllCollisionWithConvexHull(
		[]*coordinates.Geodetic{
			{
				(*startGeodeticBox.Min.Longitude() + *endGeodeticBox.Min.Longitude()) / 2,
				(*startGeodeticBox.Min.Latitude() + *endGeodeticBox.Min.Latitude()) / 2,
				(*startGeodeticBox.Min.Altitude() + *endGeodeticBox.Min.Altitude()) / 2,
			},
			{
				(*startGeodeticBox.Max.Longitude() + *endGeodeticBox.Max.Longitude()) / 2,
				(*startGeodeticBox.Max.Latitude() + *endGeodeticBox.Max.Latitude()) / 2,
				(*startGeodeticBox.Max.Altitude() + *endGeodeticBox.Max.Altitude()) / 2,
			},
		},
		0.0,
	) {
		if i >= len(expected) {
			t.Fatalf("Too many tile XYZs: %v", i)
		}
		if !reflect.DeepEqual(spatialID, expected[i]) {
			t.Fatalf("Unexpected tile XYZ: %v", spatialID)
		}

		i += 1
	}
	if i != len(expected) {
		t.Fatalf("Too few tile XYZs: %v", i)
	}
}

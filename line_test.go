package spatialID

import (
	"fmt"
	"testing"

	"github.com/trajectoryjp/geodesy_go/coordinates"
	"github.com/trajectoryjp/spatial_id_go/v4/common/enum"
	"github.com/trajectoryjp/spatial_id_go/v4/common/object"
	"github.com/trajectoryjp/spatial_id_go/v4/common/spatial"
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
	var zoom int64 = 36
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/13/10/10/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetSpatialIdsOnLine(startPoint, endPoint, zoom)

	expectVal := make([]string, 0, 2)
	expectErr := "InputValueError,入力チェックエラー"

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	if resultErr == nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：nil", expectErr)
	}
	t.Log("テスト終了")
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
	var hZoom int64 = 10
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/13/10/10/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)

	expectVal := []string{"10/10/10/10/10", "10/11/10/10/10", "10/12/10/10/10", "10/13/10/10/10"}

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
		// 戻り値のエラーインスタンスがが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultVal)
	}
	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnLine02 始点違反
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：nil, 終点：10/13/10/10/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - エラーインスタンス（InputValueErrorCode）が返却されること
func TestGetExtendedSpatialIdsOnLine02(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/13/10/10/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(nil, endPoint, hZoom, vZoom)

	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	if resultErr == nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：nil", expectErr)
	}
	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnLine03 終点違反
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10/10, 終点：nil, 水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - エラーインスタンス（InputValueErrorCode）が返却されること
func TestGetExtendedSpatialIdsOnLine03(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(startPoint, nil, hZoom, vZoom)

	expectVal := []string{}
	expectErr := "InputValueError,入力チェックエラー"

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	if resultErr == nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：nil", expectErr)
	}
	t.Log("テスト終了")
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
	var hZoom int64 = 36
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/13/10/10/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)

	expectVal := make([]string, 0, 2)
	expectErr := "InputValueError,入力チェックエラー"

	//戻り値要素数と期待値の比較
	if len(resultVal) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(resultVal))
	}

	if resultErr == nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：nil", expectErr)
	}
	t.Log("テスト終了")
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
	var hZoom int64 = 10
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)
	resultCnt := len(resultVal)

	expectVal := []string{"10/10/10/10/10"}
	expectCnt := len(expectVal)

	//戻り値要素数と期待値の比較
	if resultCnt != expectCnt {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", expectCnt, resultCnt)
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(resultVal, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, resultVal)
		}
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスがが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultVal)
	}
	t.Log("テスト終了")
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
	var hZoom int64 = 31
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("31/10/10/10/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	endCoodinate, _ := GetPointOnExtendedSpatialId("31/13/10/10/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)

	expectVal := []string{"31/10/10/10/10", "31/11/10/10/10", "31/12/10/10/10", "31/13/10/10/10"}

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
		// 戻り値のエラーインスタンスがが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultVal)
	}
	t.Log("テスト終了")
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
	var hZoom int64 = 30
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("30/10/10/10/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	endCoodinate, _ := GetPointOnExtendedSpatialId("30/13/10/10/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)

	expectVal := []string{"30/10/10/10/10", "30/11/10/10/10", "30/12/10/10/10", "30/13/10/10/10"}

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
		// 戻り値のエラーインスタンスがが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultVal)
	}
	t.Log("テスト終了")
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
	var hZoom int64 = 10
	var vZoom int64 = 34
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/34/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/13/10/34/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)

	expectVal := []string{"10/10/10/34/10", "10/11/10/34/10", "10/12/10/34/10", "10/13/10/34/10"}

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
		// 戻り値のエラーインスタンスがが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultVal)
	}
	t.Log("テスト終了")
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
	var hZoom int64 = 10
	var vZoom int64 = 33
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/33/10", enum.Center)
	startPoint, _ := object.NewPoint(startCoodinate[0].Lon(), startCoodinate[0].Lat(), startCoodinate[0].Alt())
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/13/10/33/10", enum.Center)
	endPoint, _ := object.NewPoint(endCoodinate[0].Lon(), endCoodinate[0].Lat(), endCoodinate[0].Alt())
	resultVal, resultErr := GetExtendedSpatialIdsOnLine(startPoint, endPoint, hZoom, vZoom)

	expectVal := []string{"10/10/10/33/10", "10/11/10/33/10", "10/12/10/33/10", "10/13/10/33/10"}

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
		// 戻り値のエラーインスタンスがが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultVal)
	}
	t.Log("テスト終了")
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

func TestGetSpatialIdsOnLineForTests(t *testing.T) {

	startPoint, error := object.NewPoint(139.788452, 35.670935, 100)
	if error != nil {
		t.Error(error)
	}

	endPoint, error := object.NewPoint(139.788074, 35.675711, 100)
	if error != nil {
		t.Error(error)
	}

	ids, error := GetExtendedSpatialIdsOnLine(startPoint, endPoint, 23, 23)
	if error != nil {
		t.Error(error)
	}

	for _, v := range ids {
		fmt.Printf("\"%v\",", v)
	}
}

// TestMiddleSpatialIds01 正常系動作確認 中点が始点・終点の周囲にある場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10/10, 終点：10/12/10/10/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の中点の空間IDに対する処理関数によって、空間IDのスライスに始点終点の中点の空間IDが追加されること
func TestMiddleSpatialIds01(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint := spatial.Point3{X: startCoodinate[0].Lon(), Y: startCoodinate[0].Lat(), Z: startCoodinate[0].Alt()}
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/12/10/10/10", enum.Center)
	endPoint := spatial.Point3{X: endCoodinate[0].Lon(), Y: endCoodinate[0].Lat(), Z: endCoodinate[0].Alt()}
	spatialIDs := []string{"10/10/10/10/10", "10/12/10/10/10"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/10/10/10/10", "10/11/10/10/10", "10/12/10/10/10"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds02 正常系動作確認  ベクトルの成分Xが経度閾値(境界値)未満の場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000002, Y: 0.00000001, Z: 0.001},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds02(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000002, Y: 0.00000001, Z: 0.001}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds03 正常系動作確認  ベクトルの成分Xが経度閾値(境界値)と等しい場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000003, Y: 0.00000001, Z: 0.001},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds03(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000003, Y: 0.00000001, Z: 0.001}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds04 正常系動作確認  ベクトルの成分Xが経度閾値(境界値)より大きい場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000004, Y: 0.00000001, Z: 0.001},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds04(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000004, Y: 0.00000001, Z: 0.001}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds05 正常系動作確認  ベクトルの成分Yが緯度閾値(境界値)未満の場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000001, Y: 0.00000002, Z: 0.001},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds05(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000001, Y: 0.00000002, Z: 0.001}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds06 正常系動作確認  ベクトルの成分Yが緯度閾値(境界値)と等しい場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000001, Y: 0.00000003, Z: 0.001},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds06(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000001, Y: 0.00000003, Z: 0.001}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds07 正常系動作確認  ベクトルの成分Yが緯度閾値(境界値)より大きい場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000001, Y: 0.00000004, Z: 0.001},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds07(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000001, Y: 0.00000004, Z: 0.001}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds08 正常系動作確認  ベクトルの成分Zが高さ閾値(境界値)未満の場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000001, Y: 0.00000001, Z: 0.003},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds08(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.003}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds09 正常系動作確認  ベクトルの成分Zが高さ閾値(境界値)と等しい場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000001, Y: 0.00000001, Z: 0.004},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds09(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.004}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds10 正常系動作確認  ベクトルの成分Zが高さ閾値(境界値)より大きい場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：{X: 0.00000001, Y: 0.00000001, Z: 0.001}, 終点：{X: 0.00000001, Y: 0.00000001, Z: 0.005},
//     水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 入力の始点終点の中点の空間IDが設定されること
func TestMiddleSpatialIds10(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.001}
	endPoint := spatial.Point3{X: 0.00000001, Y: 0.00000001, Z: 0.005}

	spatialIDs := []string{"10/512/511/10/0"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/512/511/10/0", "10/512/511/10/0"}

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	for _, exp := range expectVal {
		if !contains(spatialIDs, exp) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
		}
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds11 正常系動作確認  中点が始点の周囲にある場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10/10, 終点：10/7/10/10/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 終点の範囲外かつ、始点の周囲内の中点の空間ID10/9/10/10/10が設定されること
func TestMiddleSpatialIds11(t *testing.T) {
	//入力値
	var hZoom int64 = 10
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint := spatial.Point3{X: startCoodinate[0].Lon(), Y: startCoodinate[0].Lat(), Z: startCoodinate[0].Alt()}
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/7/10/10/10", enum.Center)
	endPoint := spatial.Point3{X: endCoodinate[0].Lon(), Y: endCoodinate[0].Lat(), Z: endCoodinate[0].Alt()}
	spatialIDs := []string{"10/10/10/10/10", "10/7/10/10/10"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	//期待値
	expectVal := []string{"10/10/10/10/10", "10/9/10/10/10", "10/8/10/10/10", "10/7/10/10/10"}
	expectID := "10/9/10/10/10"

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	if !contains(spatialIDs, expectID) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds12 正常系動作確認  中点が終点の周囲にある場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10/10, 終点：10/13/10/10/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 始点の範囲外かつ、終点の周囲内の中点の空間ID10/12/10/10/10が設定されること
func TestMiddleSpatialIds12(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint := spatial.Point3{X: startCoodinate[0].Lon(), Y: startCoodinate[0].Lat(), Z: startCoodinate[0].Alt()}
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/13/10/10/10", enum.Center)
	endPoint := spatial.Point3{X: endCoodinate[0].Lon(), Y: endCoodinate[0].Lat(), Z: endCoodinate[0].Alt()}
	spatialIDs := []string{"10/10/10/10/10", "10/13/10/10/10"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/10/10/10/10", "10/11/10/10/10", "10/12/10/10/10", "10/13/10/10/10"}
	expectID := "10/12/10/10/10"

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	if !contains(spatialIDs, expectID) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
	}
	t.Log("テスト終了")
}

// TestMiddleSpatialIds13 正常系動作確認  中点が始点・終点の周囲にない場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (始点：10/10/10/10/10, 終点：10/14/10/10/10, 水平方向の精度レベル:10, 垂直方向の精度レベル:10)
//
// + 確認内容
//   - 始点、終点の周囲にない中点の空間ID10/12/10/10/10が設定されること
func TestMiddleSpatialIds13(t *testing.T) {
	var hZoom int64 = 10
	var vZoom int64 = 10
	startCoodinate, _ := GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Center)
	startPoint := spatial.Point3{X: startCoodinate[0].Lon(), Y: startCoodinate[0].Lat(), Z: startCoodinate[0].Alt()}
	endCoodinate, _ := GetPointOnExtendedSpatialId("10/14/10/10/10", enum.Center)
	endPoint := spatial.Point3{X: endCoodinate[0].Lon(), Y: endCoodinate[0].Lat(), Z: endCoodinate[0].Alt()}
	spatialIDs := []string{"10/10/10/10/10", "10/14/10/10/10"}
	middleSpatialIds(startPoint, endPoint, hZoom, vZoom, LonMinima, LatMinima, AltMinima, func(spatialID string) {
		spatialIDs = append(spatialIDs, spatialID)
	})

	expectVal := []string{"10/10/10/10/10", "10/11/10/10/10", "10/12/10/10/10", "10/13/10/10/10", "10/14/10/10/10"}
	expectID := "10/12/10/10/10"

	//戻り値要素数と期待値の比較
	if len(spatialIDs) != len(expectVal) {
		t.Errorf("空間ID - 期待要素数：%v, 取得要素数：%v", len(expectVal), len(spatialIDs))
	}

	//戻り値の空間IDと期待値の比較
	if !contains(spatialIDs, expectID) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectVal, spatialIDs)
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

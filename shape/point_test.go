package shape

import (
	"math"
	"reflect"
	"testing"

	"github.com/trajectoryjp/spatial_id_go/v2/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v2/common/enum"
	"github.com/trajectoryjp/spatial_id_go/v2/common/object"

	"github.com/wroge/wgs84"
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
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1}
	var Zoom int64 = 25

	// 期待値
	expectVal := []string{"25/0/29803148/13212522"}

	// テスト対象呼び出し
	resultVal, resultErr := GetSpatialIdsOnPoints(pList, Zoom)

	// 戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
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
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1}
	var Zoom int64 = 36

	// 期待値
	expectVal := make([]string, 0, len(pList))
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetSpatialIdsOnPoints(pList, Zoom)

	// 戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
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
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1}
	var Zoom int64 = -1

	// 期待値
	expectVal := make([]string, 0, len(pList))
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetSpatialIdsOnPoints(pList, Zoom)

	// 戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetSpatialIdsOnPoints04 空間ID取得関数 エラー確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0), nil],  精度：25)
//
// + 確認内容
//   - 地理座標群の入力不備で入力チェックエラーとなること
func TestGetSpatialIdsOnPoints04(t *testing.T) {
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1, nil}
	var Zoom int64 = 25

	// 期待値
	expectVal := make([]string, 0, len(pList))
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetSpatialIdsOnPoints(pList, Zoom)

	// 戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetSpatialIdsOnPoints05 空間ID取得関数 地理座標群のサイズが0の場合の確認
// 試験詳細：
// + 試験データ
//
// + 確認内容
func TestGetSpatialIdsOnPoints05(t *testing.T) {
	//空ケースになることはないと保証されている。
	t.Log("テスト終了")
}

// TestGetSpatialIdsOnPoints06 空間ID取得関数 地理座標群のサイズが複数の場合の確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0), (139.753098, 35.685371, 0.0),
//     (139.753098, 35.685371, 0.0)], 精度：25)
//
// + 確認内容
//   - 入力座標群に対応した拡張空間IDが取得できること
func TestGetSpatialIdsOnPoints06(t *testing.T) {
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	p2, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	p3, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1, p2, p3}
	var Zoom int64 = 25

	// 期待値
	expectVal := []string{"25/0/29803148/13212522", "25/0/29803148/13212522", "25/0/29803148/13212522"}

	// テスト対象呼び出し
	resultVal, resultErr := GetSpatialIdsOnPoints(pList, Zoom)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
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
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1}
	var hZoom int64 = 18
	var vZoom int64 = 25

	// 期待値
	expectVal := []string{"18/232837/103222/25/0"}

	// テスト対象呼び出し
	resultVal, resultErr := GetExtendedSpatialIdsOnPoints(pList, hZoom, vZoom)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
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
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1}
	var hZoom int64 = 36
	var vZoom int64 = 25

	// 期待値
	expectVal := make([]string, 0, len(pList))
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetExtendedSpatialIdsOnPoints(pList, hZoom, vZoom)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
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
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1}
	var hZoom int64 = 18
	var vZoom int64 = -1

	// 期待値
	expectVal := make([]string, 0, len(pList))
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetExtendedSpatialIdsOnPoints(pList, hZoom, vZoom)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
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
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1, nil}
	var hZoom int64 = 18
	var vZoom int64 = 25

	// 期待値
	expectVal := make([]string, 0, len(pList))
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetExtendedSpatialIdsOnPoints(pList, hZoom, vZoom)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnPoints05 拡張空間ID取得関数 地理座標群のサイズが0の場合の確認
// 試験詳細：
// + 試験データ
//
// + 確認内容
func TestGetExtendedSpatialIdsOnPoints05(t *testing.T) {
	//空ケースになることはないと保証されている。
	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdsOnPoints06 拡張空間ID取得関数 地理座標群のサイズが複数の場合の確認
// 試験詳細：
// + 試験データ
//   - パターン1：
//     (地理座標群：[(139.753098, 35.685371, 0.0), (139.753098, 35.685371, 0.0),
//     (139.753098, 35.685371, 0.0)], 水平精度：18, 垂直精度：25)
//
// + 確認内容
//   - 入力座標群に対応した拡張空間IDが取得できること
func TestGetExtendedSpatialIdsOnPoints06(t *testing.T) {
	// テスト用入力パラメータ
	p1, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	p2, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	p3, _ := object.NewPoint(139.753098, 35.685371, 0.0)
	pList := []*object.Point{p1, p2, p3}
	var hZoom int64 = 18
	var vZoom int64 = 25

	// 期待値
	expectVal := []string{"18/232837/103222/25/0", "18/232837/103222/25/0", "18/232837/103222/25/0"}

	// テスト対象呼び出し
	resultVal, resultErr := GetExtendedSpatialIdsOnPoints(pList, hZoom, vZoom)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
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

	// 入力値
	inputSpatialID := "18/32/232837/103222"
	inputOption := enum.Center

	// 期待値
	expectVal := []*object.Point{}
	point, _ := object.NewPoint(139.75364685058594, 35.6857446882, 4160.0)
	expectVal = append(expectVal, point)

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnSpatialId(inputSpatialID, inputOption)

	// 戻り値のエラーと期待値の比較
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s", resultErr)
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if len(resultVal) != 1 {
		// 中心座標取得結果の確認
		t.Errorf("中心座標 - 取得値：%d", len(resultVal))
	} else if !reflect.DeepEqual(resultVal, expectVal) {
		// 中心座標取得結果の確認
		t.Errorf("中心座標 - 期待値：%+v, 取得値：%+v", expectVal[0], resultVal[0])
	}
	t.Log("テスト終了")
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

	// パターン1の入力パラメータ
	inputSpatialID01 := "3/32/0/0"
	inputOption01 := enum.Vertex

	// パターン2の入力パラメータ
	inputSpatialID02 := "3/32/8/0"
	inputOption02 := enum.Vertex

	// パターン1の取得値
	resultVal01, resultErr01 := GetPointOnSpatialId(inputSpatialID01, inputOption01)

	// パターン2の取得値
	resultVal02, resultErr02 := GetPointOnSpatialId(inputSpatialID02, inputOption02)

	// 戻り値のエラーと期待値の比較
	if resultErr01 != resultErr02 {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 取得値01：%s, 取得値02：%s", resultErr01, resultErr02)
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if len(resultVal01) != 8 {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標 - 取得値：%d", len(resultVal01))
	} else if !reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 期待値：%+v, 取得値：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 期待値：%+v, 取得値：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 期待値：%+v, 取得値：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 期待値：%+v, 取得値：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 期待値：%+v, 取得値：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 期待値：%+v, 取得値：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 期待値：%+v, 取得値：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 期待値：%+v, 取得値：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetPointOnSpatialId03 空間IDの座標取得関数 空間ID入力値不足のエラー確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID："18/32/232837", option：enum.Center (空間IDの中心座標を取得))
//
// + 確認内容
//   - 空間IDの入力不備で入力チェックエラーとなること
func TestGetPointOnSpatialId03(t *testing.T) {

	// 入力値
	inputSpatialID := "18/32/232837"
	inputOption := enum.Center

	// 期待値
	expectVal := []*object.Point{}
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnSpatialId(inputSpatialID, inputOption)

	// 戻り値のエラーと期待値の比較
	if expectErr != resultErr.Error() {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("座標 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestGetPointOnSpatialId04 空間IDの座標取得関数 空間ID精度レベル入力不備エラー確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID："36/32/232837/103222", option：enum.Center (空間IDの中心座標を取得))
//
// + 確認内容
//   - 空間IDの入力不備で入力チェックエラーとなること
func TestGetPointOnSpatialId04(t *testing.T) {

	// 入力値
	inputSpatialID := "36/32/232837/103222"
	inputOption := enum.Center

	// 期待値
	expectVal := []*object.Point{}
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnSpatialId(inputSpatialID, inputOption)

	// 戻り値のエラーと期待値の比較
	if expectErr != resultErr.Error() {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("座標 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	t.Log("テスト終了")
}

// TestGetPointOnSpatialId05 空間IDの座標取得関数 サポート対象外のオプション入力によるエラー確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID："18/32/232837/103222", option：3 (存在しないオプション入力))
//
// + 確認内容
//   - 空間IDの入力不備でオプション値の指定エラーとなること
func TestGetPointOnSpatialId05(t *testing.T) {

	// 入力値
	inputSpatialID := "18/32/232837/103222"

	// 期待値
	expectVal := []*object.Point{}
	expectErr := "OptionFailedError,オプション値の指定エラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnSpatialId(inputSpatialID, 3)

	// 戻り値のエラーと期待値の比較
	if expectErr != resultErr.Error() {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("座標 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	t.Log("テスト終了")
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

	// 入力値
	inputSpatialID := "18/232837/103222/18/32"
	inputOption := enum.Center

	// 期待値
	expectVal := []*object.Point{}
	point, _ := object.NewPoint(139.75364685058594, 35.6857446882, 4160.0)
	expectVal = append(expectVal, point)

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnExtendedSpatialId(inputSpatialID, inputOption)

	// 戻り値のエラーと期待値の比較
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s", resultErr)
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if len(resultVal) != 1 {
		// 中心座標取得結果の確認
		t.Errorf("中心座標 - 取得値：%d", len(resultVal))
	} else if !reflect.DeepEqual(resultVal, expectVal) {
		// 中心座標取得結果の確認
		t.Errorf("中心座標 - 期待値：%+v, 取得値：%+v", expectVal[0], resultVal[0])
	}
	t.Log("テスト終了")
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

	// パターン1の入力パラメータ
	inputSpatialID01 := "3/0/0/18/32"
	inputOption01 := enum.Vertex

	// パターン2の入力パラメータ
	inputSpatialID02 := "3/8/0/18/32"
	inputOption02 := enum.Vertex

	// パターン1の取得値
	resultVal01, resultErr01 := GetPointOnExtendedSpatialId(inputSpatialID01, inputOption01)

	// パターン2の取得値
	resultVal02, resultErr02 := GetPointOnExtendedSpatialId(inputSpatialID02, inputOption02)

	// 戻り値のエラーと期待値の比較
	if resultErr01 != resultErr02 {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 取得値01：%s, 取得値02：%s", resultErr01, resultErr02)
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if len(resultVal01) != 8 {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標 - 取得値：%d", len(resultVal01))
	} else if !reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 期待値：%+v, 取得値：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 期待値：%+v, 取得値：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 期待値：%+v, 取得値：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 期待値：%+v, 取得値：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 期待値：%+v, 取得値：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 期待値：%+v, 取得値：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 期待値：%+v, 取得値：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 期待値：%+v, 取得値：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetPointOnExtendedSpatialId03 拡張空間IDの座標取得関数 拡張空間ID入力値不足のエラー確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (拡張空間ID："18/232837/103222/18", option：enum.Center (拡張空間IDの中心座標を取得))
//
// + 確認内容
//   - 拡張空間IDの入力不備で入力チェックエラーとなること
func TestGetPointOnExtendedSpatialId03(t *testing.T) {

	// 入力値
	inputSpatialID := "18/232837/103222/18"
	inputOption := enum.Center

	// 期待値
	expectVal := []*object.Point{}
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnExtendedSpatialId(inputSpatialID, inputOption)

	// 戻り値のエラーと期待値の比較
	if expectErr != resultErr.Error() {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("座標 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
}

// TestGetPointOnExtendedSpatialId04 拡張空間IDの座標取得関数 拡張空間ID水平方向精度レベルの入力不備エラー確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (拡張空間ID："36/232837/103222/18/32", option：enum.Center (拡張空間IDの中心座標を取得))
//
// + 確認内容
//   - 拡張空間IDの水平方向の精度レベルの入力不備で入力チェックエラーとなること
func TestGetPointOnExtendedSpatialId04(t *testing.T) {

	// 入力値
	inputSpatialID := "36/232837/103222/18/32"
	inputOption := enum.Center

	// 期待値
	expectVal := []*object.Point{}
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnExtendedSpatialId(inputSpatialID, inputOption)

	// 戻り値のエラーと期待値の比較
	if expectErr != resultErr.Error() {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("座標 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetPointOnExtendedSpatialId05 拡張空間IDの座標取得関数 拡張空間ID垂直方向精度レベルの入力不備エラー確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (拡張空間ID："18/232837/103222/-1/32", option：enum.Center (拡張空間IDの中心座標を取得))
//
// + 確認内容
//   - 拡張空間IDの垂直方向の精度レベルの入力不備で入力チェックエラーとなること
func TestGetPointOnExtendedSpatialId05(t *testing.T) {

	// 入力値
	inputSpatialID := "18/232837/103222/-1/32"
	inputOption := enum.Center

	// 期待値
	expectVal := []*object.Point{}
	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnExtendedSpatialId(inputSpatialID, inputOption)

	// 戻り値のエラーと期待値の比較
	if expectErr != resultErr.Error() {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("座標 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetPointOnExtendedSpatialId06 拡張空間IDの座標取得関数 サポート対象外のオプション入力によるエラー確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (拡張空間ID："18/232837/103222/18/32", option：3 (存在しないオプション入力))
//
// + 確認内容
//   - 拡張空間IDの入力不備でオプション値の指定エラーとなること
func TestGetPointOnExtendedSpatialId06(t *testing.T) {

	// 入力値
	inputSpatialID := "18/232837/103222/18/32"

	// 期待値
	expectVal := []*object.Point{}
	expectErr := "OptionFailedError,オプション値の指定エラー"

	// テスト対象呼び出し
	resultVal, resultErr := GetPointOnExtendedSpatialId(inputSpatialID, 3)

	// 戻り値のエラーと期待値の比較
	if expectErr != resultErr.Error() {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(expectVal, resultVal) {
		// 期待値と戻り値に差異があった場合はログに出力
		t.Errorf("座標 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestConvertPointListToProjectedPointList01 地理座標系リストを投影座標系リストに変換する関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 地理座標リスト：[(139.753098, 35.685371, 100.0)], CRS：4326
//
// + 確認内容
//   - 入力値に対して、空間IDの中心点の投影座標が格納されたインスタンスのリストが適切に出力されること。
func TestConvertPointListToProjectedPointList01(t *testing.T) {
	// 入力値
	X := 139.753098
	Y := 35.685371
	Z := 100.0
	p1, _ := object.NewPoint(X, Y, Z)
	pointList := []*object.Point{p1}
	projectedCrs := 4326

	// 戻り値格納用変数
	expectpList := make([]*object.ProjectedPoint, 0, len(pointList))
	// 期待値
	expectpList = append(expectpList, &object.ProjectedPoint{
		X:   X,
		Y:   Y,
		Alt: Z,
	})

	// テスト対象呼び出し
	resultpList, resultErr := ConvertPointListToProjectedPointList(pointList, projectedCrs)

	// 戻り値の空間IDの各頂点、または中心点の投影座標が格納されたインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultpList, expectpList) {
		// 戻り値のリストが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID投影座標リスト - 期待値：%+v, 取得値：%+v", expectpList[0], resultpList[0])
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}
	t.Log("テスト終了")
}

// TestConvertPointListToProjectedPointList02 地理座標系リストを投影座標系リストに変換する関数 存在しないEPSGコード動作確認
// 試験詳細：
// + 試験データ
//   - 地理座標リスト：[(139.753098, 35.685371, 100.0)], CRS：1
//
// + 確認内容
//   - 存在しないEPSGコードが入力された場合のエラーインスタンスが返却されること。
func TestConvertPointListToProjectedPointList02(t *testing.T) {
	// 入力値
	X := 139.753098
	Y := 35.685371
	Z := 100.0
	p1, _ := object.NewPoint(X, Y, Z)
	pointList := []*object.Point{p1}
	projectedCrs := 1

	// 戻り値格納用変数
	expectpList := make([]*object.ProjectedPoint, 0, len(pointList))
	geoCrsCode := wgs84.EPSG().Code(consts.GeoCrs)
	proCrsCode := wgs84.EPSG().Code(projectedCrs)

	// 地理座標を投影座標に変換する。
	x, y, _, _ := wgs84.SafeTransform(geoCrsCode, proCrsCode)(X, Y, Z)
	// x,yに0以外が入ったらエラー
	if x != 0.0 {
		t.Errorf("投影座標X - 期待値：0.0, 取得値：%+v", x)
	}
	if y != 0.0 {
		t.Errorf("投影座標Y - 期待値：0.0, 取得値：%+v", y)
	}

	expectErr := "ValueConvertError,値の変換エラー"

	// テスト対象呼び出し
	resultpList, resultErr := ConvertPointListToProjectedPointList(pointList, projectedCrs)

	// 戻り値の空間IDの各頂点、または中心点の投影座標が格納されたインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultpList, expectpList) {
		// 戻り値のリストが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID投影座標リスト - 期待値：%+v, 取得値：%+v", expectpList, resultpList)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}
	t.Log("テスト終了")
}

// TestConvertPointListToProjectedPointList03 地理座標系リストを投影座標系リストに変換する関数 空入力時の動作
// 試験詳細：
// + 試験データ
//
// + 確認内容
func TestConvertPointListToProjectedPointList03(t *testing.T) {
	//空ケースになることはないと保証されている。
	t.Log("テスト終了")
}

// TestConvertProjectedPointListToPointList01 投影座標系リストを地理座標系リストに変換する関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 投影座標系リスト：[(139.753098 35.685371 100.0)], 投影座標のCRS：4326
//
// + 確認内容
//   - 入力座標群に対応した地理座標系リストが取得できること
func TestConvertProjectedPointListToPointList01(t *testing.T) {
	// 入力値
	X := 139.753098
	Y := 35.685371
	Z := 100.0
	pp1 := &object.ProjectedPoint{X: X, Y: Y, Alt: Z}
	projectedPointList := []*object.ProjectedPoint{pp1}
	projectedCrs := 4326

	// 戻り値格納用変数
	expectpList := make([]*object.Point, 0, len(projectedPointList))
	newPoint, _ := object.NewPoint(X, Y, Z)

	// 期待値
	expectpList = append(expectpList, newPoint)

	// テスト対象呼び出し
	resultpList, resultErr := ConvertProjectedPointListToPointList(projectedPointList, projectedCrs)

	// 戻り値の空間IDの各頂点、または中心点の地理座標が格納されたインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultpList, expectpList) {
		// 戻り値のリストが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID地理座標リスト - 期待値：%+v, 取得値：%+v", expectpList[0], resultpList[0])
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestConvertProjectedPointListToPointList02 投影座標系リストを地理座標系リストに変換する関数 存在しないEPSGコード動作確認
// 試験詳細：
// + 試験データ
//   - 投影座標系リスト：[(139.753098 35.685371 100.0)], 投影座標のCRS：1
//
// + 確認内容
//   - 存在しないEPSGコードが入力された場合のエラーインスタンスが返却されること。
func TestConvertProjectedPointListToPointList02(t *testing.T) {
	// 入力値
	X := 139.753098
	Y := 35.685371
	Z := 100.0
	pp1 := &object.ProjectedPoint{X: X, Y: Y, Alt: Z}
	projectedPointList := []*object.ProjectedPoint{pp1}
	projectedCrs := 1

	// 戻り値格納用変数
	expectpList := make([]*object.Point, 0, len(projectedPointList))

	geoCrsCode := wgs84.EPSG().Code(consts.GeoCrs)
	proCrsCode := wgs84.EPSG().Code(projectedCrs)

	// 投影座標を地理座標に変換する。(失敗させる)
	x, y, _, _ := wgs84.SafeTransform(proCrsCode, geoCrsCode)(X, Y, Z)
	// x,yに0以外が入ったらエラー
	if x != 0.0 {
		t.Errorf("地理座標X - 期待値：0.0, 取得値：%+v", x)
	}
	if y != 0.0 {
		t.Errorf("地理座標Y - 期待値：0.0, 取得値：%+v", y)
	}

	// 期待値
	expectErr := "ValueConvertError,値の変換エラー"

	// テスト対象呼び出し
	resultpList, resultErr := ConvertProjectedPointListToPointList(projectedPointList, projectedCrs)

	// 戻り値の空間IDの各頂点、または中心点の地理座標が格納されたインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultpList, expectpList) {
		// 戻り値のリストが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID地理座標リスト - 期待値：%+v, 取得値：%+v", expectpList, resultpList)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestConvertProjectedPointListToPointList03 投影座標系リストを地理座標系リストに変換する関数 空入力時の動作
// 試験詳細：
// + 試験データ
//
// + 確認内容
func TestConvertProjectedPointListToPointList03(t *testing.T) {
	//空ケースになることはないと保証されている。
	t.Log("テスト終了")
}

// TestCheckZoom01 入力精度チェック関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 精度：0
//
// + 確認内容
//   - 入力値に対して、0-35の範囲内の場合Trueを返却すること。
func TestCheckZoom01(t *testing.T) {
	// テスト用入力パラメータ
	var Zoom int64 = 0

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := CheckZoom(Zoom)

	// 戻り値のboolと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値のboolが期待値と異なる場合Errorをログに出力
		t.Errorf("戻り値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestCheckZoom02 入力精度チェック関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 精度：1
//
// + 確認内容
//   - 入力値に対して、0-35の範囲内の場合Trueを返却すること。
func TestCheckZoom02(t *testing.T) {
	// テスト用入力パラメータ
	var Zoom int64 = 1

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := CheckZoom(Zoom)

	// 戻り値のboolと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値のboolが期待値と異なる場合Errorをログに出力
		t.Errorf("戻り値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestCheckZoom03 入力精度チェック関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 精度：34
//
// + 確認内容
//   - 入力値に対して、0-35の範囲内の場合Trueを返却すること。
func TestCheckZoom03(t *testing.T) {
	// テスト用入力パラメータ
	var Zoom int64 = 34

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := CheckZoom(Zoom)

	// 戻り値のboolと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値のboolが期待値と異なる場合Errorをログに出力
		t.Errorf("戻り値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestCheckZoom04 入力精度チェック関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 精度：35
//
// + 確認内容
//   - 入力値に対して、0-35の範囲内の場合Trueを返却すること。
func TestCheckZoom04(t *testing.T) {
	// テスト用入力パラメータ
	var Zoom int64 = 35

	// 期待値
	expectVal := true

	// テスト対象呼び出し
	resultVal := CheckZoom(Zoom)

	// 戻り値のboolと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値のboolが期待値と異なる場合Errorをログに出力
		t.Errorf("戻り値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestCheckZoom05 入力精度チェック関数 精度下限超過動作確認
// 試験詳細：
// + 試験データ
//   - 精度：-1
//
// + 確認内容
//   - 入力値に対して、0-35の範囲外の場合Falseを返却すること。
func TestCheckZoom05(t *testing.T) {
	// テスト用入力パラメータ
	var Zoom int64 = -1

	// 期待値
	expectVal := false

	// テスト対象呼び出し
	resultVal := CheckZoom(Zoom)

	// 戻り値のboolと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値のboolが期待値と異なる場合Errorをログに出力
		t.Errorf("戻り値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestCheckZoom06 入力精度チェック関数 精度上限超過動作確認
// 試験詳細：
// + 試験データ
//   - 精度：36
//
// + 確認内容
//   - 入力値に対して、0-35の範囲外の場合Falseを返却すること。
func TestCheckZoom06(t *testing.T) {
	// テスト用入力パラメータ
	var Zoom int64 = 36

	// 期待値
	expectVal := false

	// テスト対象呼び出し
	resultVal := CheckZoom(Zoom)

	// 戻り値のboolと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値のboolが期待値と異なる場合Errorをログに出力
		t.Errorf("戻り値 - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdAttrs01 拡張空間IDフォーマットチェック関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 拡張空間ID："18/232837/103222/18/32"
//
// + 確認内容
//   - 入力拡張空間IDに対応した以下のint64型配列が取得できること
//     [水平精度, X成分, Y成分, 垂直精度, Z成分]
//     {18, 232837, 103222, 18, 32}
func TestGetExtendedSpatialIdAttrs01(t *testing.T) {
	//入力パラメータ
	spatialID := "18/232837/103222/18/32"
	// 期待値
	expectVal := []int64{18, 232837, 103222, 18, 32}

	// テスト対象呼び出し
	resultVal, resultErr := getExtendedSpatialIdAttrs(spatialID)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdAttrs02 拡張空間IDフォーマットチェック関数 拡張空間ID区切り文字不足時動作確認
// 試験詳細：
// + 試験データ
//   - 拡張空間ID：""(値未入力)
//
// + 確認内容
//   - 入力拡張空間IDの入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdAttrs02(t *testing.T) {
	//入力パラメータ
	spatialID := ""
	// 期待値
	expectVal := []int64{}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := getExtendedSpatialIdAttrs(spatialID)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdAttrs03 拡張空間IDフォーマットチェック関数 拡張空間ID文字列混入動作確認
// 試験詳細：
// + 試験データ
//   - 拡張空間ID："A/232837/103222/18/32"
//
// + 確認内容
//   - 入力拡張空間IDの入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdAttrs03(t *testing.T) {
	//入力パラメータ
	spatialID := "A/232837/103222/18/32"
	// 期待値
	expectVal := []int64{}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := getExtendedSpatialIdAttrs(spatialID)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdAttrs04 拡張空間IDフォーマットチェック関数 拡張空間ID文字列混入動作確認
// 試験詳細：
// + 試験データ
//   - 拡張空間ID："18/B/103222/18/32"
//
// + 確認内容
//   - 入力拡張空間IDの入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdAttrs04(t *testing.T) {
	//入力パラメータ
	spatialID := "18/B/103222/18/32"
	// 期待値
	expectVal := []int64{18}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := getExtendedSpatialIdAttrs(spatialID)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdAttrs05 拡張空間IDフォーマットチェック関数 拡張空間ID文字列混入動作確認
// 試験詳細：
// + 試験データ
//   - 拡張空間ID："18/232837/C/18/32"
//
// + 確認内容
//   - 入力拡張空間IDの入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdAttrs05(t *testing.T) {
	//入力パラメータ
	spatialID := "18/232837/C/18/32"
	// 期待値
	expectVal := []int64{18, 232837}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := getExtendedSpatialIdAttrs(spatialID)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdAttrs06 拡張空間IDフォーマットチェック関数 拡張空間ID文字列混入動作確認
// 試験詳細：
// + 試験データ
//   - 拡張空間ID："18/232837/103222/D/32"
//
// + 確認内容
//   - 入力拡張空間IDの入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdAttrs06(t *testing.T) {
	//入力パラメータ
	spatialID := "18/232837/103222/D/32"
	// 期待値
	expectVal := []int64{18, 232837, 103222}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := getExtendedSpatialIdAttrs(spatialID)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetExtendedSpatialIdAttrs07 拡張空間IDフォーマットチェック関数 拡張空間ID文字列混入動作確認
// 試験詳細：
// + 試験データ
//   - 拡張空間ID："18/232837/103222/18/E"
//
// + 確認内容
//   - 入力拡張空間IDの入力不備で入力チェックエラーとなること
func TestGetExtendedSpatialIdAttrs07(t *testing.T) {
	//入力パラメータ
	spatialID := "18/232837/103222/18/E"
	// 期待値
	expectVal := []int64{18, 232837, 103222, 18}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := getExtendedSpatialIdAttrs(spatialID)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}
	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestGetHorizontalTileIdOnPoint01 水平方向タイルID取得関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 経度：139.753098, 緯度：35.685371, 水平方向の精度：14)
//
// + 確認内容
//   - 入力値に対して、適切な水平方向のタイルIDを返却されること。
func TestGetHorizontalTileIdOnPoint01(t *testing.T) {
	// 入力パラメータ
	lon := 139.753098
	lat := 35.685371
	var hZoom int64 = 14

	// 期待値
	expectVal := "14/14552/6451"

	// テスト対象呼び出し
	resultVal := getHorizontalTileIdOnPoint(lon, lat, hZoom)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetHorizontalTileIdOnPoint02 水平方向タイルID取得関数 正常系動作確認(経度180度)
// 試験詳細：
// + 試験データ
//   - 経度：180.0, 緯度：85.05112877, 水平方向の精度：1)
//
// + 確認内容
//   - 入力値に対して、適切な水平方向のタイルIDを返却されること。
func TestGetHorizontalTileIdOnPoint02(t *testing.T) {
	// 入力パラメータ
	lon := 180.0
	lat := 85.05112877
	var hZoom int64 = 1

	// 期待値
	expectVal := "1/0/0"

	// テスト対象呼び出し
	resultVal := getHorizontalTileIdOnPoint(lon, lat, hZoom)

	// 戻り値の水平方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の水平方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("水平方向タイルID - 期待値：%+v, 取得値：%+v", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetVerticalTileIdOnAltitude01 垂直方向タイルID取得関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 高さ：56.9, 垂直方向の精度：25
//
// + 確認内容
//   - 入力値に対して、適切な垂直方向のタイルIDを返却されること。
func TestGetVerticalTileIdOnAltitude01(t *testing.T) {
	// テスト用入力パラメータ
	alt := 56.9
	var vzoom int64 = 25

	// 期待値
	expectVal := "25/56"

	// テスト対象呼び出し
	resultVal := getVerticalTileIdOnAltitude(alt, vzoom)

	// 戻り値の垂直方向のタイルIDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の垂直方向のタイルIDが期待値と異なる場合Errorをログに出力
		t.Errorf("垂直方向タイルID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestGetAltitudeOnVerticalIndexAndZoom01 高さ及び分解能取得関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 高さの位置：56, 垂直方向精度：25
//
// + 確認内容
//   - 入力値に対して、適切な高さの位置と垂直方向精度が返却されること。
func TestGetAltitudeOnVerticalIndexAndZoom01(t *testing.T) {
	// テスト用入力パラメータ
	var altIndex int64 = 56
	var vZoom int64 = 25

	// 期待値
	//高さ期待値
	expectAlt := float64(altIndex) * math.Pow(2, 25.0) / math.Pow(2, float64(vZoom))

	//分解能期待値
	expectResolution := math.Pow(2, 25.0) / math.Pow(2, float64(vZoom))

	// テスト対象呼び出し
	resultPoint := getAltitudeOnVerticalIndexAndZoom(altIndex, vZoom)

	// 高さと期待値の比較
	if !reflect.DeepEqual(resultPoint.Alt, expectAlt) {
		// 高さが期待値と異なる場合Errorをログに出力
		t.Errorf("高さ - 期待値：%+v, 取得値：%+v", expectAlt, resultPoint.Alt)
	}
	// 分解能と期待値の比較
	if !reflect.DeepEqual(resultPoint.Resolution, expectResolution) {
		// 分解能が期待値と異なる場合Errorをログに出力
		t.Errorf("分解能 - 期待値：%+v, 取得値：%+v", expectResolution, resultPoint.Resolution)
	}

	t.Log("テスト終了")
}

//TestGetCenterPointOnVoxelOffset01 ボクセル中心座標取得関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 経度方向のインデックス：2, 緯度方向のインデックス：2, 水平方向の精度：2,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
// + 確認内容
//   - 入力値に対して、適切な空間の中心点の座標が格納されたインスタンスを返却されること。

func TestGetCenterPointOnVoxelOffset01(t *testing.T) {

	//入力パラメータ
	var lonIndex int64 = 2
	var latIndex int64 = 2
	var hZoom int64 = 2
	vPoint := object.VerticalPoint{Alt: 56, Resolution: 1}

	//期待値
	expectPoint, _ := object.NewPoint(45, -33.2566302215, 56.5)

	// テスト対象呼び出し
	resultPoint := getCenterPointOnVoxelOffset(lonIndex, latIndex, hZoom, vPoint)

	// 空間の中心点の座標と期待値の比較
	if reflect.DeepEqual(resultPoint, expectPoint) {
		// 空間の中心点の座標が期待値と異なる場合Errorをログに出力
		t.Errorf("中心座標 - 期待値：%+v, 取得値：%+v", expectPoint, resultPoint)
	}

	t.Log("テスト終了")
}

//TestGetVertexOnVoxelOffset01 ボクセル頂点座標取得関数 正常系動作確認
// 試験詳細：
// + 試験データ
//   - 経度方向のインデックス：0, 緯度方向のインデックス：0, 水平方向の精度：0,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
// + 確認内容
//   - 入力値に対して、各頂点の座標が設定されたインスタンスのリストを返却されること。

func TestGetVertexOnVoxelOffset01(t *testing.T) {
	//入力パラメータ
	var lonIndex int64 = 0
	var latIndex int64 = 0
	var hZoom int64 = 25
	vPoint := object.VerticalPoint{Alt: 56, Resolution: 1}

	//期待値
	ep1, _ := object.NewPoint(-180.0, 85.0511287798, 56.0)
	ep2, _ := object.NewPoint(-179.99998927116394, 85.0511287798, 56.0)
	ep3, _ := object.NewPoint(-179.99998927116394, 85.0511278542, 56.0)
	ep4, _ := object.NewPoint(-180.0, 85.0511278542, 56.0)
	ep5, _ := object.NewPoint(-180.0, 85.0511287798, 57.0)
	ep6, _ := object.NewPoint(-179.99998927116394, 85.0511287798, 57.0)
	ep7, _ := object.NewPoint(-179.99998927116394, 85.0511278542, 57.0)
	ep8, _ := object.NewPoint(-180.0, 85.0511278542, 57.0)

	expectVal := []*object.Point{ep1, ep2, ep3, ep4, ep5, ep6, ep7, ep8}
	// テスト対象呼び出し
	resultVal := getVertexOnVoxelOffset(lonIndex, latIndex, hZoom, vPoint)

	// ボクセル頂点の座標と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// ボクセル頂点座標が期待値と異なる場合Errorをログに出力
		t.Errorf("頂点座標1 - 期待値：%+v, 取得値：%+v", expectVal[0], resultVal[0])
		t.Errorf("頂点座標2 - 期待値：%+v, 取得値：%+v", expectVal[1], resultVal[1])
		t.Errorf("頂点座標3 - 期待値：%+v, 取得値：%+v", expectVal[2], resultVal[2])
		t.Errorf("頂点座標4 - 期待値：%+v, 取得値：%+v", expectVal[3], resultVal[3])
		t.Errorf("頂点座標5 - 期待値：%+v, 取得値：%+v", expectVal[4], resultVal[4])
		t.Errorf("頂点座標6 - 期待値：%+v, 取得値：%+v", expectVal[5], resultVal[5])
		t.Errorf("頂点座標7 - 期待値：%+v, 取得値：%+v", expectVal[6], resultVal[6])
		t.Errorf("頂点座標8 - 期待値：%+v, 取得値：%+v", expectVal[7], resultVal[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset02 ボクセル頂点座標取得関数 境界値（緯度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値±0
//     経度方向のインデックス：0, 緯度方向のインデックス：7, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値+1
//     経度方向のインデックス：0, 緯度方向のインデックス：8, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していること
func TestGetVertexOnVoxelOffset02(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = 0
	var latIndex01 int64 = 7
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 0
	var latIndex02 int64 = 8
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset03 ボクセル頂点座標取得関数 境界値（緯度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値±0
//     経度方向のインデックス：0, 緯度方向のインデックス：7, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値-1
//     経度方向のインデックス：0, 緯度方向のインデックス：6, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していないこと
func TestGetVertexOnVoxelOffset03(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = 0
	var latIndex01 int64 = 7
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 0
	var latIndex02 int64 = 6
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset04 ボクセル頂点座標取得関数 境界値（緯度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値±0
//     経度方向のインデックス：0, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値+1
//     経度方向のインデックス：0, 緯度方向のインデックス：1, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していないこと
func TestGetVertexOnVoxelOffset04(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = 0
	var latIndex01 int64 = 0
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 0
	var latIndex02 int64 = 1
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset05 ボクセル頂点座標取得関数 境界値（緯度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値±0
//     経度方向のインデックス：0, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値-1
//     経度方向のインデックス：0, 緯度方向のインデックス：-1, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していること
func TestGetVertexOnVoxelOffset05(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = 0
	var latIndex01 int64 = 0
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 0
	var latIndex02 int64 = -1
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset06 ボクセル頂点座標取得関数 境界値（経度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値±0
//     経度方向のインデックス：6, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値-1
//     経度方向のインデックス：14, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していること
func TestGetVertexOnVoxelOffset06(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = 6
	var latIndex01 int64 = 0
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 14
	var latIndex02 int64 = 0
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset07 ボクセル頂点座標取得関数 境界値（経度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値±0
//     経度方向のインデックス：8, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値±0
//     経度方向のインデックス：0, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していること
func TestGetVertexOnVoxelOffset07(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = 8
	var latIndex01 int64 = 0
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 0
	var latIndex02 int64 = 0
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset08 ボクセル頂点座標取得関数 境界値（経度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値±0
//     経度方向のインデックス：9, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値+1
//     経度方向のインデックス：1, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していること
func TestGetVertexOnVoxelOffset08(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = 9
	var latIndex01 int64 = 0
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 1
	var latIndex02 int64 = 0
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset09 ボクセル頂点座標取得関数 境界値（経度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値-1
//     経度方向のインデックス：-1, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値±0
//     経度方向のインデックス：1, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していること
func TestGetVertexOnVoxelOffset09(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = -1
	var latIndex01 int64 = 0
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 31
	var latIndex02 int64 = 0
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if !reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestGetVertexOnVoxelOffset10 ボクセル頂点座標取得関数 境界値（経度の取得）確認
// 試験詳細：
// + 試験データ
//   - パターン1：境界値+1
//     経度方向のインデックス：0, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//   - パターン2：境界値±0
//     経度方向のインデックス：7, 緯度方向のインデックス：0, 水平方向の精度：3,
//     垂直方向インスタンス：object.VerticalPoint{Alt: 56, Resolution: 1})
//
// + 確認内容
//   - パターン1とパターン2の結果が一致していること
func TestGetVertexOnVoxelOffset10(t *testing.T) {
	// パターン1の入力パラメータ
	var lonIndex01 int64 = 0
	var latIndex01 int64 = 0
	var hZoom01 int64 = 3
	vPoint01 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン2の入力パラメータ
	var lonIndex02 int64 = 7
	var latIndex02 int64 = 0
	var hZoom02 int64 = 3
	vPoint02 := object.VerticalPoint{Alt: 56, Resolution: 1}

	// パターン1の取得値
	resultVal01 := getVertexOnVoxelOffset(lonIndex01, latIndex01, hZoom01, vPoint01)

	// パターン2の取得値
	resultVal02 := getVertexOnVoxelOffset(lonIndex02, latIndex02, hZoom02, vPoint02)

	// 戻り値のインスタンスのリストと期待値の比較
	if reflect.DeepEqual(resultVal01, resultVal02) {
		// 頂点座標取得結果の確認
		t.Errorf("頂点座標1 - 取得値01：%+v, 取得値02：%+v", resultVal01[0], resultVal02[0])
		t.Errorf("頂点座標2 - 取得値01：%+v, 取得値02：%+v", resultVal01[1], resultVal02[1])
		t.Errorf("頂点座標3 - 取得値01：%+v, 取得値02：%+v", resultVal01[2], resultVal02[2])
		t.Errorf("頂点座標4 - 取得値01：%+v, 取得値02：%+v", resultVal01[3], resultVal02[3])
		t.Errorf("頂点座標5 - 取得値01：%+v, 取得値02：%+v", resultVal01[4], resultVal02[4])
		t.Errorf("頂点座標6 - 取得値01：%+v, 取得値02：%+v", resultVal01[5], resultVal02[5])
		t.Errorf("頂点座標7 - 取得値01：%+v, 取得値02：%+v", resultVal01[6], resultVal02[6])
		t.Errorf("頂点座標8 - 取得値01：%+v, 取得値02：%+v", resultVal01[7], resultVal02[7])
	}

	t.Log("テスト終了")
}

// TestConvertSpatialIdsToExtendedSpatialIds01 空間IDフォーマット変換関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID：["18/32/232837/103222"])
//
// + 確認内容
//   - 入力値の空間IDに対して、拡張空間IDが適切に変換されること。
func TestConvertSpatialIdsToExtendedSpatialIds01(t *testing.T) {

	// 入力値
	inputSpatialIDs := []string{"18/32/232837/103222"}

	// 期待値
	expectVal := []string{"18/232837/103222/18/32"}

	// テスト対象呼び出し
	resultVal, resultErr := ConvertSpatialIdsToExtendedSpatialIds(inputSpatialIDs)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}
	t.Log("テスト終了")
}

// TestConvertSpatialIdsToExtendedSpatialIds02 空間IDフォーマット変換関数 複数入力時動作確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID：["18/32/232837/103222", "19/32/232837/103222", "20/32/232837/103222"])
//
// + 確認内容
//   - 入力値の空間IDに対して、拡張空間IDが適切に変換されること。
func TestConvertSpatialIdsToExtendedSpatialIds02(t *testing.T) {

	// 入力値
	inputSpatialIDs := []string{"18/32/232837/103222", "19/32/232837/103222", "20/32/232837/103222"}

	// 期待値
	expectVal := []string{"18/232837/103222/18/32", "19/232837/103222/19/32", "20/232837/103222/20/32"}

	// テスト対象呼び出し
	resultVal, resultErr := ConvertSpatialIdsToExtendedSpatialIds(inputSpatialIDs)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}
	t.Log("テスト終了")
}

// TestConvertSpatialIdsToExtendedSpatialIds03 空間IDフォーマット変換関数 区切り文字が少ない場合
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID：["32/232837/103222"])
//
// + 確認内容
//   - 入力値に対して、入力チェックエラーが返されること。
func TestConvertSpatialIdsToExtendedSpatialIds03(t *testing.T) {

	// 入力値
	inputSpatialIDs := []string{"32/232837/103222"}

	// 期待値
	expectVal := []string{}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := ConvertSpatialIdsToExtendedSpatialIds(inputSpatialIDs)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}
	t.Log("テスト終了")
}

// TestConvertSpatialIdsToExtendedSpatialIds04 空間IDフォーマット変換関数 区切り文字が多い場合
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID：["18/32/232837/103222/18"])
//
// + 確認内容
//   - 入力値に対して、入力チェックエラーが返されること。
func TestConvertSpatialIdsToExtendedSpatialIds04(t *testing.T) {

	// 入力値
	inputSpatialIDs := []string{"18/32/232837/103222/18"}

	// 期待値
	expectVal := []string{}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := ConvertSpatialIdsToExtendedSpatialIds(inputSpatialIDs)

	// 戻り値の拡張空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の拡張空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("拡張空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}
	t.Log("テスト終了")
}

// TestConvertSpatialIdsToExtendedSpatialIds05 空間IDフォーマット変換関数 空配列
// 試験詳細：
// + 試験データ
//
// + 確認内容
func TestConvertSpatialIdsToExtendedSpatialIds05(t *testing.T) {
	//空ケースになることはないと保証されている。
	t.Log("テスト終了")
}

// TestConvertExtendedSpatialIdsToSpatialIds01 拡張空間IDフォーマット変換関数 正常動作確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID：["18/232837/103222/18/32"])
//
// + 確認内容
//   - 入力値の拡張空間IDに対して、空間IDが適切に変換されること。
func TestConvertExtendedSpatialIdsToSpatialIds01(t *testing.T) {

	// 入力値
	inputSpatialIDs := []string{"18/232837/103222/18/32"}

	// 期待値
	expectVal := []string{"18/32/232837/103222"}

	// テスト対象呼び出し
	resultVal, resultErr := ConvertExtendedSpatialIdsToSpatialIds(inputSpatialIDs)

	// 戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}
	t.Log("テスト終了")
}

// TestConvertExtendedSpatialIdsToSpatialIds02 拡張空間IDフォーマット変換関数 複数入力時動作確認
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID：["18/232837/103222/18/32", "19/232837/103222/19/32", "20/232837/103222/20/32"])
//
// + 確認内容
//   - 入力値の空間IDに対して、拡張空間IDが適切に変換されること。
func TestConvertExtendedSpatialIdsToSpatialIds02(t *testing.T) {

	// 入力値
	inputSpatialIDs := []string{"18/232837/103222/18/32", "19/232837/103222/19/32", "20/232837/103222/20/32"}

	// 期待値
	expectVal := []string{"18/32/232837/103222", "19/32/232837/103222", "20/32/232837/103222"}

	// テスト対象呼び出し
	resultVal, resultErr := ConvertExtendedSpatialIdsToSpatialIds(inputSpatialIDs)

	// 戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}
	t.Log("テスト終了")
}

// TestConvertExtendedSpatialIdsToSpatialIds03 拡張空間IDフォーマット変換関数 区切り文字が少ない場合
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID：["18/232837/103222/18/32"])
//
// + 確認内容
//   - 入力値に対して、入力チェックエラーが返されること。
func TestConvertExtendedSpatialIdsToSpatialIds03(t *testing.T) {

	// 入力値
	inputSpatialIDs := []string{"18/232837/103222/32"}

	// 期待値
	expectVal := []string{}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := ConvertExtendedSpatialIdsToSpatialIds(inputSpatialIDs)

	// 戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}
	t.Log("テスト終了")
}

// TestConvertExtendedSpatialIdsToSpatialIds04 拡張空間IDフォーマット変換関数 区切り文字が多い場合
// 試験詳細：
// + 試験データ
//   - パターン：
//     (空間ID：["18/232837/103222/18/32/55"])
//
// + 確認内容
//   - 入力値に対して、入力チェックエラーが返されること。
func TestConvertExtendedSpatialIdsToSpatialIds04(t *testing.T) {

	// 入力値
	inputSpatialIDs := []string{"18/232837/103222/18/32/55"}

	// 期待値
	expectVal := []string{}

	expectErr := "InputValueError,入力チェックエラー"

	// テスト対象呼び出し
	resultVal, resultErr := ConvertExtendedSpatialIdsToSpatialIds(inputSpatialIDs)

	// 戻り値の空間IDと期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値の空間IDが期待値と異なる場合Errorをログに出力
		t.Errorf("空間ID - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	if resultErr.Error() != expectErr {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s", expectErr, resultErr.Error())
	}
	t.Log("テスト終了")
}

// TestConvertExtendedSpatialIdsToSpatialIds05 拡張空間IDフォーマット変換関数 空配列
// 試験詳細：
// + 試験データ
//
// + 確認内容
func TestConvertExtendedSpatialIdsToSpatialIds05(t *testing.T) {
	//空ケースになることはないと保証されている。
	t.Log("テスト終了")
}

// object ライブラリ内で共通的に使用するオブジェクトを管理するパッケージ。
package spatialID

import (
	"reflect"
	"testing"
)

const expectErr = "InputValueError,入力チェックエラー"

// TestNewPoint01 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(179.9, 85.0511287798, 0.1)])
//   + 確認内容
//     - 入力された経度、緯度、高さが設定されたPointオブジェクトが返却されること
func TestNewPoint01(t *testing.T) {
	// テスト用入力パラメータ
	lon := 179.9
	lat := 85.0511287798
	alt := 0.1

	// 期待値
	expectVal := &Point{lon, lat, alt}
	// テスト対象呼び出し
	resultVal, resultErr := NewPoint(lon, lat, alt)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultVal, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestNewPoint02 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(180.0, 85.0511287798, 0.1)])
//   + 確認内容
//     - 入力された経度、緯度、高さが設定されたPointオブジェクトが返却されること
func TestNewPoint02(t *testing.T) {
	// テスト用入力パラメータ
	lon := 180.0
	lat := 85.0511287798
	alt := 0.1

	// 期待値
	expectVal := &Point{lon, lat, alt}
	// テスト対象呼び出し
	resultVal, resultErr := NewPoint(lon, lat, alt)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultVal, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestNewPoint03 エラー系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(180.1, 85.0511287798, 0.1)])
//   + 確認内容
//     - 入力された経度、緯度、高さが設定されたPointオブジェクトが返却されること
func TestNewPoint03(t *testing.T) {
	// テスト用入力パラメータ
	lon := 180.1
	lat := 85.0511287798
	alt := 0.1

	// 期待値
	expectVal := &Point{0.0, 0.0, 0.0}

	// テスト対象呼び出し
	resultVal, resultErr := NewPoint(lon, lat, alt)

	//　正常系動作の検証
	verificationFail("Point", t, resultVal, resultErr, expectVal, expectErr)

	t.Log("テスト終了")
}

// TestNewPoint04 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//    - パターン1：
//       (Pointオブジェクト:[(180.0, 85.0511287797, 0.1)])
//   + 確認内容
//     - 入力された経度、緯度、高さが設定されたPointオブジェクトが返却されること
func TestNewPoint04(t *testing.T) {
	// テスト用入力パラメータ
	lon := 180.0
	lat := 85.0511287797
	alt := 0.1

	// 浮動小数点誤差により取得値と入力値に差が発生するため、期待値用の緯度を定義
	latExpect := 85.0511287796

	// 期待値
	expectVal := &Point{lon, latExpect, alt}
	// テスト対象呼び出し
	resultVal, resultErr := NewPoint(lon, lat, alt)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultVal, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestNewPoint05 エラー系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(180.0, 85.0511287799, 0.1)])
//   + 確認内容
//     - 入力された経度、緯度、高さが設定されたPointオブジェクトが返却されること
func TestNewPoint05(t *testing.T) {
	// テスト用入力パラメータ
	lon := 180.0
	lat := 85.0511287799
	alt := 0.1

	// 期待値
	expectVal := &Point{lon, 0.0, 0.0}

	// テスト対象呼び出し
	resultVal, resultErr := NewPoint(lon, lat, alt)

	//　検証
	verificationFail("Point", t, resultVal, resultErr, expectVal, expectErr)

	t.Log("テスト終了")
}

// TestSetLon01 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(179.0, 0.0, 0.0)])
//   + 確認内容
//     - 入力された経度がPointオブジェクトに設定されていること
func TestSetLon01(t *testing.T) {
	// テスト用入力パラメータ
	lon := 179.0

	// 期待値
	expectVal := &Point{lon, 0.0, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLon(lon)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultPoint, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestSetLon02 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(180.0, 0.0, 0.0)])
//   + 確認内容
//     -  入力された経度がPointオブジェクトに設定されていること
func TestSetLon02(t *testing.T) {
	// テスト用入力パラメータ
	lon := 180.0

	// 期待値
	expectVal := &Point{lon, 0.0, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLon(lon)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultPoint, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestSetLon03 エラー系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//     (Pointオブジェクト:[(181.0, 0.0, 0.0)])
//   + 確認内容
//     -  入力された経度がPointオブジェクトに設定されずエラーとなること
func TestSetLon03(t *testing.T) {
	// テスト用入力パラメータ
	lon := 181.0

	// 期待値
	expectVal := &Point{0.0, 0.0, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLon(lon)

	//　正常系動作の検証
	verificationFail("Point", t, resultPoint, resultErr, expectVal, expectErr)

	t.Log("テスト終了")
}

// TestSetLat01 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, 0.1, 0.0)])
//   + 確認内容
//     - 入力された緯度がPointオブジェクトに設定されていること
func TestSetLat01(t *testing.T) {
	// テスト用入力パラメータ
	lat := 0.1

	// 期待値
	expectVal := &Point{0.0, lat, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLat(lat)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultPoint, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestSetLat02 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, 0.0, 0.0)])
//   + 確認内容
//     - 入力された緯度がPointオブジェクトに設定されていること
func TestSetLat02(t *testing.T) {
	// テスト用入力パラメータ
	lat := 0.0

	// 期待値
	expectVal := &Point{0.0, lat, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLat(lat)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultPoint, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestSetLat03 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, -0.1, 0.0)])
//   + 確認内容
//     - 入力された緯度がPointオブジェクトに設定されていること
func TestSetLat03(t *testing.T) {
	// テスト用入力パラメータ
	lat := -0.1

	// 期待値
	expectVal := &Point{0.0, lat, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLat(lat)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultPoint, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestSetLat04 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, 85.0511287797, 0.0)])
//   + 確認内容
//     - 入力された緯度がPointオブジェクトに設定されていること
func TestSetLat04(t *testing.T) {
	// テスト用入力パラメータ
	lat := 85.0511287797

	// 浮動小数点誤差により取得値と入力値に差が発生するため、期待値用の緯度を定義
	latExpect := 85.0511287796

	// 期待値
	expectVal := &Point{0.0, latExpect, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLat(lat)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultPoint, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestSetLat05 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, 85.0511287798, 0.0)])
//   + 確認内容
//     - 入力された緯度がPointオブジェクトに設定されていること
func TestSetLat05(t *testing.T) {
	// テスト用入力パラメータ
	lat := 85.0511287798

	// 期待値
	expectVal := &Point{0.0, lat, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLat(lat)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultPoint, resultErr, expectVal)

	t.Log("テスト終了")
}

// TestSetLat06 エラー系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, 85.0511287799, 0.0)])
//   + 確認内容
//     - 入力された緯度がPointオブジェクトに設定されずエラーが返却されること
func TestSetLat06(t *testing.T) {
	// テスト用入力パラメータ
	lat := 85.0511287799

	// 期待値
	expectVal := &Point{0.0, 0.0, 0.0}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultErr := resultPoint.SetLat(lat)

	//　エラー系動作の検証
	verificationFail("Point", t, resultPoint, resultErr, expectVal, expectErr)

	t.Log("テスト終了")
}

// TestSetAlt01 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, 0.0, 1.0)])
//   + 確認内容
//     - 入力された高さがPointオブジェクトに設定されていること
func TestSetAlt01(t *testing.T) {
	// テスト用入力パラメータ
	alt := 1.0

	// 期待値
	expectVal := &Point{0.0, 0.0, alt}
	resultPoint := &Point{}

	// テスト対象呼び出し
	resultPoint.SetAlt(alt)

	//　正常系動作の検証
	verificationSuccess("Point", t, resultPoint, nil, expectVal)

	t.Log("テスト終了")
}

// TestLon01 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクトのlon:180.0)
//   + 確認内容
//     - 取得した経度が期待値と一致すること
func TestLon01(t *testing.T) {
	// テスト用入力パラメータ&期待値
	expectVal := 180.0

	// 期待値
	resultPoint := &Point{expectVal, 0.0, 0.0}

	// テスト対象呼び出し&正常系動作の検証
	if !reflect.DeepEqual(resultPoint.Lon(), expectVal) {
		// 戻り値のPointオブジェクトが期待値と異なる場合Errorをログに出力
		t.Errorf("lon - 期待値：%g, 取得値：%g", resultPoint.Lon(), expectVal)
	}

	t.Log("テスト終了")
}

// TestLat01 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, 85.0511287798, 0.0)])
//   + 確認内容
//     - 入力された経度、緯度、高さが設定されたPointオブジェクトが返却されること
func TestLat01(t *testing.T) {
	// テスト用入力パラメータ&期待値
	expectVal := 85.0511287798

	// 期待値
	resultPoint := &Point{0.0, expectVal, 0.0}

	// テスト対象呼び出し&正常系動作の検証
	if !reflect.DeepEqual(resultPoint.Lat(), expectVal) {
		// 戻り値のPointオブジェクトが期待値と異なる場合Errorをログに出力
		t.Errorf("lat - 期待値：%g, 取得値：%g", resultPoint.Lon(), expectVal)
	}

	t.Log("テスト終了")
}

// TestAlt01 正常系動作確認
//
// 試験詳細：
//   + 試験データ
//     - パターン1：
//       (Pointオブジェクト:[(0.0, 0.0, 1.0)])
//   + 確認内容
//     - 入力された経度、緯度、高さが設定されたPointオブジェクトが返却されること
func TestAlt01(t *testing.T) {
	// テスト用入力パラメータ&期待値
	expectVal := 1.0

	// 期待値
	resultPoint := &Point{0.0, 0.0, expectVal}

	// テスト対象呼び出し&正常系動作の検証
	if !reflect.DeepEqual(resultPoint.Alt(), expectVal) {
		// 戻り値のPointオブジェクトが期待値と異なる場合Errorをログに出力
		t.Errorf("Alt - 期待値：%g, 取得値：%g", resultPoint.Alt(), expectVal)
	}

	t.Log("テスト終了")
}

// verificationSuccess 正常系動作の検証
//
// 引数：
//
//	name：テストケース名
//	ｔ：testing
//	resultVal：試験結果
//	resultErr：エラー結果
//	expectVal：期待値
//
// 備考：
// 期待値と一致しない場合は試験NGの出力
// エラー結果はnilでない場合は試験NGの出力
func verificationSuccess(name string, t *testing.T, resultVal *Point, resultErr error, expectVal *Point) {
	// 戻り値のPointオブジェクトの値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値のPointオブジェクトが期待値と異なる場合Errorをログに出力
		t.Errorf("%s - 期待値：%g, 取得値：%g", name, expectVal, resultVal)
	}

	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}
}

// verificationFail エラー系動作の検証
//
// 引数：
//
//	name：テストケース名
//	ｔ：testing
//	resultVal：試験結果
//	resultErr：エラー結果
//	expectVal：期待値
//	expectErr：期待エラー
//
// 備考：
// 期待値と一致しない場合は試験NGの出力
// エラー結果がnilの場合は試験NGの出力
func verificationFail(name string, t *testing.T, resultVal *Point, resultErr error, expectVal *Point, expectErr string) {
	// 戻り値のPointオブジェクトの値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		// 戻り値のPointオブジェクトが期待値と異なる場合Errorをログに出力
		t.Errorf("%s - 期待値：%g, 取得値：%g", name, expectVal, resultVal)
	}

	if resultErr == nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：nil", expectErr)
	}
}

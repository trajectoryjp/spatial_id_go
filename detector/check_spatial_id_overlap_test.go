package detector

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"testing"

	sperrors "github.com/trajectoryjp/spatial_id_go/v4/common/errors"
)

// TestCheckSpatialIdsOverlap01 空間IDの重複確認関数 正常系動作確認
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"13/0/7274/3225"}, {"16/0/58198/25804"}
//
// + 確認内容
//   - 入力値から包含関係があることを確認できること
func TestCheckSpatialIdsOverlap01(t *testing.T) {
	// 入力値
	spatialId1 := "13/0/7274/3225"
	spatialId2 := "16/0/58198/25804"
	resultValue, resultErr := CheckSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := true

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsOverlap02 空間IDの重複確認関数 正常系動作確認
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"13/0/7275/3226"}, {"16/0/58198/25804"}
//
// + 確認内容
//   - 入力値から包含関係がないことを確認できること
func TestCheckSpatialIdsOverlap02(t *testing.T) {
	// 入力値
	spatialId1 := "13/0/7275/3226"
	spatialId2 := "16/0/58198/25804"
	resultValue, resultErr := CheckSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := false

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsOverlap03 空間IDの重複確認関数 空入力時動作確認
//
// 試験詳細：
//   - パターン1：
//     比較対象の空間ID：空
//
// + 確認内容
//   - エラーを取得できること
func TestCheckSpatialIdsOverlap03(t *testing.T) {
	// 入力値
	spatialId1 := ""
	spatialId2 := "16/0/58198/25804"
	resultValue, resultErr := CheckSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := false
	expectError := "InputValueError,入力チェックエラー,spatialId:  @spatialId1[0]"

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr.Error() != expectError {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectError, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsOverlap04 空間IDの重複確認関数 空間IDフォーマット不正
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"25/0/29803148/13212522/777"},{"16/0/58198/25804"}
//   - パターン2：
//     比較対象の空間ID：{"25/0/29803148/13212522"},{"16/0/58198/25804/777"}
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestCheckSpatialIdsOverlap04(t *testing.T) {
	testCases := []struct {
		spatialId1  string
		spatialId2  string
		expectValue bool
		expectError string
	}{
		{
			//入力値
			spatialId1: "25/0/29803148/13212522/777",
			spatialId2: "16/0/58198/25804",
			// 期待値
			expectValue: false,
			expectError: "InputValueError,入力チェックエラー,spatialId: 25/0/29803148/13212522/777 @spatialId1[0]",
		},
		{
			//入力値
			spatialId1: "25/0/29803148/13212522",
			spatialId2: "16/0/58198/25804/777",
			// 期待値
			expectValue: false,
			expectError: "InputValueError,入力チェックエラー,spatialId: 16/0/58198/25804/777 @spatialId2[0]",
		},
	}
	for _, testCase := range testCases {
		resultValue, resultErr := CheckSpatialIdsOverlap(testCase.spatialId1, testCase.spatialId2)

		if !reflect.DeepEqual(resultValue, testCase.expectValue) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", testCase.expectValue, resultValue)
		}
		if resultErr.Error() != testCase.expectError {
			// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
			t.Errorf("error - 期待値：%s, 取得値：%s\n", testCase.expectError, resultErr.Error())
		}
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsOverlap05 空間IDの重複確認関数 正常系動作確認
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"13/1/7274/3225"}, {"16/8/58198/25804"}
//
// + 確認内容
//   - fインデックスの包含関係があることを確認できること
func TestCheckSpatialIdsOverlap05(t *testing.T) {
	// 入力値
	spatialId1 := "13/1/7274/3225"
	spatialId2 := "16/8/58198/25804"
	resultValue, resultErr := CheckSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := true

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsOverlap06 空間IDの重複確認関数 正常系動作確認
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"13/1/7274/3225"}, {"16/17/58198/25804"}
//
// + 確認内容
//   - fインデックスの包含関係がないことを確認できること
func TestCheckSpatialIdsOverlap06(t *testing.T) {
	// 入力値
	spatialId1 := "13/1/7274/3225"
	spatialId2 := "16/17/58198/25804"
	resultValue, resultErr := CheckSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := false

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsOverlap07 空間IDの重複確認関数 正常系動作確認
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"13/-1/7274/3225"}, {"16/-8/58198/25804"}
//
// + 確認内容
//   - fインデックスが負数の場合においても入力値から包含関係があることを確認できること
func TestCheckSpatialIdsOverlap07(t *testing.T) {
	// 入力値
	spatialId1 := "13/-1/7274/3225"
	spatialId2 := "16/-8/58198/25804"
	resultValue, resultErr := CheckSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := true

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsOverlap08 空間IDの重複確認関数 正常系動作確認
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"13/-1/7274/3225"}, {"16/-10/58198/25804"}
//
// + 確認内容
//   - fインデックスに関する包含関係がないことを確認できること
func TestCheckSpatialIdsOverlap08(t *testing.T) {
	// 入力値
	spatialId1 := "13/-1/7274/3225"
	spatialId2 := "16/-10/58198/25804"
	resultValue, resultErr := CheckSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := false

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsOverlap09 空間IDの重複確認関数 空間ID高度変換失敗
//
// 試験詳細：
// + 試験データ
//   - パターン1(入力高度範囲外エラー)：
//     比較対象の空間ID：{"25/16777216/0/3225"},{"18/1/58198/25804"}
//   - パターン2(入力高度範囲外エラー)：
//     比較対象の空間ID：{"18/1/58198/25804"},{"25/16777216/0/3225"}
//   - パターン3(入力高度範囲外エラー)：
//     比較対象の空間ID：{"25/-16777217/0/3225"},{"16/8/58198/25804"}
//   - パターン4(入力高度範囲外エラー)：
//     比較対象の空間ID：{"18/1/58198/25804"},{"25/-16777217/0/3225"}
//
// + 確認内容
//   - 入力値から指定した入力チェックエラー、変換エラーを取得できること
func TestCheckSpatialIdsOverlap09(t *testing.T) {
	testCases := []struct {
		spatialId1  string
		spatialId2  string
		expectValue bool
		expectError string
	}{
		{
			// 入力空間IDのfインデックスが不正
			//入力値
			spatialId1: "25/16777216/0/3225",
			spatialId2: "18/1/58198/25804",
			// 期待値
			expectValue: false,
			expectError: "InputValueError,入力チェックエラー,outputIndex=33554432 is out of range at outputZoom=25 @spatialId1[0] = 25/16777216/0/3225",
		},
		{
			// 入力空間IDのfインデックスが不正
			//入力値
			spatialId1: "18/1/58198/25804",
			spatialId2: "25/16777216/0/3225",
			// 期待値
			expectValue: false,
			expectError: "InputValueError,入力チェックエラー,outputIndex=33554432 is out of range at outputZoom=25 @spatialId2[0] = 25/16777216/0/3225",
		},
		{
			// 入力可能な高度インデックス範囲を超えている(下限より小さい)
			//入力値
			spatialId1: "25/-16777217/0/3225",
			spatialId2: "18/1/58198/25804",
			// 期待値
			expectValue: false,
			// 高度変換が負数を許容しないため高度変換エラーになる
			expectError: "InputValueError,入力チェックエラー,outputIndex=-1 is out of range at outputZoom=25 @spatialId1[0] = 25/-16777217/0/3225",
		},
		{
			// 入力可能な高度インデックス範囲を超えている(下限より小さい)
			//入力値
			spatialId1: "18/1/58198/25804",
			spatialId2: "25/-16777217/0/3225",
			// 期待値
			expectValue: false,
			// 高度変換が負数を許容しないため高度変換エラーになる
			expectError: "InputValueError,入力チェックエラー,outputIndex=-1 is out of range at outputZoom=25 @spatialId2[0] = 25/-16777217/0/3225",
		},
	}
	for _, testCase := range testCases {
		resultValue, err := CheckSpatialIdsOverlap(testCase.spatialId1, testCase.spatialId2)
		var resultErr string
		if err != nil {
			resultErr = err.Error()
		}

		if !reflect.DeepEqual(resultValue, testCase.expectValue) {
			t.Errorf("空間ID - 期待値：%v, 取得値：%v", testCase.expectValue, resultValue)
		}
		if resultErr != testCase.expectError {
			// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
			t.Errorf("error - 期待値：%s, 取得値：%s\n", testCase.expectError, resultErr)
		}
	}

	t.Log("テスト終了")
}

// BenchmarkCheckSpatialIdsOverlap01 空間IDの重複確認関数 包含関係あり
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"13/0/7274/3225"}, {"16/0/58198/25804"}(包含関係あり)
//
// + 確認内容
//   - 入力値の先頭に包含関係があった場合の処理速度
func BenchmarkCheckSpatialIdsOverlap01(b *testing.B) {
	// 入力値
	spatialId1 := "13/0/7274/3225"
	spatialId2 := "16/0/58198/25804"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CheckSpatialIdsOverlap(spatialId1, spatialId2)
	}
	b.StopTimer()
	b.Log("テスト終了")
}

// BenchmarkCheckSpatialIdsOverlap02 空間IDの重複確認関数 包含関係なし
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"13/0/7275/3226"}, {"16/0/58198/25804"}(包含関係なし)
//
// + 確認内容
//   - 入力値に包含関係がなかった場合の処理速度
func BenchmarkCheckSpatialIdsOverlap02(b *testing.B) {
	// 入力値
	spatialId1 := "13/0/7275/3226"
	spatialId2 := "16/0/58198/25804"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CheckSpatialIdsOverlap(spatialId1, spatialId2)
	}
	b.StopTimer()
	b.Log("テスト終了")
}

// TestCheckSpatialIdsArrayOverlap01 空間ID列の重複確認関数 正常系動作確認
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID列：{"13/0/7274/3225"}, {"16/0/58198/25800", "16/0/58198/25804"}
//
// + 確認内容
//   - 入力値から包含関係があることを確認できること
func TestCheckSpatialIdsArrayOverlap01(t *testing.T) {
	// 入力値
	spatialIds1 := []string{"13/0/7274/3225"}
	spatialIds2 := []string{"16/0/58198/25800", "16/0/58198/25804"}
	resultValue, resultErr := CheckSpatialIdsArrayOverlap(spatialIds1, spatialIds2)

	// 期待値
	expectValue := true

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsArrayOverlap02 空間ID列の重複確認関数 正常系動作確認
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID列：{"13/0/7275/3226"}, {"16/0/58198/25804"}
//
// + 確認内容
//   - 入力値から包含関係がないことを確認できること
func TestCheckSpatialIdsArrayOverlap02(t *testing.T) {
	// 入力値
	spatialIds1 := []string{"13/0/7275/3226"}
	spatialIds2 := []string{"16/0/58198/25804", "16/0/58198/25805"}
	resultValue, resultErr := CheckSpatialIdsArrayOverlap(spatialIds1, spatialIds2)

	// 期待値
	expectValue := false

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsArrayOverlap03 空間ID列の重複確認関数 空入力時動作確認
//
// 試験詳細：
//   - パターン1：
//     比較対象の空間ID列：空
//
// + 確認内容
//   - エラーを取得できること
func TestCheckSpatialIdsArrayOverlap03(t *testing.T) {
	// 入力値
	spatialIds1 := []string{""}
	spatialIds2 := []string{"16/0/58198/25803", "16/0/58198/25804"}
	resultValue, resultErr := CheckSpatialIdsArrayOverlap(spatialIds1, spatialIds2)

	// 期待値
	expectValue := false
	expectError := "InputValueError,入力チェックエラー,spatialId:  @spatialId1[0]"

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr.Error() != expectError {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectError, resultErr.Error())
	}

	t.Log("テスト終了")
}

// TestCheckSpatialIdsArrayOverlap04 空間ID列の重複確認関数 空間IDフォーマット不正
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の空間ID列：{"25/0/29803148/13212522/777"},{"16/0/58198/25803", "16/0/58198/25804"}
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestCheckSpatialIdsArrayOverlap04(t *testing.T) {
	//入力値
	spatialIds1 := []string{"25/0/29803148/13212522/777"}
	spatialIds2 := []string{"16/0/58198/25803", "16/0/58198/25804"}
	resultValue, resultErr := CheckSpatialIdsArrayOverlap(spatialIds1, spatialIds2)

	// 期待値
	expectValue := false
	expectError := "InputValueError,入力チェックエラー,spatialId: 25/0/29803148/13212522/777 @spatialId1[0]"

	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr.Error() != expectError {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectError, resultErr.Error())
	}

	t.Log("テスト終了")
}

// BenchmarkCheckSpatialIdsArrayOverlap01 空間ID列の重複確認関数 空間IDの重複確認関数 配列ベンチマーク(包含関係あり)
//
// + 試験データ
//   - パターン1(包含関係あり)：
//     比較対象の空間ID列：[100]{"13/0/7274/3225"}, [100]{"16/0/58198/25804"}
//
// + 確認内容
//   - 入力値に包含関係があった場合の処理速度
func BenchmarkCheckSpatialIdsArrayOverlap01(b *testing.B) {
	// 入力値
	spatialIds1 := []string{}
	spatialIds2 := []string{}
	for i := 0; i < 100; i++ {
		spatialIds1 = append(spatialIds1, "13/0/7274/3225")
		spatialIds2 = append(spatialIds2, "16/0/58198/25804")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CheckSpatialIdsArrayOverlap(spatialIds1, spatialIds2)
	}
	b.StopTimer()
	b.Log("テスト終了")
}

// BenchmarkCheckSpatialIdsArrayOverlap02 空間ID列の重複確認関数 空間IDの重複確認関数 配列ベンチマーク(包含関係なし)
//
// + 試験データ
//   - パターン1(包含関係なし)：
//     比較対象の空間ID列：[100]{"13/0/7275/3226"}, [100]{"16/0/58198/25804"}
//
// + 確認内容
//   - 入力値に包含関係がない場合の処理速度
func BenchmarkCheckSpatialIdsArrayOverlap02(b *testing.B) {
	// 入力値
	spatialIds1 := []string{}
	spatialIds2 := []string{}
	for i := 0; i < 100; i++ {
		spatialIds1 = append(spatialIds1, "13/0/7275/3226")
		spatialIds2 = append(spatialIds2, "16/0/58198/25804")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CheckSpatialIdsArrayOverlap(spatialIds1, spatialIds2)
	}
	b.StopTimer()
	b.Log("テスト終了")
}

// TestSpatialIdOverlapDetector SpatialIdOverlapDetectorのテスト
func TestSpatialIdOverlapDetector(t *testing.T) {
	testCases := []struct {
		spatialIds1  []string
		spatialIds2  []string
		expectValue  bool
		expectError1 error
		expectError2 error
	}{
		{
			spatialIds1:  []string{"13/0/7274/3225"},
			spatialIds2:  []string{"16/0/58198/25800", "16/0/58198/25804"},
			expectValue:  true,
			expectError1: nil,
			expectError2: nil,
		},
		{
			spatialIds1:  []string{"13/0/7275/3226"},
			spatialIds2:  []string{"16/0/58198/25804", "16/0/58198/25805"},
			expectValue:  false,
			expectError1: nil,
			expectError2: nil,
		},
		{
			spatialIds1: []string{""},
			spatialIds2: []string{"16/0/58198/25803", "16/0/58198/25804"},
			expectValue: false,
			expectError1: sperrors.NewSpatialIdError(
				sperrors.InputValueErrorCode,
				"spatialId: ",
			),
			// ),
			expectError2: nil,
		},
		{
			spatialIds1:  []string{"16/0/58198/25803", "16/0/58198/25804"},
			spatialIds2:  []string{"25/0/29803148/13212522/777"},
			expectValue:  false,
			expectError1: nil,
			expectError2: sperrors.NewSpatialIdError(
				sperrors.InputValueErrorCode,
				"spatialId: 25/0/29803148/13212522/777",
			),
		},
	}

	for _, testCase := range testCases {
		for _, newSpatialIdOverlapDetector := range []func([]string) (SpatialIdOverlapDetector, error){
			NewSpatialIdGreedyOverlapDetector,
			NewSpatialIdTreeOverlapDetector,
		} {
			spatialIdOverlapDetector, resultError1 := newSpatialIdOverlapDetector(testCase.spatialIds1)
			if !errors.Is(resultError1, testCase.expectError1) {
				// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
				t.Errorf("error - 期待値：%s, 取得値：%s\n", testCase.expectError1, resultError1)
			}
			if resultError1 != nil {
				continue
			}

			resultValue, resultError2 := spatialIdOverlapDetector.IsOverlap(testCase.spatialIds2)
			if !reflect.DeepEqual(resultValue, testCase.expectValue) {
				t.Errorf("空間ID - 期待値：%v, 取得値：%v\n", testCase.expectValue, resultValue)
			}
			if !errors.Is(resultError2, testCase.expectError2) {
				// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
				t.Errorf("error - 期待値：%s, 取得値：%s\n", testCase.expectError2, resultError2)
			}
		}
	}
}

// OverlapDetectorBenchmarkCase OverlapDetectorのベンチマークケース
type OverlapDetectorBenchmarkCase struct {
	name                 string
	detectIdNumberCbrt   int
	detectedIdNumberCbrt int
	detectZoomLevel      int
	detectedZoomLevel    int
	unmatchDepth         int
}

// newOverlapDetectorBenchmarkCase ベンチマークケース名nameを受け取り、OverlapDetectorBenchmarkCaseを返す
func newOverlapDetectorBenchmarkCase(name string) OverlapDetectorBenchmarkCase {
	return OverlapDetectorBenchmarkCase{
		name:                 name,
		detectIdNumberCbrt:   1,
		detectedIdNumberCbrt: 10,
		detectZoomLevel:      23,
		detectedZoomLevel:    23,
		unmatchDepth:         18,
	}
}

// createDetectSpatialIds ベンチマークケースから検知空間IDを生成する
func (benchmarkCase *OverlapDetectorBenchmarkCase) createDetectSpatialIds() []string {
	zoomLevelDiff := max(0, benchmarkCase.detectZoomLevel-benchmarkCase.detectedZoomLevel)
	unmatchHeight := benchmarkCase.detectZoomLevel - benchmarkCase.unmatchDepth

	detectSpatialIds := []string{}
	for f := range benchmarkCase.detectIdNumberCbrt {
		for x := range benchmarkCase.detectIdNumberCbrt {
			for y := range benchmarkCase.detectIdNumberCbrt {
				detectSpatialId := fmt.Sprintf(
					"%d/%d/%d/%d",
					benchmarkCase.detectZoomLevel,
					(1<<unmatchHeight)|(f<<zoomLevelDiff),
					(1<<unmatchHeight)|(x<<zoomLevelDiff),
					(1<<unmatchHeight)|(y<<zoomLevelDiff),
				)
				detectSpatialIds = append(detectSpatialIds, detectSpatialId)
			}
		}
	}
	return detectSpatialIds
}

// createDetectedSpatialIds ベンチマークケースから被検知空間IDを生成する
func (benchmarkCase *OverlapDetectorBenchmarkCase) createDetectedSpatialIds() []string {
	detectedSpatialIds := []string{}
	for f := range benchmarkCase.detectedIdNumberCbrt {
		for x := range benchmarkCase.detectedIdNumberCbrt {
			for y := range benchmarkCase.detectedIdNumberCbrt {
				detectedSpatialId := fmt.Sprintf(
					"%d/%d/%d/%d",
					benchmarkCase.detectedZoomLevel,
					f,
					x,
					y,
				)
				detectedSpatialIds = append(detectedSpatialIds, detectedSpatialId)
			}
		}
	}
	return detectedSpatialIds
}

// BenchmarkOverlapDetectorProcessingTime OverlapDetectorの処理時間に対するベンチマークを行う
func BenchmarkOverlapDetectorProcessingTime(b *testing.B) {
	benchmarkCases := []OverlapDetectorBenchmarkCase{}
	for _, detectIdNumberCbrt := range []int{1, 10} {
		for detectedIdNumberCbrt := 6; detectedIdNumberCbrt <= 10; detectedIdNumberCbrt += 1 {
			name := fmt.Sprintf("Pattern=1,DetectIdNumberCbrt=%d,DetectedIdNumberCbrt=%d", detectIdNumberCbrt, detectedIdNumberCbrt)
			benchmarkCase := newOverlapDetectorBenchmarkCase(name)
			benchmarkCase.detectIdNumberCbrt = detectIdNumberCbrt
			benchmarkCase.detectedIdNumberCbrt = detectedIdNumberCbrt
			benchmarkCases = append(benchmarkCases, benchmarkCase)
		}
		for detectedZoomLevel := 10; detectedZoomLevel <= 35; detectedZoomLevel += 5 {
			name := fmt.Sprintf("Pattern=2,DetectIdNumberCbrt=%d,DetectedZoomLevel=%d", detectIdNumberCbrt, detectedZoomLevel)
			benchmarkCase := newOverlapDetectorBenchmarkCase(name)
			benchmarkCase.detectIdNumberCbrt = detectIdNumberCbrt
			benchmarkCase.detectedZoomLevel = detectedZoomLevel
			benchmarkCase.unmatchDepth = 5
			benchmarkCases = append(benchmarkCases, benchmarkCase)
		}
		for unmatchDepth := 5; unmatchDepth <= 30; unmatchDepth += 5 {
			name := fmt.Sprintf("Pattern=3,DetectIdNumberCbrt=%d,UnmatchDepth=%d", detectIdNumberCbrt, unmatchDepth)
			benchmarkCase := newOverlapDetectorBenchmarkCase(name)
			benchmarkCase.detectIdNumberCbrt = detectIdNumberCbrt
			benchmarkCase.detectedZoomLevel = 35
			benchmarkCase.detectZoomLevel = 35
			benchmarkCase.unmatchDepth = unmatchDepth
			benchmarkCases = append(benchmarkCases, benchmarkCase)
		}
	}

	for _, benchmarkCase := range benchmarkCases {
		detectSpatialIds := benchmarkCase.createDetectSpatialIds()
		detectedSpatialIds := benchmarkCase.createDetectedSpatialIds()

		// ベンチマーク内ではエラーが無視されるため、エラーが出ていないかの確認
		greedyDetector, err := NewSpatialIdGreedyOverlapDetector(detectedSpatialIds)
		if err != nil {
			b.Error(err)
			return
		}
		if _, err := greedyDetector.IsOverlap(detectSpatialIds); err != nil {
			b.Error(err)
			return
		}
		treeDetector, err := NewSpatialIdTreeOverlapDetector(detectedSpatialIds)
		if err != nil {
			b.Error(err)
			return
		}
		if _, err := treeDetector.IsOverlap(detectSpatialIds); err != nil {
			b.Error(err)
			return
		}

		logic_slice := []string{"greedy", "tree"}
		step_slice := []string{"construct", "search"}
		for _, logic := range logic_slice {
			for _, step := range step_slice {
				name := fmt.Sprintf("Logic=%s,Step=%s,%s", logic, step, benchmarkCase.name)

				var f func(b *testing.B)
				switch logic {
				case "greedy":
					switch step {
					case "construct":
						f = func(b *testing.B) {
							for range b.N {
								NewSpatialIdGreedyOverlapDetector(detectedSpatialIds)
							}
						}
					case "search":
						f = func(b *testing.B) {
							for range b.N {
								greedyDetector.IsOverlap(detectSpatialIds)
							}
						}
					}
				case "tree":
					switch step {
					case "construct":
						f = func(b *testing.B) {
							for range b.N {
								NewSpatialIdTreeOverlapDetector(detectedSpatialIds)
							}
						}
					case "search":
						f = func(b *testing.B) {
							for range b.N {
								treeDetector.IsOverlap(detectSpatialIds)
							}
						}
					}
				}

				b.Run(name, f)
			}
		}
	}
}

// BenchmarkOverlapDetectorMemoryUsage OverlapDetectorのメモリ使用量に対するベンチマークを行う
func BenchmarkOverlapDetectorMemoryUsage(b *testing.B) {
	benchmarkCases := []OverlapDetectorBenchmarkCase{}
	for _, detectIdNumberCbrt := range []int{1, 10} {
		for detectedIdNumberCbrt := 6; detectedIdNumberCbrt <= 10; detectedIdNumberCbrt += 1 {
			name := fmt.Sprintf("Pattern=1,DetectIdNumberCbrt=%d,DetectedIdNumberCbrt=%d", detectIdNumberCbrt, detectedIdNumberCbrt)
			benchmarkCase := newOverlapDetectorBenchmarkCase(name)
			benchmarkCase.detectIdNumberCbrt = detectIdNumberCbrt
			benchmarkCase.detectedIdNumberCbrt = detectedIdNumberCbrt
			benchmarkCases = append(benchmarkCases, benchmarkCase)
		}
		for detectedZoomLevel := 5; detectedZoomLevel <= 35; detectedZoomLevel += 5 {
			name := fmt.Sprintf("Pattern=2,DetectIdNumberCbrt=%d,DetectedZoomLevel=%d", detectIdNumberCbrt, detectedZoomLevel)
			benchmarkCase := newOverlapDetectorBenchmarkCase(name)
			benchmarkCase.detectIdNumberCbrt = detectIdNumberCbrt
			benchmarkCase.detectedZoomLevel = detectedZoomLevel
			benchmarkCase.unmatchDepth = 5
			benchmarkCases = append(benchmarkCases, benchmarkCase)
		}
	}

	var detector SpatialIdOverlapDetector // コンパイラに最適化されないように
	for _, benchmarkCase := range benchmarkCases {
		detectedSpatialIds := benchmarkCase.createDetectedSpatialIds()

		// ベンチマーク内ではエラーが無視されるため、エラーが出ていないかの確認
		var err error
		_, err = NewSpatialIdGreedyOverlapDetector(detectedSpatialIds)
		if err != nil {
			b.Error(err)
			return
		}
		_, err = NewSpatialIdTreeOverlapDetector(detectedSpatialIds)
		if err != nil {
			b.Error(err)
			return
		}

		logic_slice := []string{"greedy", "tree"}
		step := "construct"
		for _, logic := range logic_slice {
			name := fmt.Sprintf("Logic=%s,Step=%s,%s", logic, step, benchmarkCase.name)

			memAllocSum := uint64(0)
			var f func(b *testing.B)
			switch logic {
			case "greedy":
				for range 100 {
					startAllocatedMemory := getAllocatedMemory()
					detector, _ = NewSpatialIdGreedyOverlapDetector(detectedSpatialIds)
					goalAllocatedMemory := getAllocatedMemory()
					if goalAllocatedMemory < startAllocatedMemory {
						b.Error("goalAllocatedMemory < startAllocatedMemory")
						return
					}
					memAllocSum += goalAllocatedMemory - startAllocatedMemory
				}

				f = func(b *testing.B) {
					for range b.N {
						NewSpatialIdGreedyOverlapDetector(detectedSpatialIds)
					}
				}
			case "tree":
				for range 100 {
					startAllocatedMemory := getAllocatedMemory()
					detector, _ = NewSpatialIdTreeOverlapDetector(detectedSpatialIds)
					goalAllocatedMemory := getAllocatedMemory()
					memAllocSum += goalAllocatedMemory - startAllocatedMemory
					if goalAllocatedMemory < startAllocatedMemory {
						b.Error("goalAllocatedMemory < startAllocatedMemory")
						return
					}
				}

				f = func(b *testing.B) {
					for range b.N {
						NewSpatialIdTreeOverlapDetector(detectedSpatialIds)
					}
				}
			}

			fmt.Printf("BenchmarkOverlapDetectorMemoryUsage/%s,AllocatedMemoryMean=%d\n", name, memAllocSum/100)

			b.Run(name, f)
		}
	}
	fmt.Println(detector)
}

//	TestCheckExtendedSpatialIdsOverlap01 拡張空間IDの重複確認関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の拡張空間ID：{"16/58209/25805/16/1"},{"15/29104/12902/15/0"}
//
// + 確認内容
//   - 入力値から包含関係があることを確認できること
func TestCheckExtendedSpatialIdsOverlap01(t *testing.T) {
	//入力値
	spatialId1 := "16/58209/25805/16/1"
	spatialId2 := "15/29104/12902/15/0"
	resultValue, resultErr := CheckExtendedSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := true

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

//	TestCheckExtendedSpatialIdsOverlap02 拡張空間IDの重複確認関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の拡張空間ID：{"16/58210/25805/16/2"},{"15/29104/12902/15/0"}
//
// + 確認内容
//   - 入力値から包含関係がないことを確認できること
func TestCheckExtendedSpatialIdsOverlap02(t *testing.T) {
	//入力値
	spatialId1 := "16/58210/25806/16/2"
	spatialId2 := "15/29104/12902/15/0"
	resultValue, resultErr := CheckExtendedSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := false

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

//	TestCheckExtendedSpatialIdsOverlap03 拡張空間IDの重複確認関数 区切り文字数がフォーマットに従っていない場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の拡張空間ID：{"16/58210/25805/16"},{"15/29104/12902/15"}
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestCheckExtendedSpatialIdsOverlap03(t *testing.T) {
	//入力値
	spatialId1 := "16/58209/25805/16"
	spatialId2 := "15/29104/12902/15"
	resultValue, resultErr := CheckExtendedSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := false
	expectError := "invalid format. extendedSpatialId1: 16/58209/25805/16, extendedSpatialId2: 15/29104/12902/15"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr.Error() != expectError {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectError, resultErr.Error())
	}

	t.Log("テスト終了")
}

//	TestCheckExtendedSpatialIdsOverlap04 拡張空間IDの重複確認関数 int64変換時にエラーが発生した場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の拡張空間ID：{"16/58209A/25805/16/1"},{"15/29104/12902/15/0"}
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestCheckExtendedSpatialIdsOverlap04(t *testing.T) {
	//入力値
	spatialId1 := "16/58209A/25805/16/1"
	spatialId2 := "15/29104/12902/15/0"
	resultValue, resultErr := CheckExtendedSpatialIdsOverlap(spatialId1, spatialId2)

	// 期待値
	expectValue := false
	expectError := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr.Error() != expectError {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectError, resultErr.Error())
	}

	t.Log("テスト終了")
}

// BenchmarkCheckExtendedSpatialIdsOverlap01 空間IDの重複確認関数 自然数fインデックスベンチマーク
//
// + 試験データ
//   - パターン1：
//     比較対象の空間ID：{"16/58206/25805/16/1"}, {"15/29104/12902/15/0"}(包含関係あり)
//
// + 確認内容
//   - 入力値に包含関係があった場合の処理速度
func BenchmarkCheckExtendedSpatialIdsOverlap01(b *testing.B) {
	// 入力値
	spatialId1 := "16/58209/25805/16/1"
	spatialId2 := "15/29104/12902/15/0"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CheckExtendedSpatialIdsOverlap(spatialId1, spatialId2)
	}
	b.StopTimer()
	b.Log("テスト終了")
}

//	TestCheckExtendedSpatialIdsArrayOverlap01 拡張空間IDの重複確認関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の拡張空間ID：{"16/58209/25805/16/1"},{"15/29104/12902/15/0"}
//
// + 確認内容
//   - 入力値から包含関係があることを確認できること
func TestCheckExtendedSpatialIdsArrayOverlap01(t *testing.T) {
	//入力値
	spatialIds1 := []string{"16/58206/25805/16/1", "16/58209/25805/16/1"}
	spatialIds2 := []string{"15/29104/12902/15/0"}
	resultValue, resultErr := CheckExtendedSpatialIdsArrayOverlap(spatialIds1, spatialIds2)

	// 期待値
	expectValue := true

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

//	TestCheckExtendedSpatialIdsArrayOverlap02 拡張空間IDの重複確認関数 正常系動作確認
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の拡張空間ID：{"16/58210/25805/16/2"},{"15/29104/12902/15/0"}
//
// + 確認内容
//   - 入力値から包含関係がないことを確認できること
func TestCheckExtendedSpatialIdsArrayOverlap02(t *testing.T) {
	//入力値
	spatialIds1 := []string{"16/58210/25806/16/2", "16/58211/25806/16/2"}
	spatialIds2 := []string{"15/29104/12902/15/0"}
	resultValue, resultErr := CheckExtendedSpatialIdsArrayOverlap(spatialIds1, spatialIds2)

	// 期待値
	expectValue := false

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr != nil {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：nil, 取得値：%s", resultErr)
	}

	t.Log("テスト終了")
}

//	TestCheckExtendedSpatialIdsArrayOverlap03 拡張空間IDの重複確認関数 区切り文字数がフォーマットに従っていない場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の拡張空間ID：{"16/58210/25805/16"},{"15/29104/12902/15"}
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestCheckExtendedSpatialIdsArrayOverlap03(t *testing.T) {
	//入力値
	spatialIds1 := []string{"16/58209/25805/16"}
	spatialIds2 := []string{"15/29104/12902/15"}
	resultValue, resultErr := CheckExtendedSpatialIdsArrayOverlap(spatialIds1, spatialIds2)

	// 期待値
	expectValue := false
	expectError := "invalid format. extendedSpatialId1: 16/58209/25805/16, extendedSpatialId2: 15/29104/12902/15"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr.Error() != expectError {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectError, resultErr.Error())
	}

	t.Log("テスト終了")
}

//	TestCheckExtendedSpatialIdsArrayOverlap04 拡張空間IDの重複確認関数 int64変換時にエラーが発生した場合
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の拡張空間ID：{"16/58209A/25805/16/1"},{"15/29104/12902/15/0"}
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestCheckExtendedSpatialIdsArrayOverlap04(t *testing.T) {
	//入力値
	spatialIds1 := []string{"16/58209A/25805/16/1"}
	spatialIds2 := []string{"15/29104/12902/15/0"}
	resultValue, resultErr := CheckExtendedSpatialIdsArrayOverlap(spatialIds1, spatialIds2)

	// 期待値
	expectValue := false
	expectError := "InputValueError,入力チェックエラー"

	// 空間IDと期待値の比較
	if !reflect.DeepEqual(resultValue, expectValue) {
		t.Errorf("空間ID - 期待値：%v, 取得値：%v", expectValue, resultValue)
	}
	if resultErr.Error() != expectError {
		// 戻り値のエラーインスタンスが期待値と異なる場合Errorをログに出力
		t.Errorf("error - 期待値：%s, 取得値：%s\n", expectError, resultErr.Error())
	}

	t.Log("テスト終了")
}

// BenchmarkCheckExtendedSpatialIdsArrayOverlap01 空間ID列の重複確認関数 空間IDの重複確認関数 配列ベンチマーク
//
// + 試験データ
//   - パターン1(包含関係あり)：
//     比較対象の空間ID列：[10]{"16/58206/25805/16/1"}, [10]{"15/29104/12902/15/0"}
//
// + 確認内容
//   - 入力値に包含関係があった場合の処理速度
func BenchmarkCheckExtendedSpatialIdsArrayOverlap01(b *testing.B) {
	// 入力値
	spatialIds1 := []string{"16/58206/25805/16/1"}
	spatialIds2 := []string{"15/29104/12902/15/0"}
	for i := 0; i < 9; i++ {
		spatialIds1 = append(spatialIds1, "16/58206/25805/16/1")
		spatialIds2 = append(spatialIds2, "15/29104/12902/15/0")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CheckExtendedSpatialIdsArrayOverlap(spatialIds1, spatialIds2)
	}
	b.StopTimer()
	b.Log("テスト終了")
}

// getAllocatedMemory GC後の現在のメモリ使用量を取得する
func getAllocatedMemory() uint64 {
	runtime.GC()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return mem.Alloc
}

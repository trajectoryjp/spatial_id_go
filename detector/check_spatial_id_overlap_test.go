package detector

import (
	"github.com/trajectoryjp/multidimensional-radix-tree/src/tree"
	"reflect"
	"testing"
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
//   - パターン1(ズームレベル超過エラー)：
//     比較対象の空間ID：{"26/-1/7274/3225"},{"16/8/58198/25804"}
//   - パターン2(ズームレベル超過エラー)：
//     比較対象の空間ID：{"13/1/7274/3225"},{"26/-8/58198/25804"}
//   - パターン3(高度範囲外エラー)：
//     比較対象の空間ID：{"25/16777216/0/3225"},{"18/1/58198/25804"}
//   - パターン4(高度範囲外エラー)：
//     比較対象の空間ID：{"18/1/58198/25804"},{"25/16777216/0/3225"}
//   - パターン5(高度範囲外エラー)：
//     比較対象の空間ID：{"25/-16777217/0/3225"},{"16/8/58198/25804"}
//   - パターン6(高度範囲外エラー)：
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
			//入力値
			spatialId1: "26/-1/7274/3225",
			spatialId2: "16/8/58198/25804",
			// 期待値
			expectValue: false,
			expectError: "InputValueError,入力チェックエラー,zoom level 26 must be less than 25 @spatialId1[0] = 26/-1/7274/3225",
		},
		{
			//入力値
			spatialId1: "13/1/7274/3225",
			spatialId2: "26/-8/58198/25804",
			// 期待値
			expectValue: false,
			expectError: "InputValueError,入力チェックエラー,zoom level 26 must be less than 25 @spatialId2[0] = 26/-8/58198/25804",
		},
		{
			//入力値
			spatialId1: "25/16777216/0/3225",
			spatialId2: "18/1/58198/25804",
			// 期待値
			expectValue: false,
			expectError: "ValueConvertError,値の変換エラー,output index (33554432) does not exist with given zoomLevel (25) @spatialId1[0] = 25/16777216/0/3225",
		},
		{
			//入力値
			spatialId1: "18/1/58198/25804",
			spatialId2: "25/16777216/0/3225",
			// 期待値
			expectValue: false,
			expectError: "ValueConvertError,値の変換エラー,output index (33554432) does not exist with given zoomLevel (25) @spatialId2[0] = 25/16777216/0/3225",
		},
		{
			//入力値
			spatialId1: "25/-16777217/0/3225",
			spatialId2: "18/1/58198/25804",
			// 期待値
			expectValue: false,
			expectError: "InputValueError,入力チェックエラー,input f-index -16777217 is out of altitude range @spatialId1[0] = 25/-16777217/0/3225",
		},
		{
			//入力値
			spatialId1: "18/1/58198/25804",
			spatialId2: "25/-16777217/0/3225",
			// 期待値
			expectValue: false,
			expectError: "InputValueError,入力チェックエラー,input f-index -16777217 is out of altitude range @spatialId2[0] = 25/-16777217/0/3225",
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

// TestPanicCheckSpatialIdsArrayOverlapWithCache 空間ID列の重複確認関数 バイナリツリー次元不正
//
// 試験詳細：
// + 試験データ
//   - パターン1：
//     比較対象の空間ID列：2DTable(){"13/7274/3225"}, {"16/0/58198/25804"}
//
// + 確認内容
//   - 入力値からpanic発生を取得できること
func TestPanicCheckSpatialIdsArrayOverlapWithCache(t *testing.T) {
	defer func() {
		// panicが発生しなかった場合Errorをログに出力
		if err := recover(); err == nil {
			t.Errorf("error - 期待値：(panic()), 取得値：nil\n")
		} else {
			t.Logf("success paniced: %v", err)
		}
	}()
	//入力値
	spatialIds1 := tree.CreateTree(tree.Create2DTable())
	// "13/7274/3225" (Appendでpanicしないよう、2Dデータを作成)
	spatialIds1.Append(tree.Indexs{0, 7274, 3225}, tree.ZoomSetLevel(13), "2d value")
	spatialIds2 := []string{"16/0/58198/25803"}

	// panic生成
	_, _ = CheckSpatialIdsArrayOverlapWithCache(spatialIds1, spatialIds2)

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

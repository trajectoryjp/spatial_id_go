package detector

import (
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
//
// + 確認内容
//   - 入力値から入力チェックエラーを取得できること
func TestCheckSpatialIdsOverlap04(t *testing.T) {
	//入力値
	spatialId1 := "25/0/29803148/13212522/777"
	spatialId2 := "16/0/58198/25804"
	resultValue, resultErr := CheckSpatialIdsOverlap(spatialId1, spatialId2)

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

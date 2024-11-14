package spatialID

import (
	"reflect"
	"testing"
)

// TestError01 エラー詳細有りの場合確認
//
// 試験詳細
//  + 試験データ
//      code: "code", msg: "message", detail: "detail"
//  + 確認内容
//    - エラー詳細があった場合、エラーコード,エラーメッセージ,エラー詳細を返却することを確認
func TestError01(t *testing.T) {

	// エラー定義
	e := spatialIdError{code: "code", msg: "message", detail: "detail"}

	// 期待値
	expectVal := "code,message,detail"

	// テスト対象呼び出し
	resultVal := e.Error()

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("エラー定義 - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestError02 エラー詳細無しの場合確認
//
// 試験詳細
//  + 試験データ
//      code: "code", msg: "message", detail: ""
//  + 確認内容
//    - エラー詳細があった場合、エラーコード,エラーメッセージを返却することを確認
func TestError02(t *testing.T) {

	// エラー定義
	e := spatialIdError{code: "code", msg: "message", detail: ""}

	// 期待値
	expectVal := "code,message"

	// テスト対象呼び出し
	resultVal := e.Error()

	// 戻り値と期待値の比較
	if !reflect.DeepEqual(resultVal, expectVal) {
		t.Errorf("エラー定義 - 期待値：%s, 取得値：%s", expectVal, resultVal)
	}

	t.Log("テスト終了")
}

// TestNewSpatialIdError01 エラーコードがInputValueErrorCodeの場合確認
//
// 試験詳細
//  + 試験データ
//      code: InputValueErrorCode, msg: "入力チェックエラー", detail: "detail"
//  + 確認内容
//    - エラーコードがInputValueErrorCodeの場合、メッセージ文言が"入力チェックエラー"となることを確認
func TestNewSpatialIdError01(t *testing.T) {

	// 期待値
	code := InputValueErrorCode
	msg := "入力チェックエラー"
	detail := "detail"
	e := spatialIdError{code: code, msg: msg, detail: detail}

	// テスト対象呼び出し
	resultVal := NewSpatialIdError(code, detail)

	// 戻り値と期待値の比較
	if e != resultVal {
		t.Errorf("error - 期待値：%s, 取得値：%s", e, resultVal)
	}

	t.Log("テスト終了")
}

// TestNewSpatialIdError02 エラーコードがOptionFailedErrorCodeの場合確認
//
// 試験詳細
//  + 試験データ
//      code: OptionFailedErrorCode, msg: "オプション値の指定エラー", detail: "detail"
//  + 確認内容
//    - エラーコードがOptionFailedErrorCodeの場合、メッセージ文言が"オプション値の指定エラー"となることを確認
func TestNewSpatialIdError02(t *testing.T) {

	// 期待値
	code := OptionFailedErrorCode
	msg := "オプション値の指定エラー"
	detail := "detail"
	e := spatialIdError{code: code, msg: msg, detail: detail}

	// テスト対象呼び出し
	resultVal := NewSpatialIdError(code, detail)

	// 戻り値と期待値の比較
	if e != resultVal {
		t.Errorf("error - 期待値：%s, 取得値：%s", e, resultVal)
	}

	t.Log("テスト終了")
}

// TestNewSpatialIdError03 エラーコードがValueConvertErrorCodeの場合確認
//
// 試験詳細
//  + 試験データ
//      code: ValueConvertErrorCode, msg: "値の変換エラー", detail: "detail"
//  + 確認内容
//    - エラーコードがValueConvertErrorCodeの場合、メッセージ文言が"値の変換エラー"となることを確認
func TestNewSpatialIdError03(t *testing.T) {

	// 期待値
	code := ValueConvertErrorCode
	msg := "値の変換エラー"
	detail := "detail"
	e := spatialIdError{code: code, msg: msg, detail: detail}

	// テスト対象呼び出し
	resultVal := NewSpatialIdError(code, detail)

	// 戻り値と期待値の比較
	if e != resultVal {
		t.Errorf("error - 期待値：%s, 取得値：%s", e, resultVal)
	}

	t.Log("テスト終了")
}

// TestNewSpatialIdError04 エラーコードがOtherErrorCodeの場合確認
//
// 試験詳細
//  + 試験データ
//      code: OtherErrorCode, msg: "その他例外が発生", detail: "detail"
//  + 確認内容
//    - エラーコードがOtherErrorCodeの場合、メッセージ文言が"その他例外が発生"となることを確認
func TestNewSpatialIdError04(t *testing.T) {

	// 期待値
	code := OtherErrorCode
	msg := "その他例外が発生"
	detail := "detail"
	e := spatialIdError{code: code, msg: msg, detail: detail}

	// テスト対象呼び出し
	resultVal := NewSpatialIdError(code, detail)

	// 戻り値と期待値の比較
	if e != resultVal {
		t.Errorf("error - 期待値：%s, 取得値：%s", e, resultVal)
	}

	t.Log("テスト終了")
}

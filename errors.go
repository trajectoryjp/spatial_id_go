package spatialID

import (
	"fmt"
)

// spatialIdErrorCode 空間IDエラーコード定義用の型
type spatialIdErrorCode string

// spatialIdError エラー定義用の構造体
type spatialIdError struct {
	code   spatialIdErrorCode // エラーコード
	msg    string             // エラーメッセージ
	detail string             // エラー詳細
}

// 空間IDエラーコード
const (
	InputValueErrorCode   spatialIdErrorCode = "InputValueError"   // 入力チェックエラー
	OptionFailedErrorCode spatialIdErrorCode = "OptionFailedError" // オプション値の指定エラー
	ValueConvertErrorCode spatialIdErrorCode = "ValueConvertError" // 座標系変換時の変換エラー
	OtherErrorCode        spatialIdErrorCode = "OtherError"        // その他例外が発生
)

// Error errorインタフェース実装
//
// Error()メソッドを定義することでerrorインタフェースを実装。
// エラーインスタンスのフィールド情報を文字列型で返却する。
//
// 返却する文字列のフォーマットは以下の通り。
//
//	エラー詳細無しの場合：
//	 エラーコード,エラーメッセージ
//	エラー詳細有りの場合：
//	 エラーコード,エラーメッセージ,エラー詳細
//
// 戻り値：
//
//	エラーインスタンスのフィールドに設定されたパラメータを含む文字列
func (e spatialIdError) Error() string {
	if len(e.detail) > 0 {
		return fmt.Sprintf("%s,%s,%s", e.code, e.msg, e.detail)
	} else {
		return fmt.Sprintf("%s,%s", e.code, e.msg)
	}
}

// NewSpatialIdError 空間ID用エラーインスタンス返却関数
//
// 引数で指定された空間IDエラーコードに対応したエラーインスタンスを返却する。
// 引数のエラー詳細については、使用しない場合空文字の入力が必須。
//
// 引数：
//
//	err   ：空間IDエラーコード
//	detail：エラーインスタンスに設定するエラー詳細
func NewSpatialIdError(err spatialIdErrorCode, detail string) error {
	switch err {
	case InputValueErrorCode:
		return spatialIdError{code: err, msg: "入力チェックエラー", detail: detail}
	case OptionFailedErrorCode:
		return spatialIdError{code: err, msg: "オプション値の指定エラー", detail: detail}
	case ValueConvertErrorCode:
		return spatialIdError{code: err, msg: "値の変換エラー", detail: detail}
	default:
		return spatialIdError{code: err, msg: "その他例外が発生", detail: detail}
	}
}

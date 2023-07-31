// logger ロガーパッケージ
package logger

import (
	"io/ioutil"
	"strings"
	"testing"
)

//TestInit01 logディレクトリ作成試験
//
//試験詳細：
// + 試験データ
//  - パターン1
//     logインスタンスが存在する
// + 確認内容
//  - logインスタンスが存在していることを確認
func TestInit01(t *testing.T) {

	if log == nil {
		t.Errorf("ログインスタンスなし")
	}
	t.Log("テスト終了")

}

//TestDebug01 Debugログ出力試験
//
//試験詳細
// + 試験データ
//  - パターン1
//     ログレベルとデバッグログレベルが一致していること
// + 確認内容
//  - デバッグログに入力内容が出力されていること
func TestDebug01(t *testing.T) {

	message := "test_loggerMessage"
	values := "test_loggerVal"

	fileName := "./log/spatial.log"

	Debug(message, values)

	text, _ := ioutil.ReadFile(fileName)

	if string(text) == "" {
		t.Errorf("ファイルに内容が書き込まれていない、想定外のルート")
	} else {
		if !strings.Contains(string(text), message) || !strings.Contains(string(text), values) {
			t.Errorf("正しい内容が書き込まれていない、想定外の挙動")
		}
	}
	t.Log("テスト終了")
}

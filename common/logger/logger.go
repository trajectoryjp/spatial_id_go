// Package logger ロガーパッケージ
package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

// ログレベル
const (
	infoLevel  = zap.InfoLevel  // InfoLevel設定値
	debugLevel = zap.DebugLevel // DebugLevel設定値
	logLevel   = debugLevel     // ロガーに設定するログレベル
)

// ログインスタンス
var log *zap.Logger

// ログディレクトリ存在フラグ
var logDirExist = true

// 初期化関数
func init() {
	// カレントディレクトリを取得
	curDirPath, _ := os.Getwd()

	// ログディレクトリ名を定義
	logDirName := "log"

	// ログファイル名を定義
	logFileName := "spatial.log"

	// ログディレクトリのパスを定義
	logDirPath := filepath.Join(curDirPath, logDirName)

	// ログファイルのパスを定義
	logFilePath := filepath.Join(logDirPath, logFileName)

	// ロガーをDebugレベルで初期化
	config := zap.NewDevelopmentConfig()

	// logLevel の設定値に応じて、出力するログレベルを設定
	if logLevel == infoLevel {
		// ロガーをInfoレベルで初期化
		config = zap.NewProductionConfig()
	}

	// ログディレクトリが存在しない場合作成
	if f, err := os.Stat(logDirPath); os.IsNotExist(err) {
		_ = os.Mkdir(logDirPath, 0755)
	} else if !f.IsDir() {
		fmt.Println("ログディレクトリと同名のファイルが存在するため、ログが出力されません。")
		fmt.Println("ログディレクトリ名：log")
		logDirExist = false
	}

	config.OutputPaths = []string{logFilePath}

	// ログ出力の呼び出し元を出力する設定
	log, _ = config.Build(zap.AddCallerSkip(1))
}

// Debug Debugログ出力関数
//
// debugLevelのログをログファイルに出力する。
// 第1引数に入力された書式に、第2引数以降に入力されたパラメータを埋め込み、
// ログファイルに出力する。
//
// 第2引数以降が未入力の場合、第1引数がフォーマットされずにログファイルに出力される。
//
// 引数
//
//	message：ログメッセージの指定書式
//	values ：指定書式に埋め込むパラメータ
func Debug(message string, values ...any) {
	if logLevel == debugLevel && logDirExist == true {
		log.Debug(fmt.Sprintf(message, values...))
	}
}

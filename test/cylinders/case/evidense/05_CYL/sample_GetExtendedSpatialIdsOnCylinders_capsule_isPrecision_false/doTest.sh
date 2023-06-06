#!/bin/bash -u

################################################
# スクリプト共通変数定義(必須)
################################################
typeset -r script_home=$(cd $(dirname $0);pwd)
typeset -r test_home=${script_home%/*/*}
typeset -r common_template_dir="${test_home}/common/template"

# 共通環境設定読み込み
. ${test_home}/testenv

################################################
# 定数定義(固定)
################################################
typeset -r OUTPUT_DIR=${log_dir}

################################################
# 試験条件定義
################################################

################################################
# 関数定義
################################################

test_finalize() {
    # ログ収集ジョブ終了
    stop_jobs
}
trap test_finalize EXIT

################################################
# 試験準備
################################################

# go.mod
#GO_MOD_FILE="${script_home}/go.mod"
convert_conf ${common_template_dir} ${script_home}

# go.mod初期化
go mod tidy

################################################
# 試験本処理 (Main Process)
################################################
output_txt=${OUTPUT_DIR}/output.txt
expected_txt=${OUTPUT_DIR}/expected.txt


exec_test go run expect.go env.go
exec_test go run main.go env.go


cat ${OUTPUT_DIR}/output.txt | sort > ${OUTPUT_DIR}/output_sorted.txt
cat ${OUTPUT_DIR}/expected.txt | sort > ${OUTPUT_DIR}/expected_sorted.txt

diff ${OUTPUT_DIR}/output_sorted.txt ${OUTPUT_DIR}/expected_sorted.txt > ${OUTPUT_DIR}/result.diff

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
typeset -r OUTPUT_DIR=${script_home}/output
mkdir -p ${OUTPUT_DIR}
rm -rf ${OUTPUT_DIR}/*

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
convert_conf ${common_template_dir} ${script_home}

# go.mod初期化
go mod tidy

################################################
# 試験本処理 (Main Process)
################################################
spatial_txt=${OUTPUT_DIR}/spatial.txt
point_txt=${OUTPUT_DIR}/point.txt
exec_test go run main.go -o ${spatial_txt} -p ${point_txt}

exec_cmd python plot.py -i ${spatial_txt} -p ${point_txt}

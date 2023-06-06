package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	//TODO:env.goに定義した値が読まれていることを確認するためのprintです。
	// 実際に使う際には削除をお願いします。
	fmt.Println("true: %v", ProjectedCrs)
	fmt.Println("true: %v", IsCapsule)
	fmt.Println("円の中心座標: %s", CenterSpatialIDs)

	//TODO:スクリプト記述時に円柱の関数が完成していないため、
	// output.txtを出力するために最適化の関数の出力値を仮で使用しています。
	// 実際に使う際には関数、引数の修正をお願いします。
	SpatialIds := []string{"34/0/0/0", "34/1/1/1"}
	var zoom int64 = 35
	spatialIDs, err := integrate.MergeSpatialIds(SpatialIds, zoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID最適化時にエラー発生: %w", err))
		os.Exit(1)
	}

	fileName := "./logs/output.txt"
	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	for _, id := range spatialIDs {
		fp.WriteString(id)
		fp.WriteString("\n")
	}

	// 結果を出力
	fmt.Printf("最適化後の空間ID: %s\n", spatialIDs)
}

package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の空間ID
	SpatialIds := []string{"21/512/1024/2048", "21/512/1024/2049", "21/512/1024/2050", "21/512/1024/2048"}
	// 最適化精度
	var zoom int64 = 11

	// 空間ID最適化
	spatialIDs, err := integrate.MergeSpatialIds(SpatialIds, zoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID最適化時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("最適化後の空間ID: %s\n", spatialIDs)
}

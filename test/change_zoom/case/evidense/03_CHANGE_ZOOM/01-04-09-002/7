package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 精度変換対象の空間ID
	SpatialIds := []string{"34/1024/1024/1024"}
	// 変換後の精度
	var zoom int64 = 35

	// 空間ID精度変換
	spatialIDs, err := integrate.ChangeSpatialIdsZoom(SpatialIds, zoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID精度変換時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("精度変換後の空間ID: %s\n", spatialIDs)
}

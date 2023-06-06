package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 精度変換対象の拡張空間ID
	SpatialIds := []string{"12/83/80/12/80"}
	// 変換後の水平精度
	var hzoom int64 = 10
	// 変換後の垂直精度
	var vzoom int64 = 10

	// 拡張空間ID精度変換
	spatialIDs, err := integrate.ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)
	if err != nil {
		fmt.Println(fmt.Errorf("拡張空間ID精度変換時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("精度変換後の拡張空間ID: %s\n", spatialIDs)
}

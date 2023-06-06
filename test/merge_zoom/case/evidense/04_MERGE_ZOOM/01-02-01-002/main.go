package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の拡張空間ID
	SpatialIds := []string{"22/40/40/22/10", "22/41/40/22/10", "22/42/42/22/10", "22/40/43/22/10", "22/41/43/22/10", "22/42/40/22/10", "22/43/40/22/10", "22/42/41/22/10", "22/43/41/22/10", "22/40/42/22/10", "22/41/42/22/10", "22/43/42/22/10", "22/42/43/22/10", "22/43/43/22/10", "22/40/41/22/10", "22/41/41/22/10"}
	// 最適化水平精度
	var hzoom int64 = 20
	// 最適化垂直精度
	var vzoom int64 = 22

	// 拡張空間ID最適化
	spatialIDs, err := integrate.MergeExtendedSpatialIds(SpatialIds, hzoom, vzoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID最適化時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("最適化後の空間ID: %s\n", spatialIDs)
}

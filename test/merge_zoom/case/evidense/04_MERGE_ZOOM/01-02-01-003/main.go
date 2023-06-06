package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の拡張空間ID
	SpatialIds := []string{"22/10/10/22/40", "22/10/10/22/41", "22/10/10/22/42", "22/10/10/22/43"}
	// 最適化水平精度
	var hzoom int64 = 22
	// 最適化垂直精度
	var vzoom int64 = 20

	// 拡張空間ID最適化
	spatialIDs, err := integrate.MergeExtendedSpatialIds(SpatialIds, hzoom, vzoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID最適化時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("最適化後の空間ID: %s\n", spatialIDs)
}

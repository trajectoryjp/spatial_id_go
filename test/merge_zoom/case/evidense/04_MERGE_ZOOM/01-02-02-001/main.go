package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の拡張空間ID
	SpatialIds := []string{"21/20/20/22/40","21/20/20/22/41","21/20/20/22/42","21/20/20/22/43","21/20/21/22/40","21/20/21/22/41","21/20/21/22/42","21/20/21/22/43","21/21/20/22/40","21/21/20/22/41","21/21/20/22/42","21/21/20/22/43","21/21/21/22/40","21/21/21/22/41","21/21/21/22/42","21/21/21/22/43"}
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

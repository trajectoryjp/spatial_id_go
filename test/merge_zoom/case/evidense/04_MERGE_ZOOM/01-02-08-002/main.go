package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の拡張空間ID
	SpatialIds := []string{"21/20/20/21/20","21/20/20/21/21","21/21/20/21/20","21/21/20/21/21","21/20/21/21/20","21/20/21/21/21","21/21/21/21/20","21/21/21/21/21","21/23/23/21/20","21/23/23/21/21","21/22/22/21/20","21/22/22/21/21","21/23/22/21/20","21/23/22/21/21","21/22/23/21/20","21/22/23/21/21"}
	// 最適化水平精度
	var hzoom int64 = 20
	// 最適化垂直精度
	var vzoom int64 = 20

	// 拡張空間ID最適化
	spatialIDs, err := integrate.MergeExtendedSpatialIds(SpatialIds, hzoom, vzoom)
	if err != nil {
		fmt.Println(fmt.Errorf("拡張空間ID最適化時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("最適化後の拡張空間ID: %s\n", spatialIDs)
}

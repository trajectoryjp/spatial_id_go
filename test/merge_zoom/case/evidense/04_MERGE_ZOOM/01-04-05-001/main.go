package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の空間ID
	SpatialIds := []string{"21/20/20/21","21/21/20/21","21/20/21/21","21/21/21/21","21/20/20/20","21/21/20/20","21/20/21/20","21/21/21/20","21/22/20/21","21/23/20/21","21/22/21/21","21/23/21/21","21/22/20/20","21/23/20/20","21/22/21/20"}
	// 最適化精度
	var zoom int64 = 20

	// 空間ID最適化
	spatialIDs, err := integrate.MergeSpatialIds(SpatialIds, zoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID最適化時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("最適化後の空間ID: %s\n", spatialIDs)
}

package main

import (
	"os"
	"fmt"
	"spatial-id/common/enum"
	"spatial-id/shape"
)

func main() {
	// 空間ID
	spatialID := "10/5/5/10/-5"

	// ボクセル中心
	centers, err := shape.GetPointOnExtendedSpatialId(spatialID, enum.Center)
	if err != nil {
		fmt.Println(fmt.Errorf("ボクセル中心算出時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	for index, center := range centers {
		fmt.Printf("ボクセルの中心%d: %.12f\n", index+1, *center)
	}
}

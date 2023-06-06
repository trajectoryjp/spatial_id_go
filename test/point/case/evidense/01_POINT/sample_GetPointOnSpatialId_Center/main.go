package main

import (
	"os"
	"fmt"
	"spatial-id/common/enum"
	"spatial-id/shape"
)

func main() {
	// 空間ID
	spatialID := "26/2000/59606297/26425045"

	// ボクセル中心
	centers, err := shape.GetPointOnSpatialId(spatialID, enum.Center)
	if err != nil {
		fmt.Println(fmt.Errorf("ボクセル中心算出時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	for index, center := range centers {
		fmt.Printf("ボクセルの中心%d: %.12f\n", index+1, *center)
	}
}

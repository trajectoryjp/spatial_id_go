package main

import (
	"os"
	"fmt"
	"spatial-id/common/enum"
	"spatial-id/shape"
)

func main() {
	// 空間ID
	spatialID := "10/2048/2048/26/852"

	// ボクセル頂点
	centers, err := shape.GetPointOnExtendedSpatialId(spatialID, enum.Vertex)
	if err != nil {
		fmt.Println(fmt.Errorf("ボクセル頂点算出時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	for index, center := range centers {
		fmt.Printf("ボクセルの頂点%d: %.12f\n", index+1, *center)
	}
}

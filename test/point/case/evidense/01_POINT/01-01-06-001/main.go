package main

import (
	"fmt"
	"os"
	"spatial-id/common/enum"
	"spatial-id/shape"
)

func main() {
	// 空間ID
	spatialID := "18/232837/103222/18/32"

	// 水平精度
	var hZoom int64 = 18
	// 垂直精度
	var vZoom int64 = 18

	// 空間IDを出力
	fmt.Printf("投入した空間ID: %v\n", spatialID)

	// 精度を出力
	fmt.Printf("投入した水平精度:%v 垂直精度:%v\n", hZoom, vZoom)

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

	spatialIDs, err := shape.GetExtendedSpatialIdsOnPoints(centers, hZoom, vZoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("変換後の空間ID: %s\n", spatialIDs)
}

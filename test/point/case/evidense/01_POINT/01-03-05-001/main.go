package main

import (
	"fmt"
	"os"
	"spatial-id/common/enum"
	"spatial-id/shape"
)

func main() {
	// 空間ID
	spatialID := "18/32/232837/103222"

	// 精度
	var zoom int64 = 18

	// 空間IDを出力
	fmt.Printf("投入した空間ID: %v\n", spatialID)

	// 精度を出力
	fmt.Printf("投入した精度:%v\n", zoom)

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

	spatialIDs, err := shape.GetSpatialIdsOnPoints(centers, zoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("変換後の空間ID: %s\n", spatialIDs)
}

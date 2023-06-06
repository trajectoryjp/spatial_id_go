package main

import (
	"os"
	"fmt"
	"spatial-id/common/enum"
	"spatial-id/shape"
)

func main() {
	// 空間ID
	spatialID := "26/59606297/26425045/26/2000"
	fmt.Printf("変換前の空間ID: %s\n", spatialID)

	// ボクセル中心
	center, err1 := shape.GetPointOnExtendedSpatialId(spatialID, enum.Center)
	if err1 != nil {
		fmt.Println(fmt.Errorf("ボクセル中心算出時にエラー発生: %w", err1))
		os.Exit(1)
	}

	// 水平精度
	var hZoom int64 = 26
	// 垂直精度
	var vZoom int64 = 26

	// 空間ID取得
	newSpatialIDs, err2 := shape.GetExtendedSpatialIdsOnPoints(center, hZoom, vZoom)
	if err2 != nil {
		fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err2))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("変換後の空間ID: %s\n", newSpatialIDs)
}

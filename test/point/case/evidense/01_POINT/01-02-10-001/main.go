package main

import (
	"fmt"
	"os"
	"spatial-id/common/enum"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 経度
	lon := 65.5
	// 緯度
	lat := 37.2
	// 高さ
	alt := 620.0
	// 水平精度
	var hZoom int64 = 11
	// 垂直精度
	var vZoom int64 = 17

	// 地理座標
	point, _ := object.NewPoint(lon, lat, alt)
	points := []*object.Point{point}

	// 座標を出力
	fmt.Printf("投入した座標: %v\n", point)

	// 座標を出力
	fmt.Printf("投入した水平精度:%v 垂直精度:%v\n", hZoom, vZoom)

	// 空間ID取得
	spatialIDs, err := shape.GetExtendedSpatialIdsOnPoints(points, hZoom, vZoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 空間ID結果を出力
	fmt.Printf("変換後の空間ID: %s\n", spatialIDs)

	// ボクセル頂点
	for index, spatialID := range spatialIDs {
		centers, err2 := shape.GetPointOnExtendedSpatialId(spatialID, enum.Vertex)
		if err != nil {
			fmt.Println(fmt.Errorf("ボクセル頂点%d算出時にエラー発生: %w", index, err2))
			os.Exit(1)
		}

		// 頂点座標結果を出力
		for indexA, center := range centers {
			fmt.Printf("ボクセルの頂点%d: %.12f\n", indexA+1, *center)
		}
	}
}

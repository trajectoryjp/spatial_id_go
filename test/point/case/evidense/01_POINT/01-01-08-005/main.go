package main

import (
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 経度1
	lon1 := 180.0
	// 緯度1
	lat1 := -85.0511287798
	// 高さ1
	alt1 := 1.0
	// 水平精度1
	var hZoom1 int64 = 10
	// 垂直精度1
	var vZoom1 int64 = 35

	// 経度2
	lon2 := -180.0
	// 緯度2
	lat2 := 85.0511287798
	// 高さ2
	alt2 := -1300.0
	// 水平精度2
	var hZoom2 int64 = 24
	// 垂直精度2
	var vZoom2 int64 = 25

	// 地理座標1
	point1, _ := object.NewPoint(lon1, lat1, alt1)
	points1 := []*object.Point{point1}

	// 地理座標2
	point2, _ := object.NewPoint(lon2, lat2, alt2)
	points2 := []*object.Point{point2}

	// 空間ID1取得
	spatialIDs1, err := shape.GetExtendedSpatialIdsOnPoints(points1, hZoom1, vZoom1)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID1取得時にエラー発生: %w", err))
		os.Exit(1)
	}
	// 空間ID2取得
	spatialIDs2, err := shape.GetExtendedSpatialIdsOnPoints(points2, hZoom2, vZoom2)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID2取得時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("変換後の空間ID1: %s\n", spatialIDs1)

	// 結果を出力
	fmt.Printf("変換後の空間ID2: %s\n", spatialIDs2)
}

package main

import (
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 経度
	lon := 139.753098

	// 緯度1
	lat1 := 35.615309
	// 緯度2
	lat2 := 35.615310
	// 緯度3
	lat3 := 35.615311
	// 緯度4
	lat4 := 35.615312
	// 緯度5
	lat5 := 35.615313
	// 緯度6
	lat6 := 35.615314
	// 緯度7
	lat7 := 35.615315
	// 緯度8
	lat8 := 35.615316

	// 高さ
	alt := 0.0
	// 水平精度
	var hZoom int64 = 26
	// 垂直精度
	var vZoom int64 = 26

	// 地理座標
	point1, _ := object.NewPoint(lon, lat1, alt)
	point2, _ := object.NewPoint(lon, lat2, alt)
	point3, _ := object.NewPoint(lon, lat3, alt)
	point4, _ := object.NewPoint(lon, lat4, alt)
	point5, _ := object.NewPoint(lon, lat5, alt)
	point6, _ := object.NewPoint(lon, lat6, alt)
	point7, _ := object.NewPoint(lon, lat7, alt)
	point8, _ := object.NewPoint(lon, lat8, alt)

	points := []*object.Point{point1, point2, point3, point4, point5, point6, point7, point8}

	// 空間ID取得
	spatialIDs, err := shape.GetExtendedSpatialIdsOnPoints(points, hZoom, vZoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 空間ID結果を出力
	for index, spatialID := range spatialIDs {
		fmt.Printf("空間ID%d: %s\n", index+1, spatialID)
	}
}

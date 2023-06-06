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
	// 緯度
	lat := 35.685371
	// 高さ
	alt := 0.0
	// 水平精度
	var hZoom int64 = 18
	// 垂直精度
	var vZoom int64 = 36

	// 地理座標
	point, _ := object.NewPoint(lon, lat, alt)
	points := []*object.Point{point}

	// 空間ID取得
	_, err := shape.GetExtendedSpatialIdsOnPoints(points, hZoom, vZoom)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// エラーを出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

package main

import (
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 経度
	lon := 139.92271122072384
	// 緯度
	lat := 35.5610740346
	// 高さ
	alt := -0.8500000000029104
	// 精度
	var zoom int64 = 40

	// 地理座標
	point, _ := object.NewPoint(lon, lat, alt)
	points := []*object.Point{point}

	// 空間ID取得
	_, err := shape.GetSpatialIdsOnPoints(points, zoom)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// エラーを出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

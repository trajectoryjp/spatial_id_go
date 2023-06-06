package main

import (
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 経度
	lon := -21.942400
	// 緯度
	lat := 64.147209
	// 高さ
	alt := 100.0
	// 水平精度
	var hZoom int64 = 9
	// 垂直精度
	var vZoom int64 = 25

	// 地理座標
	point, _ := object.NewPoint(lon, lat, alt)
	points := []*object.Point{point}

	// 空間ID取得
	spatialIDs, err := shape.GetExtendedSpatialIdsOnPoints(points, hZoom, vZoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("変換後の空間ID: %s\n", spatialIDs)
}

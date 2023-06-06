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
	alt := 1000.0
	// 水平精度
	var hZoom int64 = 0
	// 垂直精度
	var vZoom int64 = 0

	// 地理座標
	point, _ := object.NewPoint(lon, lat, alt)
	points := []*object.Point{point}

	// 空間ID取得
	for hZoom < 36 {
		fmt.Printf("水平精度: %d 垂直精度: %d\n", hZoom, vZoom)

		spatialIDs, err := shape.GetExtendedSpatialIdsOnPoints(points, hZoom, vZoom)
		if err != nil {
			fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err))
			os.Exit(1)
		}

		hZoom++
		vZoom++
		// 結果を出力
		fmt.Printf("変換後の空間ID: %s\n", spatialIDs)
	}

}

package main

import (
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 水平精度
	var hZoom int64 = 25
	// 垂直精度
	var vZoom int64 = 25

	// 地理座標
	// 経度、緯度、高さの座標はnilとする
	points := []*object.Point{nil}

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

package main

import (
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 精度
	var zoom int64 = 25

	// 地理座標
	points := []*object.Point{nil}

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

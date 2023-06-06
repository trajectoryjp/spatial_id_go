package main

import (
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 始点
	start, _ := object.NewPoint(0.0, 0.0, 0.0)
	// 終点
	end, _ := object.NewPoint(45.0, 75.821946114, 4193404.0)
	// 水平精度
	var hZoom int64 = 36
	// 垂直精度
	var vZoom int64 = 10

	// 空間ID取得
	_, err := shape.GetExtendedSpatialIdsOnLine(start, end, hZoom, vZoom)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// エラーを出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

package main

import (
	"fmt"
	"os"
	"spatial-id/common/object"
	"spatial-id/shape"
)

func main() {
	// 始点
	start, _ := object.NewPoint(139.92271122072384, 35.5610740346, -0.8500000000029104)
	// 終点
	end, _ := object.NewPoint(139.92259973802746, 35.5608653809, -0.8500000000029104)
	// 水平精度
	var hZoom int64 = 40
	// 垂直精度
	var vZoom int64 = 40

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

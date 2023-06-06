package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の拡張空間ID
	SpatialIds := []string{"10/10/10/10/10", "/11/11/10/11"}
	// 最適化水平精度
	var hzoom int64 = 10
	// 最適化垂直精度
	var vzoom int64 = 10

	// 拡張空間ID最適化
	_, err := integrate.MergeExtendedSpatialIds(SpatialIds, hzoom, vzoom)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// エラーを出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

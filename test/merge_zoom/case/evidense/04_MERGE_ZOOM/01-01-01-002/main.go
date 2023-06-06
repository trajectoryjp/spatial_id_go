package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の拡張空間ID
	SpatialIds := []string{"10/1024/2048/10/1024", "10/1024/2048/10/1024", "11/1024/2048/11/1024", "9/1024/2048/9/1024", "11/1024/2048/11/1024"}
	// 最適化水平精度
	var hzoom int64 = -1
	// 最適化垂直精度
	var vzoom int64 = 9

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

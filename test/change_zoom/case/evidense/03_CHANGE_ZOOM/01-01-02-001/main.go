package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 精度変換対象の拡張空間ID
	SpatialIds := []string{"10/10/10/10"}
	// 変換後の水平精度
	var hzoom int64 = 11
	// 変換後の垂直精度
	var vzoom int64 = 11

	// 拡張空間ID精度変換
	_, err := integrate.ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// エラーを出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

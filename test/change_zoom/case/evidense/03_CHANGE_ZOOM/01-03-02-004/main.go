package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 精度変換対象の空間ID
	SpatialIds := []string{"/10/10/10"}
	// 変換後の精度
	var zoom int64 = 11

	// 空間ID精度変換
	_, err := integrate.ChangeSpatialIdsZoom(SpatialIds, zoom)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// エラーを出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の空間ID
	SpatialIds := []string{"21/512/1024/2048", "21/512/1024/2049", "21/512/1024/2050", "21/512/1024/2048"}
	// 最適化精度
	var zoom int64 = 40

	// 空間ID最適化
	_, err := integrate.MergeSpatialIds(SpatialIds, zoom)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// エラーを出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

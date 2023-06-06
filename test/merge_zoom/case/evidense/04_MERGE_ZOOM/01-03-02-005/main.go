package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 最適化対象の空間ID
	SpatialIds := []string{"10/10/10/10", "10/11//10"}
	// 最適化精度
	var zoom int64 = 10

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

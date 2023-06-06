package main

import (
	"os"
	"fmt"
	"spatial-id/common/object"
)

func main() {
	// 経度
	lon := 180.1
	// 緯度
	lat := 35.685371
	// 高さ
	alt := 0.0

	// 点群
	_, err := object.NewPoint(lon, lat, alt)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

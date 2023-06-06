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
	lat := 35.5610740346
	// 高さ
	alt := -0.8500000000029104

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

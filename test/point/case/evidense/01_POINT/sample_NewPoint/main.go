package main

import (
	"os"
	"fmt"
	"spatial-id/common/object"
)

func main() {
	// 経度
	lon := 139.92271122072384
	// 緯度
	lat := 35.5610740346
	// 高さ
	alt := -0.8500000000029104

	// 点群
	point, err := object.NewPoint(lon, lat, alt)
	if err != nil {
		fmt.Println(fmt.Errorf("地理座標初期化時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("地理座標の経度: %.12f\n", point.Lon())
	fmt.Printf("地理座標の緯度: %.12f\n", point.Lat())
	fmt.Printf("地理座標の高さ: %.12f\n", point.Alt())
}

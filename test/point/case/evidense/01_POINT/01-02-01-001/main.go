package main

import (
	"fmt"
	"os"
	"spatial-id/shape"
)

func main() {
	// 空間ID
	spatialID := "10/5/5/10/10"

	// ボクセル中心
	_, err := shape.GetPointOnExtendedSpatialId(spatialID, 10)
	if err == nil {
		fmt.Println("エラーが未発生")
		os.Exit(1)
	}

	// エラーを出力
	fmt.Printf("%T:", err)
	fmt.Println(err)
}

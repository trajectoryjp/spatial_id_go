package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 精度変換対象の拡張空間ID
	SpatialIds := []string{"20/10023/10023/20/10023", "20/10024/10024/20/10024", "20/10025/10025/20/10025", "20/10026/10026/20/10026", "20/10027/10027/20/10027", "22/40112/40112/22/40112", "22/40114/40114/22/40114", "22/40116/40116/22/40116", "22/40118/40118/22/40118", "22/40120/40120/22/40120"}
	// 変換後の水平精度
	var hzoom int64 = 21
	// 変換後の垂直精度
	var vzoom int64 = 21

	// 拡張空間ID精度変換
	spatialIDs, err := integrate.ChangeExtendedSpatialIdsZoom(SpatialIds, hzoom, vzoom)
	if err != nil {
		fmt.Println(fmt.Errorf("拡張空間ID精度変換時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("精度変換後の拡張空間ID: %s\n", spatialIDs)
}

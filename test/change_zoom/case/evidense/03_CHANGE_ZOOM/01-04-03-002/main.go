package main

import (
	"fmt"
	"os"
	"spatial-id/integrate"
)

func main() {
	// 精度変換対象の空間ID
	SpatialIds := []string{"20/10023/10023/10023", "20/10024/10024/10024", "20/10025/10025/10025", "20/10026/10026/10026", "20/10027/10027/10027", "22/40112/40112/40112", "22/40114/40114/40114", "22/40116/40116/40116", "22/40118/40118/40118", "22/40120/40120/40120"}
	// 変換後の精度
	var zoom int64 = 21

	// 空間ID精度変換
	spatialIDs, err := integrate.ChangeSpatialIdsZoom(SpatialIds, zoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID精度変換時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("精度変換後の空間ID: %s\n", spatialIDs)
}

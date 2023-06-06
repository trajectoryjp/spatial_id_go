package main

import (
	"fmt"
	"os"
	"spatial-id/shape"
	"spatial-id/common/object"
	"spatial-id/common/enum"
)

func main() {

	// 中心座標を取得
	pointList := []*object.Point{}
	for _, spatialid := range centers {
		center, _ := shape.GetPointOnSpatialId(spatialid, enum.Center)
		p, _ := object.NewPoint(center[0].Lon(), center[0].Lat(), center[0].Alt())
		pointList = append(pointList, p)
	}
	// 円柱の空間ID取得
	_, err := shape.GetSpatialIdsOnCylinders(pointList, Radius, zoom, IsCapsule)
	if err == nil {
		// エラーが返却されなかった場合異常終了
		fmt.Println(fmt.Errorf("入力値閾値超過値時にエラー未発生"))
		os.Exit(1)
	} else {
		// エラーが返却された場合エラーを標準出力
		fmt.Println(fmt.Errorf("エラー返却確認: %w", err))
	}
}

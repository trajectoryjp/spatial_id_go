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
		center, _ := shape.GetPointOnExtendedSpatialId(spatialid, enum.Center)
		p, _ := object.NewPoint(center[0].Lon(), center[0].Lat(), center[0].Alt())
		pointList = append(pointList, p)
	}
         fmt.Println(pointList[0])
         // 試験条件に合わせて高さを修正
         pointList[0].SetAlt(10.8)
         pointList[1].SetAlt(10.2)
	// 円柱の空間ID取得
	spatialIDs, err := shape.GetExtendedSpatialIdsOnCylinders(pointList, Radius, Hzoom, Vzoom, IsCapsule, shape.IsPrecision(false))
	if err != nil {
		fmt.Println(fmt.Errorf("円柱の空間ID取得時にエラー発生: %w", err))
		os.Exit(1)
	}

	fileName := "./log/output.txt"
	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	for _, id := range spatialIDs {
		fp.WriteString(id)
		fp.WriteString("\n")
	}

}

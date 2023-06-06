package main

import (
	"flag"
	"fmt"
	"os"
	"spatial-id/common/enum"
	"spatial-id/common/object"
	"spatial-id/shape"
	"strconv"
)

func main() {
	var (
		output      = flag.String("o", "/dev/null", "output path")
		pointOutput = flag.String("p", "/dev/null", "point output path")
	)

	// コマンドライン引数解析
	flag.Parse()

	//座標群
	points, _ := shape.GetPointOnExtendedSpatialId("10/10/10/10/10", enum.Vertex)

	// 始点
	start, _ := object.NewPoint(points[5].Lon(), points[5].Lat(), points[5].Alt())
	// 終点
	end, _ := object.NewPoint(points[1].Lon(), points[1].Lat(), points[1].Alt())
	// 水平精度
	var hZoom int64 = 10
	// 垂直精度
	var vZoom int64 = 10

	// 入力の座標出力
	pp, _ := os.Create(*pointOutput)
	defer pp.Close()

	startLon := strconv.FormatFloat(start.Lon(), 'f', -1, 64)
	startLat := strconv.FormatFloat(start.Lat(), 'f', -1, 64)
	startAlt := strconv.FormatFloat(start.Alt(), 'f', -1, 64)
	endLon := strconv.FormatFloat(end.Lon(), 'f', -1, 64)
	endLat := strconv.FormatFloat(end.Lat(), 'f', -1, 64)
	endAlt := strconv.FormatFloat(end.Alt(), 'f', -1, 64)
	pp.WriteString(startLon + "\n")
	pp.WriteString(startLat + "\n")
	pp.WriteString(startAlt + "\n")
	pp.WriteString(endLon + "\n")
	pp.WriteString(endLat + "\n")
	pp.WriteString(endAlt + "\n")

	// 空間ID取得
	spatialIDs, err := shape.GetExtendedSpatialIdsOnLine(start, end, hZoom, vZoom)
	if err != nil {
		fmt.Println(fmt.Errorf("空間ID取得時にエラー発生: %w", err))
		os.Exit(1)
	}

	// 結果を出力
	fmt.Printf("変換後の空間ID: %s\n", spatialIDs)
	fp, _ := os.Create(*output)
	defer fp.Close()

	for _, spatialID := range spatialIDs {
		fp.WriteString(spatialID + "\n")
	}
}

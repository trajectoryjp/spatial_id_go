package main

import (
	"fmt"
	"os"
	"sort"
	"spatial-id/common/enum"
	"spatial-id/common/spatial"
	"strconv"
	"strings"

	"spatial-id/common/object"
	"spatial-id/shape"

	"spatial-id/shape/physics"
)

func main() {
	//総当たりに衝突判定するid範囲の半径(実際の立体よりは大きめにすること)
	var IsRadius int64 = 7

	// 総当たりに衝突判定するid範囲
	var xID int64 = 31
	var yID int64 = 62
	var zID int64 = 11

	//円の中心座標を格納する集合初期化
	pointList := []*object.Point{}

	//円の中心座標を取得
	for _, spatialid := range CenterSpatialIDs {
		center, _ := shape.GetPointOnExtendedSpatialId(spatialid.CenterSpatialID, spatialid.InputOption)
		p, _ := object.NewPoint(center[0].Lon(), center[0].Lat(), center[0].Alt())
		pointList = append(pointList, p)
	}

	centerPoints, _ := shape.ConvertPointListToProjectedPointList(pointList, ProjectedCrs)

	//カプセルのインスタンス格納用配列を定義
	bulletObjects := make([]physics.CapsulePhysics, 0, len(centerPoints))

	var i int64 = 0
	//衝突判定結果の配列をインスタンスにしてループ
	for index, _ := range centerPoints {

		if index != len(centerPoints)-1 {
			bulletObjects = append(bulletObjects, *physics.NewCapsulePhysics(Radius,
				spatial.Point3{centerPoints[i].X, centerPoints[i].Y, centerPoints[i].Alt},
				spatial.Point3{centerPoints[i+1].X, centerPoints[i+1].Y, centerPoints[i+1].Alt}))
		}
		i++
	}

	//衝突判定する各ID範囲から半径を引き、IDの最小値を取得
	minX := xID - IsRadius
	minY := yID - IsRadius
	minZ := zID - IsRadius

	//結果を格納する集合の初期化
	ExtendedSpatialIDs := make([]string, 0)

	//取りうるIDの範囲を定義する。半径*2+1の範囲を設定。
	IDRange := IsRadius*2 + 1

	var ix int64 = 0
	var jy int64 = 0
	var kz int64 = 0

	i = 0
	//ix軸(経度)IDの範囲だけループ
	for ix = 0; ix < IDRange; ix++ {
		//jy軸(緯度)IDの範囲だけループ
		for jy = 0; jy < IDRange; jy++ {
			//kz軸(高さ)IDの範囲だけループ
			for kz = 0; kz < IDRange; kz++ {
				//fmt.Printf("消すログ:ループには入った\n")
				i++
				//ループで参照しているID及び、水平/垂直方向精度から拡張空間IDを生成
				ExtendedSpatialID := []string{fmt.Sprintf("%v/%v/%v/%v/%v", strconv.FormatInt(Hzoom, 10), strconv.FormatInt(minX+ix, 10), strconv.FormatInt(minY+jy, 10), strconv.FormatInt(Vzoom, 10), strconv.FormatInt(minZ+kz, 10))}
				//fmt.Printf("衝突候補空間ID%d:%s\n", i, ExtendedSpatialID)
				//空間IDの中心の地理座標、投影座標を取得
				projPoints, _ := spatialIdToSkspatialPoints(strings.Join(ExtendedSpatialID, ""), enum.Center)
				centerOrthPoint := projPoints[0]

				//空間IDの頂点(ボクセル8頂点)の地理座標、投影座標を取得
				projPoints, _ = spatialIdToSkspatialPoints(strings.Join(ExtendedSpatialID, ""), enum.Vertex)

				sort.Slice(projPoints, func(i, j int) bool { return projPoints[i].X < projPoints[j].X })
				var halfLon float64 = centerOrthPoint.X - projPoints[0].X

				sort.Slice(projPoints, func(i, j int) bool { return projPoints[i].Y < projPoints[j].Y })
				var halfLat float64 = centerOrthPoint.Y - projPoints[0].Y

				sort.Slice(projPoints, func(i, j int) bool { return projPoints[i].Z < projPoints[j].Z })
				var halfAlt float64 = centerOrthPoint.Z - projPoints[0].Z

				halfExtent := spatial.Vector3{halfLon, halfLat, halfAlt}

				//衝突判定を格納する変数を、False(衝突無し)で初期化
				isCollide := false

				for _, bulletObject := range bulletObjects {

					if bulletObject.IsCollideVoxel(centerOrthPoint, halfExtent) {
						isCollide = true
						break
					}
				}

				if isCollide {
					ExtendedSpatialIDs = append(ExtendedSpatialIDs, ExtendedSpatialID...)
				}
			}
		}
	}

	fileName := "./logs/expected.txt"
	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	for _, id := range ExtendedSpatialIDs {
		fp.WriteString(id)
		fp.WriteString("\n")
	}

}

//空間IDから地理座標、投影座標を取得
func spatialIdToSkspatialPoints(
	extendedSpatialId string,
	option enum.PointOption) ([]spatial.Point3, []*object.Point) {
	//点群APIの"空間IDから座標への変換"関数を呼び出して地理座標を取得
	exPoint, _ := shape.GetPointOnExtendedSpatialId(extendedSpatialId, option)
	//fmt.Printf("頂点座標個数:%d\n", len(exPoint))
	ep1, _ := object.NewPoint(exPoint[0].Lon(), exPoint[0].Lat(), exPoint[0].Alt())
	pointList := []*object.Point{ep1}
	if option == enum.Vertex {
		ep2, _ := object.NewPoint(exPoint[1].Lon(), exPoint[1].Lat(), exPoint[1].Alt())
		ep3, _ := object.NewPoint(exPoint[2].Lon(), exPoint[2].Lat(), exPoint[2].Alt())
		ep4, _ := object.NewPoint(exPoint[3].Lon(), exPoint[3].Lat(), exPoint[3].Alt())
		ep5, _ := object.NewPoint(exPoint[4].Lon(), exPoint[4].Lat(), exPoint[4].Alt())
		ep6, _ := object.NewPoint(exPoint[5].Lon(), exPoint[5].Lat(), exPoint[5].Alt())
		ep7, _ := object.NewPoint(exPoint[6].Lon(), exPoint[6].Lat(), exPoint[6].Alt())
		ep8, _ := object.NewPoint(exPoint[7].Lon(), exPoint[7].Lat(), exPoint[7].Alt())
		pointList = []*object.Point{ep1, ep2, ep3, ep4, ep5, ep6, ep7, ep8}
	}
	projectedCrs := 4326
	//点群APIの"地理座標系リストを投影座標系リストに変換"関数を呼び出して投影座標を取得
	projPoints, _ := shape.ConvertPointListToProjectedPointList(pointList, projectedCrs)

	prop1 := spatial.Point3{projPoints[0].X, projPoints[0].Y, projPoints[0].Alt}
	projPointList := []spatial.Point3{prop1}

	if option == enum.Vertex {

		prop2 := spatial.Point3{projPoints[1].X, projPoints[1].Y, projPoints[1].Alt}
		prop3 := spatial.Point3{projPoints[2].X, projPoints[2].Y, projPoints[2].Alt}
		prop4 := spatial.Point3{projPoints[3].X, projPoints[3].Y, projPoints[3].Alt}
		prop5 := spatial.Point3{projPoints[4].X, projPoints[4].Y, projPoints[4].Alt}
		prop6 := spatial.Point3{projPoints[5].X, projPoints[5].Y, projPoints[5].Alt}
		prop7 := spatial.Point3{projPoints[6].X, projPoints[6].Y, projPoints[6].Alt}
		prop8 := spatial.Point3{projPoints[7].X, projPoints[7].Y, projPoints[7].Alt}

		projPointList = []spatial.Point3{prop1, prop2, prop3, prop4, prop5, prop6, prop7, prop8}
	}

	return projPointList, pointList
}

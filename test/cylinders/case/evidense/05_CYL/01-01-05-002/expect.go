package main

import (
	"fmt"
	"strings"
	"os"
	"sort"
	"spatial-id/common/enum"
	"spatial-id/common"
	"spatial-id/common/spatial"
	"spatial-id/common/object"
	"spatial-id/shape"
	"spatial-id/shape/physics"
	"math"
)

func main() {
	//総当たりに衝突判定するid範囲の半径(実際の立体よりは大きめにすること)
	var IsRadius int64 = 3
	//取りうるIDの範囲を定義する。半径*2+1の範囲を設定。
	IDRange := IsRadius*2 + 1

	// 中心座標を取得
	pointList := []*object.Point{}
	for _, spatialid := range centers {
		center, _ := shape.GetPointOnExtendedSpatialId(spatialid, enum.Center)
		p, _ := object.NewPoint(center[0].Lon(), center[0].Lat(), center[0].Alt())
		pointList = append(pointList, p)
	}

        // 試験条件に合わせて高さを修正
        pointList[0].SetAlt(12.6)
        pointList[1].SetAlt(12.4)

	// Webメルカトル係数
	radian := common.DegreeToRadian(pointList[0].Lat())
	factor := 1 / math.Cos(radian)

	// 入力を空間IDに変換
	spatialIDs,_ := shape.GetExtendedSpatialIdsOnPoints(pointList, Hzoom, Vzoom)
	xIDs := []int64{}
	yIDs := []int64{}
	zIDs := []int64{}
	for _, spatialID := range spatialIDs {
		indexes := shape.GetVoxelIDToSpatialID(spatialID)
		xIDs = append(xIDs, indexes[0])
		yIDs = append(yIDs, indexes[1])
		zIDs = append(zIDs, indexes[2])
	}

	// 総当たりに衝突判定するid範囲
	minX, _ := common.Min(xIDs)
	minY, _ := common.Min(yIDs)
	minZ, _ := common.Min(zIDs)
	maxX, _ := common.Max(xIDs)
	maxY, _ := common.Max(yIDs)
	maxZ, _ := common.Max(zIDs)
	minX = minX - IDRange
	minY = minY - IDRange
	minZ = minZ - IDRange
	maxX = maxX + IDRange
	maxY = maxY + IDRange
	maxZ = maxZ + IDRange

	//カプセルのインスタンス格納用配列を定義
	centerPoints, _ := shape.ConvertPointListToProjectedPointList(pointList, ProjectedCrs)
	bulletObjects := []physics.Physics{}

	//衝突判定結果の配列をインスタンスにしてループ
	for index, _ := range centerPoints {

		// 点が球の場合は
		if len(centerPoints) == 1 {
			bulletObject := physics.NewSpherePhysics(
				Radius * factor,
				spatial.Point3{centerPoints[index].X, centerPoints[index].Y, centerPoints[index].Alt * factor},
			)
			bulletObjects = append(bulletObjects, *bulletObject)
		}

		// 経路のオブジェクト
		if index != len(centerPoints)-1 {
			var object physics.Physics
			if IsCapsule {
				bulletObject := physics.NewCapsulePhysics(
					Radius * factor,
					spatial.Point3{centerPoints[index].X, centerPoints[index].Y, centerPoints[index].Alt * factor},
					spatial.Point3{centerPoints[index+1].X, centerPoints[index+1].Y, centerPoints[index+1].Alt * factor},
				)
				object = *bulletObject

			} else {
				bulletObject := physics.NewCylinderPhysics(
					Radius * factor,
					spatial.Point3{centerPoints[index].X, centerPoints[index].Y, centerPoints[index].Alt * factor},
					spatial.Point3{centerPoints[index+1].X, centerPoints[index+1].Y, centerPoints[index+1].Alt * factor},
				)
				object = *bulletObject
			}

			bulletObjects = append(bulletObjects, object)
		}

		// 円柱の場合は接続点の球
		if !IsCapsule && index != 0 && index != len(centerPoints)-1 {
			bulletObject := physics.NewSpherePhysics(
				Radius * factor,
				spatial.Point3{centerPoints[index].X, centerPoints[index].Y, centerPoints[index].Alt * factor},
			)
			bulletObjects = append(bulletObjects, *bulletObject)
		}
	}

	//結果を格納する集合の初期化
	ExtendedSpatialIDs := make([]string, 0)


	//ix軸(経度)IDの範囲だけループ
	for ix := minX; ix <= maxX; ix++ {
		//jy軸(緯度)IDの範囲だけループ
		for jy := minY; jy <= maxY; jy++ {
			//kz軸(高さ)IDの範囲だけループ
			for kz := minZ; kz <= maxZ; kz++ {
				//ループで参照しているID及び、水平/垂直方向精度から拡張空間IDを生成
				ExtendedSpatialID := shape.GetSpatialIDOnAxisIDs(ix, jy, kz, Hzoom, Vzoom)
				//空間IDの中心の地理座標、投影座標を取得
				projPoints, _ := spatialIdToSkspatialPoints(ExtendedSpatialID, enum.Center, factor)
				centerOrthPoint := projPoints[0]

				//空間IDの頂点(ボクセル8頂点)の地理座標、投影座標を取得
				projPoints, _ = spatialIdToSkspatialPoints(ExtendedSpatialID, enum.Vertex, factor)

				// 対角線成分を取得
				sort.Slice(projPoints, func(i, j int) bool { return projPoints[i].X < projPoints[j].X })
				var lensLon float64 = projPoints[len(projPoints)-1].X - projPoints[0].X

				sort.Slice(projPoints, func(i, j int) bool { return projPoints[i].Y < projPoints[j].Y })
				var lensLat float64 =  projPoints[len(projPoints)-1].Y - projPoints[0].Y

				sort.Slice(projPoints, func(i, j int) bool { return projPoints[i].Z < projPoints[j].Z })
				var lensAlt float64 =  projPoints[len(projPoints)-1].Z - projPoints[0].Z

				lens := spatial.Vector3{lensLon, lensLat, lensAlt}

				//衝突判定を格納する変数を、False(衝突無し)で初期化
				isCollide := false

				for _, bulletObject := range bulletObjects {

					if bulletObject.IsCollideVoxel(centerOrthPoint, lens) {
						isCollide = true
						break
					}
				}

				if isCollide {
					ExtendedSpatialIDs = append(ExtendedSpatialIDs, ExtendedSpatialID)
				}
			}
		}
	}

	// 円柱の場合は衝突判定の内部を埋める
	if !IsCapsule {
		// 経度緯度をキーとした辞書初期化
		lotLatMap := map[string][]string{}

		// 空間IDで経度緯度が同一のものを集約
		for _, spatialID := range ExtendedSpatialIDs {
			ids := strings.Split(spatialID, "/")
			lotLatMap[ids[1]+ids[2]] = append(lotLatMap[ids[1]+ids[2]], spatialID)
		}

		// 経度緯度が同一の空間ID内の高さ間の空間IDを取得
		for _, spatialIDs := range lotLatMap {
			zIndexes := []int64{}
			baseIndexes := shape.GetVoxelIDToSpatialID(spatialIDs[0])

			// 経度緯度が同一の空間ID内の高さを集約
			for _, checkSpatialID := range spatialIDs {
				ids := shape.GetVoxelIDToSpatialID(checkSpatialID)
				zIndexes = append(zIndexes, ids[2])
			}

			// 高さの最大値・最小値を取得
			minZ, _ := common.Min(zIndexes)
			maxZ, _ := common.Max(zIndexes)
			// 高さの最大値・最小値間の空間IDを結果に追加
			for zIndex := minZ; zIndex <= maxZ; zIndex++ {
				spatialID := shape.GetSpatialIDOnAxisIDs(
					baseIndexes[0],
					baseIndexes[1],
					zIndex,
					Hzoom,
					Vzoom,
				)
				ExtendedSpatialIDs = append(ExtendedSpatialIDs, spatialID)
			}
		}

		// 空間IDの重複を削除
		ExtendedSpatialIDs = common.Unique(ExtendedSpatialIDs)
	}


	fileName := "./log/expected.txt"
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
	option enum.PointOption,
	factor float64) ([]spatial.Point3, []*object.Point) {
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
	//点群APIの"地理座標系リストを投影座標系リストに変換"関数を呼び出して投影座標を取得
	projPoints, _ := shape.ConvertPointListToProjectedPointList(pointList, ProjectedCrs)

	prop1 := spatial.Point3{projPoints[0].X, projPoints[0].Y, projPoints[0].Alt * factor}
	projPointList := []spatial.Point3{prop1}

	if option == enum.Vertex {

		prop2 := spatial.Point3{projPoints[1].X, projPoints[1].Y, projPoints[1].Alt * factor}
		prop3 := spatial.Point3{projPoints[2].X, projPoints[2].Y, projPoints[2].Alt * factor}
		prop4 := spatial.Point3{projPoints[3].X, projPoints[3].Y, projPoints[3].Alt * factor}
		prop5 := spatial.Point3{projPoints[4].X, projPoints[4].Y, projPoints[4].Alt * factor}
		prop6 := spatial.Point3{projPoints[5].X, projPoints[5].Y, projPoints[5].Alt * factor}
		prop7 := spatial.Point3{projPoints[6].X, projPoints[6].Y, projPoints[6].Alt * factor}
		prop8 := spatial.Point3{projPoints[7].X, projPoints[7].Y, projPoints[7].Alt * factor}

		projPointList = []spatial.Point3{prop1, prop2, prop3, prop4, prop5, prop6, prop7, prop8}
	}

	return projPointList, pointList
}

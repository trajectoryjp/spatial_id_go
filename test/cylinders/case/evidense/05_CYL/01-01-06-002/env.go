package main

import (
       //"spatial-id/common"
       "math"
       "spatial-id/common/enum"
       "spatial-id/shape"
)


var (
	// 円柱の接続点
	centers = []string{
		"25/29798223/13211885/25/10",
	}

        //  ボクセルの半径算出
        //  頂点座標取得
        vertexPointList, _ = shape.GetPointOnExtendedSpatialId(centers[0], enum.Vertex)
        //  # 投影座標に変換
        vertexProjectedPointList, _ = shape.ConvertPointListToProjectedPointList(vertexPointList, ProjectedCrs)

        xRange float64 = math.Abs(vertexProjectedPointList[0].X - vertexProjectedPointList[1].X)
        yRange float64 = math.Abs(vertexProjectedPointList[0].Y - vertexProjectedPointList[3].Y)
        zRange float64 = math.Abs(vertexProjectedPointList[0].Alt - vertexProjectedPointList[4].Alt)

        //   円柱の半径(単位:m)
	//Radian = common.DegreeToRadian(35.690925154500)
	//Factor = 1 / math.Cos(Radian)
        //Radius = xRange / 2 / Factor
        Radius = xRange / 2 / 1.231260308260132

	// 変換後の水平精度
	Hzoom int64 = 25
	// 変換後の垂直精度
	Vzoom int64 = 25

	//座標参照系
	ProjectedCrs = 3857

	//始点、終点が球状であるかを示す。True: カプセル / False: 円柱
	IsCapsule = true
)

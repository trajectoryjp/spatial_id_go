package main

import (
       "math"
       "spatial-id/common/enum"
       "spatial-id/shape"
)


var (
	// 円柱の接続点
	centers = []string{
		"25/29798223/13211885/25/1",
		"25/29798223/13211885/25/6",
	}

       //  ボクセルの半径算出
       //  頂点座標取得
       vertexPointList, _ = shape.GetPointOnExtendedSpatialId("25/29798223/13211885/25/1", enum.Vertex)
       //  # 投影座標に変換
       vertexProjectedPointList, _ = shape.ConvertPointListToProjectedPointList(vertexPointList, ProjectedCrs)

       yRange float64 = math.Abs(vertexProjectedPointList[0].Y - vertexProjectedPointList[3].Y)

       //   円柱の半径(単位:m)
       Radius = yRange / 2

	// 変換後の水平精度
	Hzoom int64 = 25
	// 変換後の垂直精度
	Vzoom int64 = 25

	//座標参照系
	ProjectedCrs = 3857

	//始点、終点が球状であるかを示す。True: カプセル / False: 円柱
	IsCapsule = true
)

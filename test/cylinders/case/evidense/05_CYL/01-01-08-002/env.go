package main

import (

)


var (
	// 円柱の接続点
	centers = []string{
            "25/29798223/13211885/25/10",
            "25/29798223/13211885/25/10",
	}

	// 円柱の半径(単位:m)
	Radius float64 = 0.2

	// 変換後の水平精度
	Hzoom int64 = 25
	// 変換後の垂直精度
	Vzoom int64 = 25

	//座標参照系
	ProjectedCrs = 3857

	//始点、終点が球状であるかを示す。True: カプセル / False: 円柱
	IsCapsule = true
)

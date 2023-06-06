package main

import (

)


var (
	// 円柱の接続点
	centers = []string{
            "25/1/29798223/13211885",
            "25/3/29798225/13211887",
            "25/4/29798227/13211889",
	}

	// 円柱の半径(単位:m)
	Radius float64 = 2.1

	// 変換後の精度
	zoom int64 = 25

	//座標参照系
	ProjectedCrs = 3857

	//始点、終点が球状であるかを示す。True: カプセル / False: 円柱
	IsCapsule = true
)

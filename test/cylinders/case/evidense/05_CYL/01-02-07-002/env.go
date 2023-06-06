package main

import (

)


var (
	// 円柱の接続点
	centers = []string{
            "25/12/29798223/13211885",
            "25/10/29798224/13211887",
            "25/11/29798225/13211886",
	}

	// 円柱の半径(単位:m)
	Radius float64 = 4.9

	// 変換後の精度
	zoom int64 = 25

	//座標参照系
	ProjectedCrs = 3857

	//始点、終点が球状であるかを示す。True: カプセル / False: 円柱
	IsCapsule = false
)

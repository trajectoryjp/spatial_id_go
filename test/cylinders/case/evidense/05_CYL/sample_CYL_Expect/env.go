package main

import "spatial-id/common/enum"

type Centers struct {
	CenterSpatialID string           // 円の中心拡張空間ID
	InputOption     enum.PointOption // 中心・頂点座標指定
}

var (
	// 円柱の中心の接続点
	CenterSpatialIDs = []Centers{
		Centers{"25/30/60/25/12", enum.Center},
		Centers{"25/33/64/25/10", enum.Center},
		Centers{"25/37/61/25/16", enum.Center},
	}
	// 円柱の半径(単位:m)
	Radius float64 = 2.1

	// 変換後の水平精度
	Hzoom int64 = 25
	// 変換後の垂直精度
	Vzoom int64 = 25

	//座標参照系
	ProjectedCrs = 4326

	//始点、終点が球状であるかを示す。True: カプセル / False: 円柱
	IsCapsule = true
)

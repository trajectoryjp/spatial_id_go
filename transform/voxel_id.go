package transform

import (
	// "fmt"

	"strconv"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/v3/common/consts"
)

// GetVoxelIDfromSpatialID ボクセル成分ID取得
// spatial_id_plus_goのGetVoxcelIDtoSpatialID()から移動
//
// 拡張空間IDからボクセル成分ID取得
//
// 引数：
//
//	spatialID： 取得するボクセルの拡張空間ID
//
// 戻り値：
//
//	(xインデックス, yインデックス, vインデックス)
func GetVoxelIDfromSpatialID(spatialID string) []int64 {
	ids := strings.Split(spatialID, consts.SpatialIDDelimiter)
	lonIndex, _ := strconv.ParseInt(ids[1], 10, 64)
	latIndex, _ := strconv.ParseInt(ids[2], 10, 64)
	altIndex, _ := strconv.ParseInt(ids[4], 10, 64)

	return []int64{
		lonIndex,
		latIndex,
		altIndex,
	}
}

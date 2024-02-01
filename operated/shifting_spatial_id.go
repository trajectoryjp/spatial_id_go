// Package operated 拡張空間ID操作パッケージ
package operated

import (
	"fmt"

	"github.com/trajectoryjp/spatial_id_go/common"
	"github.com/trajectoryjp/spatial_id_go/common/object"
)

// Get6spatialIdsAdjacentToFaces 拡張空間IDの面6個の拡張空間ID取得関数
//
// 拡張空間IDの面に直接、接している6個の拡張空間IDを取得する。
//
// 引数：
//
//	spatialID： 元の位置となる拡張空間ID
//
// 戻り値：
//
//	拡張空間IDスライス： []string
func Get6spatialIdsAdjacentToFaces(spatialID string) []string {
	// 返却用リスト
	spatialIDs := make([]string, 0, 6)
	var shiftIndex int64
	for shiftIndex = -1; shiftIndex < 2; shiftIndex += 2 {
		// 経度方向の移動
		xShiftID := GetShiftingSpatialID(spatialID, shiftIndex, 0, 0)
		spatialIDs = append(spatialIDs, xShiftID)
		// 緯度方向の移動
		yShiftID := GetShiftingSpatialID(spatialID, 0, shiftIndex, 0)
		spatialIDs = append(spatialIDs, yShiftID)
		// 高さ方向の移動
		vShiftID := GetShiftingSpatialID(spatialID, 0, 0, shiftIndex)
		spatialIDs = append(spatialIDs, vShiftID)
	}

	return spatialIDs
}

// Get8spatialIdsAroundHorizontal 拡張空間IDの水平方向の一周分の8個の拡張空間IDの拡張空間ID取得関数
//
// 拡張空間IDの水平方向の周囲、一周分の8個の拡張空間IDを取得する。
//
// 引数：
//
//	spatialID： 元の位置となる拡張空間ID
//
// 戻り値：
//
//	拡張空間IDスライス： []string
func Get8spatialIdsAroundHorizontal(spatialID string) []string {
	// 返却用リスト
	spatialIDs := make([]string, 0, 8)
	var shiftIndex int64
	for shiftIndex = -1; shiftIndex < 2; shiftIndex += 2 {
		// 経度方向の移動
		xShiftID := GetShiftingSpatialID(spatialID, shiftIndex, 0, 0)
		spatialIDs = append(spatialIDs, xShiftID)
		// 緯度方向の移動
		yShiftID := GetShiftingSpatialID(spatialID, 0, shiftIndex, 0)
		spatialIDs = append(spatialIDs, yShiftID)
		// 水平方向に右肩下がりの斜め方向の移動
		downShiftID := GetShiftingSpatialID(spatialID, shiftIndex, shiftIndex, 0)
		spatialIDs = append(spatialIDs, downShiftID)
		// 水平方向に右肩上がりの斜め方向の移動
		upShiftID := GetShiftingSpatialID(spatialID, shiftIndex, -shiftIndex, 0)
		spatialIDs = append(spatialIDs, upShiftID)
	}

	return spatialIDs
}

// Get26spatialIdsAroundVoxel 拡張空間IDを囲う26個の拡張空間ID取得関数
//
// 拡張空間IDを囲う26個の拡張空間IDを取得する。
//
// 引数：
//
//	spatialID： 元の位置となる拡張空間ID
//
// 戻り値：
//
//	拡張空間IDスライス： []string
func Get26spatialIdsAroundVoxel(spatialID string) []string {
	// 返却用リスト
	spatialIDs := make([]string, 0, 26)
	// 入力された拡張空間IDからみて
	//  高さが一つ分低い位置の拡張空間ID、
	//  同じ高さの拡張空間ID、
	//  一つ分高い位置の拡張空間ID
	// を取得する。
	var shiftIndex int64
	for shiftIndex = -1; shiftIndex < 2; shiftIndex++ {
		// 高さを移動した拡張空間IDを取得する。
		vShiftID := GetShiftingSpatialID(spatialID, 0, 0, shiftIndex)

		// 高さを移動が無い場合は入力元の拡張空間IDであるため、格納しない
		if shiftIndex != 0 {
			// 取得した空間idを返却用リストに格納する。
			spatialIDs = append(spatialIDs, vShiftID)
		}

		// 高さを移動した拡張空間IDの水平方向の拡張空間IDを取得し、返却用リストに格納する。
		horizon8IDs := Get8spatialIdsAroundHorizontal(vShiftID)
		spatialIDs = append(spatialIDs, horizon8IDs...)
	}

	return spatialIDs
}

// GetNspatialIdsAroundVoxels 拡張空間ID（複数）を囲う"N"個の拡張空間ID取得関数
//
// 拡張空間ID（一個以上）を囲う"N"個の拡張空間IDを取得する。
//
// 引数：
//
//	spatialIDs： 元の位置となる拡張空間IDs（スライス）
//	hLayers: 水平方向の層目（>= 1）
//	 vLayers: 垂直方向の層目（>= 1）
//
// 戻り値：
//
//	拡張空間IDスライス： []string
//	 error: エラー
func GetNspatialIdsAroundVoxcels(spatialIDs []string, hLayers, vLayers int64) ([]string, error) {

	if hLayers < 1 || vLayers < 1 {
		return nil, fmt.Errorf("both hLayers and vLayers parameters must be >= 1")
	}

	hExpandParam := hLayers * 2
	vExpandParam := vLayers * 2

	nIds := ((vExpandParam + 1) * (hExpandParam + 1) * (hExpandParam + 1)) - 1

	finalspatialIDs := make([]string, 0, int(nIds))

	var xShiftIndex int64
	var yShiftIndex int64
	var vShiftIndex int64

	// Loop through each X, Y, Z shift index and each SpatialID
	for xShiftIndex = -hLayers; xShiftIndex < hLayers+1; xShiftIndex += 1 {
		for yShiftIndex = -hLayers; yShiftIndex < hLayers+1; yShiftIndex += 1 {
			for vShiftIndex = -vLayers; vShiftIndex < vLayers+1; vShiftIndex += 1 {

				if xShiftIndex == 0 && yShiftIndex == 0 && vShiftIndex == 0 {
					continue
				}

				var shiftIDs = []string{}

				for _, spatialID := range spatialIDs {

					shiftedID := GetShiftingSpatialID(spatialID, xShiftIndex, yShiftIndex, vShiftIndex)
					shiftIDs = append(shiftIDs, shiftedID)
				}

				finalspatialIDs = append(finalspatialIDs, shiftIDs...)

			}

		}
	}

	finalspatialIDs = common.Unique(finalspatialIDs)

	return finalspatialIDs, nil
}

// GetShiftingSpatialID 拡張空間IDの移動関数
//
// 指定の数値分、移動した場合の拡張空間IDを取得する。
// 水平方向の移動は、南緯、東経方向が正、北緯、西経方向を負とする。
// 垂直方向の移動は、上空方向が正、地中方向を負とする。
//
// 引数：
//
//	spatialID： 元の位置となる拡張空間ID
//	x： 拡張空間IDを経度方向に動かす数値
//	y： 拡張空間IDを緯度方向に動かす数値
//	v： 拡張空間IDを高さ方向に動かす数値
//
// 戻り値：
//
//	指定の数値分、移動した場合の拡張空間ID：string
func GetShiftingSpatialID(spatialID string, x, y, v int64) string {
	// 拡張空間IDを分解して経度、緯度、高さの位置を取得する
	extendedSpatialID, error := object.NewExtendedSpatialID(spatialID)
	if error != nil {
		return ""
	}
	// // IDから水平方向の位置を取得する。前方2桁は精度
	lonIndex := extendedSpatialID.X()
	latIndex := extendedSpatialID.Y()
	altIndex := extendedSpatialID.Z()

	// シフト後のインデックスを計算して、設定する
	shiftXIndex := lonIndex + x
	shiftYIndex := latIndex + y
	shiftAltIndex := altIndex + v

	extendedSpatialID.SetX(shiftXIndex)
	extendedSpatialID.SetY(shiftYIndex)
	extendedSpatialID.SetZ(shiftAltIndex)

	return extendedSpatialID.ID()
}

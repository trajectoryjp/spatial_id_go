// 拡張空間IDパッケージ
package transform

import (
	"math"
	"strconv"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/v3/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v3/common/object"
	"github.com/trajectoryjp/spatial_id_go/v3/integrate"
	"github.com/trajectoryjp/spatial_id_go/v3/shape"
)

// 宣言
var (
	alt25 = math.Pow(2, 25)
)

// VerticalIndexAltitudes represents the lower limit (inclusive) and upper limit (exclusive) of a VerticalIndex
type VerticalIndexAltitudes struct {
	MinAltitude float64
	MaxAltitude float64
}

// ConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs 内部形式IDを拡張空間IDに変換する。
//
// 変換後の水平方向のIDはXYZ形式のタイルIDとなる。
// 最高高度、最低高度の両方が同じ値の場合、変換前の高さ方向のIDを拡張空間IDインデックス形式とする。
// 最高高度＞最低高度となる値が入力されている場合、変換前の高さ方向のIDを2分木におけるbit形式とする。
// 最高高度＜最低高度となる値が入力されている場合、エラーとする。
// 引数で入力する精度は出力の拡張空間IDの精度となる。
// 入力された構造体に格納された内部形式IDの組み合わせを拡張空間IDに変換する。
//
// 変換前と変換後の精度差によって出力される拡張空間IDの個数は増減する。
//
// 引数 :
//
//	quadkeyAndVerticalIDs : 変換対象の内部形式ID、精度、最高高度、最低高度が格納された構造体のポインタ。
//
//	outputHZoom : 入力値が変換後の拡張空間IDの水平精度となる。拡張空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
//	outputVZoom : 入力値が変換後の拡張空間IDの高さの精度となる。拡張空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
// 戻り値 :
//
//	拡張空間IDのスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過(出力精度)           ：出力の水平方向精度、または高さ方向の精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 精度閾値超過(入力精度)           ：入力のquadkeyの精度に 1 ～ 31、または高さの方向精度に0 ～ 35  の整数値以外が入力されていた場合。
//	 変換条件不正           ：最高高度＜最低高度となる値が入力された場合。
//	 変換条件不正           ：最高高度＞最低高度となる値が入力された場合に高さの方向のID>2^(高さ方向の精度+1)となる値が入力されていた場合。
//
// 補足事項：
//
//	高さの方向のIDについて：
//	 最高高度、最低高度の両方が同じ値の場合、変換前の高さ方向のIDを拡張空間IDインデックス形式とする。
//	 最高高度＞最低高度となる値が入力されている場合、変換前の高さ方向のIDを2分木におけるbit形式とする。
//	 最高高度＜最低高度となる値が入力されている場合、エラーとする。
//
// パラメータ変換例：
//
//	quadkeyZoom : 「6」
//	quadkey : 「2914」
//	vZoom : 「106」
//	verticalID : 「7」
//	maxHeight : 「256」
//	minHeight : 「-256」
//	outputHZoom : 「6」
//	outputVZoom : 「26」
//	を変換する場合
//
//	[]string{
//	  ["6/24/53/26/52"],["6/24/53/26/51"],["6/24/53/26/50"],["6/24/53/26/49"],["6/24/53/26/48"]
func ConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs(quadkeyAndVerticalIDs []*object.QuadkeyAndVerticalID, outputHZoom int64, outputVZoom int64) ([]string, error) {
	extendedSpatialIDs := []string{}

	// 出力精度の入力チェック
	if !extendedSpatialIDCheckZoom(outputHZoom, outputVZoom) {
		return []string{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}

	// IDを一つずつ処理する。
	for _, quadkeyAndVerticalID := range quadkeyAndVerticalIDs {
		// quadkeyの入力精度をチェックする
		if !quadkeyCheckZoom(quadkeyAndVerticalID.QuadkeyZoom(), quadkeyAndVerticalID.VZoom()) {
			return []string{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}
		// quadkeyの入力チェック
		if quadkeyAndVerticalID.Quadkey() > 4611686018427388064 {
			return []string{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}
		xIndex, yIndex := convertQuadkeyToHorizontalID(quadkeyAndVerticalID.Quadkey(), quadkeyAndVerticalID.QuadkeyZoom())
		// 最高高度と最低高度が同値の場合は初期値(0)とみなし、拡張空間IDインデックス形式とする。
		if quadkeyAndVerticalID.MaxHeight() == quadkeyAndVerticalID.MinHeight() {
			// 精度変換を実行し、それぞれのインデックスのスライスを取得する組み合わせた拡張空間IDを作成する。
			horizontalIDs := integrate.HorizontalZoom(quadkeyAndVerticalID.QuadkeyZoom(), xIndex, yIndex, outputHZoom)
			verticalIDs := integrate.VerticalZoom(quadkeyAndVerticalID.VZoom(), quadkeyAndVerticalID.VIndex(), outputVZoom)
			for _, horizontalID := range horizontalIDs {
				for _, verticalID := range verticalIDs {
					extendedSpatialIDs = append(extendedSpatialIDs, horizontalID+"/"+verticalID)
				}
			}
		} else if quadkeyAndVerticalID.MaxHeight() > quadkeyAndVerticalID.MinHeight() {
			// 高さ方向のIDの入力チェック
			if quadkeyAndVerticalID.VIndex() > int64(math.Pow(2, float64(quadkeyAndVerticalID.VZoom()+1))) {
				return []string{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
			}
			// 水平方向と高さ方向のIDをそれぞれ精度変換し、組み合わせて拡張空間IDとする
			verticalIDs := convertBitToVerticalID(quadkeyAndVerticalID.VZoom(), quadkeyAndVerticalID.VIndex(), outputVZoom, quadkeyAndVerticalID.MaxHeight(), quadkeyAndVerticalID.MinHeight())
			// 精度変換済みのIDを取得する。
			horizontalIDs := integrate.HorizontalZoom(quadkeyAndVerticalID.QuadkeyZoom(), xIndex, yIndex, outputHZoom)

			for _, horizontalID := range horizontalIDs {
				for _, verticalID := range verticalIDs {
					extendedSpatialIDs = append(extendedSpatialIDs, horizontalID+"/"+verticalID)
				}
			}
		} else {
			// 高度エラー
			return []string{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}
	}
	// 拡張空間ID文字列のスライスを返却する
	return deleteDuplicationList(extendedSpatialIDs), nil
}

// ConvertQuadkeysAndVerticalIDsToSpatialIDs 内部形式IDを空間IDに変換する。
//
// 変換後の水平方向のIDはXYZタイル形式のIDとなる。
// 最高高度、最低高度の両方が同じ値の場合、変換前の高さ方向のIDを空間IDインデックス形式とする。
// 最高高度＞最低高度となる値が入力されている場合、変換前の高さ方向のIDを2分木におけるbit形式とする。
// 最高高度＜最低高度となる値が入力されている場合、エラーとする。
// 引数で入力する精度は出力の空間IDの精度となる。
// 入力された構造体に格納された内部形式IDの組み合わせを空間IDに変換する。
//
// 変換前と変換後の精度差によって出力される空間IDの個数は増減する。
//
// 引数 :
//
//	quadkeyAndVerticalIDs : 変換対象の内部形式ID、精度、最高高度、最低高度が格納された構造体のポインタ。
//
//	outputZoom : 入力値が変換後の空間IDの精度となる。空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
// 戻り値 :
//
//	空間IDのスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過(出力精度)           ：出力の水平方向精度、または高さ方向の精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 精度閾値超過(入力精度)           ：入力のquadkeyの精度に 1 ～ 31、または高さの方向精度に0 ～ 35  の整数値以外が入力されていた場合。
//	 変換条件不正           ：最高高度＜最低高度となる値が入力された場合。
//	 変換条件不正           ：最高高度＞最低高度となる値が入力された場合に高さの方向のID>2^(高さ方向の精度+1)となる値が入力されていた場合。
//
// 補足事項：
//
//	高さの方向のIDについて：
//	 最高高度、最低高度の両方が同じ値の場合、変換前の高さ方向のIDを空間IDインデックス形式とする。
//	 最高高度＞最低高度となる値が入力されている場合、変換前の高さ方向のIDを2分木におけるbit形式とする。
//	 最高高度＜最低高度となる値が入力されている場合、エラーとする。
//
// パラメータ変換例：
//
//	quadkeyZoom : 「6」
//	quadkey : 「2914」
//	vZoom : 「106」
//	verticalID : 「7」
//	maxHeight : 「256」
//	minHeight : 「-256」
//	outputZoom : 「6」
//	を変換する場合
//
//	[]string{
//	  ["6/52/24/53"],["6/51/24/53"],["6/50/24/53"],["6/49/24/53"],["6/48/24/53"]
func ConvertQuadkeysAndVerticalIDsToSpatialIDs(quadkeyAndVerticalIDs []*object.QuadkeyAndVerticalID, outputZoom int64) ([]string, error) {
	extendedSpatialIDs, error := ConvertQuadkeysAndVerticalIDsToExtendedSpatialIDs(quadkeyAndVerticalIDs, outputZoom, outputZoom)
	spatialIDs := []string{}
	if error != nil {
		return []string{}, error
	}
	for _, extendedSpatialID := range extendedSpatialIDs {

		spatialValue := strings.Split(extendedSpatialID, "/")
		// 精度/高さのインデックス/xインデックス/yインデックス に並び替える
		spatialIDs = append(spatialIDs, spatialValue[0]+"/"+spatialValue[4]+"/"+spatialValue[1]+"/"+spatialValue[2])
	}

	return spatialIDs, nil
}

// ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs 拡張空間IDを内部形式IDに変換する。
//
// 最高高度、最低高度の両方が同じ値の場合、変換後の高さ方向のIDを拡張空間IDインデックス形式とする。
// 最高高度＞最低高度となる値が入力されている場合、変換後の高さ方向のIDを2分木におけるbit形式とする。
// 最高高度＜最低高度となる値が入力されている場合、エラーとする。
//
// 変換前と変換後の精度差によって出力される内部形式IDの個数は増減する。
//
// 引数 :
//
//	extendedSpatialIDs : 変換対象の拡張空間IDのスライス
//
//	outputHZoom : 入力値が変換後のQuadkeyの精度となる。quadkeyの精度の閾値である 1 ～ 31 の整数値を指定可能。
//
//	outputVZoom : 入力値が変換後の高さの方向IDの精度となる。0 ～ 35 の整数値を指定可能。
//
//	maxHeight : 最高高度。高さのIDを2分木におけるbit形式とする場合は最高高度＞最低高度となる値を入力する。拡張空間IDインデックス形式とする場合は最低高度と同値を入力する。
//
//	minHeight : 最低高度。高さのIDを2分木におけるbit形式とする場合は最高高度＞最低高度となる値を入力する。拡張空間IDインデックス形式とする場合は最高高度と同値を入力する。
//
// 戻り値 :
//
//	水平方向精度、[quadkey,高さのID]のスライス、垂直方向精度、最高高度、最低高度の要素を持った構造体のスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過          ：水平方向精度に 1 ～ 31、高さの方向精度が拡張空間IDインデックス形式の場合に 0 ～ 35 の整数値以外が入力されていた場合、または2分木におけるbit形式の場合に 0 ～ 35 の整数値以外が入力されていた場合。
//	 拡張空間IDフォーマット不正：拡張空間IDのフォーマットに違反する値が"変換対象の拡張空間ID"に入力されていた場合。
//	 変換条件不正          ：最高高度＜最低高度となる値が入力されていた場合。
//
// 補足事項：
//
//	精度の有効範囲について：
//	 精度「0」はquadkeyで表現できないため、エラーとする。
//
//	高さの方向のIDについて：
//	 最高高度、最低高度の両方が同じ値の場合、変換後の高さ方向のIDを拡張空間IDインデックス形式とする。
//	 最高高度＞最低高度となる値が入力されている場合、変換後の高さ方向のIDを2分木におけるbit形式とする。
//	 最高高度＜最低高度となる値が入力されている場合、エラーとする。
//
//	最高高度と最低高度について：
//	 拡張空間IDから変換した高さが、最高高度、または最低高度を超えていた場合は、エラーとせず、最大、最小のIDを返却する。
//
// パラメータ変換例：
//
//	extendedSpatialIDs :「["6/24/53/26/51"]」
//	outputHZoom : 「6」
//	outputVZoom : 「7」
//	maxHeight : 「256」
//	minHeight : 「-256」
//	を変換する場合
//
//	[]struct{
//		quadkeyZoom : 6
//		innerIDList : [[2914,75],[2914,74],[2914,73],[2914,72]]
//		vZoom : 7
//		maxHeight : 256
//		minHeight : -256
//	}(スライスの要素は順不同)
//
//	extendedSpatialIDs :「["6/24/53/26/51"]」
//	outputHZoom : 「6」
//	outputVZoom : 「26」
//	maxHeight : 「0」
//	minHeight : 「0」
//	を変換する場合
//
//	[]struct{
//		quadkeyZoom : 6
//		innerIDList : [[2914,51]]
//		vZoom : 26
//		maxHeight : 0
//		minHeight : 0
//	}(スライスの要素は順不同)
func ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs(extendedSpatialIDs []string, outputHZoom int64, outputVZoom int64, maxHeight float64, minHeight float64) ([]*object.FromExtendedSpatialIDToQuadkeyAndVerticalID, error) {
	extendedSpatialIDToQuadkeyAndVerticalID := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}
	// 精度の判定
	if !quadkeyCheckZoom(outputHZoom, outputVZoom) {
		return []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	// outputのZoomレベルが指定されている前提のため、QuadkeyとAltitudeKeyのみを比較
	deduplication := map[[2]int64]interface{}{}

	for _, spatialID := range extendedSpatialIDs {
		vIndexes := []int64{}
		quadkeies := []int64{}
		// 拡張空間IDを水平方向と垂直方向に分割する
		indexes := strings.Split(spatialID, "/")
		indexesInt := []int64{}
		for _, index := range indexes {
			value, e := strconv.ParseInt(index, 10, 64)
			if e != nil {
				return []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}, errors.NewSpatialIdError(errors.InputValueErrorCode, e.Error())
			}
			indexesInt = append(indexesInt, value)
		}

		hZoom := indexesInt[0]
		xIndex := indexesInt[1]
		yIndex := indexesInt[2]
		vZoom := indexesInt[3]
		vIndex := indexesInt[4]
		//　入力精度のチェック
		if !extendedSpatialIDCheckZoom(hZoom, vZoom) {
			return []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}
		// 水平方向の精度変換をする
		horizontalIDs := integrate.HorizontalZoom(hZoom, xIndex, yIndex, outputHZoom)
		// 精度変換後のインデックスを引数にquadkeyを取得する
		for _, horizontalID := range horizontalIDs {
			quadkey := convertHorizontalIDToQuadkey(horizontalID)
			quadkeies = append(quadkeies, quadkey)
		}
		if maxHeight == minHeight {
			// 精度変換のみ実行する。重複しているIDを削除する
			verticalIDs := deleteDuplicationList(integrate.VerticalZoom(vZoom, vIndex, outputVZoom))
			for _, verticalID := range verticalIDs {
				// 「精度/高さのインデックス」 から高さのインデックスのみを抽出する。
				vIndex, _ := strconv.ParseInt(strings.Split(verticalID, "/")[1], 10, 64)
				vIndexes = append(vIndexes, vIndex)
			}

		} else if maxHeight > minHeight {
			// Bit形式の数値に変換する。
			vIndexes = convertVerticallIDToBit(vZoom, vIndex, outputVZoom, maxHeight, minHeight)
		} else {
			// 高度の入力値エラー
			return extendedSpatialIDToQuadkeyAndVerticalID,
				errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}
		// 水平方向と垂直方向の組み合わせを作成する
		idList := [][2]int64{}
		for _, quadkey := range quadkeies {
			for _, vIndex := range vIndexes {
				newID := [2]int64{quadkey, vIndex}
				if _, ok := deduplication[newID]; ok {
					continue
				} else {
					deduplication[newID] = new(interface{})
				}
				idList = append(idList, [2]int64{quadkey, vIndex})
			}
		}
		if len(idList) == 0 {
			continue
		}
		newQuadkeyAndVerticalID := object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(outputHZoom, idList, outputVZoom, maxHeight, minHeight)
		extendedSpatialIDToQuadkeyAndVerticalID = append(extendedSpatialIDToQuadkeyAndVerticalID, newQuadkeyAndVerticalID)
	}

	// 構造体のスライスを返却する。
	return extendedSpatialIDToQuadkeyAndVerticalID, nil
}

// func ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDsV2(extendedSpatialIDs []string, outputHZoom int64, outputVZoom int64, outputMaxHeight int64, outputMinHeight int64) ([]*object.FromExtendedSpatialIDToQuadkeyAndVerticalID, error) {
// 	extendedSpatialIDToQuadkeyAndVerticalID := []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}

// 	// validate zoom levels
// 	if !quadkeyCheckZoom(outputHZoom, outputVZoom) {
// 		return []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
// 	}
// 	// outputのZoomレベルが指定されている前提のため、QuadkeyとAltitudeKeyのみを比較
// 	deduplication := map[[2]int64]interface{}{}

// 	for _, idString := range extendedSpatialIDs {
// 		vIndexes := []int64{}
// 		quadkeys := []int64{}

// 		currentID, error := object.NewExtendedSpatialID(idString)
// 		if error != nil {
// 			return nil, error
// 		}

// 		// hZoom := indexesInt[0]
// 		// xIndex := indexesInt[1]
// 		// yIndex := indexesInt[2]
// 		// vZoom := indexesInt[3]
// 		// vIndex := indexesInt[4]

// 		// check zoom of currentID
// 		if !extendedSpatialIDCheckZoom(currentID.HZoom(), currentID.VZoom()) {
// 			return []*object.FromExtendedSpatialIDToQuadkeyAndVerticalID{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
// 		}
// 		// A. convert horizontal IDs to quadkeys to fit output Horizontal Zoom Level
// 		horizontalIDs := integrate.HorizontalZoom(currentID.HZoom(), currentID.X(), currentID.Y(), outputHZoom)

// 		for _, horizontalID := range horizontalIDs {
// 			quadkey := convertHorizontalIDToQuadkey(horizontalID)
// 			quadkeys = append(quadkeys, quadkey)
// 		}

// 		// B. convert vertical IDs to fit Output Vertical Zoom Level
// 		// Here there are three possible cases and the expected output of each
// 		// 1) outputMaxHeight == outputMinHeight: use input verticalIds
// 		// 2) outputMaxHeight > outputMinHeight: use convertBitIndex to find outputBitIndex
// 		// 3) outputMaxHeight < outputMinHeight: should not be possible. Return an error.
// 		if outputMaxHeight == outputMinHeight {
// 			// 精度変換のみ実行する。重複しているIDを削除する
// 			verticalIDs := deleteDuplicationList(integrate.VerticalZoom(currentID.VZoom(), currentID.Z(), outputVZoom))
// 			for _, verticalID := range verticalIDs {
// 				// 「精度/高さのインデックス」 から高さのインデックスのみを抽出する。
// 				vIndex, _ := strconv.ParseInt(strings.Split(verticalID, "/")[1], 10, 64)
// 				vIndexes = append(vIndexes, vIndex)
// 			}

// 		} else if outputMaxHeight > outputMinHeight {
// 			vIndexes, error = convertVerticalIndex(currentID.VZoom(), currentID.Z(), outputVZoom, outputMaxHeight, outputMinHeight)
// 			if error != nil {
// 				return nil, error
// 			}
// 		} else {
// 			return extendedSpatialIDToQuadkeyAndVerticalID,
// 				errors.NewSpatialIdError(errors.InputValueErrorCode, "")
// 		}
// 		// 水平方向と垂直方向の組み合わせを作成する
// 		idList := [][2]int64{}
// 		for _, quadkey := range quadkeys {
// 			for _, vIndex := range vIndexes {
// 				newID := [2]int64{quadkey, vIndex}
// 				if _, ok := deduplication[newID]; ok {
// 					continue
// 				} else {
// 					deduplication[newID] = new(interface{})
// 				}
// 				idList = append(idList, [2]int64{quadkey, vIndex})
// 			}
// 		}
// 		if len(idList) == 0 {
// 			continue
// 		}
// 		newQuadkeyAndVerticalID := object.NewFromExtendedSpatialIDToQuadkeyAndVerticalID(outputHZoom, idList, outputVZoom, float64(outputMaxHeight), float64(outputMinHeight))
// 		extendedSpatialIDToQuadkeyAndVerticalID = append(extendedSpatialIDToQuadkeyAndVerticalID, newQuadkeyAndVerticalID)
// 	}

// 	// 構造体のスライスを返却する。
// 	return extendedSpatialIDToQuadkeyAndVerticalID, nil
// }

// ConvertSpatialIDsToQuadkeysAndVerticalIDs 空間IDを内部形式IDに変換する。
//
// 最高高度、最低高度の両方が同じ値の場合、変換後の高さ方向のIDを空間IDインデックス形式とする。
// 最高高度＞最低高度となる値が入力されている場合、変換後の高さ方向のIDを2分木におけるbit形式とする。
// 最高高度＜最低高度となる値が入力されている場合、エラーとする。
//
// 変換前と変換後の精度差によって出力される内部形式IDの個数は増減する。
//
// 引数 :
//
//	spatialIDs : 変換対象の空間IDのスライス
//
//	1HZoom : 入力値が変換後のQuadkeyの精度となる。quadkeyの精度の閾値である 1 ～ 31 の整数値を指定可能。
//
//	outputVZoom : 入力値が変換後の高さの方向IDの精度となる。0 ～ 35 の整数値を指定可能。
//
//	maxHeight : 最高高度。高さのIDを2分木におけるbit形式とする場合は最高高度＞最低高度となる値を入力する。空間IDインデックス形式とする場合は最低高度と同値を入力する。
//
//	minHeight : 最低高度。高さのIDを2分木におけるbit形式とする場合は最高高度＞最低高度となる値を入力する。空間IDインデックス形式とする場合は最高高度と同値を入力する。
//
// 戻り値 :
//
//	水平方向精度、[quadkey,高さのID]のスライス、垂直方向精度、最高高度、最低高度の要素を持った構造体のスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過          ：水平方向精度に 1 ～ 31、高さの方向精度が空間IDインデックス形式の場合に 0 ～ 35 の整数値以外が入力されていた場合、または2分木におけるbit形式の場合に 0 ～ 35 の整数値以外が入力されていた場合。
//	 空間IDフォーマット不正：空間IDのフォーマットに違反する値が"変換対象の空間ID"に入力されていた場合。
//	 変換条件不正          ：最高高度＜最低高度となる値が入力されていた場合。
//
// 補足事項：
//
//	精度の有効範囲について：
//	 精度「0」はquadkeyで表現できないため、エラーとする。
//
//	高さの方向のIDについて：
//	 最高高度、最低高度の両方が同じ値の場合、変換後の高さ方向のIDを空間IDインデックス形式とする。
//	 最高高度＞最低高度となる値が入力されている場合、変換後の高さ方向のIDを2分木におけるbit形式とする。
//	 最高高度＜最低高度となる値が入力されている場合、エラーとする。
//
//	最高高度と最低高度について：
//	 空間IDから変換した高さが、最高高度、または最低高度を超えていた場合は、エラーとせず、最大、最小のIDを返却する。
//
// パラメータ変換例：
//
//	spatialIDs :「["6/51/24/53"]」
//	outputHZoom : 「6」
//	outputVZoom : 「7」
//	maxHeight : 「256」
//	minHeight : 「-256」
//	を変換する場合
//
//	[]struct{
//		quadkeyZoom : 6
//		innerIDList : [[2914,75],[2914,74],[2914,73],[2914,72]]
//		vZoom : 7
//		maxHeight : 256
//		minHeight : -256
//	}(スライスの要素は順不同)
//
//	spatialIDs :「["6/51/24/53"]」
//	outputHZoom : 「6」
//	outputVZoom : 「26」
//	maxHeight : 「0」
//	minHeight : 「0」
//	を変換する場合
//
//	[]struct{
//		quadkeyZoom : 6
//		innerIDList : [[2914,51]]
//		vZoom : 26
//		maxHeight : 0
//		minHeight : 0
//	}(スライスの要素は順不同)
func ConvertSpatialIDsToQuadkeysAndVerticalIDs(spatialIDs []string, outputHZoom int64, outputVZoom int64, maxHeight float64, minHeight float64) ([]*object.FromExtendedSpatialIDToQuadkeyAndVerticalID, error) {
	extendedSpatialIDs := []string{}
	for _, spatialID := range spatialIDs {
		spatialIDValue := strings.Split(spatialID, "/")
		// 水平精度/xインデックス/yインデックス/垂直精度/高さのインデックス に並び替える
		extendedSpatialIDs = append(extendedSpatialIDs, spatialIDValue[0]+"/"+spatialIDValue[2]+"/"+spatialIDValue[3]+"/"+spatialIDValue[0]+"/"+spatialIDValue[1])
	}
	return ConvertExtendedSpatialIDsToQuadkeysAndVerticalIDs(extendedSpatialIDs, outputHZoom, outputVZoom, maxHeight, minHeight)
}

// 文字列のスライスから重複した値を削除する。
//
// 引数 :
//
//	duplication_list : 削除対象のスライス
//
// 戻り値 :
//
//	重複を削除した状態のスライス
func deleteDuplicationList(duplicationList []string) []string {
	// 重複の削除
	mList := make(map[string]struct{})
	for _, value := range duplicationList {
		mList[value] = struct{}{}
	}
	returnList := []string{}
	for value := range mList {
		returnList = append(returnList, value)
	}
	return returnList
}

// 経度と緯度のインデックス値からquadkeyを計算する。
//
// 引数 :
//
//	horizontalID : 拡張空間IDの水平部分のID
//
// 戻り値 :
//
//	quadkey
func convertHorizontalIDToQuadkey(horizontalID string) int64 {

	var quadkey int64
	var i int64
	indexes := strings.Split(horizontalID, "/")
	hZoom, _ := strconv.ParseInt(indexes[0], 10, 64)
	xIndexTmp, _ := strconv.ParseInt(indexes[1], 10, 64)
	yIndexTmp, _ := strconv.ParseInt(indexes[2], 10, 64)

	//　インデックス/2が0よりも大きい場合にループを継続する
	// 入力の精度の桁数になった場合、処理を終了する
	for i = 0; xIndexTmp > 0 && i < hZoom; i++ {
		// 商と余りを取得
		mx := xIndexTmp % 2
		x := mx
		xIndexTmp = xIndexTmp / 2
		quadkey += x << (i * 2)
	}
	for i = 0; yIndexTmp > 0 && i < hZoom; i++ {
		my := yIndexTmp % 2
		y := my * 2
		yIndexTmp = yIndexTmp / 2
		quadkey += y << (i * 2)
	}

	return quadkey
}

// quadkeyからXYZタイル形式のインデックスを計算する。
//
// 引数 :
//
//	quadkey : quadkey
//	zoom : 精度
//
// 戻り値 :
//
//	経度方向のインデックス,緯度方向のインデックス
func convertQuadkeyToHorizontalID(quadkey int64, zoom int64) (int64, int64) {

	var x, y int64
	// 4進数文字列に変換
	quadkeyStr := strconv.FormatInt(int64(quadkey), 4)
	for i, s := range strings.Split(quadkeyStr, "") {
		bitX := x << 1
		bitY := y << 1

		if s == "1" {
			bitX++
			bitY += 0
		} else if s == "2" {
			x += 0
			bitY++
		} else if s == "3" {
			bitX++
			bitY++
		} else {
			bitX += 0
			bitY += 0
		}
		x = bitX
		y = bitY
		// 入力の精度に達していた場合、処理を終了する
		if i == int(zoom-1) {
			break
		}

	}
	// Xインデックス、Yインデックスを返却する
	return x, y
}

// 拡張空間ID形式の高さのインデックスをbit形式のインデックスに変換する。
//
// 引数 :
//
//	vZoom : 垂直精度
//
//	vIndex : 拡張空間ID形式の高さのインデックス
//
//	outputZoom : 出力の精度
//
//	maxHeight : 最高高度
//
//	minHeight : 最低高度
//
// 戻り値 :
//
//	2分木によるbit形式のインデックスのスライス
func convertVerticallIDToBit(vZoom int64, vIndex int64, outputZoom int64, maxHeight float64, minHeight float64) []int64 {

	// 拡張空間IDボクセルの高さの最大値と最小値を計算する。
	spatialIDMaxHeight := float64(vIndex+1) * alt25 / float64(math.Pow(2, float64(vZoom)))
	spatialIDMinHeight := float64(vIndex) * alt25 / float64(math.Pow(2, float64(vZoom)))

	// bitの取得
	maxBitIndex := calcBitIndex(spatialIDMaxHeight, outputZoom, maxHeight, minHeight)
	minBitIndex := calcBitIndex(spatialIDMinHeight, outputZoom, maxHeight, minHeight)

	// インデックスが同値の場合、後続の処理は不要。
	if maxBitIndex == minBitIndex {
		return []int64{maxBitIndex}
	}

	// bitをintにキャスト
	bitIndexes := []int64{
		maxBitIndex, minBitIndex,
	}

	// インデックスの隙間を補完し、IDとして返却用変数に格納する。
	for i := minBitIndex + 1; i < maxBitIndex; i++ {
		bitIndexes = append(bitIndexes, i)
	}

	return bitIndexes

}

func convertVerticalIndex(inputIndex int64, inputZoom int64, outputZoom int64, zoomScalar int64, offset int64) ([]int64, error) {

	var (
		outputIndexes                        []int64
		indexAltitues, currentIndexAltitudes *VerticalIndexAltitudes
		error                                error
	)

	// return altitues of original index
	indexAltitues, error = returnAltitudesOfVerticalIndexB(inputIndex, inputZoom, 0, 0)
	if error != nil {
		return nil, error
	}

	// determine the upper and lower index bounds to search for matches in height solution space
	lowerBounds, error := calculateMinVerticalIndex(inputIndex, inputZoom, outputZoom, zoomScalar, offset)
	if error != nil {
		return nil, error
	}
	upperBounds, error := calculateMinVerticalIndex(inputIndex+1, inputZoom, outputZoom, zoomScalar, offset)
	if error != nil {
		return nil, error
	}

	// solve by checking for each index between lower and upper bounds
	for i := lowerBounds; i <= upperBounds; i++ {

		currentIndexAltitudes, error = returnAltitudesOfVerticalIndexB(int64(i), outputZoom, zoomScalar, 0)
		if error != nil {
			return nil, error
		}

		if currentIndexAltitudes.MinAltitude >= indexAltitues.MinAltitude &&
			currentIndexAltitudes.MaxAltitude <= indexAltitues.MaxAltitude {
			outputIndexes = append(outputIndexes, int64(i))
		}

	}

	return outputIndexes, nil

}

// converts a vertical index from one set of zoom parameters to another disregarding the floor() cacluation. This creates a simplier system of equations where the solution set for height is a single variable. However, this does not describe the full solution set of height since we have excluded the floor calculation; it describes the condition where m = x, given m = floor(x) if and only if m <= x < m +1;
func calculateMinVerticalIndex(inputIndex int64, inputZoom int64, outputZoom int64, zoomScalar int64, offset int64) (int64, error) {

	outputIndex := float64(offset) + float64(inputIndex)*float64(calculateVerticalResolution((outputZoom-inputZoom+zoomScalar)))

	return int64(outputIndex), nil

}

// returns the min and max altitudes of a given vertical index, zoomLevel, zoomScalar, and offset (add alpha)
func returnAltitudesOfVerticalIndex(index int64, zoomLevel int64, zoomScalar int64, offset int64) (*VerticalIndexAltitudes, error) {

	MinAltitude := float64(offset) + float64(index)*math.Pow(2, float64(25-zoomLevel-zoomScalar))
	MaxAltitude := float64(offset) + float64(index+1)*math.Pow(2, float64(25-zoomLevel-zoomScalar))
	return &VerticalIndexAltitudes{
		MinAltitude: MinAltitude,
		MaxAltitude: MaxAltitude,
	}, nil
}

// for use in situations where h is a multidimensional set (subtract alpha)
func returnAltitudesOfVerticalIndexB(index int64, zoom int64, zoomScalar int64, offset int64) (*VerticalIndexAltitudes, error) {

	MinAltitude := (float64(index) * math.Pow(2, float64(25)-float64(zoom)-float64(zoomScalar))) - float64(offset)
	MaxAltitude := (float64(index+1) * math.Pow(2, float64(25)-float64(zoom)-float64(zoomScalar))) - float64(offset)

	return &VerticalIndexAltitudes{
		MinAltitude: MinAltitude,
		MaxAltitude: MaxAltitude,
	}, nil
}

// returns the number of indexes from global min to global max altitudes
func calculateVerticalResolution(zoomLevel int64) int64 {
	verticalResolution := math.Pow(2, float64(zoomLevel))
	return int64(verticalResolution)
}

// VoxelHeight = [spatialIDMaxHeight] - [spatialIDMinHeight], or
//
//	= [float64(vIndex+1) * alt25 / float64(math.Pow(2, float64(vZoom))) ] -
//	         float64(vIndex) * alt25 / float64(math.Pow(2, float64(vZoom)))
//	= (globalMaxHeight - globalMinHeight) / float64(math.Pow(2, float64(vZoom)))
func calculateVoxelHeight(vZoom int64, globalMaxHeight float64, globalMinHeight float64) float64 {
	return (globalMaxHeight - globalMinHeight) / float64(calculateVerticalResolution(vZoom))
}

// 高さのbit形式のインデックスを計算する。
//
// 引数 :
//
//	altitude : 高さ
//
//	outputZoom : 出力の精度
//
//	maxHeight : 最高高度
//
//	minHeight : 最低高度
//
// 戻り値 :
//
//	2分木によるbit形式のインデックス
func calcBitIndex(altitude float64, outputZoom int64, maxHeight float64, minHeight float64) int64 {
	var bitIndex int64
	var i int64
	// 出力の精度の回数ループする。
	for i = 0; i < outputZoom; i++ {
		bit := bitIndex << 1 // bit=0, same type as bitIndex
		// 境界の高さを計算する
		borderHeight := (maxHeight-minHeight)/2 + minHeight
		if altitude >= borderHeight {
			bit++
			// 境界値を最低高度に置き換える。
			minHeight = borderHeight
		} else {
			bit += 0
			// 境界値を最高高度に置き換える。
			maxHeight = borderHeight
		}
		bitIndex = bit
	}
	return bitIndex
}

// bit形式のIDを拡張空間ID形式に変換する。
//
// 引数 :
//
//	verticalID : 2分木によるbit形式のID
//
//	outputZoom: 出力の精度
//
//	maxHeight : 最高高度
//
//	minHeight : 最低高度
//
// 戻り値 :
//
//	拡張空間IDインデックス形式の高さの方向のID
func convertBitToVerticalID(vZoom int64, vIndex int64, outputZoom int64, maxHeight float64, minHeight float64) []string {
	// 精度あたりの1ボクセルの高さを計算する。
	voxelHeight := (maxHeight - minHeight) / float64(math.Pow(2, float64(vZoom)))

	// 2分木によるBit化形式のボクセルの高さを計算する。
	maxAltitude := float64(vIndex+1)*voxelHeight + minHeight
	minAltitude := float64(vIndex)*voxelHeight + minHeight

	//高さを拡張空間IDに変換する。
	maxOutputPoint, _ := object.NewPoint(0.0, 0.0, maxAltitude)
	minOutputPoint, _ := object.NewPoint(0.0, 0.0, minAltitude)
	outputPoints := make([]*object.Point, 0, 2)
	outputPoints = append(outputPoints, maxOutputPoint)
	outputPoints = append(outputPoints, minOutputPoint)
	// 高さのID抽出用スライス
	verticalIndexes := []string{}
	vIndexes := []string{}
	// 高さと出力の制度を引数に拡張空間IDインデックス形式のIDを取得する。
	spatialIDs, _ := shape.GetExtendedSpatialIdsOnPoints(outputPoints, 0, outputZoom)

	// 2分木によるBit形式のIDボクセルの最高高度と最低高度を拡張空間IDにする
	for _, id := range spatialIDs {
		// 高さ成分のみを取り出す
		vID := strings.ReplaceAll(id, "0/0/0/", "")
		// インデックスを取得
		vIndexes = append(vIndexes, strings.Split(vID, "/")[1])
	}
	maxIndex, _ := strconv.ParseInt(vIndexes[0], 10, 64)
	minIndex, _ := strconv.ParseInt(vIndexes[1], 10, 64)
	verticalIndexes = append(verticalIndexes, strconv.Itoa(int(outputZoom))+"/"+vIndexes[0])
	verticalIndexes = append(verticalIndexes, strconv.Itoa(int(outputZoom))+"/"+vIndexes[1])

	// 拡張空間IDインデックスの隙間を補完する。同値の場合は補完をスキップする。
	if maxIndex != minIndex {
		for i := minIndex + 1; i < maxIndex; i++ {
			verticalIndexes = append(verticalIndexes, strconv.FormatInt(outputZoom, 10)+"/"+strconv.FormatInt(i, 10))
		}
	}
	return verticalIndexes
}

// quadkeyCheckZoom quadkey用入力精度チェック関数
//
// 入力の精度が水平方向は1-31、垂直方向は0-35の範囲内か判定をする。
//
// 引数：
//
//	hZoom：チェック対象の水平精度
//
//	vZoom：チェック対象の垂直精度
//
// 戻り値：
//
//	入力された精度が水平方向は1-31、垂直方向は0-35の範囲内の場合True、それ以外の場合Falseを返却する。
func quadkeyCheckZoom(hZoom int64, vZoom int64) bool {
	// 水平/垂直精度が範囲内に収まっている場合はTrueを返却
	return (1 <= hZoom && hZoom <= 31) && (0 <= vZoom && vZoom <= 35)
}

// extendedSpatialIDCheckZoom extendedSpatialID用入力精度チェック関数
//
// 入力の精度が水平方向は0-35、垂直方向は0-35の範囲内か判定をする。
//
// 引数：
//
//	hZoom：チェック対象の水平精度
//
//	vZoom：チェック対象の垂直精度
//
// 戻り値：
//
//	入力された精度が水平方向は0-35、垂直方向は0-35の範囲内の場合True、それ以外の場合Falseを返却する。
func extendedSpatialIDCheckZoom(hZoom int64, vZoom int64) bool {
	// 水平/垂直精度が範囲内に収まっている場合はTrueを返却
	return (0 <= hZoom && hZoom <= 35) && (0 <= vZoom && vZoom <= 35)
}

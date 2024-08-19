// 拡張空間IDパッケージ
package transform

import (
	"fmt"
	"golang.org/x/exp/maps"
	"math"
	"strconv"
	"strings"

	"github.com/trajectoryjp/spatial_id_go/v4/common"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v4/common/object"
	"github.com/trajectoryjp/spatial_id_go/v4/integrate"
	"github.com/trajectoryjp/spatial_id_go/v4/shape"
)

// 宣言
var (
	alt25              = math.Pow(2, 25)
	zOriginValue int64 = 25
)

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
	// outputのZoomレベルが指定されている前提のため、QuadkeyとAltitudekeyのみを比較
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

// ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys 拡張空間IDをquadkeyとaltitudekeyに変換する。
//
// 変換前と変換後の精度差によって出力される内部形式IDの個数は増減する。
//
// 引数 :
//
//		extendedSpatialIDs : 変換対象の拡張空間IDのスライス
//
//	 outputQuadkeyZoom : 入力値が変換後のquadkeyの精度となる。quadkeyの精度の閾値である 1 ～ 31 の整数値を指定可能。
//
//	 outputAltitudekeyZoom : 入力値が変換後のaltitudekeyの精度となる。
//
//	 zBaseExponent : altitudekeyの高さが1mとなるズームレベル
//
//	 zBaseOffset : ズームレベルがzBaseExponentで高度0mにおけるaltitudekey
//
// 戻り値 :
//
//	quadkeyの精度、[quadkey,altitudekey]のスライス、altitudekeyの精度、altitudekeyの高さが1mとなるズームレベル、ズームレベルがzBaseExponentで高度0mにおけるaltitudekeyの要素を持った構造体のスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過          ：水平方向精度に 1 ～ 31の整数値以外が入力されていた場合。
//	 拡張空間IDフォーマット不正：拡張空間IDのフォーマットに違反する値が"変換対象の拡張空間ID"に入力されていた場合。
func ConvertExtendedSpatialIDsToQuadkeysAndAltitudekeys(extendedSpatialIDs []string, outputQuadkeyZoom int64, outputAltitudekeyZoom int64, zBaseExponent int64, zBaseOffset int64) ([]*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey, error) {
	extendedSpatialIDToQuadkeyAndAltitudekey := []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{}

	// validate zoom levels
	if !quadkeyCheckZoom(outputQuadkeyZoom, outputAltitudekeyZoom) {
		return []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
	}
	// outputのZoomレベルが指定されている前提のため、QuadkeyとAltitudekeyのみを比較
	duplicate := map[[2]int64]interface{}{}

	for _, idString := range extendedSpatialIDs {
		var altitudeKeys []int64
		quadkeys := []int64{}

		currentID, error := object.NewExtendedSpatialID(idString)
		if error != nil {
			return nil, error
		}

		// check zoom of currentID
		if !extendedSpatialIDCheckZoom(currentID.HZoom(), currentID.VZoom()) {
			return []*object.FromExtendedSpatialIDToQuadkeyAndAltitudekey{}, errors.NewSpatialIdError(errors.InputValueErrorCode, "")
		}
		// A. convert horizontal IDs to quadkeys to fit output Horizontal Zoom Level
		horizontalIDs := integrate.HorizontalZoom(currentID.HZoom(), currentID.X(), currentID.Y(), outputQuadkeyZoom)

		for _, horizontalID := range horizontalIDs {
			quadkey := convertHorizontalIDToQuadkey(horizontalID)
			quadkeys = append(quadkeys, quadkey)
		}

		// B. convert vertical IDs to fit Output Vertical Zoom Level
		altitudeKeys, error = ConvertZToAltitudekey(currentID.Z(), currentID.VZoom(), outputAltitudekeyZoom, zBaseExponent, zBaseOffset)
		if error != nil {
			return nil, error
		}

		// 水平方向と垂直方向の組み合わせを作成する
		idList := [][2]int64{}
		for _, quadkey := range quadkeys {
			for _, altitudeKey := range altitudeKeys {
				newID := [2]int64{quadkey, altitudeKey}
				if _, ok := duplicate[newID]; ok {
					continue
				} else {
					duplicate[newID] = new(interface{})
				}
				idList = append(idList, [2]int64{quadkey, altitudeKey})
			}
		}
		if len(idList) == 0 {
			continue
		}

		newQuadkeyAndVerticalID := object.NewFromExtendedSpatialIDToQuadkeyAndAltitudekey(
			outputQuadkeyZoom,
			idList,
			outputAltitudekeyZoom,
			zBaseExponent,
			zBaseOffset,
		)

		extendedSpatialIDToQuadkeyAndAltitudekey = append(extendedSpatialIDToQuadkeyAndAltitudekey, newQuadkeyAndVerticalID)
	}

	// 構造体のスライスを返却する。
	return extendedSpatialIDToQuadkeyAndAltitudekey, nil
}

// ConvertExtendedSpatialIDToSpatialIDs 拡張空間IDを空間IDへ変換する
//
// 引数 :
//
//	extendedSpatialID : 変換対象の拡張空間ID構造体
//
// 戻り値 :
//
//	変換後の空間IDのスライス
//
// 補足事項：
//
//	入力拡張空間IDの垂直精度と水平精度の間で差がある場合、差が1増えるごとに4のべき乗で出力空間ID数が増加する。
//	そのため、精度差が大きすぎると変換後の空間ID数は大幅に増大する。
//	動作環境によってはメモリ不足となる可能性があるため、注意すること。
//
//	変換利用例：
//	1. 水平精度の方が低い場合
//	入力
//	ExtendedSpatialID{
//		hZoom: 6,
//		x:     24,
//		y:     49,
//		vZoom: 7,
//		z:     0,
//	}
//
//	出力
//	[]string{"7/0/48/98", "7/0/48/99", "7/0/49/98", "7/0/49/99"}
//
//	2. 垂直精度の方が低い場合
//	入力
//	ExtendedSpatialID{
//		hZoom: 7,
//		x:     24,
//		y:     53,
//		vZoom: 6,
//		z:     24,
//	}
//
//	出力
//	[]string{"7/48/24/53", "7/49/24/53"}
//
//	3. 水平精度、垂直精度に差がない場合
//	入力
//	ExtendedSpatialID{
//		hZoom: 6,
//		x:     24,
//		y:     49,
//		vZoom: 6,
//		z:     0,
//	}
//
//	出力
//	[]string{"6/0/24/49"}
func ConvertExtendedSpatialIDToSpatialIDs(extendedSpatialID *object.ExtendedSpatialID) []string {
	spatialIds := []string{}
	// 精度が高い方へズームレベルを上げる
	switch {
	case extendedSpatialID.HZoom() < extendedSpatialID.VZoom():
		targetZoomLevel := extendedSpatialID.VZoom()
		xMin, yMin, xMax, yMax := integrate.HorizontalZoomMinMax(extendedSpatialID.HZoom(), extendedSpatialID.X(), extendedSpatialID.Y(), targetZoomLevel)
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				spatialIds = append(spatialIds, fmt.Sprintf("%v/%v/%v/%v", targetZoomLevel, extendedSpatialID.Z(), x, y))
			}
		}
	case extendedSpatialID.HZoom() > extendedSpatialID.VZoom():
		targetZoomLevel := extendedSpatialID.HZoom()
		verticalIds := integrate.VerticalZoom(extendedSpatialID.VZoom(), extendedSpatialID.Z(), targetZoomLevel)
		for _, verticalId := range verticalIds {
			// "z/f" + "x/y"
			spatialIds = append(spatialIds, fmt.Sprintf("%v/%v/%v", verticalId, extendedSpatialID.X(), extendedSpatialID.Y()))
		}
	// ズームレベルが等しい場合は直接変換可
	case extendedSpatialID.HZoom() == extendedSpatialID.VZoom():
		spatialId := fmt.Sprintf("%v/%v/%v/%v", extendedSpatialID.HZoom(), extendedSpatialID.Z(), extendedSpatialID.X(), extendedSpatialID.Y())
		spatialIds = append(spatialIds, spatialId)
	}
	return spatialIds
}

// ConvertTileXYZToExtendedSpatialIDs
// TileXYZ形式から拡張空間ID形式へ変換する
//
// TileXYZのx,yは同一の意味のまま拡張空間IDのX,Yに変換される。
// このため出力の水平精度は入力のものが用いられる(ただし精度チェックは行われる)。
// TileXYZのzから拡張空間ID垂直インデックス(f)へは高度変換が用いられ変化する。
// 引数 :
//
//	request : 変換対象のTileXYZ構造体のスライス
//	zBaseExponent： TileXYZのzの高さが1mとなるズームレベル
//	zBaseOffset： ズームレベルがzBaseExponentのとき高度0mにおけるTileXYZのz
//	outputVZoom : 入力値が変換後の拡張空間IDの高さの精度となる。拡張空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
// 戻り値 :
//
//	変換後の拡張空間IDのスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過(出力精度)：出力の水平方向精度、または高さ方向の精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 z高度範囲外：変換前のzがその垂直ズームレベルにおける高度範囲外である場合。
//	 拡張空間ID高度範囲外：変換後の拡張空間ID高度がその垂直ズームレベルにおける高度範囲外である場合。
//
// 補足事項：
//
//	高さについて：
//	 TileXYZ形式のzと拡張空間ID形式垂直インデックスは高度基準が異なる(同じにすることも可能)
//	 引数のzBaseExponentとzBaseOffsetで高度基準を定義する
//	 TileXYZ内のデータもしくは引数が次の状態のとき、入力TileXYZ数に対して出力拡張空間ID垂直インデックス数が増加する
//	 - vZoomがzBaseExponentまたは25(空間IDの高度基準におけるzBaseExponent)より大きい場合
//	 - zBaseOffsetが2のべき乗でない場合
//
// 変換利用例：
//
// 1. 入力TileXYZのzが出力拡張空間ID垂直インデックスに対応する場合
//
//	[]TileXYZ{
//	    {
//	        hZoom : 20
//	        x 85263
//	        y 65423
//	        vZoom 23
//	        z 0
//	    }
//	},
//	zBaseExponent 25,
//	zBaseOffset 8,
//	outputVZoom 23
//
//	extendedSpatialIDs :["20/85263/65423/23/-2"]
//
// 2. vZoomが25より大きい場合
//
//	[]TileXYZ{
//	    {
//	        hZoom : 20
//	        x 85263
//	        y 65423
//	        vZoom 26
//	        z 3
//	    }
//	},
//	zBaseExponent 25,
//	zBaseOffset -2,
//	outputVZoom 26
//
//	extendedSpatialIDs :["20/85263/65423/26/7", "20/85263/65423/26/7]
//
// 3. zBaseOffsetが2のべき乗でない場合
//
//	[]TileXYZ{
//	    {
//	        hZoom : 20
//	        x 85263
//	        y 65423
//	        vZoom 23
//	        z 0
//	    }
//	},
//	zBaseExponent 25,
//	zBaseOffset 7,
//	outputVZoom 23
//
//	extendedSpatialIDs :["20/85263/65423/23/-2", "20/85263/65423/23/-1"]
func ConvertTileXYZToExtendedSpatialIDs(request []*object.TileXYZ, zBaseExponent uint16, zBaseOffset int64, outputVZoom uint16) ([]object.ExtendedSpatialID, error) {

	extendedSpatialIDsMap := make(map[int64]object.ExtendedSpatialID)
	extendedSpatialIDs := []object.ExtendedSpatialID{}

	for _, r := range request {
		if !extendedSpatialIDCheckZoom(int64(r.HZoom()), int64(outputVZoom)) {
			return nil, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("extendSpatialID zoom level must be in 0-35: hZoom=%v, vZoom=%v", r.HZoom(), outputVZoom))
		}

		zMin, zMax, err := ConvertAltitudeKeyToZ(r.Z(), int64(r.VZoom()), int64(outputVZoom), int64(zBaseExponent), zBaseOffset)
		if err != nil {
			return nil, err
		}

		for z := zMin; z <= zMax; z++ {
			// 重複排除
			if _, exists := extendedSpatialIDsMap[z]; exists {
				continue
			}
			extendedSpatialID := new(object.ExtendedSpatialID)
			extendedSpatialID.SetX(r.X())
			extendedSpatialID.SetY(r.Y())
			extendedSpatialID.SetZ(z)
			extendedSpatialID.SetZoom(int64(r.HZoom()), int64(outputVZoom))
			extendedSpatialIDsMap[z] = *extendedSpatialID
		}
	}
	for _, extendedSpatialID := range maps.Values(extendedSpatialIDsMap) {
		extendedSpatialIDs = append(extendedSpatialIDs, extendedSpatialID)
	}

	return extendedSpatialIDs, nil
}

// ConvertTileXYZToSpatialIDs
// TileXYZ形式から空間ID形式へ変換する
// ConvertTileXYZToExtendedSpatialIDs と ConvertExtendedSpatialIDToSpatialIDs の組み合わせであるため、詳細はそちらを参照
//
// TileXYZのx,yは水平精度と垂直精度のうち高い方の空間IDのX,Yに変換される。
// TileXYZのZから拡張空間ID垂直インデックス(f)へは高度変換が用いられ変化する。
//
// 引数 :
//
//	request : 変換対象のTileXYZ構造体のスライス
//	zBaseExponent： zの高さが1mとなるズームレベル
//	zBaseOffset： ズームレベルがzBaseExponentのとき高度0mにおけるvZoom
//	outputVZoom : 拡張空間ID変換中の拡張空間IDの高さの精度指定。空間ID変換時、水平精度の方が高い場合この値は使われない。拡張空間IDの精度の閾値である 0 ～ 35 の整数値を指定可能。
//
// 戻り値 :
//
//	変換後の空間IDのスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 精度閾値超過(出力精度)：出力の水平方向精度、または高さ方向の精度に 0 ～ 35 の整数値以外が入力されていた場合。
//	 z高度範囲外：変換前のzがその垂直ズームレベルにおける高度範囲外である場合。
//	 拡張空間ID高度範囲外：拡張空間ID変換時の拡張空間ID高度がその垂直ズームレベルにおける高度範囲外である場合。
//
// 補足事項：
//
//	入力のoutputVZoomは次の2項目と関係しており、出力空間ID数に影響を与える
//	1. TileXYZの垂直精度, zBaseExponent, zBaseOffset: 条件により拡張空間ID変換後の拡張空間ID数が増加する(詳しくは ConvertTileXYZToExtendedSpatialIDs の例を参照)
//	2. TileXYZの水平精度の差: 差が1増えるごとに1で出力された拡張空間ID数の4のべき乗で出力空間ID数が増加する。
//	そのため、これら2項目の値によっては変換後の空間ID数は大幅に増大する。
//	動作環境によってはメモリ不足となる可能性があるため、注意すること。
func ConvertTileXYZToSpatialIDs(request []*object.TileXYZ, zBaseExponent uint16, zBaseOffset int64, outputVZoom uint16) ([]string, error) {
	var outputData []string
	outputExtendedSpatialIds, err := ConvertTileXYZToExtendedSpatialIDs(
		request, zBaseExponent, zBaseOffset, outputVZoom,
	)
	if err != nil {
		return nil, err
	}
	for _, out := range outputExtendedSpatialIds {
		spatialIds := ConvertExtendedSpatialIDToSpatialIDs(&out)
		outputData = append(outputData, spatialIds...)
	}
	return outputData, nil
}

// ConvertAltitudeKeyToZ altitudekeyを(拡張)空間IDのz成分に高度変換する。
//
// 変換前と変換後の精度差によって出力されるaltitudekeyの個数は増減する。
//
// 引数 :
//
//	altitudekey : 変換前のaltitudekeyの値
//
//	altitudekeyZoom : 変換前のaltitudekeyの精度指定
//
//	outputZoom : 変換対象の拡張空間ID垂直精度の指定
//
//	zBaseExponent : 変換対象の拡張空間IDの高さが1mとなるズームレベル
//
//	zBaseOffset : ズームレベルがzBaseExponentのとき、高度0mにおける拡張空間IDの垂直インデックス値
//
// 戻り値 :
//
//	変換後のaltitudekeyのスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力インデックス不正       ：inputIndexにそのズームレベル(inputZoom)で存在しないインデックス値が入力されていた場合。
//	 出力インデックス不正       ：出力altitudekeyが出力ズームレベル(outputZoom)で存在しないインデックス値になった場合。
func ConvertAltitudeKeyToZ(altitudekey int64, altitudekeyZoomLevel int64, outputZoom int64, zBaseExponent int64, zBaseOffset int64) (int64, int64, error) {
	// 1. check that the input index exists in the input system
	inputResolution := common.CalculateArithmeticShift(1, altitudekeyZoomLevel)

	maxInputIndex := inputResolution - 1
	minInputIndex := int64(0)

	if altitudekey > maxInputIndex || altitudekey < minInputIndex {
		return 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, "input index does not exist")
	}

	// 2. Calculate internal index
	zoomDifference := zBaseExponent - altitudekeyZoomLevel

	internalMinIndex := common.CalculateArithmeticShift(altitudekey, zoomDifference)
	internalMaxIndex := internalMinIndex
	if zoomDifference > 0 {
		internalMaxIndex = common.CalculateArithmeticShift(altitudekey+1, zoomDifference) - 1
	}
	// 3. Calculate outputMinIndex
	outputZoomDifference := outputZoom - 25
	outputMinIndex := common.CalculateArithmeticShift(internalMinIndex-zBaseOffset, outputZoomDifference)
	outputMaxIndex := common.CalculateArithmeticShift(internalMaxIndex-zBaseOffset, outputZoomDifference)
	if outputZoomDifference > 0 {
		outputMaxIndex = common.CalculateArithmeticShift(internalMaxIndex-zBaseOffset+1, outputZoomDifference) - 1
	}

	// 4. Check to make sure outputMinIndex exists in the output system
	outputResolution := common.CalculateArithmeticShift(1, outputZoom)

	maxOutputIndex := outputResolution - 1
	minOutputIndex := -outputResolution

	if outputMaxIndex > maxOutputIndex || outputMinIndex < minOutputIndex {
		return 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, "output index does not exist with given outputZoom, zBaseExponent, and zBaseOffset")
	}

	return outputMinIndex, outputMaxIndex, nil
}

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
//	outputHZoom : 入力値が変換後のQuadkeyの精度となる。quadkeyの精度の閾値である 1 ～ 31 の整数値を指定可能。
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

// ConvertZToAltitudekey (拡張)空間IDのz成分をaltitudekeyに高度変換する。
//
// 変換前と変換後の精度差によって出力されるaltitudekeyの個数は増減する。
//
// 引数 :
//
//	inputIndex : 変換対象の(拡張)空間IDのz成分(fインデックス)
//
//	inputZoom : 変換対象の(拡張)空間IDのズームレベル(zインデックス)
//
//	outputZoom : 変換後のaltitudekeyの精度指定
//
//	zBaseExponent : 変換後のaltitudekeyの高さが1mとなるズームレベル
//
//	zBaseOffset : ズームレベルがzBaseExponentのとき、高度0mにおけるaltitudekeyのインデックス値
//
// 戻り値 :
//
//	変換後のaltitudekeyのスライス
//
// 戻り値(エラー) :
//
//	以下の条件に当てはまる場合、エラーインスタンスが返却される。
//	 入力インデックス不正       ：inputIndexにそのズームレベル(inputZoom)で存在しないインデックス値が入力されていた場合。
//	 出力インデックス不正       ：出力altitudekeyが出力ズームレベル(outputZoom)で存在しないインデックス値になった場合。
func ConvertZToAltitudekey(inputIndex int64, inputZoom int64, outputZoom int64, zBaseExponent int64, zBaseOffset int64) ([]int64, error) {

	var (
		outputIndexes []int64
		error         error
	)

	// determine the upper and lower index bounds to search for matches in height solution space
	lowerBound, error := convertZToMinAltitudekey(inputIndex, inputZoom, outputZoom, zBaseExponent, zBaseOffset)
	if error != nil {
		return nil, error
	}
	upperBound, error := convertZToMinAltitudekey(inputIndex+1, inputZoom, outputZoom, zBaseExponent, zBaseOffset)
	if error != nil {
		return nil, error
	}

	// Determine the vertical index/indices to return.
	// a) always return the lowerBound index. Regardless of the difference between the inputZoom and outputZoom,
	// mathematically the altitude associated with the lower bounds will always satisfy the solution set.
	// b) cycle through indices from lowerBounds+1 to upperBounds with i to find any possible additional indexes
	// that satisfy the solution set.
	outputIndexes = append(outputIndexes, lowerBound)

	for i := lowerBound + 1; i < upperBound; i++ {

		outputIndexes = append(outputIndexes, int64(i))
	}

	return outputIndexes, nil

}

func convertZToMinAltitudekey(inputIndex int64, inputZoom int64, outputZoom int64, zBaseExponent int64, zBaseOffset int64) (int64, error) {

	// 1. check that the input index exists in the input system
	inputResolution := common.CalculateArithmeticShift(1, inputZoom)

	maxInputIndex := inputResolution - 1
	minInputIndex := -inputResolution

	if inputIndex > maxInputIndex || inputIndex < minInputIndex {
		return 0, errors.NewSpatialIdError(errors.InputValueErrorCode, "input index does not exist")
	}

	// 2. Calculate outputIndex
	outputIndex := common.CalculateArithmeticShift(inputIndex, -(inputZoom - zOriginValue))
	outputIndex += zBaseOffset
	outputIndex = common.CalculateArithmeticShift(outputIndex, (outputZoom - zBaseExponent))

	// 3. Check to make sure outputIndex exists in the output system
	outputResolution := common.CalculateArithmeticShift(1, outputZoom)

	maxOutputIndex := outputResolution - 1
	minOutputIndex := int64(0)

	if outputIndex > maxOutputIndex || outputIndex < minOutputIndex {
		return 0, errors.NewSpatialIdError(errors.InputValueErrorCode, "output index does not exist with given outputZoom, zBaseExponent, and zBaseOffset")
	}

	return outputIndex, nil

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

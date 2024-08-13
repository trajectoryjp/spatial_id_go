package detector

import (
	"fmt"
	"github.com/trajectoryjp/multidimensional-radix-tree/src/tree"
	"github.com/trajectoryjp/spatial_id_go/v4/common/consts"
	"github.com/trajectoryjp/spatial_id_go/v4/common/errors"
	"github.com/trajectoryjp/spatial_id_go/v4/integrate"
	"github.com/trajectoryjp/spatial_id_go/v4/transform"
	"strconv"
	"strings"
)

// CheckSpatialIdsOverlap 2つの空間ID重複の判定関数
//
// 比較対象として入力された2つの空間IDのズームレベルを変換して揃え、重複の判定を行う。
//
// 重複がある場合true、ない場合falseが返却される。
//
// 引数の入力値が空間IDフォーマットの仕様に違反していた場合、エラーインスタンスが返却される。
//
// 引数:
//
//	spatialId1, spatialId2: 重複判定対象の空間ID。ズームレベルが異なっている入力も許容。
//
// 戻り値:
//
//	bool:
//		重複の有無が返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。ただしこのときbool値にfalseが返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。
func CheckSpatialIdsOverlap(spatialId1 string, spatialId2 string) (bool, error) {
	ids1 := []string{spatialId1}
	ids2 := []string{spatialId2}
	return CheckSpatialIdsArrayOverlap(ids1, ids2)
}

// CheckSpatialIdsArrayOverlap 2つの空間ID重複の判定関数
//
// 比較対象として入力された2つの空間ID列の重複の判定を行う。
//
// 重複がある場合true、ない場合falseが返却される。
//
// 引数の入力値が空間IDフォーマットの仕様に違反していた場合、エラーインスタンスが返却される。
//
// 引数:
//
//	spatialIds1, spatialIds2: 重複判定対象の空間ID列。ズームレベルが異なっている入力も許容。
//
// ただし高度は16,777,216未満 〜 -16,777,216m以上に制限される(この範囲で高度変換を行う)
//
// 戻り値:
//
//	bool:
//		重複の有無が返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。このときbool値はfalseで返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。
//	 		空間ID高度範囲外：高度範囲外の空間IDが入力されていた場合。
func CheckSpatialIdsArrayOverlap(spatialIds1 []string, spatialIds2 []string) (bool, error) {
	// f,x,yで3次元分の2分木
	tr := tree.CreateTree(tree.Create3DTable())
	// spatialIds1から各要素をradix treeに格納
	for indexSpatialId1, spatialId1 := range spatialIds1 {
		zoom1, f1, x1, y1, err := getSpatialIdAttrs(spatialId1)
		if err != nil {
			return false, fmt.Errorf("%w @spatialId1[%v]", err, indexSpatialId1)
		}
		convertedFIndex, errAltConversion := transform.AddZBaseOffsetToZ(int64(f1), uint8(zoom1), consts.ZBaseOffsetForNegativeFIndex)
		if convertedFIndex < 0 {
			return false, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("input f-index %v is out of altitude range @spatialId1[%v] = %v", f1, indexSpatialId1, spatialId1))
		}
		if errAltConversion != nil {
			return false, fmt.Errorf("%w @spatialId1[%v] = %v", errAltConversion, indexSpatialId1, spatialId1)
		}
		index1 := tree.Indexs{convertedFIndex, int64(x1), int64(y1)}
		tr.Append(index1, tree.ZoomSetLevel(zoom1), spatialId1)
	}
	// spatialIds2から各要素の取り出し
	for indexSpatialId2, spatialId2 := range spatialIds2 {
		zoom2, f2, x2, y2, err := getSpatialIdAttrs(spatialId2)
		if err != nil {
			return false, fmt.Errorf("%w @spatialId2[%v]", err, indexSpatialId2)
		}
		// 取り出した要素の比較
		// 高度インデックスをオフセット変換のみ実行して自然数にする
		convertedFIndex2, errAltConversion := transform.AddZBaseOffsetToZ(int64(f2), uint8(zoom2), consts.ZBaseOffsetForNegativeFIndex)
		if convertedFIndex2 < 0 {
			return false, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("input f-index %v is out of altitude range @spatialId2[%v] = %v", f2, indexSpatialId2, spatialId2))
		}
		if errAltConversion != nil {
			return false, fmt.Errorf("%w @spatialId2[%v] = %v", errAltConversion, indexSpatialId2, spatialId2)
		}
		result := tr.IsOverlap(tree.Indexs{convertedFIndex2, int64(x2), int64(y2)}, tree.ZoomSetLevel(zoom2))
		if result {
			// 重複判定時、trueとnilを返却
			return result, nil
		}
	}

	return false, nil
}

// getSpatialIdAttrs 空間IDフォーマットチェック関数
//
// 入力された空間IDの各成分を格納した配列を返却する。
// 以下の条件に当てはまる場合はフォーマット違反となりエラーインスタンスが返却される。
//
//	・空間IDの各成分に数値が入力されていない場合
//	・区切り文字の数が3つで無い場合
//
// 引数：
//
//	spatialId：空間ID
//
// 戻り値：
//
//	空間IDに含まれる精度(ズームレベル), F, X, Y成分を格納した以下のtuple
//	 (精度(ズームレベル), F成分, X成分, Y成分)
//	入力された拡張空間IDのフォーマットが不正な場合、戻り値を0にしエラーインスタンスを返却。
func getSpatialIdAttrs(spatialId string) (int, int, int, int, error) {
	// 分割
	spatialIdAttributes := strings.Split(spatialId, consts.SpatialIDDelimiter)
	if len(spatialIdAttributes) != 4 {
		// 不正形式(要素数)
		return 0, 0, 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("spatialId: %v", spatialId))
	}
	var errNumberConversion error
	zoom, errNumberConversion := strconv.Atoi(spatialIdAttributes[0])
	f, errNumberConversion := strconv.Atoi(spatialIdAttributes[1])
	x, errNumberConversion := strconv.Atoi(spatialIdAttributes[2])
	y, errNumberConversion := strconv.Atoi(spatialIdAttributes[3])
	// 不正形式(数値)
	if errNumberConversion != nil {
		return 0, 0, 0, 0, errors.NewSpatialIdError(errors.InputValueErrorCode, fmt.Sprintf("spatialId: %v", spatialId))
	}
	return zoom, f, x, y, nil
}

// CheckExtendedSpatialIdsOverlap 2つの拡張空間IDの重複の判定関数
//
// 比較対象として入力された2つの拡張空間IDのズームレベルを変換して揃え、重複の判定を行う。
//
// 重複がある場合true、ない場合falseが返却される。
//
// 引数の入力値が拡張空間IDフォーマットの仕様に違反していた場合、エラーインスタンスが返却される。
//
// 引数:
//
//	extendedSpatialId1, extendedSpatialId2 : 重複判定対象の拡張空間ID。ズームレベルが異なっている入力も許容。
//
// 戻り値:
//
//	bool:
//		重複の有無がbool値で返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。ただしこのときfalseも返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。
func CheckExtendedSpatialIdsOverlap(extendedSpatialId1 string, extendedSpatialId2 string) (bool, error) {
	// ズームレベルの取り出し
	arr1 := strings.Split(extendedSpatialId1, "/")
	arr2 := strings.Split(extendedSpatialId2, "/")
	if len(arr1) != 5 || len(arr2) != 5 {
		// 不正形式
		return false, fmt.Errorf("invalid format. extendedSpatialId1: %v, extendedSpatialId2: %v", extendedSpatialId1, extendedSpatialId2)
	}

	// 変換後の水平方向ズームレベルの決定
	horizontalZoom1, _ := strconv.Atoi(arr1[0])
	horizontalZoom2, _ := strconv.Atoi(arr2[0])
	// zoomレベルで場合分け
	targetHorizontalZoom := horizontalZoom1
	if horizontalZoom1 > horizontalZoom2 {
		targetHorizontalZoom = horizontalZoom2
	}

	// 変換後の垂直方向のズームレベルの設定
	verticalZoom1, _ := strconv.Atoi(arr1[3])
	verticalZoom2, _ := strconv.Atoi(arr2[3])
	targetVerticalZoom := verticalZoom1
	// zoomレベルで場合分け
	if verticalZoom1 > verticalZoom2 {
		targetVerticalZoom = verticalZoom2
	}

	// 変換後のZoomレベルを指定して変換
	extendedSpatialIds1, err := integrate.ChangeExtendedSpatialIdsZoom([]string{extendedSpatialId1}, int64(targetHorizontalZoom), int64(targetVerticalZoom))
	if err != nil {
		// 変換エラー
		return false, err
	}
	extendedSpatialIds2, err := integrate.ChangeExtendedSpatialIdsZoom([]string{extendedSpatialId2}, int64(targetHorizontalZoom), int64(targetVerticalZoom))
	if err != nil {
		// 変換エラー
		return false, err
	}

	return extendedSpatialIds1[0] == extendedSpatialIds2[0], nil

}

// CheckExtendedSpatialIdsArrayOverlap 2つの拡張空間IDの重複の判定関数
//
// 比較対象として入力された2つの拡張空間ID列の重複の判定を行う。
//
// 重複がある場合true、ない場合falseが返却される。
//
// 引数の入力値が拡張空間IDフォーマットの仕様に違反していた場合、エラーインスタンスが返却される。
//
// 引数:
//
//	extendedSpatialIds1, extendedSpatialIds2 : 重複判定対象の拡張空間ID列。ズームレベルが異なっている入力も許容。
//
// 戻り値:
//
//	bool:
//		重複の有無がbool値で返却される。true: 重複あり false: 重複なし
//
//	error:
//		以下の条件に当てはまる場合、エラーインスタンスが返却される。ただしこのときfalseも返却される。
//	 		空間IDフォーマット不正：空間IDのフォーマットに違反する値が"重複判定対象空間ID"に入力されていた場合。

func CheckExtendedSpatialIdsArrayOverlap(extendedSpatialIds1 []string, extendedSpatialIds2 []string) (bool, error) {
	// spatialIds1から各要素の取り出し
	for _, extendedSpatialId1 := range extendedSpatialIds1 {
		// spatialIds2から各要素の取り出し
		for _, extendedSpatialId2 := range extendedSpatialIds2 {
			// 取り出した要素の比較
			result, err := CheckExtendedSpatialIdsOverlap(extendedSpatialId1, extendedSpatialId2)
			if err != nil {
				// エラー発生時、falseとerrorを返却
				return false, err
			}
			if result {
				// 重複判定時、trueとnilを返却
				return result, nil
			}
		}
	}
	return false, nil
}
